// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hosting

import (
	"context"
	"reflect"

	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an hosting privatedatabase user grant.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-ovh/sdk/go/ovh/Hosting"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := Hosting.GetPrivateDatabaseUserGrant(ctx, &hosting.GetPrivateDatabaseUserGrantArgs{
//				DatabaseName: "XXXXXX",
//				ServiceName:  "XXXXXX",
//				UserName:     "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPrivateDatabaseUserGrant(ctx *pulumi.Context, args *LookupPrivateDatabaseUserGrantArgs, opts ...pulumi.InvokeOption) (*LookupPrivateDatabaseUserGrantResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateDatabaseUserGrantResult
	err := ctx.Invoke("ovh:Hosting/getPrivateDatabaseUserGrant:getPrivateDatabaseUserGrant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateDatabaseUserGrant.
type LookupPrivateDatabaseUserGrantArgs struct {
	// The database name on which grant the user
	DatabaseName string `pulumi:"databaseName"`
	// The internal name of your private database
	ServiceName string `pulumi:"serviceName"`
	// The user name
	UserName string `pulumi:"userName"`
}

// A collection of values returned by getPrivateDatabaseUserGrant.
type LookupPrivateDatabaseUserGrantResult struct {
	// Creation date of the database
	CreationDate string `pulumi:"creationDate"`
	DatabaseName string `pulumi:"databaseName"`
	// Grant name
	Grant string `pulumi:"grant"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	ServiceName string `pulumi:"serviceName"`
	UserName    string `pulumi:"userName"`
}

func LookupPrivateDatabaseUserGrantOutput(ctx *pulumi.Context, args LookupPrivateDatabaseUserGrantOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateDatabaseUserGrantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateDatabaseUserGrantResult, error) {
			args := v.(LookupPrivateDatabaseUserGrantArgs)
			r, err := LookupPrivateDatabaseUserGrant(ctx, &args, opts...)
			var s LookupPrivateDatabaseUserGrantResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateDatabaseUserGrantResultOutput)
}

// A collection of arguments for invoking getPrivateDatabaseUserGrant.
type LookupPrivateDatabaseUserGrantOutputArgs struct {
	// The database name on which grant the user
	DatabaseName pulumi.StringInput `pulumi:"databaseName"`
	// The internal name of your private database
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// The user name
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (LookupPrivateDatabaseUserGrantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDatabaseUserGrantArgs)(nil)).Elem()
}

// A collection of values returned by getPrivateDatabaseUserGrant.
type LookupPrivateDatabaseUserGrantResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateDatabaseUserGrantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDatabaseUserGrantResult)(nil)).Elem()
}

func (o LookupPrivateDatabaseUserGrantResultOutput) ToLookupPrivateDatabaseUserGrantResultOutput() LookupPrivateDatabaseUserGrantResultOutput {
	return o
}

func (o LookupPrivateDatabaseUserGrantResultOutput) ToLookupPrivateDatabaseUserGrantResultOutputWithContext(ctx context.Context) LookupPrivateDatabaseUserGrantResultOutput {
	return o
}

// Creation date of the database
func (o LookupPrivateDatabaseUserGrantResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserGrantResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

func (o LookupPrivateDatabaseUserGrantResultOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserGrantResult) string { return v.DatabaseName }).(pulumi.StringOutput)
}

// Grant name
func (o LookupPrivateDatabaseUserGrantResultOutput) Grant() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserGrantResult) string { return v.Grant }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrivateDatabaseUserGrantResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserGrantResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateDatabaseUserGrantResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserGrantResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupPrivateDatabaseUserGrantResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserGrantResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateDatabaseUserGrantResultOutput{})
}
