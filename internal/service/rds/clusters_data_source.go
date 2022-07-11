package rds

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
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
						"arn": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"cluster_identifier": {
							Type:     schema.TypeString,
							Required: true,
						},

						"availability_zones": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Computed: true,
							Set:      schema.HashString,
						},

						"backtrack_window": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"backup_retention_period": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"cluster_members": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Computed: true,
							Set:      schema.HashString,
						},

						"cluster_resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"database_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"db_subnet_group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"db_cluster_parameter_group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"enabled_cloudwatch_logs_exports": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},

						"endpoint": {
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

						"final_snapshot_identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"iam_database_authentication_enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"iam_roles": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},

						"kms_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"master_username": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"preferred_backup_window": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"preferred_maintenance_window": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"port": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"reader_endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"hosted_zone_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"replication_source_identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"storage_encrypted": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"tags": tftags.TagsSchemaComputed(),

						"vpc_security_group_ids": {
							Type:     schema.TypeSet,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Set:      schema.HashString,
						},
					},
				},
			},
		},
	}
}

func dataSourceClustersRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).RDSConn
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig
	var clusters []interface{}
	var clusters_paginate []*rds.DBCluster

	params := &rds.DescribeDBClustersInput{}
	log.Printf("[DEBUG] Reading RDS Cluster: %s", params)

	for true {

		resp, err := conn.DescribeDBClusters(params)
		if err != nil {
			return fmt.Errorf("Error retrieving RDS cluster: %w", err)
		}

		clusters_paginate = append(clusters_paginate, resp.DBClusters...)

		if resp.Marker != nil {
			params.Marker = resp.Marker
		} else {
			break
		}
	}

	if clusters_paginate == nil || len(clusters_paginate) < 1 || clusters_paginate[0] == nil {
		d.SetId("db_clusters")
		d.Set("clusters", clusters)
		return nil
	}

	// var dbc *rds.DBCluster
	filter := d.Get("filter").(*schema.Set).List()

	for _, dbc := range clusters_paginate {

		cluster := make(map[string]interface{})

		cluster["availability_zones"] = aws.StringValueSlice(dbc.AvailabilityZones)
		cluster["arn"] = dbc.DBClusterArn
		cluster["backtrack_window"] = dbc.BacktrackWindow
		cluster["backup_retention_period"] = dbc.BackupRetentionPeriod
		cluster["cluster_identifier"] = dbc.DBClusterIdentifier

		var cm []string
		for _, m := range dbc.DBClusterMembers {
			cm = append(cm, aws.StringValue(m.DBInstanceIdentifier))
		}

		cluster["cluster_members"] = cm
		cluster["cluster_resource_id"] = dbc.DbClusterResourceId

		// Only set the DatabaseName if it is not nil. There is a known API bug where
		// RDS accepts a DatabaseName but does not return it, causing a perpetual
		// diff.
		//	See https://github.com/hashicorp/terraform/issues/4671 for backstory
		if dbc.DatabaseName != nil {
			cluster["database_name"] = dbc.DatabaseName
		}

		cluster["db_cluster_parameter_group_name"] = dbc.DBClusterParameterGroup
		cluster["db_subnet_group_name"] = dbc.DBSubnetGroup
		cluster["enabled_cloudwatch_logs_exports"] = aws.StringValueSlice(dbc.EnabledCloudwatchLogsExports)
		cluster["endpoint"] = dbc.Endpoint
		cluster["engine_version"] = dbc.EngineVersion
		cluster["engine"] = dbc.Engine
		cluster["hosted_zone_id"] = dbc.HostedZoneId
		cluster["iam_database_authentication_enabled"] = dbc.IAMDatabaseAuthenticationEnabled

		var roles []string
		for _, r := range dbc.AssociatedRoles {
			roles = append(roles, aws.StringValue(r.RoleArn))
		}

		cluster["iam_roles"] = roles
		cluster["kms_key_id"] = dbc.KmsKeyId
		cluster["master_username"] = dbc.MasterUsername
		cluster["port"] = dbc.Port
		cluster["preferred_backup_window"] = dbc.PreferredBackupWindow
		cluster["preferred_maintenance_window"] = dbc.PreferredMaintenanceWindow
		cluster["reader_endpoint"] = dbc.ReaderEndpoint
		cluster["replication_source_identifier"] = dbc.ReplicationSourceIdentifier
		cluster["storage_encrypted"] = dbc.StorageEncrypted

		var vpcg []string
		for _, g := range dbc.VpcSecurityGroups {
			vpcg = append(vpcg, aws.StringValue(g.VpcSecurityGroupId))
		}

		cluster["vpc_security_group_ids"] = vpcg

		tags, err := ListTags(conn, *dbc.DBClusterArn)

		if err != nil {
			return fmt.Errorf("error listing tags for RDS Cluster (%s): %w", *dbc.DBClusterArn, err)
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

	d.SetId("db_clusters")
	d.Set("clusters", clusters)

	return nil
}
