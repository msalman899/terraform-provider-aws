package rds

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

func DataSourceRds() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRdsRead,

		Schema: map[string]*schema.Schema{
			"rds_identifiers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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

func dataSourceRdsRead(d *schema.ResourceData, meta interface{}) error {
	conn := meta.(*conns.AWSClient).RDSConn
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig
	var instances []interface{}
	var instances_paginate []*rds.DBInstance

	var clusters []interface{}
	var clusters_paginate []*rds.DBCluster

	rds_identifiers := d.Get("rds_identifiers").([]interface{})

	for _, dbId := range rds_identifiers {
		opts := &rds.DescribeDBInstancesInput{
			DBInstanceIdentifier: aws.String(dbId.(string)),
		}

		ins_resp, err := conn.DescribeDBInstances(opts)
		if err != nil && !strings.Contains(err.Error(), "DBInstanceNotFound") {
			return err
		}

		instances_paginate = append(instances_paginate, ins_resp.DBInstances...)

		params := &rds.DescribeDBClustersInput{
			DBClusterIdentifier: aws.String(dbId.(string)),
		}
		cls_resp, err := conn.DescribeDBClusters(params)
		if err != nil && !strings.Contains(err.Error(), "DBClusterNotFound") {
			return err
		}

		clusters_paginate = append(clusters_paginate, cls_resp.DBClusters...)

	}

	for _, dbInstance := range instances_paginate {

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

		instances = append(instances, instance)

	}

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

		clusters = append(clusters, cluster)

	}

	d.SetId("db_clusters")
	d.Set("clusters", clusters)

	d.SetId("db_instances")
	d.Set("instances", instances)

	return nil
}
