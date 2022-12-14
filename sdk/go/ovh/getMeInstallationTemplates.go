// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of custom installation templates available for dedicated servers.
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
//			_, err := ovh.GetMeInstallationTemplates(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMeInstallationTemplates(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetMeInstallationTemplatesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMeInstallationTemplatesResult
	err := ctx.Invoke("ovh:index/getMeInstallationTemplates:getMeInstallationTemplates", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getMeInstallationTemplates.
type GetMeInstallationTemplatesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of custom installation templates IDs available for dedicated servers.
	Results []string `pulumi:"results"`
}
