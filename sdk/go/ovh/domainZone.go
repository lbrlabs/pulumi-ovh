// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a domain zone.
//
// ## Important
//
// This resource is in beta state. Use with caution.
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
//			}, nil)
//			if err != nil {
//				return err
//			}
//			zoneOrderCartProductPlan, err := ovh.GetOrderCartProductPlan(ctx, &GetOrderCartProductPlanArgs{
//				CartId:        mycart.Id,
//				PriceCapacity: "renew",
//				Product:       "dns",
//				PlanCode:      "zone",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ovh.NewDomainZone(ctx, "zoneDomainZone", &ovh.DomainZoneArgs{
//				OvhSubsidiary: pulumi.String(mycart.OvhSubsidiary),
//				PaymentMean:   pulumi.String("fidelity"),
//				Plan: &DomainZonePlanArgs{
//					Duration:    pulumi.String(zoneOrderCartProductPlan.SelectedPrices[0].Duration),
//					PlanCode:    pulumi.String(zoneOrderCartProductPlan.PlanCode),
//					PricingMode: pulumi.String(zoneOrderCartProductPlan.SelectedPrices[0].PricingMode),
//					Configurations: DomainZonePlanConfigurationArray{
//						&DomainZonePlanConfigurationArgs{
//							Label: pulumi.String("zone"),
//							Value: pulumi.String("myzone.mydomain.com"),
//						},
//						&DomainZonePlanConfigurationArgs{
//							Label: pulumi.String("template"),
//							Value: pulumi.String("minimized"),
//						},
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
type DomainZone struct {
	pulumi.CustomResourceState

	// Is DNSSEC supported by this zone
	DnssecSupported pulumi.BoolOutput `pulumi:"dnssecSupported"`
	// hasDnsAnycast flag of the DNS zone
	HasDnsAnycast pulumi.BoolOutput `pulumi:"hasDnsAnycast"`
	// Last update date of the DNS zone
	LastUpdate pulumi.StringOutput `pulumi:"lastUpdate"`
	// Zone name
	Name pulumi.StringOutput `pulumi:"name"`
	// Name servers that host the DNS zone
	NameServers pulumi.StringArrayOutput `pulumi:"nameServers"`
	// Details about an Order
	Orders DomainZoneOrderArrayOutput `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringOutput `pulumi:"ovhSubsidiary"`
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean pulumi.StringOutput `pulumi:"paymentMean"`
	// Product Plan to order
	Plan DomainZonePlanOutput `pulumi:"plan"`
	// Product Plan to order
	PlanOptions DomainZonePlanOptionArrayOutput `pulumi:"planOptions"`
}

// NewDomainZone registers a new resource with the given unique name, arguments, and options.
func NewDomainZone(ctx *pulumi.Context,
	name string, args *DomainZoneArgs, opts ...pulumi.ResourceOption) (*DomainZone, error) {
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
	var resource DomainZone
	err := ctx.RegisterResource("ovh:index/domainZone:DomainZone", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainZone gets an existing DomainZone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainZone(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainZoneState, opts ...pulumi.ResourceOption) (*DomainZone, error) {
	var resource DomainZone
	err := ctx.ReadResource("ovh:index/domainZone:DomainZone", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainZone resources.
type domainZoneState struct {
	// Is DNSSEC supported by this zone
	DnssecSupported *bool `pulumi:"dnssecSupported"`
	// hasDnsAnycast flag of the DNS zone
	HasDnsAnycast *bool `pulumi:"hasDnsAnycast"`
	// Last update date of the DNS zone
	LastUpdate *string `pulumi:"lastUpdate"`
	// Zone name
	Name *string `pulumi:"name"`
	// Name servers that host the DNS zone
	NameServers []string `pulumi:"nameServers"`
	// Details about an Order
	Orders []DomainZoneOrder `pulumi:"orders"`
	// Ovh Subsidiary
	OvhSubsidiary *string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean *string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan *DomainZonePlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []DomainZonePlanOption `pulumi:"planOptions"`
}

type DomainZoneState struct {
	// Is DNSSEC supported by this zone
	DnssecSupported pulumi.BoolPtrInput
	// hasDnsAnycast flag of the DNS zone
	HasDnsAnycast pulumi.BoolPtrInput
	// Last update date of the DNS zone
	LastUpdate pulumi.StringPtrInput
	// Zone name
	Name pulumi.StringPtrInput
	// Name servers that host the DNS zone
	NameServers pulumi.StringArrayInput
	// Details about an Order
	Orders DomainZoneOrderArrayInput
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringPtrInput
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean pulumi.StringPtrInput
	// Product Plan to order
	Plan DomainZonePlanPtrInput
	// Product Plan to order
	PlanOptions DomainZonePlanOptionArrayInput
}

func (DomainZoneState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainZoneState)(nil)).Elem()
}

type domainZoneArgs struct {
	// Ovh Subsidiary
	OvhSubsidiary string `pulumi:"ovhSubsidiary"`
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean string `pulumi:"paymentMean"`
	// Product Plan to order
	Plan DomainZonePlan `pulumi:"plan"`
	// Product Plan to order
	PlanOptions []DomainZonePlanOption `pulumi:"planOptions"`
}

// The set of arguments for constructing a DomainZone resource.
type DomainZoneArgs struct {
	// Ovh Subsidiary
	OvhSubsidiary pulumi.StringInput
	// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
	PaymentMean pulumi.StringInput
	// Product Plan to order
	Plan DomainZonePlanInput
	// Product Plan to order
	PlanOptions DomainZonePlanOptionArrayInput
}

func (DomainZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainZoneArgs)(nil)).Elem()
}

type DomainZoneInput interface {
	pulumi.Input

	ToDomainZoneOutput() DomainZoneOutput
	ToDomainZoneOutputWithContext(ctx context.Context) DomainZoneOutput
}

func (*DomainZone) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainZone)(nil)).Elem()
}

func (i *DomainZone) ToDomainZoneOutput() DomainZoneOutput {
	return i.ToDomainZoneOutputWithContext(context.Background())
}

func (i *DomainZone) ToDomainZoneOutputWithContext(ctx context.Context) DomainZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneOutput)
}

// DomainZoneArrayInput is an input type that accepts DomainZoneArray and DomainZoneArrayOutput values.
// You can construct a concrete instance of `DomainZoneArrayInput` via:
//
//	DomainZoneArray{ DomainZoneArgs{...} }
type DomainZoneArrayInput interface {
	pulumi.Input

	ToDomainZoneArrayOutput() DomainZoneArrayOutput
	ToDomainZoneArrayOutputWithContext(context.Context) DomainZoneArrayOutput
}

type DomainZoneArray []DomainZoneInput

func (DomainZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainZone)(nil)).Elem()
}

func (i DomainZoneArray) ToDomainZoneArrayOutput() DomainZoneArrayOutput {
	return i.ToDomainZoneArrayOutputWithContext(context.Background())
}

func (i DomainZoneArray) ToDomainZoneArrayOutputWithContext(ctx context.Context) DomainZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneArrayOutput)
}

// DomainZoneMapInput is an input type that accepts DomainZoneMap and DomainZoneMapOutput values.
// You can construct a concrete instance of `DomainZoneMapInput` via:
//
//	DomainZoneMap{ "key": DomainZoneArgs{...} }
type DomainZoneMapInput interface {
	pulumi.Input

	ToDomainZoneMapOutput() DomainZoneMapOutput
	ToDomainZoneMapOutputWithContext(context.Context) DomainZoneMapOutput
}

type DomainZoneMap map[string]DomainZoneInput

func (DomainZoneMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainZone)(nil)).Elem()
}

func (i DomainZoneMap) ToDomainZoneMapOutput() DomainZoneMapOutput {
	return i.ToDomainZoneMapOutputWithContext(context.Background())
}

func (i DomainZoneMap) ToDomainZoneMapOutputWithContext(ctx context.Context) DomainZoneMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneMapOutput)
}

type DomainZoneOutput struct{ *pulumi.OutputState }

func (DomainZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainZone)(nil)).Elem()
}

func (o DomainZoneOutput) ToDomainZoneOutput() DomainZoneOutput {
	return o
}

func (o DomainZoneOutput) ToDomainZoneOutputWithContext(ctx context.Context) DomainZoneOutput {
	return o
}

// Is DNSSEC supported by this zone
func (o DomainZoneOutput) DnssecSupported() pulumi.BoolOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.BoolOutput { return v.DnssecSupported }).(pulumi.BoolOutput)
}

// hasDnsAnycast flag of the DNS zone
func (o DomainZoneOutput) HasDnsAnycast() pulumi.BoolOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.BoolOutput { return v.HasDnsAnycast }).(pulumi.BoolOutput)
}

// Last update date of the DNS zone
func (o DomainZoneOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringOutput { return v.LastUpdate }).(pulumi.StringOutput)
}

// Zone name
func (o DomainZoneOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Name servers that host the DNS zone
func (o DomainZoneOutput) NameServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringArrayOutput { return v.NameServers }).(pulumi.StringArrayOutput)
}

// Details about an Order
func (o DomainZoneOutput) Orders() DomainZoneOrderArrayOutput {
	return o.ApplyT(func(v *DomainZone) DomainZoneOrderArrayOutput { return v.Orders }).(DomainZoneOrderArrayOutput)
}

// Ovh Subsidiary
func (o DomainZoneOutput) OvhSubsidiary() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringOutput { return v.OvhSubsidiary }).(pulumi.StringOutput)
}

// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
func (o DomainZoneOutput) PaymentMean() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZone) pulumi.StringOutput { return v.PaymentMean }).(pulumi.StringOutput)
}

// Product Plan to order
func (o DomainZoneOutput) Plan() DomainZonePlanOutput {
	return o.ApplyT(func(v *DomainZone) DomainZonePlanOutput { return v.Plan }).(DomainZonePlanOutput)
}

// Product Plan to order
func (o DomainZoneOutput) PlanOptions() DomainZonePlanOptionArrayOutput {
	return o.ApplyT(func(v *DomainZone) DomainZonePlanOptionArrayOutput { return v.PlanOptions }).(DomainZonePlanOptionArrayOutput)
}

type DomainZoneArrayOutput struct{ *pulumi.OutputState }

func (DomainZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainZone)(nil)).Elem()
}

func (o DomainZoneArrayOutput) ToDomainZoneArrayOutput() DomainZoneArrayOutput {
	return o
}

func (o DomainZoneArrayOutput) ToDomainZoneArrayOutputWithContext(ctx context.Context) DomainZoneArrayOutput {
	return o
}

func (o DomainZoneArrayOutput) Index(i pulumi.IntInput) DomainZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainZone {
		return vs[0].([]*DomainZone)[vs[1].(int)]
	}).(DomainZoneOutput)
}

type DomainZoneMapOutput struct{ *pulumi.OutputState }

func (DomainZoneMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainZone)(nil)).Elem()
}

func (o DomainZoneMapOutput) ToDomainZoneMapOutput() DomainZoneMapOutput {
	return o
}

func (o DomainZoneMapOutput) ToDomainZoneMapOutputWithContext(ctx context.Context) DomainZoneMapOutput {
	return o
}

func (o DomainZoneMapOutput) MapIndex(k pulumi.StringInput) DomainZoneOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainZone {
		return vs[0].(map[string]*DomainZone)[vs[1].(string)]
	}).(DomainZoneOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneInput)(nil)).Elem(), &DomainZone{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneArrayInput)(nil)).Elem(), DomainZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneMapInput)(nil)).Elem(), DomainZoneMap{})
	pulumi.RegisterOutputType(DomainZoneOutput{})
	pulumi.RegisterOutputType(DomainZoneArrayOutput{})
	pulumi.RegisterOutputType(DomainZoneMapOutput{})
}
