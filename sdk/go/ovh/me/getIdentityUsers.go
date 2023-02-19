// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve list of user logins of the account's identity users.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Me.GetIdentityUsers(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetIdentityUsers(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetIdentityUsersResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetIdentityUsersResult
	err := ctx.Invoke("ovh:Me/getIdentityUsers:getIdentityUsers", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getIdentityUsers.
type GetIdentityUsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of the user's logins of all the identity users.
	Users []string `pulumi:"users"`
}
