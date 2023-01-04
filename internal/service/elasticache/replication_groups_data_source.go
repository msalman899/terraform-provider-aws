package elasticache

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func DataSourceReplicationGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceReplicationGroupsRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},

						"values": {
							Type:     schema.TypeList,
							Required: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"replication_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"replication_group_id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateReplicationGroupID,
						},
						"replication_group_description": {
							Type:       schema.TypeString,
							Computed:   true,
							Deprecated: "Use description instead",
						},
						"arn": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_token_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"automatic_failover_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"configuration_endpoint_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"primary_endpoint_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"reader_endpoint_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"num_cache_clusters": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"num_node_groups": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"number_cache_clusters": {
							Type:       schema.TypeInt,
							Computed:   true,
							Deprecated: "Use num_cache_clusters instead",
						},
						"member_clusters": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"multi_az_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"node_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"replicas_per_node_group": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"log_delivery_configuration": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"destination": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"log_format": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"log_type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"snapshot_window": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"snapshot_retention_limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"tags": tftags.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceReplicationGroupsRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ElastiCacheConn
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	var replication_groups []interface{}
	replication_groups_paginate, err := FindReplicationGroups(conn)
	if err != nil {
		return fmt.Errorf("error reading ElastiCache Replication Group %w", err)
	}

	if replication_groups_paginate == nil || len(replication_groups_paginate) < 1 || replication_groups_paginate[0] == nil {
		d.SetId("redis_replication_groups")
		d.Set("replication_groups", replication_groups)
		return nil
	}

	filter := d.Get("filter").(*schema.Set).List()

	for _, rg := range replication_groups_paginate {

		replication_group := make(map[string]interface{})

		replication_group["replication_group_id"] = aws.StringValue(rg.ReplicationGroupId)
		replication_group["description"] = rg.Description
		replication_group["replication_group_description"] = rg.Description
		replication_group["arn"] = rg.ARN
		replication_group["auth_token_enabled"] = rg.AuthTokenEnabled

		if rg.AutomaticFailover != nil {
			switch aws.StringValue(rg.AutomaticFailover) {
			case elasticache.AutomaticFailoverStatusDisabled, elasticache.AutomaticFailoverStatusDisabling:
				replication_group["automatic_failover_enabled"] = false
			case elasticache.AutomaticFailoverStatusEnabled, elasticache.AutomaticFailoverStatusEnabling:
				replication_group["automatic_failover_enabled"] = true
			}
		}

		if rg.MultiAZ != nil {
			switch strings.ToLower(aws.StringValue(rg.MultiAZ)) {
			case elasticache.MultiAZStatusEnabled:
				replication_group["multi_az_enabled"] = true
			case elasticache.MultiAZStatusDisabled:
				replication_group["multi_az_enabled"] = false
			default:
				log.Printf("Unknown MultiAZ state %q", aws.StringValue(rg.MultiAZ))
			}
		}

		if rg.ConfigurationEndpoint != nil {
			replication_group["port"] = rg.ConfigurationEndpoint.Port
			replication_group["configuration_endpoint_address"] = rg.ConfigurationEndpoint.Address
		} else {
			if rg.NodeGroups == nil {
				d.SetId("")
				return fmt.Errorf("ElastiCache Replication Group (%s) doesn't have node groups", aws.StringValue(rg.ReplicationGroupId))
			}
			replication_group["port"] = rg.NodeGroups[0].PrimaryEndpoint.Port
			replication_group["primary_endpoint_address"] = rg.NodeGroups[0].PrimaryEndpoint.Address
			replication_group["reader_endpoint_address"] = rg.NodeGroups[0].ReaderEndpoint.Address
		}

		replication_group["num_cache_clusters"] = len(rg.MemberClusters)
		replication_group["number_cache_clusters"] = len(rg.MemberClusters)
		replication_group["member_clusters"] = flex.FlattenStringList(rg.MemberClusters)

		replication_group["node_type"] = rg.CacheNodeType
		replication_group["num_node_groups"] = len(rg.NodeGroups)
		replication_group["replicas_per_node_group"] = len(rg.NodeGroups[0].NodeGroupMembers) - 1
		replication_group["log_delivery_configuration"] = flattenLogDeliveryConfigurations(rg.LogDeliveryConfigurations)
		replication_group["snapshot_window"] = rg.SnapshotWindow
		replication_group["snapshot_retention_limit"] = rg.SnapshotRetentionLimit

		tags, err := ListTags(conn, aws.StringValue(rg.ARN))

		if err != nil {
			log.Printf("[WARN] error listing tags for Elasticache Replication Group %s", err)
		}

		replication_group["tags"] = tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()

		if len(filter) > 0 {
			match_count := 0

			for _, v := range filter {
				var filter_tag_name string
				var filter_tag_values []string

				filter_map := v.(map[string]interface{})
				filter_tag_name = filter_map["name"].(string)
				for _, e := range filter_map["values"].([]interface{}) {
					filter_tag_values = append(filter_tag_values, e.(string))
				}

				if tagValue, isTagKeyPresent := tags[filter_tag_name]; isTagKeyPresent {
					if contains(filter_tag_values, *tagValue.Value) {
						match_count++
					}
				}

			}

			if match_count == len(filter) {
				replication_groups = append(replication_groups, replication_group)
			}

		} else {
			replication_groups = append(replication_groups, replication_group)
		}
	}

	d.SetId("redis_replication_groups")
	d.Set("replication_groups", replication_groups)

	return nil
}
