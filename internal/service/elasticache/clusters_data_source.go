package elasticache

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

func DataSourceClusters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceClustersRead,

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
			"clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_id": {
							Type:     schema.TypeString,
							Required: true,
							StateFunc: func(v interface{}) string {
								value := v.(string)
								return strings.ToLower(value)
							},
						},

						"node_type": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"num_cache_nodes": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"subnet_group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"engine": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"engine_version": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"parameter_group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"replication_group_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"security_group_names": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},

						"security_group_ids": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
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
						"maintenance_window": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"snapshot_window": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"snapshot_retention_limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"availability_zone": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"notification_topic_arn": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"configuration_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"cluster_address": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"arn": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"cache_nodes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"address": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"availability_zone": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"tags": tftags.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceClustersRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).ElastiCacheConn
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	var clusters []interface{}
	// clusterID := d.Get("cluster_id").(string)
	clusters_paginate, err := FindCacheClusters(conn)
	// if tfresource.NotFound(err) {
	// 	return fmt.Errorf("Your query returned no results. Please change your search criteria and try again")
	// }
	if err != nil {
		return fmt.Errorf("error reading ElastiCache Cache Cluster: %w", err)
	}

	if clusters_paginate == nil || len(clusters_paginate) < 1 || clusters_paginate[0] == nil {
		d.SetId("redis_clusters")
		d.Set("clusters", clusters)
		return nil
	}

	filter := d.Get("filter").(*schema.Set).List()

	for _, cc := range clusters_paginate {

		cluster := make(map[string]interface{})

		// d.SetId(aws.StringValue(cc.CacheClusterId))

		cluster["cluster_id"] = cc.CacheClusterId
		cluster["node_type"] = cc.CacheNodeType
		cluster["num_cache_nodes"] = cc.NumCacheNodes
		cluster["subnet_group_name"] = cc.CacheSubnetGroupName
		cluster["engine"] = cc.Engine
		cluster["engine_version"] = cc.EngineVersion
		cluster["security_group_names"] = flattenSecurityGroupNames(cc.CacheSecurityGroups)
		cluster["security_group_ids"] = flattenSecurityGroupIDs(cc.SecurityGroups)

		if cc.CacheParameterGroup != nil {
			cluster["parameter_group_name"] = cc.CacheParameterGroup.CacheParameterGroupName
		}

		if cc.ReplicationGroupId != nil {
			cluster["replication_group_id"] = cc.ReplicationGroupId
		}

		cluster["log_delivery_configuration"] = flattenLogDeliveryConfigurations(cc.LogDeliveryConfigurations)
		cluster["maintenance_window"] = cc.PreferredMaintenanceWindow
		cluster["snapshot_window"] = cc.SnapshotWindow
		cluster["snapshot_retention_limit"] = cc.SnapshotRetentionLimit
		cluster["availability_zone"] = cc.PreferredAvailabilityZone

		if cc.NotificationConfiguration != nil {
			if aws.StringValue(cc.NotificationConfiguration.TopicStatus) == "active" {
				cluster["notification_topic_arn"] = cc.NotificationConfiguration.TopicArn
			}
		}

		if cc.ConfigurationEndpoint != nil {
			cluster["port"] = cc.ConfigurationEndpoint.Port
			cluster["configuration_endpoint"] = aws.String(fmt.Sprintf("%s:%d", *cc.ConfigurationEndpoint.Address, *cc.ConfigurationEndpoint.Port))
			cluster["cluster_address"] = aws.String(*cc.ConfigurationEndpoint.Address)
		}

		cacheNodeData, err := setCacheNodeDataCust(cc)

		if err != nil {
			return err
		}

		cluster["cache_nodes"] = cacheNodeData
		cluster["arn"] = cc.ARN

		tags, err := ListTags(conn, aws.StringValue(cc.ARN))

		if err != nil && !verify.CheckISOErrorTagsUnsupported(conn.PartitionID, err) {
			return fmt.Errorf("error listing tags for Elasticache Cluster (%s): %w", d.Id(), err)
		}

		if err != nil {
			log.Printf("[WARN] error listing tags for Elasticache Cluster (%s): %s", d.Id(), err)
		}

		cluster["tags"] = tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()

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
				clusters = append(clusters, cluster)
			}

		} else {
			clusters = append(clusters, cluster)
		}
	}

	d.SetId("redis_clusters")
	d.Set("clusters", clusters)

	return nil
}
