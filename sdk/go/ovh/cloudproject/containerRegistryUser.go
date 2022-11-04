// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a user for a container registry associated with a public cloud project.
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
//			_, err = CloudProject.GetContainerRegistry(ctx, &cloudproject.GetContainerRegistryArgs{
//				ServiceName: "XXXXXX",
//				RegistryId:  "yyyy",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = CloudProject.NewContainerRegistryUser(ctx, "user", &CloudProject.ContainerRegistryUserArgs{
//				ServiceName: pulumi.Any(ovh_cloud_project_containerregistry.Registry.Service_name),
//				RegistryId:  pulumi.Any(ovh_cloud_project_containerregistry.Registry.Id),
//				Email:       pulumi.String("foo@bar.com"),
//				Login:       pulumi.String("foobar"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type ContainerRegistryUser struct {
	pulumi.CustomResourceState

	// User email
	Email pulumi.StringOutput `pulumi:"email"`
	// Registry name
	Login pulumi.StringOutput `pulumi:"login"`
	// (Sensitive) User password
	Password pulumi.StringOutput `pulumi:"password"`
	// Registry ID
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// User name
	User pulumi.StringOutput `pulumi:"user"`
}

// NewContainerRegistryUser registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistryUser(ctx *pulumi.Context,
	name string, args *ContainerRegistryUserArgs, opts ...pulumi.ResourceOption) (*ContainerRegistryUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Login == nil {
		return nil, errors.New("invalid value for required argument 'Login'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource ContainerRegistryUser
	err := ctx.RegisterResource("ovh:CloudProject/containerRegistryUser:ContainerRegistryUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistryUser gets an existing ContainerRegistryUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistryUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryUserState, opts ...pulumi.ResourceOption) (*ContainerRegistryUser, error) {
	var resource ContainerRegistryUser
	err := ctx.ReadResource("ovh:CloudProject/containerRegistryUser:ContainerRegistryUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistryUser resources.
type containerRegistryUserState struct {
	// User email
	Email *string `pulumi:"email"`
	// Registry name
	Login *string `pulumi:"login"`
	// (Sensitive) User password
	Password *string `pulumi:"password"`
	// Registry ID
	RegistryId *string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// User name
	User *string `pulumi:"user"`
}

type ContainerRegistryUserState struct {
	// User email
	Email pulumi.StringPtrInput
	// Registry name
	Login pulumi.StringPtrInput
	// (Sensitive) User password
	Password pulumi.StringPtrInput
	// Registry ID
	RegistryId pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// User name
	User pulumi.StringPtrInput
}

func (ContainerRegistryUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryUserState)(nil)).Elem()
}

type containerRegistryUserArgs struct {
	// User email
	Email string `pulumi:"email"`
	// Registry name
	Login string `pulumi:"login"`
	// Registry ID
	RegistryId string `pulumi:"registryId"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ContainerRegistryUser resource.
type ContainerRegistryUserArgs struct {
	// User email
	Email pulumi.StringInput
	// Registry name
	Login pulumi.StringInput
	// Registry ID
	RegistryId pulumi.StringInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (ContainerRegistryUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryUserArgs)(nil)).Elem()
}

type ContainerRegistryUserInput interface {
	pulumi.Input

	ToContainerRegistryUserOutput() ContainerRegistryUserOutput
	ToContainerRegistryUserOutputWithContext(ctx context.Context) ContainerRegistryUserOutput
}

func (*ContainerRegistryUser) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryUser)(nil)).Elem()
}

func (i *ContainerRegistryUser) ToContainerRegistryUserOutput() ContainerRegistryUserOutput {
	return i.ToContainerRegistryUserOutputWithContext(context.Background())
}

func (i *ContainerRegistryUser) ToContainerRegistryUserOutputWithContext(ctx context.Context) ContainerRegistryUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryUserOutput)
}

// ContainerRegistryUserArrayInput is an input type that accepts ContainerRegistryUserArray and ContainerRegistryUserArrayOutput values.
// You can construct a concrete instance of `ContainerRegistryUserArrayInput` via:
//
//	ContainerRegistryUserArray{ ContainerRegistryUserArgs{...} }
type ContainerRegistryUserArrayInput interface {
	pulumi.Input

	ToContainerRegistryUserArrayOutput() ContainerRegistryUserArrayOutput
	ToContainerRegistryUserArrayOutputWithContext(context.Context) ContainerRegistryUserArrayOutput
}

type ContainerRegistryUserArray []ContainerRegistryUserInput

func (ContainerRegistryUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistryUser)(nil)).Elem()
}

func (i ContainerRegistryUserArray) ToContainerRegistryUserArrayOutput() ContainerRegistryUserArrayOutput {
	return i.ToContainerRegistryUserArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryUserArray) ToContainerRegistryUserArrayOutputWithContext(ctx context.Context) ContainerRegistryUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryUserArrayOutput)
}

// ContainerRegistryUserMapInput is an input type that accepts ContainerRegistryUserMap and ContainerRegistryUserMapOutput values.
// You can construct a concrete instance of `ContainerRegistryUserMapInput` via:
//
//	ContainerRegistryUserMap{ "key": ContainerRegistryUserArgs{...} }
type ContainerRegistryUserMapInput interface {
	pulumi.Input

	ToContainerRegistryUserMapOutput() ContainerRegistryUserMapOutput
	ToContainerRegistryUserMapOutputWithContext(context.Context) ContainerRegistryUserMapOutput
}

type ContainerRegistryUserMap map[string]ContainerRegistryUserInput

func (ContainerRegistryUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistryUser)(nil)).Elem()
}

func (i ContainerRegistryUserMap) ToContainerRegistryUserMapOutput() ContainerRegistryUserMapOutput {
	return i.ToContainerRegistryUserMapOutputWithContext(context.Background())
}

func (i ContainerRegistryUserMap) ToContainerRegistryUserMapOutputWithContext(ctx context.Context) ContainerRegistryUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryUserMapOutput)
}

type ContainerRegistryUserOutput struct{ *pulumi.OutputState }

func (ContainerRegistryUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryUser)(nil)).Elem()
}

func (o ContainerRegistryUserOutput) ToContainerRegistryUserOutput() ContainerRegistryUserOutput {
	return o
}

func (o ContainerRegistryUserOutput) ToContainerRegistryUserOutputWithContext(ctx context.Context) ContainerRegistryUserOutput {
	return o
}

// User email
func (o ContainerRegistryUserOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryUser) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// Registry name
func (o ContainerRegistryUserOutput) Login() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryUser) pulumi.StringOutput { return v.Login }).(pulumi.StringOutput)
}

// (Sensitive) User password
func (o ContainerRegistryUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Registry ID
func (o ContainerRegistryUserOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryUser) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o ContainerRegistryUserOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryUser) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// User name
func (o ContainerRegistryUserOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryUser) pulumi.StringOutput { return v.User }).(pulumi.StringOutput)
}

type ContainerRegistryUserArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistryUser)(nil)).Elem()
}

func (o ContainerRegistryUserArrayOutput) ToContainerRegistryUserArrayOutput() ContainerRegistryUserArrayOutput {
	return o
}

func (o ContainerRegistryUserArrayOutput) ToContainerRegistryUserArrayOutputWithContext(ctx context.Context) ContainerRegistryUserArrayOutput {
	return o
}

func (o ContainerRegistryUserArrayOutput) Index(i pulumi.IntInput) ContainerRegistryUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerRegistryUser {
		return vs[0].([]*ContainerRegistryUser)[vs[1].(int)]
	}).(ContainerRegistryUserOutput)
}

type ContainerRegistryUserMapOutput struct{ *pulumi.OutputState }

func (ContainerRegistryUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistryUser)(nil)).Elem()
}

func (o ContainerRegistryUserMapOutput) ToContainerRegistryUserMapOutput() ContainerRegistryUserMapOutput {
	return o
}

func (o ContainerRegistryUserMapOutput) ToContainerRegistryUserMapOutputWithContext(ctx context.Context) ContainerRegistryUserMapOutput {
	return o
}

func (o ContainerRegistryUserMapOutput) MapIndex(k pulumi.StringInput) ContainerRegistryUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerRegistryUser {
		return vs[0].(map[string]*ContainerRegistryUser)[vs[1].(string)]
	}).(ContainerRegistryUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryUserInput)(nil)).Elem(), &ContainerRegistryUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryUserArrayInput)(nil)).Elem(), ContainerRegistryUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryUserMapInput)(nil)).Elem(), ContainerRegistryUserMap{})
	pulumi.RegisterOutputType(ContainerRegistryUserOutput{})
	pulumi.RegisterOutputType(ContainerRegistryUserArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryUserMapOutput{})
}