// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get a OVHcloud Managed Kubernetes Service cluster.
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
//			myKubeCluster, err := CloudProject.GetKube(ctx, &cloudproject.GetKubeArgs{
//				ServiceName: "XXXXXX",
//				KubeId:      "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("version", myKubeCluster.Version)
//			return nil
//		})
//	}
//
// ```
func LookupKube(ctx *pulumi.Context, args *LookupKubeArgs, opts ...pulumi.InvokeOption) (*LookupKubeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupKubeResult
	err := ctx.Invoke("ovh:CloudProject/getKube:getKube", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKube.
type LookupKubeArgs struct {
	// Kubernetes API server customization
	CustomizationApiservers []GetKubeCustomizationApiserver `pulumi:"customizationApiservers"`
	// Kubernetes kube-proxy customization
	CustomizationKubeProxy *GetKubeCustomizationKubeProxy `pulumi:"customizationKubeProxy"`
	// **Deprecated** (Optional) Use `customizationApiserver` and `customizationKubeProxy` instead. Kubernetes cluster customization
	//
	// Deprecated: Use customization_apiserver instead
	Customizations []GetKubeCustomization `pulumi:"customizations"`
	// The id of the managed kubernetes cluster.
	KubeId string `pulumi:"kubeId"`
	// Selected mode for kube-proxy.
	KubeProxyMode *string `pulumi:"kubeProxyMode"`
	// The name of the managed kubernetes cluster.
	Name *string `pulumi:"name"`
	// The OVHcloud public cloud region ID of the managed kubernetes cluster.
	Region *string `pulumi:"region"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
	// Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
	UpdatePolicy *string `pulumi:"updatePolicy"`
	// Kubernetes version of the managed kubernetes cluster.
	Version *string `pulumi:"version"`
}

// A collection of values returned by getKube.
type LookupKubeResult struct {
	// True if control-plane is up-to-date.
	ControlPlaneIsUpToDate bool `pulumi:"controlPlaneIsUpToDate"`
	// Kubernetes API server customization
	CustomizationApiservers []GetKubeCustomizationApiserver `pulumi:"customizationApiservers"`
	// Kubernetes kube-proxy customization
	CustomizationKubeProxy *GetKubeCustomizationKubeProxy `pulumi:"customizationKubeProxy"`
	// **Deprecated** (Optional) Use `customizationApiserver` and `customizationKubeProxy` instead. Kubernetes cluster customization
	//
	// Deprecated: Use customization_apiserver instead
	Customizations []GetKubeCustomization `pulumi:"customizations"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// True if all nodes and control-plane are up-to-date.
	IsUpToDate bool `pulumi:"isUpToDate"`
	// See Argument Reference above.
	KubeId string `pulumi:"kubeId"`
	// Selected mode for kube-proxy.
	KubeProxyMode *string `pulumi:"kubeProxyMode"`
	// The name of the managed kubernetes cluster.
	Name *string `pulumi:"name"`
	// Kubernetes versions available for upgrade.
	NextUpgradeVersions []string `pulumi:"nextUpgradeVersions"`
	// Cluster nodes URL.
	NodesUrl string `pulumi:"nodesUrl"`
	// OpenStack private network (or vrack) ID to use.
	PrivateNetworkId string `pulumi:"privateNetworkId"`
	// The OVHcloud public cloud region ID of the managed kubernetes cluster.
	Region *string `pulumi:"region"`
	// See Argument Reference above.
	ServiceName string `pulumi:"serviceName"`
	// Cluster status. Should be normally set to 'READY'.
	Status string `pulumi:"status"`
	// Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
	UpdatePolicy *string `pulumi:"updatePolicy"`
	// Management URL of your cluster.
	Url string `pulumi:"url"`
	// Kubernetes version of the managed kubernetes cluster.
	Version *string `pulumi:"version"`
}

func LookupKubeOutput(ctx *pulumi.Context, args LookupKubeOutputArgs, opts ...pulumi.InvokeOption) LookupKubeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKubeResult, error) {
			args := v.(LookupKubeArgs)
			r, err := LookupKube(ctx, &args, opts...)
			var s LookupKubeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKubeResultOutput)
}

// A collection of arguments for invoking getKube.
type LookupKubeOutputArgs struct {
	// Kubernetes API server customization
	CustomizationApiservers GetKubeCustomizationApiserverArrayInput `pulumi:"customizationApiservers"`
	// Kubernetes kube-proxy customization
	CustomizationKubeProxy GetKubeCustomizationKubeProxyPtrInput `pulumi:"customizationKubeProxy"`
	// **Deprecated** (Optional) Use `customizationApiserver` and `customizationKubeProxy` instead. Kubernetes cluster customization
	//
	// Deprecated: Use customization_apiserver instead
	Customizations GetKubeCustomizationArrayInput `pulumi:"customizations"`
	// The id of the managed kubernetes cluster.
	KubeId pulumi.StringInput `pulumi:"kubeId"`
	// Selected mode for kube-proxy.
	KubeProxyMode pulumi.StringPtrInput `pulumi:"kubeProxyMode"`
	// The name of the managed kubernetes cluster.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The OVHcloud public cloud region ID of the managed kubernetes cluster.
	Region pulumi.StringPtrInput `pulumi:"region"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
	UpdatePolicy pulumi.StringPtrInput `pulumi:"updatePolicy"`
	// Kubernetes version of the managed kubernetes cluster.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (LookupKubeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubeArgs)(nil)).Elem()
}

// A collection of values returned by getKube.
type LookupKubeResultOutput struct{ *pulumi.OutputState }

func (LookupKubeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKubeResult)(nil)).Elem()
}

func (o LookupKubeResultOutput) ToLookupKubeResultOutput() LookupKubeResultOutput {
	return o
}

func (o LookupKubeResultOutput) ToLookupKubeResultOutputWithContext(ctx context.Context) LookupKubeResultOutput {
	return o
}

// True if control-plane is up-to-date.
func (o LookupKubeResultOutput) ControlPlaneIsUpToDate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupKubeResult) bool { return v.ControlPlaneIsUpToDate }).(pulumi.BoolOutput)
}

// Kubernetes API server customization
func (o LookupKubeResultOutput) CustomizationApiservers() GetKubeCustomizationApiserverArrayOutput {
	return o.ApplyT(func(v LookupKubeResult) []GetKubeCustomizationApiserver { return v.CustomizationApiservers }).(GetKubeCustomizationApiserverArrayOutput)
}

// Kubernetes kube-proxy customization
func (o LookupKubeResultOutput) CustomizationKubeProxy() GetKubeCustomizationKubeProxyPtrOutput {
	return o.ApplyT(func(v LookupKubeResult) *GetKubeCustomizationKubeProxy { return v.CustomizationKubeProxy }).(GetKubeCustomizationKubeProxyPtrOutput)
}

// **Deprecated** (Optional) Use `customizationApiserver` and `customizationKubeProxy` instead. Kubernetes cluster customization
//
// Deprecated: Use customization_apiserver instead
func (o LookupKubeResultOutput) Customizations() GetKubeCustomizationArrayOutput {
	return o.ApplyT(func(v LookupKubeResult) []GetKubeCustomization { return v.Customizations }).(GetKubeCustomizationArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupKubeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeResult) string { return v.Id }).(pulumi.StringOutput)
}

// True if all nodes and control-plane are up-to-date.
func (o LookupKubeResultOutput) IsUpToDate() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupKubeResult) bool { return v.IsUpToDate }).(pulumi.BoolOutput)
}

// See Argument Reference above.
func (o LookupKubeResultOutput) KubeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeResult) string { return v.KubeId }).(pulumi.StringOutput)
}

// Selected mode for kube-proxy.
func (o LookupKubeResultOutput) KubeProxyMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeResult) *string { return v.KubeProxyMode }).(pulumi.StringPtrOutput)
}

// The name of the managed kubernetes cluster.
func (o LookupKubeResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Kubernetes versions available for upgrade.
func (o LookupKubeResultOutput) NextUpgradeVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKubeResult) []string { return v.NextUpgradeVersions }).(pulumi.StringArrayOutput)
}

// Cluster nodes URL.
func (o LookupKubeResultOutput) NodesUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeResult) string { return v.NodesUrl }).(pulumi.StringOutput)
}

// OpenStack private network (or vrack) ID to use.
func (o LookupKubeResultOutput) PrivateNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeResult) string { return v.PrivateNetworkId }).(pulumi.StringOutput)
}

// The OVHcloud public cloud region ID of the managed kubernetes cluster.
func (o LookupKubeResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

// See Argument Reference above.
func (o LookupKubeResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Cluster status. Should be normally set to 'READY'.
func (o LookupKubeResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeResult) string { return v.Status }).(pulumi.StringOutput)
}

// Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
func (o LookupKubeResultOutput) UpdatePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeResult) *string { return v.UpdatePolicy }).(pulumi.StringPtrOutput)
}

// Management URL of your cluster.
func (o LookupKubeResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKubeResult) string { return v.Url }).(pulumi.StringOutput)
}

// Kubernetes version of the managed kubernetes cluster.
func (o LookupKubeResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKubeResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKubeResultOutput{})
}
