// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package hosting

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an hosting privatedatabase user.
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
//			_, err := Hosting.GetPrivateDatabaseUser(ctx, &hosting.GetPrivateDatabaseUserArgs{
//				ServiceName: "XXXXXX",
//				UserName:    "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupPrivateDatabaseUser(ctx *pulumi.Context, args *LookupPrivateDatabaseUserArgs, opts ...pulumi.InvokeOption) (*LookupPrivateDatabaseUserResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupPrivateDatabaseUserResult
	err := ctx.Invoke("ovh:Hosting/getPrivateDatabaseUser:getPrivateDatabaseUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateDatabaseUser.
type LookupPrivateDatabaseUserArgs struct {
	// The internal name of your private database
	ServiceName string `pulumi:"serviceName"`
	// User name
	UserName string `pulumi:"userName"`
}

// A collection of values returned by getPrivateDatabaseUser.
type LookupPrivateDatabaseUserResult struct {
	// Creation date of the database
	CreationDate string `pulumi:"creationDate"`
	// Users granted to this database
	Databases []GetPrivateDatabaseUserDatabase `pulumi:"databases"`
	// The provider-assigned unique ID for this managed resource.
	Id          string `pulumi:"id"`
	ServiceName string `pulumi:"serviceName"`
	UserName    string `pulumi:"userName"`
}

func LookupPrivateDatabaseUserOutput(ctx *pulumi.Context, args LookupPrivateDatabaseUserOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateDatabaseUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateDatabaseUserResult, error) {
			args := v.(LookupPrivateDatabaseUserArgs)
			r, err := LookupPrivateDatabaseUser(ctx, &args, opts...)
			var s LookupPrivateDatabaseUserResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateDatabaseUserResultOutput)
}

// A collection of arguments for invoking getPrivateDatabaseUser.
type LookupPrivateDatabaseUserOutputArgs struct {
	// The internal name of your private database
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// User name
	UserName pulumi.StringInput `pulumi:"userName"`
}

func (LookupPrivateDatabaseUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDatabaseUserArgs)(nil)).Elem()
}

// A collection of values returned by getPrivateDatabaseUser.
type LookupPrivateDatabaseUserResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateDatabaseUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateDatabaseUserResult)(nil)).Elem()
}

func (o LookupPrivateDatabaseUserResultOutput) ToLookupPrivateDatabaseUserResultOutput() LookupPrivateDatabaseUserResultOutput {
	return o
}

func (o LookupPrivateDatabaseUserResultOutput) ToLookupPrivateDatabaseUserResultOutputWithContext(ctx context.Context) LookupPrivateDatabaseUserResultOutput {
	return o
}

// Creation date of the database
func (o LookupPrivateDatabaseUserResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Users granted to this database
func (o LookupPrivateDatabaseUserResultOutput) Databases() GetPrivateDatabaseUserDatabaseArrayOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserResult) []GetPrivateDatabaseUserDatabase { return v.Databases }).(GetPrivateDatabaseUserDatabaseArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPrivateDatabaseUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrivateDatabaseUserResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o LookupPrivateDatabaseUserResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateDatabaseUserResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateDatabaseUserResultOutput{})
}
