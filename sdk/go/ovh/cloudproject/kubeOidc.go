// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an OIDC configuration in an OVHcloud Managed Kubernetes cluster.
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
//			_, err := CloudProject.NewKubeOidc(ctx, "my-oidc", &CloudProject.KubeOidcArgs{
//				ServiceName:        pulumi.Any(_var.Projectid),
//				KubeId:             pulumi.Any(ovh_cloud_project_kube.Mykube.Id),
//				ClientId:           pulumi.String("xxx"),
//				IssuerUrl:          pulumi.String("https://ovh.com"),
//				OidcUsernameClaim:  pulumi.String("an-email"),
//				OidcUsernamePrefix: pulumi.String("ovh:"),
//				OidcGroupsClaims: pulumi.StringArray{
//					pulumi.String("groups"),
//				},
//				OidcGroupsPrefix: pulumi.String("ovh:"),
//				OidcRequiredClaims: pulumi.StringArray{
//					pulumi.String("claim1=val1"),
//				},
//				OidcSigningAlgs: pulumi.StringArray{
//					pulumi.String("RS512"),
//				},
//				OidcCaContent: pulumi.String("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUZhekNDQTFPZ0F3SUJBZ0lVYm9YRkZrL1hCQmdQUUI4UHlqbkttUGVWekNjd0RRWUpLb1pJaHZjTkFRRUwKQlFBd1JURUxNQWtHQTFVRUJoTUNRVlV4RXpBUkJnTlZCQWdNQ2xOdmJXVXRVM1JoZEdVeElUQWZCZ05WQkFvTQpHRWx1ZEdWeWJtVjBJRmRwWkdkcGRITWdVSFI1SUV4MFpEQWVGdzB5TWpFd01UUXdOalE0TlROYUZ3MHlNekV3Ck1UUXdOalE0TlROYU1FVXhDekFKQmdOVkJBWVRBa0ZWTVJNd0VRWURWUVFJREFwVGIyMWxMVk4wWVhSbE1TRXcKSHdZRFZRUUtEQmhKYm5SbGNtNWxkQ0JYYVdSbmFYUnpJRkIwZVNCTWRHUXdnZ0lpTUEwR0NTcUdTSWIzRFFFQgpBUVVBQTRJQ0R3QXdnZ0lLQW9JQ0FRQytPMk53bGx2QTQyT05SUHMyZWlqTUp2UHhpN21RblVSS3FrOHJEV1VkCkwzZU0yM1JXeVhtS1AydDQ5Zi9LVGsweEZNVStOSTUzTEhwWmh6N3NpK3dEUFUvWWZWSS9rQmZsRm8zeVZCMSsKZWdCSnpyNGIrQ3FoaWlCUkh0Vm5LblFKUmdvOVJjVkxhRm82UEY0N1V0UWJ2bWVuNGdERnExVkYwVHhUdnFMdwpIMzRZL0U2QUJsSlZnWFBzaWQzNm54eTErNnlKV05vRXNVekFiekpWMHhzTGhxc2hOazA0TWx4YnBhcG1XcEUxCmFFMHRIZGpjUlI3Y1dTRUUwMnRSQzNYL2tSNjBKb3MxR0N0Y0ZQTTVIN3NjOFBXNFRUem1EWWhOeDRiVjV4T28KU0xYRnI5ajBzZEgxbm1wSlI1dWxJT2dPTWV3MHA2d3JOYVV2MGpxc1hzdVdqMVpxdTRLRi81aEQ3azVhRlhKNQpjYWNTUi9mRWxreW1uZis0eHZFOG8wdkRWNFR5NHo3K3lSS1U0clZvZFNBZWZIN3lqeitLV1RRck96L0lHU2NwCmV1YTdqV0hRMDdMYWxyTjV2b0tFaU1JM3MrWjhzeUdVUGVyYXQwdzJMWlc3NnhxVGl4R002clZxUldxVlQ4L1oKQTJMMEc4WGRvNTZvV2lFYVF5RkJtRDFnMXU2UEsvTmFGVDI1L2tTNWJ1dnF5L1dLVGt0UVNhNHNZc1ZLbUlQTQp0Zys0NUZ2aFErNkRuQzd0TmVnaTZDTkdTb0w0R1dPOEE5UDZRNjE5RkJJZ1VjcGpFMTgvUHpQOEJmcTAxajhnCjZmdm1jNkVPMkxHVHhDcW1DbVp0TnI3OCtQaUxkMHZIY3pqY3E3NzhiNW5WRXRpUVNRQkUyb0ozTVlIZUFIUUkKYVFJREFRQUJvMU13VVRBZEJnTlZIUTRFRmdRVUpaMUhlVmx1U3pjY0U2NEZQYWtuNkRBWnhmSXdId1lEVlIwagpCQmd3Rm9BVUpaMUhlVmx1U3pjY0U2NEZQYWtuNkRBWnhmSXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QU5CZ2txCmhraUc5dzBCQVFzRkFBT0NBZ0VBQlhNSlU2MjJZVFZVNnZ1K2svNnkwMGNaWlRmVnZtdVJMOXhTcWxVM0I1QmQKVWdyVWx1TmdjN2dhUUlrYzkvWmh2MnhNd0xxUldMWEhiTWx1NkNvdkNiVTVpeWt0NHVWMnl5UzlZYWhmVVRNVQo3TVE0WFRta2hoS0dGbWZBQ2QzTUVwRE55T3hmWXh0UVBwM1NZT2IxRGFKMmUwY01Gc081bytORGQ5aFVBVzFoCjFLMjMwQnZzYldYYVo4MStIdTU4U1BsYTM5R3FMTG85MzR6dEs4WkRWNFRGTVJxMnNVQ1cxcWFidDh5ejd2RzAKSGV3dXdxelRwR1lTSFI1U0ZvMm45R0xKVUN4SnhxcDlOWVJjMlhUdXRUdkJESzVPMXFZZEJaQzd6cmcxSnczawp2SjI4UGx2TzBQRE42ZVlUdElJdC9yU05ZbW56eVVNRTRYREt0di9KRitLZWZNSWxDTkpzZDRHYXVTdlo5M1NOClhINmcrNEZvRkp4UzNxRmZ0WEc4czNRNnppNzNLRzh5UHZVNHU0WmZNRGd2aG92L0V5YkNLWUpFdVVZSlJWNGEKbmc3cWh3NDBabXQ0eWNCRzU5a2tFSGhNYWtxTWpPaUNkV2x4MEVjZXIxcEFGT1pqN3o1NktURXIxa0ZwUHVaRApjVER5SnNwTjh6dm9CQ0l1ancvQjR6S3kyWStOQitRR1p3dXhyTk9mRGR6ek9yQUE1Ym9OS2gwUUh4c0RxNTExClFaU3hCR21EcGJzN2QzMUQvQll3WEhIUWdwb3FoVUU5dFBGSThpN0pkM2FyeXZCdHlnTWlxSmt1VlRFVk1Ta0UKNTZ0VnFsMjlXenFhRXNrbDN3VUlmczVKKzN3RzRPcWNxRDdXaGQxWUtnc0VUMjdFTWlqVXZIYzQ4TXE0bU1rPQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="),
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
// OVHcloud Managed Kubernetes Service cluster OIDC can be imported using the tenant `service_name` and cluster id `kube_id` separated by "/" E.g., bash
//
// ```sh
//
//	$ pulumi import ovh:CloudProject/kubeOidc:KubeOidc my-oidc service_name/kube_id
//
// ```
type KubeOidc struct {
	pulumi.CustomResourceState

	// The OIDC client ID.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// The OIDC issuer url.
	IssuerUrl pulumi.StringOutput `pulumi:"issuerUrl"`
	// The ID of the managed kubernetes cluster. **Changing this value recreates the resource.**
	KubeId             pulumi.StringOutput      `pulumi:"kubeId"`
	OidcCaContent      pulumi.StringPtrOutput   `pulumi:"oidcCaContent"`
	OidcGroupsClaims   pulumi.StringArrayOutput `pulumi:"oidcGroupsClaims"`
	OidcGroupsPrefix   pulumi.StringPtrOutput   `pulumi:"oidcGroupsPrefix"`
	OidcRequiredClaims pulumi.StringArrayOutput `pulumi:"oidcRequiredClaims"`
	OidcSigningAlgs    pulumi.StringArrayOutput `pulumi:"oidcSigningAlgs"`
	OidcUsernameClaim  pulumi.StringPtrOutput   `pulumi:"oidcUsernameClaim"`
	OidcUsernamePrefix pulumi.StringPtrOutput   `pulumi:"oidcUsernamePrefix"`
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewKubeOidc registers a new resource with the given unique name, arguments, and options.
func NewKubeOidc(ctx *pulumi.Context,
	name string, args *KubeOidcArgs, opts ...pulumi.ResourceOption) (*KubeOidc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientId == nil {
		return nil, errors.New("invalid value for required argument 'ClientId'")
	}
	if args.IssuerUrl == nil {
		return nil, errors.New("invalid value for required argument 'IssuerUrl'")
	}
	if args.KubeId == nil {
		return nil, errors.New("invalid value for required argument 'KubeId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource KubeOidc
	err := ctx.RegisterResource("ovh:CloudProject/kubeOidc:KubeOidc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubeOidc gets an existing KubeOidc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubeOidc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubeOidcState, opts ...pulumi.ResourceOption) (*KubeOidc, error) {
	var resource KubeOidc
	err := ctx.ReadResource("ovh:CloudProject/kubeOidc:KubeOidc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KubeOidc resources.
type kubeOidcState struct {
	// The OIDC client ID.
	ClientId *string `pulumi:"clientId"`
	// The OIDC issuer url.
	IssuerUrl *string `pulumi:"issuerUrl"`
	// The ID of the managed kubernetes cluster. **Changing this value recreates the resource.**
	KubeId             *string  `pulumi:"kubeId"`
	OidcCaContent      *string  `pulumi:"oidcCaContent"`
	OidcGroupsClaims   []string `pulumi:"oidcGroupsClaims"`
	OidcGroupsPrefix   *string  `pulumi:"oidcGroupsPrefix"`
	OidcRequiredClaims []string `pulumi:"oidcRequiredClaims"`
	OidcSigningAlgs    []string `pulumi:"oidcSigningAlgs"`
	OidcUsernameClaim  *string  `pulumi:"oidcUsernameClaim"`
	OidcUsernamePrefix *string  `pulumi:"oidcUsernamePrefix"`
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName *string `pulumi:"serviceName"`
}

type KubeOidcState struct {
	// The OIDC client ID.
	ClientId pulumi.StringPtrInput
	// The OIDC issuer url.
	IssuerUrl pulumi.StringPtrInput
	// The ID of the managed kubernetes cluster. **Changing this value recreates the resource.**
	KubeId             pulumi.StringPtrInput
	OidcCaContent      pulumi.StringPtrInput
	OidcGroupsClaims   pulumi.StringArrayInput
	OidcGroupsPrefix   pulumi.StringPtrInput
	OidcRequiredClaims pulumi.StringArrayInput
	OidcSigningAlgs    pulumi.StringArrayInput
	OidcUsernameClaim  pulumi.StringPtrInput
	OidcUsernamePrefix pulumi.StringPtrInput
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringPtrInput
}

func (KubeOidcState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeOidcState)(nil)).Elem()
}

type kubeOidcArgs struct {
	// The OIDC client ID.
	ClientId string `pulumi:"clientId"`
	// The OIDC issuer url.
	IssuerUrl string `pulumi:"issuerUrl"`
	// The ID of the managed kubernetes cluster. **Changing this value recreates the resource.**
	KubeId             string   `pulumi:"kubeId"`
	OidcCaContent      *string  `pulumi:"oidcCaContent"`
	OidcGroupsClaims   []string `pulumi:"oidcGroupsClaims"`
	OidcGroupsPrefix   *string  `pulumi:"oidcGroupsPrefix"`
	OidcRequiredClaims []string `pulumi:"oidcRequiredClaims"`
	OidcSigningAlgs    []string `pulumi:"oidcSigningAlgs"`
	OidcUsernameClaim  *string  `pulumi:"oidcUsernameClaim"`
	OidcUsernamePrefix *string  `pulumi:"oidcUsernamePrefix"`
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a KubeOidc resource.
type KubeOidcArgs struct {
	// The OIDC client ID.
	ClientId pulumi.StringInput
	// The OIDC issuer url.
	IssuerUrl pulumi.StringInput
	// The ID of the managed kubernetes cluster. **Changing this value recreates the resource.**
	KubeId             pulumi.StringInput
	OidcCaContent      pulumi.StringPtrInput
	OidcGroupsClaims   pulumi.StringArrayInput
	OidcGroupsPrefix   pulumi.StringPtrInput
	OidcRequiredClaims pulumi.StringArrayInput
	OidcSigningAlgs    pulumi.StringArrayInput
	OidcUsernameClaim  pulumi.StringPtrInput
	OidcUsernamePrefix pulumi.StringPtrInput
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringInput
}

func (KubeOidcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeOidcArgs)(nil)).Elem()
}

type KubeOidcInput interface {
	pulumi.Input

	ToKubeOidcOutput() KubeOidcOutput
	ToKubeOidcOutputWithContext(ctx context.Context) KubeOidcOutput
}

func (*KubeOidc) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeOidc)(nil)).Elem()
}

func (i *KubeOidc) ToKubeOidcOutput() KubeOidcOutput {
	return i.ToKubeOidcOutputWithContext(context.Background())
}

func (i *KubeOidc) ToKubeOidcOutputWithContext(ctx context.Context) KubeOidcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeOidcOutput)
}

// KubeOidcArrayInput is an input type that accepts KubeOidcArray and KubeOidcArrayOutput values.
// You can construct a concrete instance of `KubeOidcArrayInput` via:
//
//	KubeOidcArray{ KubeOidcArgs{...} }
type KubeOidcArrayInput interface {
	pulumi.Input

	ToKubeOidcArrayOutput() KubeOidcArrayOutput
	ToKubeOidcArrayOutputWithContext(context.Context) KubeOidcArrayOutput
}

type KubeOidcArray []KubeOidcInput

func (KubeOidcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeOidc)(nil)).Elem()
}

func (i KubeOidcArray) ToKubeOidcArrayOutput() KubeOidcArrayOutput {
	return i.ToKubeOidcArrayOutputWithContext(context.Background())
}

func (i KubeOidcArray) ToKubeOidcArrayOutputWithContext(ctx context.Context) KubeOidcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeOidcArrayOutput)
}

// KubeOidcMapInput is an input type that accepts KubeOidcMap and KubeOidcMapOutput values.
// You can construct a concrete instance of `KubeOidcMapInput` via:
//
//	KubeOidcMap{ "key": KubeOidcArgs{...} }
type KubeOidcMapInput interface {
	pulumi.Input

	ToKubeOidcMapOutput() KubeOidcMapOutput
	ToKubeOidcMapOutputWithContext(context.Context) KubeOidcMapOutput
}

type KubeOidcMap map[string]KubeOidcInput

func (KubeOidcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeOidc)(nil)).Elem()
}

func (i KubeOidcMap) ToKubeOidcMapOutput() KubeOidcMapOutput {
	return i.ToKubeOidcMapOutputWithContext(context.Background())
}

func (i KubeOidcMap) ToKubeOidcMapOutputWithContext(ctx context.Context) KubeOidcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeOidcMapOutput)
}

type KubeOidcOutput struct{ *pulumi.OutputState }

func (KubeOidcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeOidc)(nil)).Elem()
}

func (o KubeOidcOutput) ToKubeOidcOutput() KubeOidcOutput {
	return o
}

func (o KubeOidcOutput) ToKubeOidcOutputWithContext(ctx context.Context) KubeOidcOutput {
	return o
}

// The OIDC client ID.
func (o KubeOidcOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// The OIDC issuer url.
func (o KubeOidcOutput) IssuerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringOutput { return v.IssuerUrl }).(pulumi.StringOutput)
}

// The ID of the managed kubernetes cluster. **Changing this value recreates the resource.**
func (o KubeOidcOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringOutput { return v.KubeId }).(pulumi.StringOutput)
}

func (o KubeOidcOutput) OidcCaContent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringPtrOutput { return v.OidcCaContent }).(pulumi.StringPtrOutput)
}

func (o KubeOidcOutput) OidcGroupsClaims() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringArrayOutput { return v.OidcGroupsClaims }).(pulumi.StringArrayOutput)
}

func (o KubeOidcOutput) OidcGroupsPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringPtrOutput { return v.OidcGroupsPrefix }).(pulumi.StringPtrOutput)
}

func (o KubeOidcOutput) OidcRequiredClaims() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringArrayOutput { return v.OidcRequiredClaims }).(pulumi.StringArrayOutput)
}

func (o KubeOidcOutput) OidcSigningAlgs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringArrayOutput { return v.OidcSigningAlgs }).(pulumi.StringArrayOutput)
}

func (o KubeOidcOutput) OidcUsernameClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringPtrOutput { return v.OidcUsernameClaim }).(pulumi.StringPtrOutput)
}

func (o KubeOidcOutput) OidcUsernamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringPtrOutput { return v.OidcUsernamePrefix }).(pulumi.StringPtrOutput)
}

// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
func (o KubeOidcOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *KubeOidc) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type KubeOidcArrayOutput struct{ *pulumi.OutputState }

func (KubeOidcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KubeOidc)(nil)).Elem()
}

func (o KubeOidcArrayOutput) ToKubeOidcArrayOutput() KubeOidcArrayOutput {
	return o
}

func (o KubeOidcArrayOutput) ToKubeOidcArrayOutputWithContext(ctx context.Context) KubeOidcArrayOutput {
	return o
}

func (o KubeOidcArrayOutput) Index(i pulumi.IntInput) KubeOidcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KubeOidc {
		return vs[0].([]*KubeOidc)[vs[1].(int)]
	}).(KubeOidcOutput)
}

type KubeOidcMapOutput struct{ *pulumi.OutputState }

func (KubeOidcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KubeOidc)(nil)).Elem()
}

func (o KubeOidcMapOutput) ToKubeOidcMapOutput() KubeOidcMapOutput {
	return o
}

func (o KubeOidcMapOutput) ToKubeOidcMapOutputWithContext(ctx context.Context) KubeOidcMapOutput {
	return o
}

func (o KubeOidcMapOutput) MapIndex(k pulumi.StringInput) KubeOidcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KubeOidc {
		return vs[0].(map[string]*KubeOidc)[vs[1].(string)]
	}).(KubeOidcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeOidcInput)(nil)).Elem(), &KubeOidc{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeOidcArrayInput)(nil)).Elem(), KubeOidcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeOidcMapInput)(nil)).Elem(), KubeOidcMap{})
	pulumi.RegisterOutputType(KubeOidcOutput{})
	pulumi.RegisterOutputType(KubeOidcArrayOutput{})
	pulumi.RegisterOutputType(KubeOidcMapOutput{})
}
