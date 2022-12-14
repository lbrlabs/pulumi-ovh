// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a OVH IP reverse.
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
//			_, err := ovh.NewIpReverse(ctx, "test", &ovh.IpReverseArgs{
//				Ip:        pulumi.String("192.0.2.0/24"),
//				IpReverse: pulumi.String("192.0.2.1"),
//				Reverse:   pulumi.String("example.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type IpReverse struct {
	pulumi.CustomResourceState

	// The IP block to which the IP belongs
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The IP to set the reverse of
	IpReverse pulumi.StringOutput `pulumi:"ipReverse"`
	// The value of the reverse
	Reverse pulumi.StringOutput `pulumi:"reverse"`
}

// NewIpReverse registers a new resource with the given unique name, arguments, and options.
func NewIpReverse(ctx *pulumi.Context,
	name string, args *IpReverseArgs, opts ...pulumi.ResourceOption) (*IpReverse, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.IpReverse == nil {
		return nil, errors.New("invalid value for required argument 'IpReverse'")
	}
	if args.Reverse == nil {
		return nil, errors.New("invalid value for required argument 'Reverse'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource IpReverse
	err := ctx.RegisterResource("ovh:index/ipReverse:IpReverse", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIpReverse gets an existing IpReverse resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIpReverse(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IpReverseState, opts ...pulumi.ResourceOption) (*IpReverse, error) {
	var resource IpReverse
	err := ctx.ReadResource("ovh:index/ipReverse:IpReverse", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IpReverse resources.
type ipReverseState struct {
	// The IP block to which the IP belongs
	Ip *string `pulumi:"ip"`
	// The IP to set the reverse of
	IpReverse *string `pulumi:"ipReverse"`
	// The value of the reverse
	Reverse *string `pulumi:"reverse"`
}

type IpReverseState struct {
	// The IP block to which the IP belongs
	Ip pulumi.StringPtrInput
	// The IP to set the reverse of
	IpReverse pulumi.StringPtrInput
	// The value of the reverse
	Reverse pulumi.StringPtrInput
}

func (IpReverseState) ElementType() reflect.Type {
	return reflect.TypeOf((*ipReverseState)(nil)).Elem()
}

type ipReverseArgs struct {
	// The IP block to which the IP belongs
	Ip string `pulumi:"ip"`
	// The IP to set the reverse of
	IpReverse string `pulumi:"ipReverse"`
	// The value of the reverse
	Reverse string `pulumi:"reverse"`
}

// The set of arguments for constructing a IpReverse resource.
type IpReverseArgs struct {
	// The IP block to which the IP belongs
	Ip pulumi.StringInput
	// The IP to set the reverse of
	IpReverse pulumi.StringInput
	// The value of the reverse
	Reverse pulumi.StringInput
}

func (IpReverseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ipReverseArgs)(nil)).Elem()
}

type IpReverseInput interface {
	pulumi.Input

	ToIpReverseOutput() IpReverseOutput
	ToIpReverseOutputWithContext(ctx context.Context) IpReverseOutput
}

func (*IpReverse) ElementType() reflect.Type {
	return reflect.TypeOf((**IpReverse)(nil)).Elem()
}

func (i *IpReverse) ToIpReverseOutput() IpReverseOutput {
	return i.ToIpReverseOutputWithContext(context.Background())
}

func (i *IpReverse) ToIpReverseOutputWithContext(ctx context.Context) IpReverseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpReverseOutput)
}

// IpReverseArrayInput is an input type that accepts IpReverseArray and IpReverseArrayOutput values.
// You can construct a concrete instance of `IpReverseArrayInput` via:
//
//	IpReverseArray{ IpReverseArgs{...} }
type IpReverseArrayInput interface {
	pulumi.Input

	ToIpReverseArrayOutput() IpReverseArrayOutput
	ToIpReverseArrayOutputWithContext(context.Context) IpReverseArrayOutput
}

type IpReverseArray []IpReverseInput

func (IpReverseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpReverse)(nil)).Elem()
}

func (i IpReverseArray) ToIpReverseArrayOutput() IpReverseArrayOutput {
	return i.ToIpReverseArrayOutputWithContext(context.Background())
}

func (i IpReverseArray) ToIpReverseArrayOutputWithContext(ctx context.Context) IpReverseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpReverseArrayOutput)
}

// IpReverseMapInput is an input type that accepts IpReverseMap and IpReverseMapOutput values.
// You can construct a concrete instance of `IpReverseMapInput` via:
//
//	IpReverseMap{ "key": IpReverseArgs{...} }
type IpReverseMapInput interface {
	pulumi.Input

	ToIpReverseMapOutput() IpReverseMapOutput
	ToIpReverseMapOutputWithContext(context.Context) IpReverseMapOutput
}

type IpReverseMap map[string]IpReverseInput

func (IpReverseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpReverse)(nil)).Elem()
}

func (i IpReverseMap) ToIpReverseMapOutput() IpReverseMapOutput {
	return i.ToIpReverseMapOutputWithContext(context.Background())
}

func (i IpReverseMap) ToIpReverseMapOutputWithContext(ctx context.Context) IpReverseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpReverseMapOutput)
}

type IpReverseOutput struct{ *pulumi.OutputState }

func (IpReverseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpReverse)(nil)).Elem()
}

func (o IpReverseOutput) ToIpReverseOutput() IpReverseOutput {
	return o
}

func (o IpReverseOutput) ToIpReverseOutputWithContext(ctx context.Context) IpReverseOutput {
	return o
}

// The IP block to which the IP belongs
func (o IpReverseOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *IpReverse) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The IP to set the reverse of
func (o IpReverseOutput) IpReverse() pulumi.StringOutput {
	return o.ApplyT(func(v *IpReverse) pulumi.StringOutput { return v.IpReverse }).(pulumi.StringOutput)
}

// The value of the reverse
func (o IpReverseOutput) Reverse() pulumi.StringOutput {
	return o.ApplyT(func(v *IpReverse) pulumi.StringOutput { return v.Reverse }).(pulumi.StringOutput)
}

type IpReverseArrayOutput struct{ *pulumi.OutputState }

func (IpReverseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IpReverse)(nil)).Elem()
}

func (o IpReverseArrayOutput) ToIpReverseArrayOutput() IpReverseArrayOutput {
	return o
}

func (o IpReverseArrayOutput) ToIpReverseArrayOutputWithContext(ctx context.Context) IpReverseArrayOutput {
	return o
}

func (o IpReverseArrayOutput) Index(i pulumi.IntInput) IpReverseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IpReverse {
		return vs[0].([]*IpReverse)[vs[1].(int)]
	}).(IpReverseOutput)
}

type IpReverseMapOutput struct{ *pulumi.OutputState }

func (IpReverseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IpReverse)(nil)).Elem()
}

func (o IpReverseMapOutput) ToIpReverseMapOutput() IpReverseMapOutput {
	return o
}

func (o IpReverseMapOutput) ToIpReverseMapOutputWithContext(ctx context.Context) IpReverseMapOutput {
	return o
}

func (o IpReverseMapOutput) MapIndex(k pulumi.StringInput) IpReverseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IpReverse {
		return vs[0].(map[string]*IpReverse)[vs[1].(string)]
	}).(IpReverseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IpReverseInput)(nil)).Elem(), &IpReverse{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpReverseArrayInput)(nil)).Elem(), IpReverseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpReverseMapInput)(nil)).Elem(), IpReverseMap{})
	pulumi.RegisterOutputType(IpReverseOutput{})
	pulumi.RegisterOutputType(IpReverseArrayOutput{})
	pulumi.RegisterOutputType(IpReverseMapOutput{})
}
