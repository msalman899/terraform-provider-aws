package rds

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func DataSourceInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInstancesRead,

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
			"instances": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tags": tftags.TagsSchemaComputed(),

						"address": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"db_instance_identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"allocated_storage": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"auto_minor_version_upgrade": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"availability_zone": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"backup_retention_period": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"db_cluster_identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"db_instance_arn": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"db_instance_class": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"db_name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"db_parameter_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},

						"db_security_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},

						"db_subnet_group": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"db_instance_port": {
							Type:     schema.TypeInt,
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

						"hosted_zone_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"iops": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"kms_key_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"license_model": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"master_username": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"monitoring_interval": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"monitoring_role_arn": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"multi_az": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"option_group_memberships": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},

						"port": {
							Type:     schema.TypeInt,
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

						"publicly_accessible": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"resource_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"storage_encrypted": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"storage_type": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"timezone": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"vpc_security_groups": {
							Type:     schema.TypeList,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},

						"replicate_source_db": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"ca_cert_identifier": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceInstancesRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).RDSConn
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	opts := &rds.DescribeDBInstancesInput{}

	resp, err := conn.DescribeDBInstances(opts)
	if err != nil {
		return err
	}

	if resp == nil || len(resp.DBInstances) < 1 || resp.DBInstances[0] == nil {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	filter := d.Get("filter").(*schema.Set).List()

	var instances []interface{}
	for _, dbInstance := range resp.DBInstances {

		instance := make(map[string]interface{})

		instance["allocated_storage"] = dbInstance.AllocatedStorage
		instance["auto_minor_version_upgrade"] = dbInstance.AutoMinorVersionUpgrade
		instance["availability_zone"] = dbInstance.AvailabilityZone
		instance["backup_retention_period"] = dbInstance.BackupRetentionPeriod
		instance["db_cluster_identifier"] = dbInstance.DBClusterIdentifier
		instance["db_instance_arn"] = dbInstance.DBInstanceArn
		instance["db_instance_class"] = dbInstance.DBInstanceClass
		instance["db_name"] = dbInstance.DBName
		instance["resource_id"] = dbInstance.DbiResourceId
		instance["allocated_storage"] = dbInstance.AllocatedStorage
		instance["auto_minor_version_upgrade"] = dbInstance.AutoMinorVersionUpgrade
		instance["availability_zone"] = dbInstance.AvailabilityZone
		instance["backup_retention_period"] = dbInstance.BackupRetentionPeriod
		instance["db_cluster_identifier"] = dbInstance.DBClusterIdentifier
		instance["db_instance_arn"] = dbInstance.DBInstanceArn
		instance["db_instance_class"] = dbInstance.DBInstanceClass
		instance["db_instance_identifier"] = dbInstance.DBInstanceIdentifier
		instance["resource_id"] = dbInstance.DbiResourceId

		var parameterGroups []string
		for _, v := range dbInstance.DBParameterGroups {
			parameterGroups = append(parameterGroups, aws.StringValue(v.DBParameterGroupName))
		}
		instance["db_parameter_groups"] = parameterGroups

		var dbSecurityGroups []string
		for _, v := range dbInstance.DBSecurityGroups {
			dbSecurityGroups = append(dbSecurityGroups, aws.StringValue(v.DBSecurityGroupName))
		}
		instance["db_security_groups"] = dbSecurityGroups

		if dbInstance.DBSubnetGroup != nil {
			instance["db_subnet_group"] = dbInstance.DBSubnetGroup.DBSubnetGroupName
		} else {
			instance["db_subnet_group"] = ""
		}

		instance["db_instance_port"] = dbInstance.DbInstancePort
		instance["engine"] = dbInstance.Engine
		instance["engine_version"] = dbInstance.EngineVersion
		instance["iops"] = dbInstance.Iops
		instance["kms_key_id"] = dbInstance.KmsKeyId
		instance["license_model"] = dbInstance.LicenseModel
		instance["master_username"] = dbInstance.MasterUsername
		instance["monitoring_interval"] = dbInstance.MonitoringInterval
		instance["monitoring_role_arn"] = dbInstance.MonitoringRoleArn
		instance["multi_az"] = dbInstance.MultiAZ

		// Per AWS SDK Go docs:
		// The endpoint might not be shown for instances whose status is creating.
		if dbEndpoint := dbInstance.Endpoint; dbEndpoint != nil {
			instance["address"] = dbEndpoint.Address
			instance["port"] = dbEndpoint.Port
			instance["hosted_zone_id"] = dbEndpoint.HostedZoneId
			instance["endpoint"] = fmt.Sprintf("%s:%d", aws.StringValue(dbEndpoint.Address), aws.Int64Value(dbEndpoint.Port))

		} else {
			instance["address"] = nil
			instance["port"] = nil
			instance["hosted_zone_id"] = nil
			instance["endpoint"] = nil
		}

		instance["enabled_cloudwatch_logs_exports"] = aws.StringValueSlice(dbInstance.EnabledCloudwatchLogsExports)

		var optionGroups []string
		for _, v := range dbInstance.OptionGroupMemberships {
			optionGroups = append(optionGroups, aws.StringValue(v.OptionGroupName))
		}

		instance["option_group_memberships"] = optionGroups
		instance["preferred_backup_window"] = dbInstance.PreferredBackupWindow
		instance["preferred_maintenance_window"] = dbInstance.PreferredMaintenanceWindow
		instance["publicly_accessible"] = dbInstance.PubliclyAccessible
		instance["storage_encrypted"] = dbInstance.StorageEncrypted
		instance["storage_type"] = dbInstance.StorageType
		instance["timezone"] = dbInstance.Timezone
		instance["replicate_source_db"] = dbInstance.ReadReplicaSourceDBInstanceIdentifier
		instance["ca_cert_identifier"] = dbInstance.CACertificateIdentifier

		var vpcSecurityGroups []string
		for _, v := range dbInstance.VpcSecurityGroups {
			vpcSecurityGroups = append(vpcSecurityGroups, aws.StringValue(v.VpcSecurityGroupId))
		}

		instance["vpc_security_groups"] = vpcSecurityGroups

		tags, err := ListTags(conn, *dbInstance.DBInstanceArn)

		if err != nil {
			return fmt.Errorf("error listing tags for RDS DB Instance (%s): %w", *dbInstance.DBInstanceArn, err)
		}

		instance["tags"] = tags.IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()

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
				instances = append(instances, instance)
			}

		} else {
			instances = append(instances, instance)
		}

	}

	d.SetId("db_instances")
	d.Set("instances", instances)

	return nil
}
