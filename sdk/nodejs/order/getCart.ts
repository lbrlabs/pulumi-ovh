// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to create a temporary order cart to retrieve information order cart products.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mycart = ovh.Order.getCart({
 *     description: "my cart",
 *     ovhSubsidiary: "fr",
 * });
 * ```
 */
export function getCart(args: GetCartArgs, opts?: pulumi.InvokeOptions): Promise<GetCartResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Order/getCart:getCart", {
        "assign": args.assign,
        "description": args.description,
        "expire": args.expire,
        "ovhSubsidiary": args.ovhSubsidiary,
    }, opts);
}

/**
 * A collection of arguments for invoking getCart.
 */
export interface GetCartArgs {
    /**
     * Assign a shopping cart to an loggedin client. Values can be `true` or `false`.
     */
    assign?: boolean;
    /**
     * Description of your cart
     */
    description?: string;
    /**
     * Expiration time (format: 2006-01-02T15:04:05+00:00)
     */
    expire?: string;
    /**
     * OVHcloud Subsidiary
     */
    ovhSubsidiary: string;
}

/**
 * A collection of values returned by getCart.
 */
export interface GetCartResult {
    readonly assign?: boolean;
    /**
     * Cart identifier
     */
    readonly cartId: string;
    readonly description?: string;
    readonly expire: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Items of your cart
     */
    readonly items: number[];
    readonly ovhSubsidiary: string;
    /**
     * Indicates if the cart has already been validated
     */
    readonly readOnly: boolean;
}
/**
 * Use this data source to create a temporary order cart to retrieve information order cart products.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mycart = ovh.Order.getCart({
 *     description: "my cart",
 *     ovhSubsidiary: "fr",
 * });
 * ```
 */
export function getCartOutput(args: GetCartOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCartResult> {
    return pulumi.output(args).apply((a: any) => getCart(a, opts))
}

/**
 * A collection of arguments for invoking getCart.
 */
export interface GetCartOutputArgs {
    /**
     * Assign a shopping cart to an loggedin client. Values can be `true` or `false`.
     */
    assign?: pulumi.Input<boolean>;
    /**
     * Description of your cart
     */
    description?: pulumi.Input<string>;
    /**
     * Expiration time (format: 2006-01-02T15:04:05+00:00)
     */
    expire?: pulumi.Input<string>;
    /**
     * OVHcloud Subsidiary
     */
    ovhSubsidiary: pulumi.Input<string>;
}
