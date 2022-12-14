// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a backend server group (farm) to be used by loadbalancing frontend(s)
type IpLoadBalancingTcpFarm struct {
	pulumi.CustomResourceState

	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
	Balance pulumi.StringPtrOutput `pulumi:"balance"`
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Port for backends to recieve traffic on.
	Port pulumi.IntPtrOutput `pulumi:"port"`
	// define a backend healthcheck probe
	Probe IpLoadBalancingTcpFarmProbePtrOutput `pulumi:"probe"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Stickiness type. No stickiness if null (`sourceIp`)
	Stickiness pulumi.StringPtrOutput `pulumi:"stickiness"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.IntPtrOutput `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewIpLoadBalancingTcpFarm registers a new resource with the given unique name, arguments, and options.
func NewIpLoadBalancingTcpFarm(ctx *pulumi.Context,
	name string, args *IpLoadBalancingTcpFarmArgs, opts ...pulumi.ResourceOption) (*IpLoadBalancingTcpFarm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IpLoadBalancingTcpFarm
	err := ctx.RegisterResource("ovh:index/ipLoadBalancingTcpFarm:IpLoadBalancingTcpFarm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpLoadBalancingTcpFarm gets an existing IpLoadBalancingTcpFarm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpLoadBalancingTcpFarm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpLoadBalancingTcpFarmState, opts ...pulumi.ResourceOption) (*IpLoadBalancingTcpFarm, error) {
	var resource IpLoadBalancingTcpFarm
	err := ctx.ReadResource("ovh:index/ipLoadBalancingTcpFarm:IpLoadBalancingTcpFarm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpLoadBalancingTcpFarm resources.
type ipLoadBalancingTcpFarmState struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
	Balance *string `pulumi:"balance"`
	// Readable label for loadbalancer farm
	DisplayName *string `pulumi:"displayName"`
	// Port for backends to recieve traffic on.
	Port *int `pulumi:"port"`
	// define a backend healthcheck probe
	Probe *IpLoadBalancingTcpFarmProbe `pulumi:"probe"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Stickiness type. No stickiness if null (`sourceIp`)
	Stickiness *string `pulumi:"stickiness"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId *int `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone *string `pulumi:"zone"`
}

type IpLoadBalancingTcpFarmState struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
	Balance pulumi.StringPtrInput
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrInput
	// Port for backends to recieve traffic on.
	Port pulumi.IntPtrInput
	// define a backend healthcheck probe
	Probe IpLoadBalancingTcpFarmProbePtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Stickiness type. No stickiness if null (`sourceIp`)
	Stickiness pulumi.StringPtrInput
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.IntPtrInput
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone pulumi.StringPtrInput
}

func (IpLoadBalancingTcpFarmState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingTcpFarmState)(nil)).Elem()
}

type ipLoadBalancingTcpFarmArgs struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
	Balance *string `pulumi:"balance"`
	// Readable label for loadbalancer farm
	DisplayName *string `pulumi:"displayName"`
	// Port for backends to recieve traffic on.
	Port *int `pulumi:"port"`
	// define a backend healthcheck probe
	Probe *IpLoadBalancingTcpFarmProbe `pulumi:"probe"`
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Stickiness type. No stickiness if null (`sourceIp`)
	Stickiness *string `pulumi:"stickiness"`
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId *int `pulumi:"vrackNetworkId"`
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a IpLoadBalancingTcpFarm resource.
type IpLoadBalancingTcpFarmArgs struct {
	// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
	Balance pulumi.StringPtrInput
	// Readable label for loadbalancer farm
	DisplayName pulumi.StringPtrInput
	// Port for backends to recieve traffic on.
	Port pulumi.IntPtrInput
	// define a backend healthcheck probe
	Probe IpLoadBalancingTcpFarmProbePtrInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput
	// Stickiness type. No stickiness if null (`sourceIp`)
	Stickiness pulumi.StringPtrInput
	// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
	VrackNetworkId pulumi.IntPtrInput
	// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
	Zone pulumi.StringInput
}

func (IpLoadBalancingTcpFarmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingTcpFarmArgs)(nil)).Elem()
}

type IpLoadBalancingTcpFarmInput interface {
	pulumi.Input

	ToIpLoadBalancingTcpFarmOutput() IpLoadBalancingTcpFarmOutput
	ToIpLoadBalancingTcpFarmOutputWithContext(ctx context.Context) IpLoadBalancingTcpFarmOutput
}

func (*IpLoadBalancingTcpFarm) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancingTcpFarm)(nil)).Elem()
}

func (i *IpLoadBalancingTcpFarm) ToIpLoadBalancingTcpFarmOutput() IpLoadBalancingTcpFarmOutput {
	return i.ToIpLoadBalancingTcpFarmOutputWithContext(context.Background())
}

func (i *IpLoadBalancingTcpFarm) ToIpLoadBalancingTcpFarmOutputWithContext(ctx context.Context) IpLoadBalancingTcpFarmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingTcpFarmOutput)
}

// IpLoadBalancingTcpFarmArrayInput is an input type that accepts IpLoadBalancingTcpFarmArray and IpLoadBalancingTcpFarmArrayOutput values.
// You can construct a concrete instance of `IpLoadBalancingTcpFarmArrayInput` via:
//
//	IpLoadBalancingTcpFarmArray{ IpLoadBalancingTcpFarmArgs{...} }
type IpLoadBalancingTcpFarmArrayInput interface {
	pulumi.Input

	ToIpLoadBalancingTcpFarmArrayOutput() IpLoadBalancingTcpFarmArrayOutput
	ToIpLoadBalancingTcpFarmArrayOutputWithContext(context.Context) IpLoadBalancingTcpFarmArrayOutput
}

type IpLoadBalancingTcpFarmArray []IpLoadBalancingTcpFarmInput

func (IpLoadBalancingTcpFarmArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancingTcpFarm)(nil)).Elem()
}

func (i IpLoadBalancingTcpFarmArray) ToIpLoadBalancingTcpFarmArrayOutput() IpLoadBalancingTcpFarmArrayOutput {
	return i.ToIpLoadBalancingTcpFarmArrayOutputWithContext(context.Background())
}

func (i IpLoadBalancingTcpFarmArray) ToIpLoadBalancingTcpFarmArrayOutputWithContext(ctx context.Context) IpLoadBalancingTcpFarmArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingTcpFarmArrayOutput)
}

// IpLoadBalancingTcpFarmMapInput is an input type that accepts IpLoadBalancingTcpFarmMap and IpLoadBalancingTcpFarmMapOutput values.
// You can construct a concrete instance of `IpLoadBalancingTcpFarmMapInput` via:
//
//	IpLoadBalancingTcpFarmMap{ "key": IpLoadBalancingTcpFarmArgs{...} }
type IpLoadBalancingTcpFarmMapInput interface {
	pulumi.Input

	ToIpLoadBalancingTcpFarmMapOutput() IpLoadBalancingTcpFarmMapOutput
	ToIpLoadBalancingTcpFarmMapOutputWithContext(context.Context) IpLoadBalancingTcpFarmMapOutput
}

type IpLoadBalancingTcpFarmMap map[string]IpLoadBalancingTcpFarmInput

func (IpLoadBalancingTcpFarmMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancingTcpFarm)(nil)).Elem()
}

func (i IpLoadBalancingTcpFarmMap) ToIpLoadBalancingTcpFarmMapOutput() IpLoadBalancingTcpFarmMapOutput {
	return i.ToIpLoadBalancingTcpFarmMapOutputWithContext(context.Background())
}

func (i IpLoadBalancingTcpFarmMap) ToIpLoadBalancingTcpFarmMapOutputWithContext(ctx context.Context) IpLoadBalancingTcpFarmMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingTcpFarmMapOutput)
}

type IpLoadBalancingTcpFarmOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingTcpFarmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancingTcpFarm)(nil)).Elem()
}

func (o IpLoadBalancingTcpFarmOutput) ToIpLoadBalancingTcpFarmOutput() IpLoadBalancingTcpFarmOutput {
	return o
}

func (o IpLoadBalancingTcpFarmOutput) ToIpLoadBalancingTcpFarmOutputWithContext(ctx context.Context) IpLoadBalancingTcpFarmOutput {
	return o
}

// Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
func (o IpLoadBalancingTcpFarmOutput) Balance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFarm) pulumi.StringPtrOutput { return v.Balance }).(pulumi.StringPtrOutput)
}

// Readable label for loadbalancer farm
func (o IpLoadBalancingTcpFarmOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFarm) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Port for backends to recieve traffic on.
func (o IpLoadBalancingTcpFarmOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFarm) pulumi.IntPtrOutput { return v.Port }).(pulumi.IntPtrOutput)
}

// define a backend healthcheck probe
func (o IpLoadBalancingTcpFarmOutput) Probe() IpLoadBalancingTcpFarmProbePtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFarm) IpLoadBalancingTcpFarmProbePtrOutput { return v.Probe }).(IpLoadBalancingTcpFarmProbePtrOutput)
}

// The internal name of your IP load balancing
func (o IpLoadBalancingTcpFarmOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFarm) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Stickiness type. No stickiness if null (`sourceIp`)
func (o IpLoadBalancingTcpFarmOutput) Stickiness() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFarm) pulumi.StringPtrOutput { return v.Stickiness }).(pulumi.StringPtrOutput)
}

// Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
func (o IpLoadBalancingTcpFarmOutput) VrackNetworkId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFarm) pulumi.IntPtrOutput { return v.VrackNetworkId }).(pulumi.IntPtrOutput)
}

// Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
func (o IpLoadBalancingTcpFarmOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancingTcpFarm) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type IpLoadBalancingTcpFarmArrayOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingTcpFarmArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancingTcpFarm)(nil)).Elem()
}

func (o IpLoadBalancingTcpFarmArrayOutput) ToIpLoadBalancingTcpFarmArrayOutput() IpLoadBalancingTcpFarmArrayOutput {
	return o
}

func (o IpLoadBalancingTcpFarmArrayOutput) ToIpLoadBalancingTcpFarmArrayOutputWithContext(ctx context.Context) IpLoadBalancingTcpFarmArrayOutput {
	return o
}

func (o IpLoadBalancingTcpFarmArrayOutput) Index(i pulumi.IntInput) IpLoadBalancingTcpFarmOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpLoadBalancingTcpFarm {
		return vs[0].([]*IpLoadBalancingTcpFarm)[vs[1].(int)]
	}).(IpLoadBalancingTcpFarmOutput)
}

type IpLoadBalancingTcpFarmMapOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingTcpFarmMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancingTcpFarm)(nil)).Elem()
}

func (o IpLoadBalancingTcpFarmMapOutput) ToIpLoadBalancingTcpFarmMapOutput() IpLoadBalancingTcpFarmMapOutput {
	return o
}

func (o IpLoadBalancingTcpFarmMapOutput) ToIpLoadBalancingTcpFarmMapOutputWithContext(ctx context.Context) IpLoadBalancingTcpFarmMapOutput {
	return o
}

func (o IpLoadBalancingTcpFarmMapOutput) MapIndex(k pulumi.StringInput) IpLoadBalancingTcpFarmOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpLoadBalancingTcpFarm {
		return vs[0].(map[string]*IpLoadBalancingTcpFarm)[vs[1].(string)]
	}).(IpLoadBalancingTcpFarmOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingTcpFarmInput)(nil)).Elem(), &IpLoadBalancingTcpFarm{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingTcpFarmArrayInput)(nil)).Elem(), IpLoadBalancingTcpFarmArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingTcpFarmMapInput)(nil)).Elem(), IpLoadBalancingTcpFarmMap{})
	pulumi.RegisterOutputType(IpLoadBalancingTcpFarmOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingTcpFarmArrayOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingTcpFarmMapOutput{})
}
