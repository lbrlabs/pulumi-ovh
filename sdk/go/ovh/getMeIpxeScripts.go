// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve a list of the names of the account's IPXE Scripts.
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
//			_, err := ovh.GetMeIpxeScripts(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMeIpxeScripts(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetMeIpxeScriptsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetMeIpxeScriptsResult
	err := ctx.Invoke("ovh:index/getMeIpxeScripts:getMeIpxeScripts", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getMeIpxeScripts.
type GetMeIpxeScriptsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of the names of all the IPXE Scripts.
	Results []string `pulumi:"results"`
}
