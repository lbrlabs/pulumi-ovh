// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use the list of IP restrictions associated with a public cloud project.
//
// ## Example Usage
//
// To get the list of IP restriction on a database cluster service:
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
//			iprestrictions, err := CloudProjectDatabase.GetIpRestrictions(ctx, &cloudprojectdatabase.GetIpRestrictionsArgs{
//				ServiceName: "XXXXXX",
//				Engine:      "YYYY",
//				ClusterId:   "ZZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("ips", iprestrictions.Ips)
//			return nil
//		})
//	}
//
// ```
func GetIpRestrictions(ctx *pulumi.Context, args *GetIpRestrictionsArgs, opts ...pulumi.InvokeOption) (*GetIpRestrictionsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIpRestrictionsResult
	err := ctx.Invoke("ovh:CloudProjectDatabase/getIpRestrictions:getIpRestrictions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpRestrictions.
type GetIpRestrictionsArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// The engine of the database cluster you want to list IP restrictions. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine string `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getIpRestrictions.
type GetIpRestrictionsResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// See Argument Reference above.
	Engine string `pulumi:"engine"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of IP restriction of the database associated with the project.
	Ips []string `pulumi:"ips"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
}

func GetIpRestrictionsOutput(ctx *pulumi.Context, args GetIpRestrictionsOutputArgs, opts ...pulumi.InvokeOption) GetIpRestrictionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIpRestrictionsResult, error) {
			args := v.(GetIpRestrictionsArgs)
			r, err := GetIpRestrictions(ctx, &args, opts...)
			var s GetIpRestrictionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIpRestrictionsResultOutput)
}

// A collection of arguments for invoking getIpRestrictions.
type GetIpRestrictionsOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// The engine of the database cluster you want to list IP restrictions. To get a full list of available engine visit:
	// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
	Engine pulumi.StringInput `pulumi:"engine"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetIpRestrictionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpRestrictionsArgs)(nil)).Elem()
}

// A collection of values returned by getIpRestrictions.
type GetIpRestrictionsResultOutput struct{ *pulumi.OutputState }

func (GetIpRestrictionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpRestrictionsResult)(nil)).Elem()
}

func (o GetIpRestrictionsResultOutput) ToGetIpRestrictionsResultOutput() GetIpRestrictionsResultOutput {
	return o
}

func (o GetIpRestrictionsResultOutput) ToGetIpRestrictionsResultOutputWithContext(ctx context.Context) GetIpRestrictionsResultOutput {
	return o
}

// See Argument Reference above.
func (o GetIpRestrictionsResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpRestrictionsResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetIpRestrictionsResultOutput) Engine() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpRestrictionsResult) string { return v.Engine }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpRestrictionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpRestrictionsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of IP restriction of the database associated with the project.
func (o GetIpRestrictionsResultOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetIpRestrictionsResult) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o GetIpRestrictionsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpRestrictionsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpRestrictionsResultOutput{})
}