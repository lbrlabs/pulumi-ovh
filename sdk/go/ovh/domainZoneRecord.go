// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a OVH domain zone record.
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
//			_, err := ovh.NewDomainZoneRecord(ctx, "test", &ovh.DomainZoneRecordArgs{
//				Fieldtype: pulumi.String("A"),
//				Subdomain: pulumi.String("test"),
//				Target:    pulumi.String("0.0.0.0"),
//				Ttl:       pulumi.Int(3600),
//				Zone:      pulumi.String("testdemo.ovh"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # OVH record can be imported using the `id` and the `zone`, eg
//
// ```sh
//
//	$ pulumi import ovh:index/domainZoneRecord:DomainZoneRecord test 1234OVH_ID.zone.tld
//
// ```
type DomainZoneRecord struct {
	pulumi.CustomResourceState

	// The type of the record
	Fieldtype pulumi.StringOutput `pulumi:"fieldtype"`
	// The name of the record
	Subdomain pulumi.StringPtrOutput `pulumi:"subdomain"`
	// The value of the record
	Target pulumi.StringOutput `pulumi:"target"`
	// The TTL of the record
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
	// The domain to add the record to
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewDomainZoneRecord registers a new resource with the given unique name, arguments, and options.
func NewDomainZoneRecord(ctx *pulumi.Context,
	name string, args *DomainZoneRecordArgs, opts ...pulumi.ResourceOption) (*DomainZoneRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Fieldtype == nil {
		return nil, errors.New("invalid value for required argument 'Fieldtype'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DomainZoneRecord
	err := ctx.RegisterResource("ovh:index/domainZoneRecord:DomainZoneRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomainZoneRecord gets an existing DomainZoneRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomainZoneRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainZoneRecordState, opts ...pulumi.ResourceOption) (*DomainZoneRecord, error) {
	var resource DomainZoneRecord
	err := ctx.ReadResource("ovh:index/domainZoneRecord:DomainZoneRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DomainZoneRecord resources.
type domainZoneRecordState struct {
	// The type of the record
	Fieldtype *string `pulumi:"fieldtype"`
	// The name of the record
	Subdomain *string `pulumi:"subdomain"`
	// The value of the record
	Target *string `pulumi:"target"`
	// The TTL of the record
	Ttl *int `pulumi:"ttl"`
	// The domain to add the record to
	Zone *string `pulumi:"zone"`
}

type DomainZoneRecordState struct {
	// The type of the record
	Fieldtype pulumi.StringPtrInput
	// The name of the record
	Subdomain pulumi.StringPtrInput
	// The value of the record
	Target pulumi.StringPtrInput
	// The TTL of the record
	Ttl pulumi.IntPtrInput
	// The domain to add the record to
	Zone pulumi.StringPtrInput
}

func (DomainZoneRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainZoneRecordState)(nil)).Elem()
}

type domainZoneRecordArgs struct {
	// The type of the record
	Fieldtype string `pulumi:"fieldtype"`
	// The name of the record
	Subdomain *string `pulumi:"subdomain"`
	// The value of the record
	Target string `pulumi:"target"`
	// The TTL of the record
	Ttl *int `pulumi:"ttl"`
	// The domain to add the record to
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a DomainZoneRecord resource.
type DomainZoneRecordArgs struct {
	// The type of the record
	Fieldtype pulumi.StringInput
	// The name of the record
	Subdomain pulumi.StringPtrInput
	// The value of the record
	Target pulumi.StringInput
	// The TTL of the record
	Ttl pulumi.IntPtrInput
	// The domain to add the record to
	Zone pulumi.StringInput
}

func (DomainZoneRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainZoneRecordArgs)(nil)).Elem()
}

type DomainZoneRecordInput interface {
	pulumi.Input

	ToDomainZoneRecordOutput() DomainZoneRecordOutput
	ToDomainZoneRecordOutputWithContext(ctx context.Context) DomainZoneRecordOutput
}

func (*DomainZoneRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainZoneRecord)(nil)).Elem()
}

func (i *DomainZoneRecord) ToDomainZoneRecordOutput() DomainZoneRecordOutput {
	return i.ToDomainZoneRecordOutputWithContext(context.Background())
}

func (i *DomainZoneRecord) ToDomainZoneRecordOutputWithContext(ctx context.Context) DomainZoneRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneRecordOutput)
}

// DomainZoneRecordArrayInput is an input type that accepts DomainZoneRecordArray and DomainZoneRecordArrayOutput values.
// You can construct a concrete instance of `DomainZoneRecordArrayInput` via:
//
//	DomainZoneRecordArray{ DomainZoneRecordArgs{...} }
type DomainZoneRecordArrayInput interface {
	pulumi.Input

	ToDomainZoneRecordArrayOutput() DomainZoneRecordArrayOutput
	ToDomainZoneRecordArrayOutputWithContext(context.Context) DomainZoneRecordArrayOutput
}

type DomainZoneRecordArray []DomainZoneRecordInput

func (DomainZoneRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainZoneRecord)(nil)).Elem()
}

func (i DomainZoneRecordArray) ToDomainZoneRecordArrayOutput() DomainZoneRecordArrayOutput {
	return i.ToDomainZoneRecordArrayOutputWithContext(context.Background())
}

func (i DomainZoneRecordArray) ToDomainZoneRecordArrayOutputWithContext(ctx context.Context) DomainZoneRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneRecordArrayOutput)
}

// DomainZoneRecordMapInput is an input type that accepts DomainZoneRecordMap and DomainZoneRecordMapOutput values.
// You can construct a concrete instance of `DomainZoneRecordMapInput` via:
//
//	DomainZoneRecordMap{ "key": DomainZoneRecordArgs{...} }
type DomainZoneRecordMapInput interface {
	pulumi.Input

	ToDomainZoneRecordMapOutput() DomainZoneRecordMapOutput
	ToDomainZoneRecordMapOutputWithContext(context.Context) DomainZoneRecordMapOutput
}

type DomainZoneRecordMap map[string]DomainZoneRecordInput

func (DomainZoneRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainZoneRecord)(nil)).Elem()
}

func (i DomainZoneRecordMap) ToDomainZoneRecordMapOutput() DomainZoneRecordMapOutput {
	return i.ToDomainZoneRecordMapOutputWithContext(context.Background())
}

func (i DomainZoneRecordMap) ToDomainZoneRecordMapOutputWithContext(ctx context.Context) DomainZoneRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainZoneRecordMapOutput)
}

type DomainZoneRecordOutput struct{ *pulumi.OutputState }

func (DomainZoneRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainZoneRecord)(nil)).Elem()
}

func (o DomainZoneRecordOutput) ToDomainZoneRecordOutput() DomainZoneRecordOutput {
	return o
}

func (o DomainZoneRecordOutput) ToDomainZoneRecordOutputWithContext(ctx context.Context) DomainZoneRecordOutput {
	return o
}

// The type of the record
func (o DomainZoneRecordOutput) Fieldtype() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZoneRecord) pulumi.StringOutput { return v.Fieldtype }).(pulumi.StringOutput)
}

// The name of the record
func (o DomainZoneRecordOutput) Subdomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainZoneRecord) pulumi.StringPtrOutput { return v.Subdomain }).(pulumi.StringPtrOutput)
}

// The value of the record
func (o DomainZoneRecordOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZoneRecord) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

// The TTL of the record
func (o DomainZoneRecordOutput) Ttl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DomainZoneRecord) pulumi.IntPtrOutput { return v.Ttl }).(pulumi.IntPtrOutput)
}

// The domain to add the record to
func (o DomainZoneRecordOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainZoneRecord) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type DomainZoneRecordArrayOutput struct{ *pulumi.OutputState }

func (DomainZoneRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DomainZoneRecord)(nil)).Elem()
}

func (o DomainZoneRecordArrayOutput) ToDomainZoneRecordArrayOutput() DomainZoneRecordArrayOutput {
	return o
}

func (o DomainZoneRecordArrayOutput) ToDomainZoneRecordArrayOutputWithContext(ctx context.Context) DomainZoneRecordArrayOutput {
	return o
}

func (o DomainZoneRecordArrayOutput) Index(i pulumi.IntInput) DomainZoneRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DomainZoneRecord {
		return vs[0].([]*DomainZoneRecord)[vs[1].(int)]
	}).(DomainZoneRecordOutput)
}

type DomainZoneRecordMapOutput struct{ *pulumi.OutputState }

func (DomainZoneRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DomainZoneRecord)(nil)).Elem()
}

func (o DomainZoneRecordMapOutput) ToDomainZoneRecordMapOutput() DomainZoneRecordMapOutput {
	return o
}

func (o DomainZoneRecordMapOutput) ToDomainZoneRecordMapOutputWithContext(ctx context.Context) DomainZoneRecordMapOutput {
	return o
}

func (o DomainZoneRecordMapOutput) MapIndex(k pulumi.StringInput) DomainZoneRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DomainZoneRecord {
		return vs[0].(map[string]*DomainZoneRecord)[vs[1].(string)]
	}).(DomainZoneRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneRecordInput)(nil)).Elem(), &DomainZoneRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneRecordArrayInput)(nil)).Elem(), DomainZoneRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DomainZoneRecordMapInput)(nil)).Elem(), DomainZoneRecordMap{})
	pulumi.RegisterOutputType(DomainZoneRecordOutput{})
	pulumi.RegisterOutputType(DomainZoneRecordArrayOutput{})
	pulumi.RegisterOutputType(DomainZoneRecordMapOutput{})
}
