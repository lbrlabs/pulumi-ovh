// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a OVHcloud Managed Kubernetes node pool.
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
//			nodepool, err := CloudProject.GetKubeIpNodePool(ctx, &cloudproject.GetKubeIpNodePoolArgs{
//				ServiceName: "XXXXXX",
//				KubeId:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
//				Name:        "xxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("maxNodes", nodepool.MaxNodes)
//			return nil
//		})
//	}
//
// ```
func GetKubeIpNodePool(ctx *pulumi.Context, args *GetKubeIpNodePoolArgs, opts ...pulumi.InvokeOption) (*GetKubeIpNodePoolResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetKubeIpNodePoolResult
	err := ctx.Invoke("ovh:CloudProject/getKubeIpNodePool:getKubeIpNodePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubeIpNodePool.
type GetKubeIpNodePoolArgs struct {
	// The id of the managed kubernetes cluster.
	KubeId string `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes *int `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes *int `pulumi:"minNodes"`
	// The name of the node pool.
	Name string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getKubeIpNodePool.
type GetKubeIpNodePoolResult struct {
	// (Optional) should the pool use the anti-affinity feature. Default to `false`.
	AntiAffinity bool `pulumi:"antiAffinity"`
	// (Optional) Enable auto-scaling for the pool. Default to `false`.
	Autoscale bool `pulumi:"autoscale"`
	// Number of nodes which are actually ready in the pool
	AvailableNodes int `pulumi:"availableNodes"`
	// Creation date
	CreatedAt string `pulumi:"createdAt"`
	// Number of nodes present in the pool
	CurrentNodes int `pulumi:"currentNodes"`
	// Number of nodes you desire in the pool
	DesiredNodes int `pulumi:"desiredNodes"`
	// Flavor name
	Flavor string `pulumi:"flavor"`
	// a valid OVHcloud public cloud flavor ID in which the nodes will be started.
	// Ex: "b2-7". Changing this value recreates the resource.
	// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
	FlavorName string `pulumi:"flavorName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	KubeId string `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes int `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes int `pulumi:"minNodes"`
	// (Optional) should the nodes be billed on a monthly basis. Default to `false`.
	MonthlyBilled bool `pulumi:"monthlyBilled"`
	// (Optional) The name of the nodepool.
	// Changing this value recreates the resource.
	// Warning: "_" char is not allowed!
	Name string `pulumi:"name"`
	// Project id
	ProjectId string `pulumi:"projectId"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// Status describing the state between number of nodes wanted and available ones
	SizeStatus string `pulumi:"sizeStatus"`
	// Current status
	Status string `pulumi:"status"`
	// Number of nodes with the latest version installed in the pool
	UpToDateNodes int `pulumi:"upToDateNodes"`
	// Last update date
	UpdatedAt string `pulumi:"updatedAt"`
}

func GetKubeIpNodePoolOutput(ctx *pulumi.Context, args GetKubeIpNodePoolOutputArgs, opts ...pulumi.InvokeOption) GetKubeIpNodePoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetKubeIpNodePoolResult, error) {
			args := v.(GetKubeIpNodePoolArgs)
			r, err := GetKubeIpNodePool(ctx, &args, opts...)
			var s GetKubeIpNodePoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetKubeIpNodePoolResultOutput)
}

// A collection of arguments for invoking getKubeIpNodePool.
type GetKubeIpNodePoolOutputArgs struct {
	// The id of the managed kubernetes cluster.
	KubeId pulumi.StringInput `pulumi:"kubeId"`
	// maximum number of nodes allowed in the pool.
	// Setting `desiredNodes` over this value will raise an error.
	MaxNodes pulumi.IntPtrInput `pulumi:"maxNodes"`
	// minimum number of nodes allowed in the pool.
	// Setting `desiredNodes` under this value will raise an error.
	MinNodes pulumi.IntPtrInput `pulumi:"minNodes"`
	// The name of the node pool.
	Name pulumi.StringInput `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetKubeIpNodePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubeIpNodePoolArgs)(nil)).Elem()
}

// A collection of values returned by getKubeIpNodePool.
type GetKubeIpNodePoolResultOutput struct{ *pulumi.OutputState }

func (GetKubeIpNodePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetKubeIpNodePoolResult)(nil)).Elem()
}

func (o GetKubeIpNodePoolResultOutput) ToGetKubeIpNodePoolResultOutput() GetKubeIpNodePoolResultOutput {
	return o
}

func (o GetKubeIpNodePoolResultOutput) ToGetKubeIpNodePoolResultOutputWithContext(ctx context.Context) GetKubeIpNodePoolResultOutput {
	return o
}

// (Optional) should the pool use the anti-affinity feature. Default to `false`.
func (o GetKubeIpNodePoolResultOutput) AntiAffinity() pulumi.BoolOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) bool { return v.AntiAffinity }).(pulumi.BoolOutput)
}

// (Optional) Enable auto-scaling for the pool. Default to `false`.
func (o GetKubeIpNodePoolResultOutput) Autoscale() pulumi.BoolOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) bool { return v.Autoscale }).(pulumi.BoolOutput)
}

// Number of nodes which are actually ready in the pool
func (o GetKubeIpNodePoolResultOutput) AvailableNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) int { return v.AvailableNodes }).(pulumi.IntOutput)
}

// Creation date
func (o GetKubeIpNodePoolResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// Number of nodes present in the pool
func (o GetKubeIpNodePoolResultOutput) CurrentNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) int { return v.CurrentNodes }).(pulumi.IntOutput)
}

// Number of nodes you desire in the pool
func (o GetKubeIpNodePoolResultOutput) DesiredNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) int { return v.DesiredNodes }).(pulumi.IntOutput)
}

// Flavor name
func (o GetKubeIpNodePoolResultOutput) Flavor() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.Flavor }).(pulumi.StringOutput)
}

// a valid OVHcloud public cloud flavor ID in which the nodes will be started.
// Ex: "b2-7". Changing this value recreates the resource.
// You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
func (o GetKubeIpNodePoolResultOutput) FlavorName() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.FlavorName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetKubeIpNodePoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetKubeIpNodePoolResultOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.KubeId }).(pulumi.StringOutput)
}

// maximum number of nodes allowed in the pool.
// Setting `desiredNodes` over this value will raise an error.
func (o GetKubeIpNodePoolResultOutput) MaxNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) int { return v.MaxNodes }).(pulumi.IntOutput)
}

// minimum number of nodes allowed in the pool.
// Setting `desiredNodes` under this value will raise an error.
func (o GetKubeIpNodePoolResultOutput) MinNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) int { return v.MinNodes }).(pulumi.IntOutput)
}

// (Optional) should the nodes be billed on a monthly basis. Default to `false`.
func (o GetKubeIpNodePoolResultOutput) MonthlyBilled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) bool { return v.MonthlyBilled }).(pulumi.BoolOutput)
}

// (Optional) The name of the nodepool.
// Changing this value recreates the resource.
// Warning: "_" char is not allowed!
func (o GetKubeIpNodePoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.Name }).(pulumi.StringOutput)
}

// Project id
func (o GetKubeIpNodePoolResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetKubeIpNodePoolResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Status describing the state between number of nodes wanted and available ones
func (o GetKubeIpNodePoolResultOutput) SizeStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.SizeStatus }).(pulumi.StringOutput)
}

// Current status
func (o GetKubeIpNodePoolResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.Status }).(pulumi.StringOutput)
}

// Number of nodes with the latest version installed in the pool
func (o GetKubeIpNodePoolResultOutput) UpToDateNodes() pulumi.IntOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) int { return v.UpToDateNodes }).(pulumi.IntOutput)
}

// Last update date
func (o GetKubeIpNodePoolResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetKubeIpNodePoolResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetKubeIpNodePoolResultOutput{})
}
