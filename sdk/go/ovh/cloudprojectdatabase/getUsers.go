// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of users of a database cluster associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/CloudProjectDatabase"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			users, err := CloudProjectDatabase.GetUsers(ctx, &cloudprojectdatabase.GetUsersArgs{
//				ServiceName: "XXXX",
//				Engine:      "YYYY",
//				ClusterId:   "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("userIds", users.UserIds)
//			return nil
//		})
//	}
//
// ```
func GetUsers(ctx *pulumi.Context, args *GetUsersArgs, opts ...pulumi.InvokeOption) (*GetUsersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetUsersResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getUsers:getUsers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUsers.
type GetUsersArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// The engine of the database cluster you want to list users. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine string `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getUsers.
type GetUsersResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// The list of users ids of the database cluster associated with the project.
	UserIds []string `pulumi:"userIds"`
}

func GetUsersOutput(ctx *pulumi.Context, args GetUsersOutputArgs, opts ...pulumi.InvokeOption) GetUsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUsersResult, error) {
			args := v.(GetUsersArgs)
			r, err := GetUsers(ctx, &args, opts...)
			var s GetUsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetUsersResultOutput)
}

// A collection of arguments for invoking getUsers.
type GetUsersOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The engine of the database cluster you want to list users. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringInput `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetUsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersArgs)(nil)).Elem()
}

// A collection of values returned by getUsers.
type GetUsersResultOutput struct{ *pulumi.OutputState }

func (GetUsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUsersResult)(nil)).Elem()
}

func (o GetUsersResultOutput) ToGetUsersResultOutput() GetUsersResultOutput {
	return o
}

func (o GetUsersResultOutput) ToGetUsersResultOutputWithContext(ctx context.Context) GetUsersResultOutput {
	return o
}

// See Argument Reference above.
func (o GetUsersResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetUsersResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetUsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetUsersResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetUsersResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// The list of users ids of the database cluster associated with the project.
func (o GetUsersResultOutput) UserIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetUsersResult) []string { return v.UserIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUsersResultOutput{})
}
