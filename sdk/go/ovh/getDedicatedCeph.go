// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a dedicated CEPH.
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
//			_, err := ovh.GetDedicatedCeph(ctx, &GetDedicatedCephArgs{
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
func GetDedicatedCeph(ctx *pulumi.Context, args *GetDedicatedCephArgs, opts ...pulumi.InvokeOption) (*GetDedicatedCephResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetDedicatedCephResult
	err := ctx.Invoke("ovh:index/getDedicatedCeph:getDedicatedCeph", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDedicatedCeph.
type GetDedicatedCephArgs struct {
	// CEPH cluster version
	CephVersion *string `pulumi:"cephVersion"`
	// The service name of the dedicated CEPH cluster.
	ServiceName string `pulumi:"serviceName"`
	// the status of the service
	Status *string `pulumi:"status"`
}

// A collection of values returned by getDedicatedCeph.
type GetDedicatedCephResult struct {
	// list of CEPH monitors IPs
	CephMons []string `pulumi:"cephMons"`
	// CEPH cluster version
	CephVersion string `pulumi:"cephVersion"`
	// CRUSH algorithm settings. Possible values
	// * OPTIMAL
	// * DEFAULT
	// * LEGACY
	// * BOBTAIL
	// * ARGONAUT
	// * FIREFLY
	// * HAMMER
	// * JEWEL
	CrushTunables string `pulumi:"crushTunables"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// CEPH cluster label
	Label string `pulumi:"label"`
	// cluster region
	Region      string `pulumi:"region"`
	ServiceName string `pulumi:"serviceName"`
	// Cluster size in TB
	Size float64 `pulumi:"size"`
	// the state of the cluster
	State string `pulumi:"state"`
	// the status of the service
	Status string `pulumi:"status"`
}

func GetDedicatedCephOutput(ctx *pulumi.Context, args GetDedicatedCephOutputArgs, opts ...pulumi.InvokeOption) GetDedicatedCephResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDedicatedCephResult, error) {
			args := v.(GetDedicatedCephArgs)
			r, err := GetDedicatedCeph(ctx, &args, opts...)
			var s GetDedicatedCephResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDedicatedCephResultOutput)
}

// A collection of arguments for invoking getDedicatedCeph.
type GetDedicatedCephOutputArgs struct {
	// CEPH cluster version
	CephVersion pulumi.StringPtrInput `pulumi:"cephVersion"`
	// The service name of the dedicated CEPH cluster.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// the status of the service
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (GetDedicatedCephOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDedicatedCephArgs)(nil)).Elem()
}

// A collection of values returned by getDedicatedCeph.
type GetDedicatedCephResultOutput struct{ *pulumi.OutputState }

func (GetDedicatedCephResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDedicatedCephResult)(nil)).Elem()
}

func (o GetDedicatedCephResultOutput) ToGetDedicatedCephResultOutput() GetDedicatedCephResultOutput {
	return o
}

func (o GetDedicatedCephResultOutput) ToGetDedicatedCephResultOutputWithContext(ctx context.Context) GetDedicatedCephResultOutput {
	return o
}

// list of CEPH monitors IPs
func (o GetDedicatedCephResultOutput) CephMons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDedicatedCephResult) []string { return v.CephMons }).(pulumi.StringArrayOutput)
}

// CEPH cluster version
func (o GetDedicatedCephResultOutput) CephVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetDedicatedCephResult) string { return v.CephVersion }).(pulumi.StringOutput)
}

// CRUSH algorithm settings. Possible values
// * OPTIMAL
// * DEFAULT
// * LEGACY
// * BOBTAIL
// * ARGONAUT
// * FIREFLY
// * HAMMER
// * JEWEL
func (o GetDedicatedCephResultOutput) CrushTunables() pulumi.StringOutput {
	return o.ApplyT(func(v GetDedicatedCephResult) string { return v.CrushTunables }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDedicatedCephResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDedicatedCephResult) string { return v.Id }).(pulumi.StringOutput)
}

// CEPH cluster label
func (o GetDedicatedCephResultOutput) Label() pulumi.StringOutput {
	return o.ApplyT(func(v GetDedicatedCephResult) string { return v.Label }).(pulumi.StringOutput)
}

// cluster region
func (o GetDedicatedCephResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetDedicatedCephResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetDedicatedCephResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetDedicatedCephResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Cluster size in TB
func (o GetDedicatedCephResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v GetDedicatedCephResult) float64 { return v.Size }).(pulumi.Float64Output)
}

// the state of the cluster
func (o GetDedicatedCephResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v GetDedicatedCephResult) string { return v.State }).(pulumi.StringOutput)
}

// the status of the service
func (o GetDedicatedCephResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetDedicatedCephResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDedicatedCephResultOutput{})
}
