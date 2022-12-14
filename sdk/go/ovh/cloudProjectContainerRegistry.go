// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a container registry associated with a public cloud project.
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
//			regcap, err := ovh.GetCloudProjectCapabilitiesContainerFilter(ctx, &GetCloudProjectCapabilitiesContainerFilterArgs{
//				ServiceName: "XXXXXX",
//				PlanName:    "SMALL",
//				Region:      "GRA",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ovh.NewCloudProjectContainerRegistry(ctx, "reg", &ovh.CloudProjectContainerRegistryArgs{
//				ServiceName: pulumi.String(regcap.ServiceName),
//				PlanId:      pulumi.String(regcap.Id),
//				Region:      pulumi.String(regcap.Region),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type CloudProjectContainerRegistry struct {
	pulumi.CustomResourceState

	// Plan creation date
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Registry name
	Name pulumi.StringOutput `pulumi:"name"`
	// Plan ID of the registry
	PlanId pulumi.StringOutput `pulumi:"planId"`
	// Plan of the registry
	Plans CloudProjectContainerRegistryPlanArrayOutput `pulumi:"plans"`
	// Project ID of your registry
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Region of the registry
	Region pulumi.StringOutput `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Current size of the registry (bytes)
	Size pulumi.IntOutput `pulumi:"size"`
	// Registry status
	Status pulumi.StringOutput `pulumi:"status"`
	// Registry last update date
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Access url of the registry
	Url pulumi.StringOutput `pulumi:"url"`
	// Version of your registry
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewCloudProjectContainerRegistry registers a new resource with the given unique name, arguments, and options.
func NewCloudProjectContainerRegistry(ctx *pulumi.Context,
	name string, args *CloudProjectContainerRegistryArgs, opts ...pulumi.ResourceOption) (*CloudProjectContainerRegistry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CloudProjectContainerRegistry
	err := ctx.RegisterResource("ovh:index/cloudProjectContainerRegistry:CloudProjectContainerRegistry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudProjectContainerRegistry gets an existing CloudProjectContainerRegistry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudProjectContainerRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudProjectContainerRegistryState, opts ...pulumi.ResourceOption) (*CloudProjectContainerRegistry, error) {
	var resource CloudProjectContainerRegistry
	err := ctx.ReadResource("ovh:index/cloudProjectContainerRegistry:CloudProjectContainerRegistry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudProjectContainerRegistry resources.
type cloudProjectContainerRegistryState struct {
	// Plan creation date
	CreatedAt *string `pulumi:"createdAt"`
	// Registry name
	Name *string `pulumi:"name"`
	// Plan ID of the registry
	PlanId *string `pulumi:"planId"`
	// Plan of the registry
	Plans []CloudProjectContainerRegistryPlan `pulumi:"plans"`
	// Project ID of your registry
	ProjectId *string `pulumi:"projectId"`
	// Region of the registry
	Region *string `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Current size of the registry (bytes)
	Size *int `pulumi:"size"`
	// Registry status
	Status *string `pulumi:"status"`
	// Registry last update date
	UpdatedAt *string `pulumi:"updatedAt"`
	// Access url of the registry
	Url *string `pulumi:"url"`
	// Version of your registry
	Version *string `pulumi:"version"`
}

type CloudProjectContainerRegistryState struct {
	// Plan creation date
	CreatedAt pulumi.StringPtrInput
	// Registry name
	Name pulumi.StringPtrInput
	// Plan ID of the registry
	PlanId pulumi.StringPtrInput
	// Plan of the registry
	Plans CloudProjectContainerRegistryPlanArrayInput
	// Project ID of your registry
	ProjectId pulumi.StringPtrInput
	// Region of the registry
	Region pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Current size of the registry (bytes)
	Size pulumi.IntPtrInput
	// Registry status
	Status pulumi.StringPtrInput
	// Registry last update date
	UpdatedAt pulumi.StringPtrInput
	// Access url of the registry
	Url pulumi.StringPtrInput
	// Version of your registry
	Version pulumi.StringPtrInput
}

func (CloudProjectContainerRegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectContainerRegistryState)(nil)).Elem()
}

type cloudProjectContainerRegistryArgs struct {
	// Registry name
	Name *string `pulumi:"name"`
	// Plan ID of the registry
	PlanId *string `pulumi:"planId"`
	// Region of the registry
	Region string `pulumi:"region"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a CloudProjectContainerRegistry resource.
type CloudProjectContainerRegistryArgs struct {
	// Registry name
	Name pulumi.StringPtrInput
	// Plan ID of the registry
	PlanId pulumi.StringPtrInput
	// Region of the registry
	Region pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (CloudProjectContainerRegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudProjectContainerRegistryArgs)(nil)).Elem()
}

type CloudProjectContainerRegistryInput interface {
	pulumi.Input

	ToCloudProjectContainerRegistryOutput() CloudProjectContainerRegistryOutput
	ToCloudProjectContainerRegistryOutputWithContext(ctx context.Context) CloudProjectContainerRegistryOutput
}

func (*CloudProjectContainerRegistry) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectContainerRegistry)(nil)).Elem()
}

func (i *CloudProjectContainerRegistry) ToCloudProjectContainerRegistryOutput() CloudProjectContainerRegistryOutput {
	return i.ToCloudProjectContainerRegistryOutputWithContext(context.Background())
}

func (i *CloudProjectContainerRegistry) ToCloudProjectContainerRegistryOutputWithContext(ctx context.Context) CloudProjectContainerRegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerRegistryOutput)
}

// CloudProjectContainerRegistryArrayInput is an input type that accepts CloudProjectContainerRegistryArray and CloudProjectContainerRegistryArrayOutput values.
// You can construct a concrete instance of `CloudProjectContainerRegistryArrayInput` via:
//
//	CloudProjectContainerRegistryArray{ CloudProjectContainerRegistryArgs{...} }
type CloudProjectContainerRegistryArrayInput interface {
	pulumi.Input

	ToCloudProjectContainerRegistryArrayOutput() CloudProjectContainerRegistryArrayOutput
	ToCloudProjectContainerRegistryArrayOutputWithContext(context.Context) CloudProjectContainerRegistryArrayOutput
}

type CloudProjectContainerRegistryArray []CloudProjectContainerRegistryInput

func (CloudProjectContainerRegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectContainerRegistry)(nil)).Elem()
}

func (i CloudProjectContainerRegistryArray) ToCloudProjectContainerRegistryArrayOutput() CloudProjectContainerRegistryArrayOutput {
	return i.ToCloudProjectContainerRegistryArrayOutputWithContext(context.Background())
}

func (i CloudProjectContainerRegistryArray) ToCloudProjectContainerRegistryArrayOutputWithContext(ctx context.Context) CloudProjectContainerRegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerRegistryArrayOutput)
}

// CloudProjectContainerRegistryMapInput is an input type that accepts CloudProjectContainerRegistryMap and CloudProjectContainerRegistryMapOutput values.
// You can construct a concrete instance of `CloudProjectContainerRegistryMapInput` via:
//
//	CloudProjectContainerRegistryMap{ "key": CloudProjectContainerRegistryArgs{...} }
type CloudProjectContainerRegistryMapInput interface {
	pulumi.Input

	ToCloudProjectContainerRegistryMapOutput() CloudProjectContainerRegistryMapOutput
	ToCloudProjectContainerRegistryMapOutputWithContext(context.Context) CloudProjectContainerRegistryMapOutput
}

type CloudProjectContainerRegistryMap map[string]CloudProjectContainerRegistryInput

func (CloudProjectContainerRegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectContainerRegistry)(nil)).Elem()
}

func (i CloudProjectContainerRegistryMap) ToCloudProjectContainerRegistryMapOutput() CloudProjectContainerRegistryMapOutput {
	return i.ToCloudProjectContainerRegistryMapOutputWithContext(context.Background())
}

func (i CloudProjectContainerRegistryMap) ToCloudProjectContainerRegistryMapOutputWithContext(ctx context.Context) CloudProjectContainerRegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudProjectContainerRegistryMapOutput)
}

type CloudProjectContainerRegistryOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerRegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudProjectContainerRegistry)(nil)).Elem()
}

func (o CloudProjectContainerRegistryOutput) ToCloudProjectContainerRegistryOutput() CloudProjectContainerRegistryOutput {
	return o
}

func (o CloudProjectContainerRegistryOutput) ToCloudProjectContainerRegistryOutputWithContext(ctx context.Context) CloudProjectContainerRegistryOutput {
	return o
}

// Plan creation date
func (o CloudProjectContainerRegistryOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Registry name
func (o CloudProjectContainerRegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Plan ID of the registry
func (o CloudProjectContainerRegistryOutput) PlanId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.StringOutput { return v.PlanId }).(pulumi.StringOutput)
}

// Plan of the registry
func (o CloudProjectContainerRegistryOutput) Plans() CloudProjectContainerRegistryPlanArrayOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) CloudProjectContainerRegistryPlanArrayOutput { return v.Plans }).(CloudProjectContainerRegistryPlanArrayOutput)
}

// Project ID of your registry
func (o CloudProjectContainerRegistryOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Region of the registry
func (o CloudProjectContainerRegistryOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o CloudProjectContainerRegistryOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Current size of the registry (bytes)
func (o CloudProjectContainerRegistryOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

// Registry status
func (o CloudProjectContainerRegistryOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Registry last update date
func (o CloudProjectContainerRegistryOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Access url of the registry
func (o CloudProjectContainerRegistryOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// Version of your registry
func (o CloudProjectContainerRegistryOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudProjectContainerRegistry) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type CloudProjectContainerRegistryArrayOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerRegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudProjectContainerRegistry)(nil)).Elem()
}

func (o CloudProjectContainerRegistryArrayOutput) ToCloudProjectContainerRegistryArrayOutput() CloudProjectContainerRegistryArrayOutput {
	return o
}

func (o CloudProjectContainerRegistryArrayOutput) ToCloudProjectContainerRegistryArrayOutputWithContext(ctx context.Context) CloudProjectContainerRegistryArrayOutput {
	return o
}

func (o CloudProjectContainerRegistryArrayOutput) Index(i pulumi.IntInput) CloudProjectContainerRegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudProjectContainerRegistry {
		return vs[0].([]*CloudProjectContainerRegistry)[vs[1].(int)]
	}).(CloudProjectContainerRegistryOutput)
}

type CloudProjectContainerRegistryMapOutput struct{ *pulumi.OutputState }

func (CloudProjectContainerRegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudProjectContainerRegistry)(nil)).Elem()
}

func (o CloudProjectContainerRegistryMapOutput) ToCloudProjectContainerRegistryMapOutput() CloudProjectContainerRegistryMapOutput {
	return o
}

func (o CloudProjectContainerRegistryMapOutput) ToCloudProjectContainerRegistryMapOutputWithContext(ctx context.Context) CloudProjectContainerRegistryMapOutput {
	return o
}

func (o CloudProjectContainerRegistryMapOutput) MapIndex(k pulumi.StringInput) CloudProjectContainerRegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudProjectContainerRegistry {
		return vs[0].(map[string]*CloudProjectContainerRegistry)[vs[1].(string)]
	}).(CloudProjectContainerRegistryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectContainerRegistryInput)(nil)).Elem(), &CloudProjectContainerRegistry{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectContainerRegistryArrayInput)(nil)).Elem(), CloudProjectContainerRegistryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudProjectContainerRegistryMapInput)(nil)).Elem(), CloudProjectContainerRegistryMap{})
	pulumi.RegisterOutputType(CloudProjectContainerRegistryOutput{})
	pulumi.RegisterOutputType(CloudProjectContainerRegistryArrayOutput{})
	pulumi.RegisterOutputType(CloudProjectContainerRegistryMapOutput{})
}
