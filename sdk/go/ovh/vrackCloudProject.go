// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Attach a Public Cloud Project to a VRack.
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
//			_, err := ovh.NewVrackCloudProject(ctx, "vcp", &ovh.VrackCloudProjectArgs{
//				ProjectId:   pulumi.String("67890"),
//				ServiceName: pulumi.String("12345"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type VrackCloudProject struct {
	pulumi.CustomResourceState

	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewVrackCloudProject registers a new resource with the given unique name, arguments, and options.
func NewVrackCloudProject(ctx *pulumi.Context,
	name string, args *VrackCloudProjectArgs, opts ...pulumi.ResourceOption) (*VrackCloudProject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource VrackCloudProject
	err := ctx.RegisterResource("ovh:index/vrackCloudProject:VrackCloudProject", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVrackCloudProject gets an existing VrackCloudProject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVrackCloudProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VrackCloudProjectState, opts ...pulumi.ResourceOption) (*VrackCloudProject, error) {
	var resource VrackCloudProject
	err := ctx.ReadResource("ovh:index/vrackCloudProject:VrackCloudProject", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VrackCloudProject resources.
type vrackCloudProjectState struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ProjectId *string `pulumi:"projectId"`
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

type VrackCloudProjectState struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ProjectId pulumi.StringPtrInput
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (VrackCloudProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackCloudProjectState)(nil)).Elem()
}

type vrackCloudProjectArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ProjectId string `pulumi:"projectId"`
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a VrackCloudProject resource.
type VrackCloudProjectArgs struct {
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ProjectId pulumi.StringInput
	// The id of the vrack. If omitted,
	// the `OVH_VRACK_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (VrackCloudProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vrackCloudProjectArgs)(nil)).Elem()
}

type VrackCloudProjectInput interface {
	pulumi.Input

	ToVrackCloudProjectOutput() VrackCloudProjectOutput
	ToVrackCloudProjectOutputWithContext(ctx context.Context) VrackCloudProjectOutput
}

func (*VrackCloudProject) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackCloudProject)(nil)).Elem()
}

func (i *VrackCloudProject) ToVrackCloudProjectOutput() VrackCloudProjectOutput {
	return i.ToVrackCloudProjectOutputWithContext(context.Background())
}

func (i *VrackCloudProject) ToVrackCloudProjectOutputWithContext(ctx context.Context) VrackCloudProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackCloudProjectOutput)
}

// VrackCloudProjectArrayInput is an input type that accepts VrackCloudProjectArray and VrackCloudProjectArrayOutput values.
// You can construct a concrete instance of `VrackCloudProjectArrayInput` via:
//
//	VrackCloudProjectArray{ VrackCloudProjectArgs{...} }
type VrackCloudProjectArrayInput interface {
	pulumi.Input

	ToVrackCloudProjectArrayOutput() VrackCloudProjectArrayOutput
	ToVrackCloudProjectArrayOutputWithContext(context.Context) VrackCloudProjectArrayOutput
}

type VrackCloudProjectArray []VrackCloudProjectInput

func (VrackCloudProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrackCloudProject)(nil)).Elem()
}

func (i VrackCloudProjectArray) ToVrackCloudProjectArrayOutput() VrackCloudProjectArrayOutput {
	return i.ToVrackCloudProjectArrayOutputWithContext(context.Background())
}

func (i VrackCloudProjectArray) ToVrackCloudProjectArrayOutputWithContext(ctx context.Context) VrackCloudProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackCloudProjectArrayOutput)
}

// VrackCloudProjectMapInput is an input type that accepts VrackCloudProjectMap and VrackCloudProjectMapOutput values.
// You can construct a concrete instance of `VrackCloudProjectMapInput` via:
//
//	VrackCloudProjectMap{ "key": VrackCloudProjectArgs{...} }
type VrackCloudProjectMapInput interface {
	pulumi.Input

	ToVrackCloudProjectMapOutput() VrackCloudProjectMapOutput
	ToVrackCloudProjectMapOutputWithContext(context.Context) VrackCloudProjectMapOutput
}

type VrackCloudProjectMap map[string]VrackCloudProjectInput

func (VrackCloudProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrackCloudProject)(nil)).Elem()
}

func (i VrackCloudProjectMap) ToVrackCloudProjectMapOutput() VrackCloudProjectMapOutput {
	return i.ToVrackCloudProjectMapOutputWithContext(context.Background())
}

func (i VrackCloudProjectMap) ToVrackCloudProjectMapOutputWithContext(ctx context.Context) VrackCloudProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VrackCloudProjectMapOutput)
}

type VrackCloudProjectOutput struct{ *pulumi.OutputState }

func (VrackCloudProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VrackCloudProject)(nil)).Elem()
}

func (o VrackCloudProjectOutput) ToVrackCloudProjectOutput() VrackCloudProjectOutput {
	return o
}

func (o VrackCloudProjectOutput) ToVrackCloudProjectOutputWithContext(ctx context.Context) VrackCloudProjectOutput {
	return o
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o VrackCloudProjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackCloudProject) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The id of the vrack. If omitted,
// the `OVH_VRACK_SERVICE` environment variable is used.
func (o VrackCloudProjectOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *VrackCloudProject) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type VrackCloudProjectArrayOutput struct{ *pulumi.OutputState }

func (VrackCloudProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VrackCloudProject)(nil)).Elem()
}

func (o VrackCloudProjectArrayOutput) ToVrackCloudProjectArrayOutput() VrackCloudProjectArrayOutput {
	return o
}

func (o VrackCloudProjectArrayOutput) ToVrackCloudProjectArrayOutputWithContext(ctx context.Context) VrackCloudProjectArrayOutput {
	return o
}

func (o VrackCloudProjectArrayOutput) Index(i pulumi.IntInput) VrackCloudProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VrackCloudProject {
		return vs[0].([]*VrackCloudProject)[vs[1].(int)]
	}).(VrackCloudProjectOutput)
}

type VrackCloudProjectMapOutput struct{ *pulumi.OutputState }

func (VrackCloudProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VrackCloudProject)(nil)).Elem()
}

func (o VrackCloudProjectMapOutput) ToVrackCloudProjectMapOutput() VrackCloudProjectMapOutput {
	return o
}

func (o VrackCloudProjectMapOutput) ToVrackCloudProjectMapOutputWithContext(ctx context.Context) VrackCloudProjectMapOutput {
	return o
}

func (o VrackCloudProjectMapOutput) MapIndex(k pulumi.StringInput) VrackCloudProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VrackCloudProject {
		return vs[0].(map[string]*VrackCloudProject)[vs[1].(string)]
	}).(VrackCloudProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VrackCloudProjectInput)(nil)).Elem(), &VrackCloudProject{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrackCloudProjectArrayInput)(nil)).Elem(), VrackCloudProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VrackCloudProjectMapInput)(nil)).Elem(), VrackCloudProjectMap{})
	pulumi.RegisterOutputType(VrackCloudProjectOutput{})
	pulumi.RegisterOutputType(VrackCloudProjectArrayOutput{})
	pulumi.RegisterOutputType(VrackCloudProjectMapOutput{})
}
