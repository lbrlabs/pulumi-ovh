// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information of order cart product products.
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
//			mycart, err := ovh.GetOrderCart(ctx, &GetOrderCartArgs{
//				OvhSubsidiary: "fr",
//				Description:   pulumi.StringRef("my cart"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = ovh.LookupOrderCartProduct(ctx, &GetOrderCartProductArgs{
//				CartId:  mycart.Id,
//				Product: "...",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOrderCartProduct(ctx *pulumi.Context, args *LookupOrderCartProductArgs, opts ...pulumi.InvokeOption) (*LookupOrderCartProductResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupOrderCartProductResult
	err := ctx.Invoke("ovh:index/getOrderCartProduct:getOrderCartProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrderCartProduct.
type LookupOrderCartProductArgs struct {
	// Cart identifier
	CartId string `pulumi:"cartId"`
	// product
	Product string `pulumi:"product"`
}

// A collection of values returned by getOrderCartProduct.
type LookupOrderCartProductResult struct {
	CartId string `pulumi:"cartId"`
	// The provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Product string `pulumi:"product"`
	// products results
	Results []GetOrderCartProductResult `pulumi:"results"`
}

func LookupOrderCartProductOutput(ctx *pulumi.Context, args LookupOrderCartProductOutputArgs, opts ...pulumi.InvokeOption) LookupOrderCartProductResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOrderCartProductResult, error) {
			args := v.(LookupOrderCartProductArgs)
			r, err := LookupOrderCartProduct(ctx, &args, opts...)
			var s LookupOrderCartProductResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOrderCartProductResultOutput)
}

// A collection of arguments for invoking getOrderCartProduct.
type LookupOrderCartProductOutputArgs struct {
	// Cart identifier
	CartId pulumi.StringInput `pulumi:"cartId"`
	// product
	Product pulumi.StringInput `pulumi:"product"`
}

func (LookupOrderCartProductOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderCartProductArgs)(nil)).Elem()
}

// A collection of values returned by getOrderCartProduct.
type LookupOrderCartProductResultOutput struct{ *pulumi.OutputState }

func (LookupOrderCartProductResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrderCartProductResult)(nil)).Elem()
}

func (o LookupOrderCartProductResultOutput) ToLookupOrderCartProductResultOutput() LookupOrderCartProductResultOutput {
	return o
}

func (o LookupOrderCartProductResultOutput) ToLookupOrderCartProductResultOutputWithContext(ctx context.Context) LookupOrderCartProductResultOutput {
	return o
}

func (o LookupOrderCartProductResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderCartProductResult) string { return v.CartId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrderCartProductResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderCartProductResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrderCartProductResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrderCartProductResult) string { return v.Product }).(pulumi.StringOutput)
}

// products results
func (o LookupOrderCartProductResultOutput) Results() GetOrderCartProductResultArrayOutput {
	return o.ApplyT(func(v LookupOrderCartProductResult) []GetOrderCartProductResult { return v.Results }).(GetOrderCartProductResultArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrderCartProductResultOutput{})
}
