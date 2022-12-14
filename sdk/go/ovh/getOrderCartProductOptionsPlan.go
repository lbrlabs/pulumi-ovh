// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information of order cart product options plan.
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
//			_, err = ovh.GetOrderCartProductOptionsPlan(ctx, &GetOrderCartProductOptionsPlanArgs{
//				CartId:          mycart.Id,
//				PriceCapacity:   "renew",
//				Product:         "cloud",
//				PlanCode:        "project",
//				OptionsPlanCode: "vrack",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetOrderCartProductOptionsPlan(ctx *pulumi.Context, args *GetOrderCartProductOptionsPlanArgs, opts ...pulumi.InvokeOption) (*GetOrderCartProductOptionsPlanResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetOrderCartProductOptionsPlanResult
	err := ctx.Invoke("ovh:index/getOrderCartProductOptionsPlan:getOrderCartProductOptionsPlan", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrderCartProductOptionsPlan.
type GetOrderCartProductOptionsPlanArgs struct {
	// Cart identifier
	CartId string `pulumi:"cartId"`
	// Catalog name
	CatalogName *string `pulumi:"catalogName"`
	// options plan code.
	OptionsPlanCode string `pulumi:"optionsPlanCode"`
	// Product offer identifier
	PlanCode string `pulumi:"planCode"`
	// Capacity of the pricing (type of pricing)
	PriceCapacity string `pulumi:"priceCapacity"`
	// Product
	Product string `pulumi:"product"`
}

// A collection of values returned by getOrderCartProductOptionsPlan.
type GetOrderCartProductOptionsPlanResult struct {
	CartId      string  `pulumi:"cartId"`
	CatalogName *string `pulumi:"catalogName"`
	// Define if options of this family are exclusive with each other
	Exclusive bool `pulumi:"exclusive"`
	// Option family
	Family string `pulumi:"family"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Define if an option of this family is mandatory
	Mandatory       bool   `pulumi:"mandatory"`
	OptionsPlanCode string `pulumi:"optionsPlanCode"`
	// Product offer identifier
	PlanCode      string `pulumi:"planCode"`
	PriceCapacity string `pulumi:"priceCapacity"`
	// Prices of the product offer
	Prices  []GetOrderCartProductOptionsPlanPrice `pulumi:"prices"`
	Product string                                `pulumi:"product"`
	// Name of the product
	ProductName string `pulumi:"productName"`
	// Product type
	ProductType string `pulumi:"productType"`
	// Selected Price according to capacity
	SelectedPrices []GetOrderCartProductOptionsPlanSelectedPrice `pulumi:"selectedPrices"`
}

func GetOrderCartProductOptionsPlanOutput(ctx *pulumi.Context, args GetOrderCartProductOptionsPlanOutputArgs, opts ...pulumi.InvokeOption) GetOrderCartProductOptionsPlanResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetOrderCartProductOptionsPlanResult, error) {
			args := v.(GetOrderCartProductOptionsPlanArgs)
			r, err := GetOrderCartProductOptionsPlan(ctx, &args, opts...)
			var s GetOrderCartProductOptionsPlanResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetOrderCartProductOptionsPlanResultOutput)
}

// A collection of arguments for invoking getOrderCartProductOptionsPlan.
type GetOrderCartProductOptionsPlanOutputArgs struct {
	// Cart identifier
	CartId pulumi.StringInput `pulumi:"cartId"`
	// Catalog name
	CatalogName pulumi.StringPtrInput `pulumi:"catalogName"`
	// options plan code.
	OptionsPlanCode pulumi.StringInput `pulumi:"optionsPlanCode"`
	// Product offer identifier
	PlanCode pulumi.StringInput `pulumi:"planCode"`
	// Capacity of the pricing (type of pricing)
	PriceCapacity pulumi.StringInput `pulumi:"priceCapacity"`
	// Product
	Product pulumi.StringInput `pulumi:"product"`
}

func (GetOrderCartProductOptionsPlanOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrderCartProductOptionsPlanArgs)(nil)).Elem()
}

// A collection of values returned by getOrderCartProductOptionsPlan.
type GetOrderCartProductOptionsPlanResultOutput struct{ *pulumi.OutputState }

func (GetOrderCartProductOptionsPlanResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOrderCartProductOptionsPlanResult)(nil)).Elem()
}

func (o GetOrderCartProductOptionsPlanResultOutput) ToGetOrderCartProductOptionsPlanResultOutput() GetOrderCartProductOptionsPlanResultOutput {
	return o
}

func (o GetOrderCartProductOptionsPlanResultOutput) ToGetOrderCartProductOptionsPlanResultOutputWithContext(ctx context.Context) GetOrderCartProductOptionsPlanResultOutput {
	return o
}

func (o GetOrderCartProductOptionsPlanResultOutput) CartId() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) string { return v.CartId }).(pulumi.StringOutput)
}

func (o GetOrderCartProductOptionsPlanResultOutput) CatalogName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) *string { return v.CatalogName }).(pulumi.StringPtrOutput)
}

// Define if options of this family are exclusive with each other
func (o GetOrderCartProductOptionsPlanResultOutput) Exclusive() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) bool { return v.Exclusive }).(pulumi.BoolOutput)
}

// Option family
func (o GetOrderCartProductOptionsPlanResultOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) string { return v.Family }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetOrderCartProductOptionsPlanResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) string { return v.Id }).(pulumi.StringOutput)
}

// Define if an option of this family is mandatory
func (o GetOrderCartProductOptionsPlanResultOutput) Mandatory() pulumi.BoolOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) bool { return v.Mandatory }).(pulumi.BoolOutput)
}

func (o GetOrderCartProductOptionsPlanResultOutput) OptionsPlanCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) string { return v.OptionsPlanCode }).(pulumi.StringOutput)
}

// Product offer identifier
func (o GetOrderCartProductOptionsPlanResultOutput) PlanCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) string { return v.PlanCode }).(pulumi.StringOutput)
}

func (o GetOrderCartProductOptionsPlanResultOutput) PriceCapacity() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) string { return v.PriceCapacity }).(pulumi.StringOutput)
}

// Prices of the product offer
func (o GetOrderCartProductOptionsPlanResultOutput) Prices() GetOrderCartProductOptionsPlanPriceArrayOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) []GetOrderCartProductOptionsPlanPrice { return v.Prices }).(GetOrderCartProductOptionsPlanPriceArrayOutput)
}

func (o GetOrderCartProductOptionsPlanResultOutput) Product() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) string { return v.Product }).(pulumi.StringOutput)
}

// Name of the product
func (o GetOrderCartProductOptionsPlanResultOutput) ProductName() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) string { return v.ProductName }).(pulumi.StringOutput)
}

// Product type
func (o GetOrderCartProductOptionsPlanResultOutput) ProductType() pulumi.StringOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) string { return v.ProductType }).(pulumi.StringOutput)
}

// Selected Price according to capacity
func (o GetOrderCartProductOptionsPlanResultOutput) SelectedPrices() GetOrderCartProductOptionsPlanSelectedPriceArrayOutput {
	return o.ApplyT(func(v GetOrderCartProductOptionsPlanResult) []GetOrderCartProductOptionsPlanSelectedPrice {
		return v.SelectedPrices
	}).(GetOrderCartProductOptionsPlanSelectedPriceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOrderCartProductOptionsPlanResultOutput{})
}
