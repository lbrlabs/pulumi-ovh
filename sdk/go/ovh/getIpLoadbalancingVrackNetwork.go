// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the details of Vrack network available for your IPLoadbalancer associated with your OVH account.
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
//			_, err := ovh.GetIpLoadbalancingVrackNetwork(ctx, &GetIpLoadbalancingVrackNetworkArgs{
//				ServiceName:    "XXXXXX",
//				VrackNetworkId: "yyy",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetIpLoadbalancingVrackNetwork(ctx *pulumi.Context, args *GetIpLoadbalancingVrackNetworkArgs, opts ...pulumi.InvokeOption) (*GetIpLoadbalancingVrackNetworkResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIpLoadbalancingVrackNetworkResult
	err := ctx.Invoke("ovh:index/getIpLoadbalancingVrackNetwork:getIpLoadbalancingVrackNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpLoadbalancingVrackNetwork.
type GetIpLoadbalancingVrackNetworkArgs struct {
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Internal Load Balancer identifier of the vRack private network
	VrackNetworkId int `pulumi:"vrackNetworkId"`
}

// A collection of values returned by getIpLoadbalancingVrackNetwork.
type GetIpLoadbalancingVrackNetworkResult struct {
	// Human readable name for your vrack network
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp       string `pulumi:"natIp"`
	ServiceName string `pulumi:"serviceName"`
	// IP block of the private network in the vRack
	Subnet string `pulumi:"subnet"`
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan           int `pulumi:"vlan"`
	VrackNetworkId int `pulumi:"vrackNetworkId"`
}

func GetIpLoadbalancingVrackNetworkOutput(ctx *pulumi.Context, args GetIpLoadbalancingVrackNetworkOutputArgs, opts ...pulumi.InvokeOption) GetIpLoadbalancingVrackNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetIpLoadbalancingVrackNetworkResult, error) {
			args := v.(GetIpLoadbalancingVrackNetworkArgs)
			r, err := GetIpLoadbalancingVrackNetwork(ctx, &args, opts...)
			var s GetIpLoadbalancingVrackNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetIpLoadbalancingVrackNetworkResultOutput)
}

// A collection of arguments for invoking getIpLoadbalancingVrackNetwork.
type GetIpLoadbalancingVrackNetworkOutputArgs struct {
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Internal Load Balancer identifier of the vRack private network
	VrackNetworkId pulumi.IntInput `pulumi:"vrackNetworkId"`
}

func (GetIpLoadbalancingVrackNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpLoadbalancingVrackNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getIpLoadbalancingVrackNetwork.
type GetIpLoadbalancingVrackNetworkResultOutput struct{ *pulumi.OutputState }

func (GetIpLoadbalancingVrackNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIpLoadbalancingVrackNetworkResult)(nil)).Elem()
}

func (o GetIpLoadbalancingVrackNetworkResultOutput) ToGetIpLoadbalancingVrackNetworkResultOutput() GetIpLoadbalancingVrackNetworkResultOutput {
	return o
}

func (o GetIpLoadbalancingVrackNetworkResultOutput) ToGetIpLoadbalancingVrackNetworkResultOutputWithContext(ctx context.Context) GetIpLoadbalancingVrackNetworkResultOutput {
	return o
}

// Human readable name for your vrack network
func (o GetIpLoadbalancingVrackNetworkResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworkResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetIpLoadbalancingVrackNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
func (o GetIpLoadbalancingVrackNetworkResultOutput) NatIp() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworkResult) string { return v.NatIp }).(pulumi.StringOutput)
}

func (o GetIpLoadbalancingVrackNetworkResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworkResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// IP block of the private network in the vRack
func (o GetIpLoadbalancingVrackNetworkResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworkResult) string { return v.Subnet }).(pulumi.StringOutput)
}

// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
func (o GetIpLoadbalancingVrackNetworkResultOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworkResult) int { return v.Vlan }).(pulumi.IntOutput)
}

func (o GetIpLoadbalancingVrackNetworkResultOutput) VrackNetworkId() pulumi.IntOutput {
	return o.ApplyT(func(v GetIpLoadbalancingVrackNetworkResult) int { return v.VrackNetworkId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIpLoadbalancingVrackNetworkResultOutput{})
}
