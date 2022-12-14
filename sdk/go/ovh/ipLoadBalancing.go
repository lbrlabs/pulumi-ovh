// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Orders an IP load balancing.
//
// ## Important
//
// This resource orders an OVH product for a long period of time and may generate heavy costs !
// Use with caution.
//
// __NOTE__ 1: the "default-payment-mean" will scan your registered bank accounts, credit card and paypal payment means to find your default payment mean.
//
// __NOTE__ 2: this resource is in beta state. Use with caution.
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
//				Description:   pulumi.StringRef("mycart"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			iplb, err := ovh.GetOrderCartProductPlan(ctx, &GetOrderCartProductPlanArgs{
//				CartId:        mycart.Id,
//				PriceCapacity: "renew",
//				Product:       "ipLoadbalancing",
//				PlanCode:      "iplb-lb1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			bhs, err := ovh.GetOrderCartProductOptionsPlan(ctx, &GetOrderCartProductOptionsPlanArgs{
//				CartId:          iplb.CartId,
//				PriceCapacity:   iplb.PriceCapacity,
//				Product:         iplb.Product,
//				PlanCode:        iplb.PlanCode,
//				OptionsPlanCode: "iplb-zone-lb1-rbx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ovh.NewIpLoadBalancing(ctx, "iplb-lb1", &ovh.IpLoadBalancingArgs{
//				OvhSubsidiary: pulumi.String(mycart.OvhSubsidiary),
//				DisplayName:   pulumi.String("my ip loadbalancing"),
//				PaymentMean:   pulumi.String("ovh-account"),
//				Plan: &IpLoadBalancingPlanArgs{
//					Duration:    pulumi.String(iplb.SelectedPrices[0].Duration),
//					PlanCode:    pulumi.String(iplb.PlanCode),
//					PricingMode: pulumi.String(iplb.SelectedPrices[0].PricingMode),
//				},
//				PlanOptions: IpLoadBalancingPlanOptionArray{
//					&IpLoadBalancingPlanOptionArgs{
//						Duration:    pulumi.String(bhs.SelectedPrices[0].Duration),
//						PlanCode:    pulumi.String(bhs.PlanCode),
//						PricingMode: pulumi.String(bhs.SelectedPrices[0].PricingMode),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type IpLoadBalancing struct {
	pulumi.CustomResourceState

	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Your IP load balancing
	IpLoadbalancing pulumi.StringOutput `pulumi:"ipLoadbalancing"`
	// The IPV4 associated to your IP load balancing
	Ipv4 pulumi.StringOutput `pulumi:"ipv4"`
	// The IPV6 associated to your IP load balancing. DEPRECATED.
	Ipv6 pulumi.StringOutput `pulumi:"ipv6"`
	// The metrics token associated with your IP load balancing
	MetricsToken pulumi.StringOutput `pulumi:"metricsToken"`
	// The offer of your IP load balancing
	Offer pulumi.StringOutput `pulumi:"offer"`
	// Available additional zone for your Load Balancer
	OrderableZones IpLoadBalancingOrderableZoneArrayOutput `pulumi:"orderableZones"`
	// Details about an Order
	Orders IpLoadBalancingOrderArrayOutput `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringOutput `pulumi:"ovhSubsidiary"`
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean pulumi.StringOutput `pulumi:"paymentMean"`
	// Product Plan to order
	Plan IpLoadBalancingPlanOutput `pulumi:"plan"`
	// Product Plan to order
	PlanOptions IpLoadBalancingPlanOptionArrayOutput `pulumi:"planOptions"`
	// The internal name of your IP load balancing
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
	SslConfiguration pulumi.StringOutput `pulumi:"sslConfiguration"`
	// Current state of your IP
	State pulumi.StringOutput `pulumi:"state"`
	// Vrack eligibility
	VrackEligibility pulumi.BoolOutput `pulumi:"vrackEligibility"`
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName pulumi.StringOutput `pulumi:"vrackName"`
	// Location where your service is
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewIpLoadBalancing registers a new resource with the given unique name, arguments, and options.
func NewIpLoadBalancing(ctx *pulumi.Context,
	name string, args *IpLoadBalancingArgs, opts ...pulumi.ResourceOption) (*IpLoadBalancing, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OvhSubsidiary == nil {
		return nil, errors.New("invalid value for required argument 'OvhSubsidiary'")
	}
	if args.PaymentMean == nil {
		return nil, errors.New("invalid value for required argument 'PaymentMean'")
	}
	if args.Plan == nil {
		return nil, errors.New("invalid value for required argument 'Plan'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IpLoadBalancing
	err := ctx.RegisterResource("ovh:index/ipLoadBalancing:IpLoadBalancing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpLoadBalancing gets an existing IpLoadBalancing resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpLoadBalancing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpLoadBalancingState, opts ...pulumi.ResourceOption) (*IpLoadBalancing, error) {
	var resource IpLoadBalancing
	err := ctx.ReadResource("ovh:index/ipLoadBalancing:IpLoadBalancing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpLoadBalancing resources.
type ipLoadBalancingState struct {
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName *string `pulumi:"displayName"`
	// Your IP load balancing
	IpLoadbalancing *string `pulumi:"ipLoadbalancing"`
	// The IPV4 associated to your IP load balancing
	Ipv4 *string `pulumi:"ipv4"`
	// The IPV6 associated to your IP load balancing. DEPRECATED.
	Ipv6 *string `pulumi:"ipv6"`
	// The metrics token associated with your IP load balancing
	MetricsToken *string `pulumi:"metricsToken"`
	// The offer of your IP load balancing
	Offer *string `pulumi:"offer"`
	// Available additional zone for your Load Balancer
	OrderableZones []IpLoadBalancingOrderableZone `pulumi:"orderableZones"`
	// Details about an Order
	Orders []IpLoadBalancingOrder `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *IpLoadBalancingPlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []IpLoadBalancingPlanOption `pulumi:"planOptions"`
	// The internal name of your IP load balancing
	ServiceName *string `pulumi:"serviceName"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
	SslConfiguration *string `pulumi:"sslConfiguration"`
	// Current state of your IP
	State *string `pulumi:"state"`
	// Vrack eligibility
	VrackEligibility *bool `pulumi:"vrackEligibility"`
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName *string `pulumi:"vrackName"`
	// Location where your service is
	Zones []string `pulumi:"zones"`
}

type IpLoadBalancingState struct {
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName pulumi.StringPtrInput
	// Your IP load balancing
	IpLoadbalancing pulumi.StringPtrInput
	// The IPV4 associated to your IP load balancing
	Ipv4 pulumi.StringPtrInput
	// The IPV6 associated to your IP load balancing. DEPRECATED.
	Ipv6 pulumi.StringPtrInput
	// The metrics token associated with your IP load balancing
	MetricsToken pulumi.StringPtrInput
	// The offer of your IP load balancing
	Offer pulumi.StringPtrInput
	// Available additional zone for your Load Balancer
	OrderableZones IpLoadBalancingOrderableZoneArrayInput
	// Details about an Order
	Orders IpLoadBalancingOrderArrayInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan IpLoadBalancingPlanPtrInput
	// Product Plan to order
	PlanOptions IpLoadBalancingPlanOptionArrayInput
	// The internal name of your IP load balancing
	ServiceName pulumi.StringPtrInput
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
	SslConfiguration pulumi.StringPtrInput
	// Current state of your IP
	State pulumi.StringPtrInput
	// Vrack eligibility
	VrackEligibility pulumi.BoolPtrInput
	// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
	VrackName pulumi.StringPtrInput
	// Location where your service is
	Zones pulumi.StringArrayInput
}

func (IpLoadBalancingState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingState)(nil)).Elem()
}

type ipLoadBalancingArgs struct {
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName *string `pulumi:"displayName"`
	// Ovh Subsidiary
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan IpLoadBalancingPlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []IpLoadBalancingPlanOption `pulumi:"planOptions"`
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
	SslConfiguration *string `pulumi:"sslConfiguration"`
}

// The set of arguments for constructing a IpLoadBalancing resource.
type IpLoadBalancingArgs struct {
	// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
	DisplayName pulumi.StringPtrInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringInput
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean pulumi.StringInput
	// Product Plan to order
	Plan IpLoadBalancingPlanInput
	// Product Plan to order
	PlanOptions IpLoadBalancingPlanOptionArrayInput
	// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
	SslConfiguration pulumi.StringPtrInput
}

func (IpLoadBalancingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipLoadBalancingArgs)(nil)).Elem()
}

type IpLoadBalancingInput interface {
	pulumi.Input

	ToIpLoadBalancingOutput() IpLoadBalancingOutput
	ToIpLoadBalancingOutputWithContext(ctx context.Context) IpLoadBalancingOutput
}

func (*IpLoadBalancing) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancing)(nil)).Elem()
}

func (i *IpLoadBalancing) ToIpLoadBalancingOutput() IpLoadBalancingOutput {
	return i.ToIpLoadBalancingOutputWithContext(context.Background())
}

func (i *IpLoadBalancing) ToIpLoadBalancingOutputWithContext(ctx context.Context) IpLoadBalancingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingOutput)
}

// IpLoadBalancingArrayInput is an input type that accepts IpLoadBalancingArray and IpLoadBalancingArrayOutput values.
// You can construct a concrete instance of `IpLoadBalancingArrayInput` via:
//
//	IpLoadBalancingArray{ IpLoadBalancingArgs{...} }
type IpLoadBalancingArrayInput interface {
	pulumi.Input

	ToIpLoadBalancingArrayOutput() IpLoadBalancingArrayOutput
	ToIpLoadBalancingArrayOutputWithContext(context.Context) IpLoadBalancingArrayOutput
}

type IpLoadBalancingArray []IpLoadBalancingInput

func (IpLoadBalancingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancing)(nil)).Elem()
}

func (i IpLoadBalancingArray) ToIpLoadBalancingArrayOutput() IpLoadBalancingArrayOutput {
	return i.ToIpLoadBalancingArrayOutputWithContext(context.Background())
}

func (i IpLoadBalancingArray) ToIpLoadBalancingArrayOutputWithContext(ctx context.Context) IpLoadBalancingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingArrayOutput)
}

// IpLoadBalancingMapInput is an input type that accepts IpLoadBalancingMap and IpLoadBalancingMapOutput values.
// You can construct a concrete instance of `IpLoadBalancingMapInput` via:
//
//	IpLoadBalancingMap{ "key": IpLoadBalancingArgs{...} }
type IpLoadBalancingMapInput interface {
	pulumi.Input

	ToIpLoadBalancingMapOutput() IpLoadBalancingMapOutput
	ToIpLoadBalancingMapOutputWithContext(context.Context) IpLoadBalancingMapOutput
}

type IpLoadBalancingMap map[string]IpLoadBalancingInput

func (IpLoadBalancingMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancing)(nil)).Elem()
}

func (i IpLoadBalancingMap) ToIpLoadBalancingMapOutput() IpLoadBalancingMapOutput {
	return i.ToIpLoadBalancingMapOutputWithContext(context.Background())
}

func (i IpLoadBalancingMap) ToIpLoadBalancingMapOutputWithContext(ctx context.Context) IpLoadBalancingMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpLoadBalancingMapOutput)
}

type IpLoadBalancingOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpLoadBalancing)(nil)).Elem()
}

func (o IpLoadBalancingOutput) ToIpLoadBalancingOutput() IpLoadBalancingOutput {
	return o
}

func (o IpLoadBalancingOutput) ToIpLoadBalancingOutputWithContext(ctx context.Context) IpLoadBalancingOutput {
	return o
}

// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
func (o IpLoadBalancingOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Your IP load balancing
func (o IpLoadBalancingOutput) IpLoadbalancing() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.IpLoadbalancing }).(pulumi.StringOutput)
}

// The IPV4 associated to your IP load balancing
func (o IpLoadBalancingOutput) Ipv4() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.Ipv4 }).(pulumi.StringOutput)
}

// The IPV6 associated to your IP load balancing. DEPRECATED.
func (o IpLoadBalancingOutput) Ipv6() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.Ipv6 }).(pulumi.StringOutput)
}

// The metrics token associated with your IP load balancing
func (o IpLoadBalancingOutput) MetricsToken() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.MetricsToken }).(pulumi.StringOutput)
}

// The offer of your IP load balancing
func (o IpLoadBalancingOutput) Offer() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.Offer }).(pulumi.StringOutput)
}

// Available additional zone for your Load Balancer
func (o IpLoadBalancingOutput) OrderableZones() IpLoadBalancingOrderableZoneArrayOutput {
	return o.ApplyT(func(v *IpLoadBalancing) IpLoadBalancingOrderableZoneArrayOutput { return v.OrderableZones }).(IpLoadBalancingOrderableZoneArrayOutput)
}

// Details about an Order
func (o IpLoadBalancingOutput) Orders() IpLoadBalancingOrderArrayOutput {
	return o.ApplyT(func(v *IpLoadBalancing) IpLoadBalancingOrderArrayOutput { return v.Orders }).(IpLoadBalancingOrderArrayOutput)
}

// Ovh Subsidiary
func (o IpLoadBalancingOutput) OvhSubsidiary() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.OvhSubsidiary }).(pulumi.StringOutput)
}

// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
func (o IpLoadBalancingOutput) PaymentMean() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.PaymentMean }).(pulumi.StringOutput)
}

// Product Plan to order
func (o IpLoadBalancingOutput) Plan() IpLoadBalancingPlanOutput {
	return o.ApplyT(func(v *IpLoadBalancing) IpLoadBalancingPlanOutput { return v.Plan }).(IpLoadBalancingPlanOutput)
}

// Product Plan to order
func (o IpLoadBalancingOutput) PlanOptions() IpLoadBalancingPlanOptionArrayOutput {
	return o.ApplyT(func(v *IpLoadBalancing) IpLoadBalancingPlanOptionArrayOutput { return v.PlanOptions }).(IpLoadBalancingPlanOptionArrayOutput)
}

// The internal name of your IP load balancing
func (o IpLoadBalancingOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
func (o IpLoadBalancingOutput) SslConfiguration() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.SslConfiguration }).(pulumi.StringOutput)
}

// Current state of your IP
func (o IpLoadBalancingOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Vrack eligibility
func (o IpLoadBalancingOutput) VrackEligibility() pulumi.BoolOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.BoolOutput { return v.VrackEligibility }).(pulumi.BoolOutput)
}

// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
func (o IpLoadBalancingOutput) VrackName() pulumi.StringOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringOutput { return v.VrackName }).(pulumi.StringOutput)
}

// Location where your service is
func (o IpLoadBalancingOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IpLoadBalancing) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

type IpLoadBalancingArrayOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpLoadBalancing)(nil)).Elem()
}

func (o IpLoadBalancingArrayOutput) ToIpLoadBalancingArrayOutput() IpLoadBalancingArrayOutput {
	return o
}

func (o IpLoadBalancingArrayOutput) ToIpLoadBalancingArrayOutputWithContext(ctx context.Context) IpLoadBalancingArrayOutput {
	return o
}

func (o IpLoadBalancingArrayOutput) Index(i pulumi.IntInput) IpLoadBalancingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpLoadBalancing {
		return vs[0].([]*IpLoadBalancing)[vs[1].(int)]
	}).(IpLoadBalancingOutput)
}

type IpLoadBalancingMapOutput struct{ *pulumi.OutputState }

func (IpLoadBalancingMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpLoadBalancing)(nil)).Elem()
}

func (o IpLoadBalancingMapOutput) ToIpLoadBalancingMapOutput() IpLoadBalancingMapOutput {
	return o
}

func (o IpLoadBalancingMapOutput) ToIpLoadBalancingMapOutputWithContext(ctx context.Context) IpLoadBalancingMapOutput {
	return o
}

func (o IpLoadBalancingMapOutput) MapIndex(k pulumi.StringInput) IpLoadBalancingOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpLoadBalancing {
		return vs[0].(map[string]*IpLoadBalancing)[vs[1].(string)]
	}).(IpLoadBalancingOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingInput)(nil)).Elem(), &IpLoadBalancing{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingArrayInput)(nil)).Elem(), IpLoadBalancingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpLoadBalancingMapInput)(nil)).Elem(), IpLoadBalancingMap{})
	pulumi.RegisterOutputType(IpLoadBalancingOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingArrayOutput{})
	pulumi.RegisterOutputType(IpLoadBalancingMapOutput{})
}
