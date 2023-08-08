// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve the list of the account's identity groups
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Dbaas.GetMeIdentityGroups(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetMeIdentityGroups(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetMeIdentityGroupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMeIdentityGroupsResult
	err := ctx.Invoke("ovh:Dbaas/getMeIdentityGroups:getMeIdentityGroups", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getMeIdentityGroups.
type GetMeIdentityGroupsResult struct {
	// The list of the group names of all the identity groups.
	Groups []string `pulumi:"groups"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
}