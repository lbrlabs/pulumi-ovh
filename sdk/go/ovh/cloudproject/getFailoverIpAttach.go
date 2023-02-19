// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the details of a failover ip address of a service in a public cloud project.
//
// ## Example Usage
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
//			_, err := CloudProject.GetFailoverIpAttach(ctx, &cloudproject.GetFailoverIpAttachArgs{
//				Ip:          pulumi.StringRef("XXXXXX"),
//				ServiceName: "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFailoverIpAttach(ctx *pulumi.Context, args *LookupFailoverIpAttachArgs, opts ...pulumi.InvokeOption) (*LookupFailoverIpAttachResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupFailoverIpAttachResult
	err := ctx.Invoke("ovh:CloudProject/getFailoverIpAttach:getFailoverIpAttach", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFailoverIpAttach.
type LookupFailoverIpAttachArgs struct {
	// The IP block
	Block         *string `pulumi:"block"`
	ContinentCode *string `pulumi:"continentCode"`
	GeoLoc        *string `pulumi:"geoLoc"`
	// The failover ip address to query
	Ip       *string `pulumi:"ip"`
	RoutedTo *string `pulumi:"routedTo"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getFailoverIpAttach.
type LookupFailoverIpAttachResult struct {
	// The IP block
	Block         string `pulumi:"block"`
	ContinentCode string `pulumi:"continentCode"`
	GeoLoc        string `pulumi:"geoLoc"`
	// The Ip id
	Id string `pulumi:"id"`
	// The Ip Address
	Ip string `pulumi:"ip"`
	// Current operation progress in percent
	Progress    int    `pulumi:"progress"`
	RoutedTo    string `pulumi:"routedTo"`
	ServiceName string `pulumi:"serviceName"`
	// Ip status, can be `ok` or `operationPending`
	Status  string `pulumi:"status"`
	SubType string `pulumi:"subType"`
}

func LookupFailoverIpAttachOutput(ctx *pulumi.Context, args LookupFailoverIpAttachOutputArgs, opts ...pulumi.InvokeOption) LookupFailoverIpAttachResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFailoverIpAttachResult, error) {
			args := v.(LookupFailoverIpAttachArgs)
			r, err := LookupFailoverIpAttach(ctx, &args, opts...)
			var s LookupFailoverIpAttachResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFailoverIpAttachResultOutput)
}

// A collection of arguments for invoking getFailoverIpAttach.
type LookupFailoverIpAttachOutputArgs struct {
	// The IP block
	Block         pulumi.StringPtrInput `pulumi:"block"`
	ContinentCode pulumi.StringPtrInput `pulumi:"continentCode"`
	GeoLoc        pulumi.StringPtrInput `pulumi:"geoLoc"`
	// The failover ip address to query
	Ip       pulumi.StringPtrInput `pulumi:"ip"`
	RoutedTo pulumi.StringPtrInput `pulumi:"routedTo"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupFailoverIpAttachOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFailoverIpAttachArgs)(nil)).Elem()
}

// A collection of values returned by getFailoverIpAttach.
type LookupFailoverIpAttachResultOutput struct{ *pulumi.OutputState }

func (LookupFailoverIpAttachResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFailoverIpAttachResult)(nil)).Elem()
}

func (o LookupFailoverIpAttachResultOutput) ToLookupFailoverIpAttachResultOutput() LookupFailoverIpAttachResultOutput {
	return o
}

func (o LookupFailoverIpAttachResultOutput) ToLookupFailoverIpAttachResultOutputWithContext(ctx context.Context) LookupFailoverIpAttachResultOutput {
	return o
}

// The IP block
func (o LookupFailoverIpAttachResultOutput) Block() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.Block }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) ContinentCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.ContinentCode }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) GeoLoc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.GeoLoc }).(pulumi.StringOutput)
}

// The Ip id
func (o LookupFailoverIpAttachResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Ip Address
func (o LookupFailoverIpAttachResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.Ip }).(pulumi.StringOutput)
}

// Current operation progress in percent
func (o LookupFailoverIpAttachResultOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) int { return v.Progress }).(pulumi.IntOutput)
}

func (o LookupFailoverIpAttachResultOutput) RoutedTo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.RoutedTo }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Ip status, can be `ok` or `operationPending`
func (o LookupFailoverIpAttachResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupFailoverIpAttachResultOutput) SubType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFailoverIpAttachResult) string { return v.SubType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFailoverIpAttachResultOutput{})
}
