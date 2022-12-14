// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a user of a postgresql cluster associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			pguser, err := ovh.LookupCloudProjectDatabasePostgresSqlUser(ctx, &GetCloudProjectDatabasePostgresSqlUserArgs{
//				ServiceName: "XXX",
//				ClusterId:   "YYY",
//				Name:        "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("pguserRoles", pguser.Roles)
//			return nil
//		})
//	}
//
// ```
func LookupCloudProjectDatabasePostgresSqlUser(ctx *pulumi.Context, args *LookupCloudProjectDatabasePostgresSqlUserArgs, opts ...pulumi.InvokeOption) (*LookupCloudProjectDatabasePostgresSqlUserResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupCloudProjectDatabasePostgresSqlUserResult
	err := ctx.Invoke("ovh:index/getCloudProjectDatabasePostgresSqlUser:getCloudProjectDatabasePostgresSqlUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudProjectDatabasePostgresSqlUser.
type LookupCloudProjectDatabasePostgresSqlUserArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// Name of the user.
	Name string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getCloudProjectDatabasePostgresSqlUser.
type LookupCloudProjectDatabasePostgresSqlUserResult struct {
	ClusterId string `pulumi:"clusterId"`
	// Date of the creation of the user.
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Name of the user.
	Name string `pulumi:"name"`
	// Roles the user belongs to.
	Roles       []string `pulumi:"roles"`
	ServiceName string   `pulumi:"serviceName"`
	// Current status of the user.
	Status string `pulumi:"status"`
}

func LookupCloudProjectDatabasePostgresSqlUserOutput(ctx *pulumi.Context, args LookupCloudProjectDatabasePostgresSqlUserOutputArgs, opts ...pulumi.InvokeOption) LookupCloudProjectDatabasePostgresSqlUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCloudProjectDatabasePostgresSqlUserResult, error) {
			args := v.(LookupCloudProjectDatabasePostgresSqlUserArgs)
			r, err := LookupCloudProjectDatabasePostgresSqlUser(ctx, &args, opts...)
			var s LookupCloudProjectDatabasePostgresSqlUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCloudProjectDatabasePostgresSqlUserResultOutput)
}

// A collection of arguments for invoking getCloudProjectDatabasePostgresSqlUser.
type LookupCloudProjectDatabasePostgresSqlUserOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Name of the user.
	Name pulumi.StringInput `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupCloudProjectDatabasePostgresSqlUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectDatabasePostgresSqlUserArgs)(nil)).Elem()
}

// A collection of values returned by getCloudProjectDatabasePostgresSqlUser.
type LookupCloudProjectDatabasePostgresSqlUserResultOutput struct{ *pulumi.OutputState }

func (LookupCloudProjectDatabasePostgresSqlUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudProjectDatabasePostgresSqlUserResult)(nil)).Elem()
}

func (o LookupCloudProjectDatabasePostgresSqlUserResultOutput) ToLookupCloudProjectDatabasePostgresSqlUserResultOutput() LookupCloudProjectDatabasePostgresSqlUserResultOutput {
	return o
}

func (o LookupCloudProjectDatabasePostgresSqlUserResultOutput) ToLookupCloudProjectDatabasePostgresSqlUserResultOutputWithContext(ctx context.Context) LookupCloudProjectDatabasePostgresSqlUserResultOutput {
	return o
}

func (o LookupCloudProjectDatabasePostgresSqlUserResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabasePostgresSqlUserResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Date of the creation of the user.
func (o LookupCloudProjectDatabasePostgresSqlUserResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabasePostgresSqlUserResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCloudProjectDatabasePostgresSqlUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabasePostgresSqlUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the user.
func (o LookupCloudProjectDatabasePostgresSqlUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabasePostgresSqlUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// Roles the user belongs to.
func (o LookupCloudProjectDatabasePostgresSqlUserResultOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabasePostgresSqlUserResult) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

func (o LookupCloudProjectDatabasePostgresSqlUserResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabasePostgresSqlUserResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the user.
func (o LookupCloudProjectDatabasePostgresSqlUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudProjectDatabasePostgresSqlUserResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudProjectDatabasePostgresSqlUserResultOutput{})
}
