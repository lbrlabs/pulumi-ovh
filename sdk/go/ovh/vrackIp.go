// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attach a Ip block to a VRack.
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
//			mycart, err := ovh.GetOrderCart(ctx, &GetOrderCartArgs{
//				OvhSubsidiary: "fr",
//				Description:   pulumi.StringRef("my cart"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vrackOrderCartProductPlan, err := ovh.GetOrderCartProductPlan(ctx, &GetOrderCartProductPlanArgs{
//				CartId:        mycart.Id,
//				PriceCapacity: "renew",
//				Product:       "vrack",
//				PlanCode:      "vrack",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vrackVrack, err := ovh.NewVrack(ctx, "vrackVrack", &ovh.VrackArgs{
//				Description:   pulumi.String(mycart.Description),
//				OvhSubsidiary: pulumi.String(mycart.OvhSubsidiary),
//				PaymentMean:   pulumi.String("fidelity"),
//				Plan: &VrackPlanArgs{
//					Duration:    pulumi.String(vrackOrderCartProductPlan.SelectedPrices[0].Duration),
//					PlanCode:    pulumi.String(vrackOrderCartProductPlan.PlanCode),
//					PricingMode: pulumi.String(vrackOrderCartProductPlan.SelectedPrices[0].PricingMode),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			ipblockOrderCartProductPlan, err := ovh.GetOrderCartProductPlan(ctx, &GetOrderCartProductPlanArgs{
//				CartId:        mycart.Id,
//				PriceCapacity: "renew",
//				Product:       "ip",
//				PlanCode:      "ip-v4-s30-ripe",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ipblockIpService, err := ovh.NewIpService(ctx, "ipblockIpService", &ovh.IpServiceArgs{
//				OvhSubsidiary: pulumi.String(mycart.OvhSubsidiary),
//				PaymentMean:   pulumi.String("ovh-account"),
//				Description:   pulumi.String(mycart.Description),
//				Plan: &IpServicePlanArgs{
//					Duration:    pulumi.String(ipblockOrderCartProductPlan.SelectedPrices[0].Duration),
//					PlanCode:    pulumi.String(ipblockOrderCartProductPlan.PlanCode),
//					PricingMode: pulumi.String(ipblockOrderCartProductPlan.SelectedPrices[0].PricingMode),
//					Configurations: IpServicePlanConfigurationArray{
//						&IpServicePlanConfigurationArgs{
//							Label: pulumi.String("country"),
//							Value: pulumi.String("FR"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ovh.NewVrackIp(ctx, "vrackblock", &ovh.VrackIpArgs{
//				ServiceName: vrackVrack.ServiceName,
//				Block:       ipblockIpService.Ip,
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type VrackIp struct {
	pulumi.CustomResourceState

	// Your IP block.
	Block pulumi.StringOutput `pulumi:"block"`
	// Your gateway
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Your IP block
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The internal name of your vrack
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Where you want your block announced on the network
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewVrackIp registers a new resource with the given unique name, arguments, and options.
func NewVrackIp(ctx *pulumi.Context,
	name string, args *VrackIpArgs, opts ...pulumi.ResourceOption) (*VrackIp, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Block == nil {
		return nil, errors.New("invalid value for required argument 'Block'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VrackIp
	err := ctx.RegisterResource("ovh:index/vrackIp:VrackIp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVrackIp gets an existing VrackIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVrackIp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VrackIpState, opts ...pulumi.ResourceOption) (*VrackIp, error) {
	var resource VrackIp
	err := ctx.ReadResource("ovh:index/vrackIp:VrackIp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VrackIp resources.
type vrackIpState struct {
	// Your IP block.
	Block *string `pulumi:"block"`
	// Your gateway
	Gateway *string `pulumi:"gateway"`
	// Your IP block
	Ip *string `pulumi:"ip"`
	// The internal name of your vrack
	ServiceName *string `pulumi:"serviceName"`
	// Where you want your block announced on the network
	Zone *string `pulumi:"zone"`
}

type VrackIpState struct {
	// Your IP block.
	Block pulumi.StringPtrInput
	// Your gateway
	Gateway pulumi.StringPtrInput
	// Your IP block
	Ip pulumi.StringPtrInput
	// The internal name of your vrack
	ServiceName pulumi.StringPtrInput
	// Where you want your block announced on the network
	Zone pulumi.StringPtrInput
}

func (VrackIpState) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackIpState)(nil)).Elem()
}

type vrackIpArgs struct {
	// Your IP block.
	Block string `pulumi:"block"`
	// The internal name of your vrack
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a VrackIp resource.
type VrackIpArgs struct {
	// Your IP block.
	Block pulumi.StringInput
	// The internal name of your vrack
	ServiceName pulumi.StringInput
}

func (VrackIpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackIpArgs)(nil)).Elem()
}

type VrackIpInput interface {
	pulumi.Input

	ToVrackIpOutput() VrackIpOutput
	ToVrackIpOutputWithContext(ctx context.Context) VrackIpOutput
}

func (*VrackIp) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackIp)(nil)).Elem()
}

func (i *VrackIp) ToVrackIpOutput() VrackIpOutput {
	return i.ToVrackIpOutputWithContext(context.Background())
}

func (i *VrackIp) ToVrackIpOutputWithContext(ctx context.Context) VrackIpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackIpOutput)
}

// VrackIpArrayInput is an input type that accepts VrackIpArray and VrackIpArrayOutput values.
// You can construct a concrete instance of `VrackIpArrayInput` via:
//
//	VrackIpArray{ VrackIpArgs{...} }
type VrackIpArrayInput interface {
	pulumi.Input

	ToVrackIpArrayOutput() VrackIpArrayOutput
	ToVrackIpArrayOutputWithContext(context.Context) VrackIpArrayOutput
}

type VrackIpArray []VrackIpInput

func (VrackIpArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrackIp)(nil)).Elem()
}

func (i VrackIpArray) ToVrackIpArrayOutput() VrackIpArrayOutput {
	return i.ToVrackIpArrayOutputWithContext(context.Background())
}

func (i VrackIpArray) ToVrackIpArrayOutputWithContext(ctx context.Context) VrackIpArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackIpArrayOutput)
}

// VrackIpMapInput is an input type that accepts VrackIpMap and VrackIpMapOutput values.
// You can construct a concrete instance of `VrackIpMapInput` via:
//
//	VrackIpMap{ "key": VrackIpArgs{...} }
type VrackIpMapInput interface {
	pulumi.Input

	ToVrackIpMapOutput() VrackIpMapOutput
	ToVrackIpMapOutputWithContext(context.Context) VrackIpMapOutput
}

type VrackIpMap map[string]VrackIpInput

func (VrackIpMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrackIp)(nil)).Elem()
}

func (i VrackIpMap) ToVrackIpMapOutput() VrackIpMapOutput {
	return i.ToVrackIpMapOutputWithContext(context.Background())
}

func (i VrackIpMap) ToVrackIpMapOutputWithContext(ctx context.Context) VrackIpMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackIpMapOutput)
}

type VrackIpOutput struct{ *pulumi.OutputState }

func (VrackIpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackIp)(nil)).Elem()
}

func (o VrackIpOutput) ToVrackIpOutput() VrackIpOutput {
	return o
}

func (o VrackIpOutput) ToVrackIpOutputWithContext(ctx context.Context) VrackIpOutput {
	return o
}

// Your IP block.
func (o VrackIpOutput) Block() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackIp) pulumi.StringOutput { return v.Block }).(pulumi.StringOutput)
}

// Your gateway
func (o VrackIpOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackIp) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Your IP block
func (o VrackIpOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackIp) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The internal name of your vrack
func (o VrackIpOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackIp) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Where you want your block announced on the network
func (o VrackIpOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackIp) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type VrackIpArrayOutput struct{ *pulumi.OutputState }

func (VrackIpArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrackIp)(nil)).Elem()
}

func (o VrackIpArrayOutput) ToVrackIpArrayOutput() VrackIpArrayOutput {
	return o
}

func (o VrackIpArrayOutput) ToVrackIpArrayOutputWithContext(ctx context.Context) VrackIpArrayOutput {
	return o
}

func (o VrackIpArrayOutput) Index(i pulumi.IntInput) VrackIpOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VrackIp {
		return vs[0].([]*VrackIp)[vs[1].(int)]
	}).(VrackIpOutput)
}

type VrackIpMapOutput struct{ *pulumi.OutputState }

func (VrackIpMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrackIp)(nil)).Elem()
}

func (o VrackIpMapOutput) ToVrackIpMapOutput() VrackIpMapOutput {
	return o
}

func (o VrackIpMapOutput) ToVrackIpMapOutputWithContext(ctx context.Context) VrackIpMapOutput {
	return o
}

func (o VrackIpMapOutput) MapIndex(k pulumi.StringInput) VrackIpOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VrackIp {
		return vs[0].(map[string]*VrackIp)[vs[1].(string)]
	}).(VrackIpOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VrackIpInput)(nil)).Elem(), &VrackIp{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrackIpArrayInput)(nil)).Elem(), VrackIpArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrackIpMapInput)(nil)).Elem(), VrackIpMap{})
	pulumi.RegisterOutputType(VrackIpOutput{})
	pulumi.RegisterOutputType(VrackIpArrayOutput{})
	pulumi.RegisterOutputType(VrackIpMapOutput{})
}
