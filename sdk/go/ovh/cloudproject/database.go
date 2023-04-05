// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// Minimum settings for each engine (region choice is up to the user):
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := CloudProject.NewDatabase(ctx, "cassandradb", &CloudProject.DatabaseArgs{
//				ServiceName: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Description: pulumi.String("my-first-cassandra"),
//				Engine:      pulumi.String("cassandra"),
//				Version:     pulumi.String("4.0"),
//				Plan:        pulumi.String("essential"),
//				Nodes: cloudproject.DatabaseNodeArray{
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("BHS"),
//					},
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("BHS"),
//					},
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("BHS"),
//					},
//				},
//				Flavor: pulumi.String("db1-4"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = CloudProject.NewDatabase(ctx, "kafkadb", &CloudProject.DatabaseArgs{
//				ServiceName:  pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Description:  pulumi.String("my-first-kafka"),
//				Engine:       pulumi.String("kafka"),
//				Version:      pulumi.String("3.1"),
//				Plan:         pulumi.String("business"),
//				KafkaRestApi: pulumi.Bool(true),
//				Nodes: cloudproject.DatabaseNodeArray{
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("DE"),
//					},
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("DE"),
//					},
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("DE"),
//					},
//				},
//				Flavor: pulumi.String("db1-4"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = CloudProject.NewDatabase(ctx, "m3db", &CloudProject.DatabaseArgs{
//				ServiceName: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Description: pulumi.String("my-first-m3db"),
//				Engine:      pulumi.String("m3db"),
//				Version:     pulumi.String("1.2"),
//				Plan:        pulumi.String("essential"),
//				Nodes: cloudproject.DatabaseNodeArray{
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("BHS"),
//					},
//				},
//				Flavor: pulumi.String("db1-7"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = CloudProject.NewDatabase(ctx, "mongodb", &CloudProject.DatabaseArgs{
//				ServiceName: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Description: pulumi.String("my-first-mongodb"),
//				Engine:      pulumi.String("mongodb"),
//				Version:     pulumi.String("5.0"),
//				Plan:        pulumi.String("essential"),
//				Nodes: cloudproject.DatabaseNodeArray{
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("GRA"),
//					},
//				},
//				Flavor: pulumi.String("db1-2"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = CloudProject.NewDatabase(ctx, "mysqldb", &CloudProject.DatabaseArgs{
//				ServiceName: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Description: pulumi.String("my-first-mysql"),
//				Engine:      pulumi.String("mysql"),
//				Version:     pulumi.String("8"),
//				Plan:        pulumi.String("essential"),
//				Nodes: cloudproject.DatabaseNodeArray{
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("SBG"),
//					},
//				},
//				Flavor: pulumi.String("db1-4"),
//				AdvancedConfiguration: pulumi.StringMap{
//					"mysql.sql_mode":                pulumi.String("ANSI,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION,NO_ZERO_DATE,NO_ZERO_IN_DATE,STRICT_ALL_TABLES"),
//					"mysql.sql_require_primary_key": pulumi.String("true"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = CloudProject.NewDatabase(ctx, "opensearchdb", &CloudProject.DatabaseArgs{
//				ServiceName:           pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Description:           pulumi.String("my-first-opensearch"),
//				Engine:                pulumi.String("opensearch"),
//				Version:               pulumi.String("1"),
//				Plan:                  pulumi.String("essential"),
//				OpensearchAclsEnabled: pulumi.Bool(true),
//				Nodes: cloudproject.DatabaseNodeArray{
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("UK"),
//					},
//				},
//				Flavor: pulumi.String("db1-4"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = CloudProject.NewDatabase(ctx, "pgsqldb", &CloudProject.DatabaseArgs{
//				ServiceName: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Description: pulumi.String("my-first-postgresql"),
//				Engine:      pulumi.String("postgresql"),
//				Version:     pulumi.String("14"),
//				Plan:        pulumi.String("essential"),
//				Nodes: cloudproject.DatabaseNodeArray{
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("WAW"),
//					},
//				},
//				Flavor: pulumi.String("db1-4"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = CloudProject.NewDatabase(ctx, "redisdb", &CloudProject.DatabaseArgs{
//				ServiceName: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Description: pulumi.String("my-first-redis"),
//				Engine:      pulumi.String("redis"),
//				Version:     pulumi.String("6.2"),
//				Plan:        pulumi.String("essential"),
//				Nodes: cloudproject.DatabaseNodeArray{
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("BHS"),
//					},
//				},
//				Flavor: pulumi.String("db1-4"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// To deploy a business PostgreSQL service with two nodes on public network:
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := CloudProject.NewDatabase(ctx, "postgresql", &CloudProject.DatabaseArgs{
//				Description: pulumi.String("my-first-postgresql"),
//				Engine:      pulumi.String("postgresql"),
//				Flavor:      pulumi.String("db1-15"),
//				Nodes: cloudproject.DatabaseNodeArray{
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("GRA"),
//					},
//					&cloudproject.DatabaseNodeArgs{
//						Region: pulumi.String("GRA"),
//					},
//				},
//				Plan:        pulumi.String("business"),
//				ServiceName: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Version:     pulumi.String("14"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// To deploy an enterprise MongoDB service with three nodes on private network:
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := CloudProject.NewDatabase(ctx, "mongodb", &CloudProject.DatabaseArgs{
//				Description: pulumi.String("my-first-mongodb"),
//				Engine:      pulumi.String("mongodb"),
//				Flavor:      pulumi.String("db1-30"),
//				Nodes: cloudproject.DatabaseNodeArray{
//					&cloudproject.DatabaseNodeArgs{
//						NetworkId: pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//						Region:    pulumi.String("SBG"),
//						SubnetId:  pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//					},
//					&cloudproject.DatabaseNodeArgs{
//						NetworkId: pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//						Region:    pulumi.String("SBG"),
//						SubnetId:  pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//					},
//					&cloudproject.DatabaseNodeArgs{
//						NetworkId: pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//						Region:    pulumi.String("SBG"),
//						SubnetId:  pulumi.String("XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"),
//					},
//				},
//				Plan:        pulumi.String("enterprise"),
//				ServiceName: pulumi.String("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
//				Version:     pulumi.String("5.0"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// OVHcloud Managed database clusters can be imported using the `service_name`, `engine`, `id` of the cluster, separated by "/" E.g., bash
//
// ```sh
//
//	$ pulumi import ovh:CloudProject/database:Database my_database_cluster service_name/engine/id
//
// ```
type Database struct {
	pulumi.CustomResourceState

	// Advanced configuration key / value.
	AdvancedConfiguration pulumi.StringMapOutput `pulumi:"advancedConfiguration"`
	// Time on which backups start every day.
	BackupTime pulumi.StringOutput `pulumi:"backupTime"`
	// Date of the creation of the cluster.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Small description of the database service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The disk size (in GB) of the database service.
	DiskSize pulumi.IntOutput `pulumi:"diskSize"`
	// Defines the disk type of the database service.
	DiskType pulumi.StringOutput `pulumi:"diskType"`
	// List of all endpoints objects of the service.
	Endpoints DatabaseEndpointArrayOutput `pulumi:"endpoints"`
	// The database engine you want to deploy. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringOutput `pulumi:"engine"`
	// A valid OVHcloud public cloud database flavor name in which the nodes will be started.
	// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
	// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
	Flavor pulumi.StringOutput `pulumi:"flavor"`
	// Defines whether the REST API is enabled on a kafka cluster
	KafkaRestApi pulumi.BoolPtrOutput `pulumi:"kafkaRestApi"`
	// Time on which maintenances can start every day.
	MaintenanceTime pulumi.StringOutput `pulumi:"maintenanceTime"`
	// Type of network of the cluster.
	NetworkType pulumi.StringOutput `pulumi:"networkType"`
	// List of nodes object.
	// Multi region cluster are not yet available, all node should be identical.
	Nodes DatabaseNodeArrayOutput `pulumi:"nodes"`
	// Defines whether the ACLs are enabled on an OpenSearch cluster
	OpensearchAclsEnabled pulumi.BoolPtrOutput `pulumi:"opensearchAclsEnabled"`
	// Plan of the cluster.
	// Enum: "essential", "business", "enterprise".
	Plan pulumi.StringOutput `pulumi:"plan"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Current status of the cluster.
	Status pulumi.StringOutput `pulumi:"status"`
	// The version of the engine in which the service should be deployed
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewDatabase registers a new resource with the given unique name, arguments, and options.
func NewDatabase(ctx *pulumi.Context,
	name string, args *DatabaseArgs, opts ...pulumi.ResourceOption) (*Database, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Engine == nil {
		return nil, errors.New("invalid value for required argument 'Engine'")
	}
	if args.Flavor == nil {
		return nil, errors.New("invalid value for required argument 'Flavor'")
	}
	if args.Nodes == nil {
		return nil, errors.New("invalid value for required argument 'Nodes'")
	}
	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Database
	err := ctx.RegisterResource("ovh:CloudProject/database:Database", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabase gets an existing Database resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseState, opts ...pulumi.ResourceOption) (*Database, error) {
	var resource Database
	err := ctx.ReadResource("ovh:CloudProject/database:Database", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Database resources.
type databaseState struct {
	// Advanced configuration key / value.
	AdvancedConfiguration map[string]string `pulumi:"advancedConfiguration"`
	// Time on which backups start every day.
	BackupTime *string `pulumi:"backupTime"`
	// Date of the creation of the cluster.
	CreatedAt *string `pulumi:"createdAt"`
	// Small description of the database service.
	Description *string `pulumi:"description"`
	// The disk size (in GB) of the database service.
	DiskSize *int `pulumi:"diskSize"`
	// Defines the disk type of the database service.
	DiskType *string `pulumi:"diskType"`
	// List of all endpoints objects of the service.
	Endpoints []DatabaseEndpoint `pulumi:"endpoints"`
	// The database engine you want to deploy. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine *string `pulumi:"engine"`
	// A valid OVHcloud public cloud database flavor name in which the nodes will be started.
	// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
	// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
	Flavor *string `pulumi:"flavor"`
	// Defines whether the REST API is enabled on a kafka cluster
	KafkaRestApi *bool `pulumi:"kafkaRestApi"`
	// Time on which maintenances can start every day.
	MaintenanceTime *string `pulumi:"maintenanceTime"`
	// Type of network of the cluster.
	NetworkType *string `pulumi:"networkType"`
	// List of nodes object.
	// Multi region cluster are not yet available, all node should be identical.
	Nodes []DatabaseNode `pulumi:"nodes"`
	// Defines whether the ACLs are enabled on an OpenSearch cluster
	OpensearchAclsEnabled *bool `pulumi:"opensearchAclsEnabled"`
	// Plan of the cluster.
	// Enum: "essential", "business", "enterprise".
	Plan *string `pulumi:"plan"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Current status of the cluster.
	Status *string `pulumi:"status"`
	// The version of the engine in which the service should be deployed
	Version *string `pulumi:"version"`
}

type DatabaseState struct {
	// Advanced configuration key / value.
	AdvancedConfiguration pulumi.StringMapInput
	// Time on which backups start every day.
	BackupTime pulumi.StringPtrInput
	// Date of the creation of the cluster.
	CreatedAt pulumi.StringPtrInput
	// Small description of the database service.
	Description pulumi.StringPtrInput
	// The disk size (in GB) of the database service.
	DiskSize pulumi.IntPtrInput
	// Defines the disk type of the database service.
	DiskType pulumi.StringPtrInput
	// List of all endpoints objects of the service.
	Endpoints DatabaseEndpointArrayInput
	// The database engine you want to deploy. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringPtrInput
	// A valid OVHcloud public cloud database flavor name in which the nodes will be started.
	// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
	// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
	Flavor pulumi.StringPtrInput
	// Defines whether the REST API is enabled on a kafka cluster
	KafkaRestApi pulumi.BoolPtrInput
	// Time on which maintenances can start every day.
	MaintenanceTime pulumi.StringPtrInput
	// Type of network of the cluster.
	NetworkType pulumi.StringPtrInput
	// List of nodes object.
	// Multi region cluster are not yet available, all node should be identical.
	Nodes DatabaseNodeArrayInput
	// Defines whether the ACLs are enabled on an OpenSearch cluster
	OpensearchAclsEnabled pulumi.BoolPtrInput
	// Plan of the cluster.
	// Enum: "essential", "business", "enterprise".
	Plan pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Current status of the cluster.
	Status pulumi.StringPtrInput
	// The version of the engine in which the service should be deployed
	Version pulumi.StringPtrInput
}

func (DatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseState)(nil)).Elem()
}

type databaseArgs struct {
	// Advanced configuration key / value.
	AdvancedConfiguration map[string]string `pulumi:"advancedConfiguration"`
	// Small description of the database service.
	Description *string `pulumi:"description"`
	// The disk size (in GB) of the database service.
	DiskSize *int `pulumi:"diskSize"`
	// The database engine you want to deploy. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine string `pulumi:"engine"`
	// A valid OVHcloud public cloud database flavor name in which the nodes will be started.
	// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
	// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
	Flavor string `pulumi:"flavor"`
	// Defines whether the REST API is enabled on a kafka cluster
	KafkaRestApi *bool `pulumi:"kafkaRestApi"`
	// List of nodes object.
	// Multi region cluster are not yet available, all node should be identical.
	Nodes []DatabaseNode `pulumi:"nodes"`
	// Defines whether the ACLs are enabled on an OpenSearch cluster
	OpensearchAclsEnabled *bool `pulumi:"opensearchAclsEnabled"`
	// Plan of the cluster.
	// Enum: "essential", "business", "enterprise".
	Plan string `pulumi:"plan"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// The version of the engine in which the service should be deployed
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a Database resource.
type DatabaseArgs struct {
	// Advanced configuration key / value.
	AdvancedConfiguration pulumi.StringMapInput
	// Small description of the database service.
	Description pulumi.StringPtrInput
	// The disk size (in GB) of the database service.
	DiskSize pulumi.IntPtrInput
	// The database engine you want to deploy. To get a full list of available engine visit.
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringInput
	// A valid OVHcloud public cloud database flavor name in which the nodes will be started.
	// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
	// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
	Flavor pulumi.StringInput
	// Defines whether the REST API is enabled on a kafka cluster
	KafkaRestApi pulumi.BoolPtrInput
	// List of nodes object.
	// Multi region cluster are not yet available, all node should be identical.
	Nodes DatabaseNodeArrayInput
	// Defines whether the ACLs are enabled on an OpenSearch cluster
	OpensearchAclsEnabled pulumi.BoolPtrInput
	// Plan of the cluster.
	// Enum: "essential", "business", "enterprise".
	Plan pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
	// The version of the engine in which the service should be deployed
	Version pulumi.StringInput
}

func (DatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseArgs)(nil)).Elem()
}

type DatabaseInput interface {
	pulumi.Input

	ToDatabaseOutput() DatabaseOutput
	ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput
}

func (*Database) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (i *Database) ToDatabaseOutput() DatabaseOutput {
	return i.ToDatabaseOutputWithContext(context.Background())
}

func (i *Database) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseOutput)
}

// DatabaseArrayInput is an input type that accepts DatabaseArray and DatabaseArrayOutput values.
// You can construct a concrete instance of `DatabaseArrayInput` via:
//
//	DatabaseArray{ DatabaseArgs{...} }
type DatabaseArrayInput interface {
	pulumi.Input

	ToDatabaseArrayOutput() DatabaseArrayOutput
	ToDatabaseArrayOutputWithContext(context.Context) DatabaseArrayOutput
}

type DatabaseArray []DatabaseInput

func (DatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (i DatabaseArray) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return i.ToDatabaseArrayOutputWithContext(context.Background())
}

func (i DatabaseArray) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseArrayOutput)
}

// DatabaseMapInput is an input type that accepts DatabaseMap and DatabaseMapOutput values.
// You can construct a concrete instance of `DatabaseMapInput` via:
//
//	DatabaseMap{ "key": DatabaseArgs{...} }
type DatabaseMapInput interface {
	pulumi.Input

	ToDatabaseMapOutput() DatabaseMapOutput
	ToDatabaseMapOutputWithContext(context.Context) DatabaseMapOutput
}

type DatabaseMap map[string]DatabaseInput

func (DatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (i DatabaseMap) ToDatabaseMapOutput() DatabaseMapOutput {
	return i.ToDatabaseMapOutputWithContext(context.Background())
}

func (i DatabaseMap) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseMapOutput)
}

type DatabaseOutput struct{ *pulumi.OutputState }

func (DatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Database)(nil)).Elem()
}

func (o DatabaseOutput) ToDatabaseOutput() DatabaseOutput {
	return o
}

func (o DatabaseOutput) ToDatabaseOutputWithContext(ctx context.Context) DatabaseOutput {
	return o
}

// Advanced configuration key / value.
func (o DatabaseOutput) AdvancedConfiguration() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Database) pulumi.StringMapOutput { return v.AdvancedConfiguration }).(pulumi.StringMapOutput)
}

// Time on which backups start every day.
func (o DatabaseOutput) BackupTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.BackupTime }).(pulumi.StringOutput)
}

// Date of the creation of the cluster.
func (o DatabaseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Small description of the database service.
func (o DatabaseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The disk size (in GB) of the database service.
func (o DatabaseOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v *Database) pulumi.IntOutput { return v.DiskSize }).(pulumi.IntOutput)
}

// Defines the disk type of the database service.
func (o DatabaseOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.DiskType }).(pulumi.StringOutput)
}

// List of all endpoints objects of the service.
func (o DatabaseOutput) Endpoints() DatabaseEndpointArrayOutput {
	return o.ApplyT(func(v *Database) DatabaseEndpointArrayOutput { return v.Endpoints }).(DatabaseEndpointArrayOutput)
}

// The database engine you want to deploy. To get a full list of available engine visit.
// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
func (o DatabaseOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Engine }).(pulumi.StringOutput)
}

// A valid OVHcloud public cloud database flavor name in which the nodes will be started.
// Ex: "db1-7". Changing this value upgrade the nodes with the new flavor.
// You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/
func (o DatabaseOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Flavor }).(pulumi.StringOutput)
}

// Defines whether the REST API is enabled on a kafka cluster
func (o DatabaseOutput) KafkaRestApi() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolPtrOutput { return v.KafkaRestApi }).(pulumi.BoolPtrOutput)
}

// Time on which maintenances can start every day.
func (o DatabaseOutput) MaintenanceTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.MaintenanceTime }).(pulumi.StringOutput)
}

// Type of network of the cluster.
func (o DatabaseOutput) NetworkType() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.NetworkType }).(pulumi.StringOutput)
}

// List of nodes object.
// Multi region cluster are not yet available, all node should be identical.
func (o DatabaseOutput) Nodes() DatabaseNodeArrayOutput {
	return o.ApplyT(func(v *Database) DatabaseNodeArrayOutput { return v.Nodes }).(DatabaseNodeArrayOutput)
}

// Defines whether the ACLs are enabled on an OpenSearch cluster
func (o DatabaseOutput) OpensearchAclsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Database) pulumi.BoolPtrOutput { return v.OpensearchAclsEnabled }).(pulumi.BoolPtrOutput)
}

// Plan of the cluster.
// Enum: "essential", "business", "enterprise".
func (o DatabaseOutput) Plan() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Plan }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o DatabaseOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the cluster.
func (o DatabaseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The version of the engine in which the service should be deployed
func (o DatabaseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Database) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type DatabaseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Database)(nil)).Elem()
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutput() DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) ToDatabaseArrayOutputWithContext(ctx context.Context) DatabaseArrayOutput {
	return o
}

func (o DatabaseArrayOutput) Index(i pulumi.IntInput) DatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Database {
		return vs[0].([]*Database)[vs[1].(int)]
	}).(DatabaseOutput)
}

type DatabaseMapOutput struct{ *pulumi.OutputState }

func (DatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Database)(nil)).Elem()
}

func (o DatabaseMapOutput) ToDatabaseMapOutput() DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) ToDatabaseMapOutputWithContext(ctx context.Context) DatabaseMapOutput {
	return o
}

func (o DatabaseMapOutput) MapIndex(k pulumi.StringInput) DatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Database {
		return vs[0].(map[string]*Database)[vs[1].(string)]
	}).(DatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseInput)(nil)).Elem(), &Database{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseArrayInput)(nil)).Elem(), DatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseMapInput)(nil)).Elem(), DatabaseMap{})
	pulumi.RegisterOutputType(DatabaseOutput{})
	pulumi.RegisterOutputType(DatabaseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseMapOutput{})
}
