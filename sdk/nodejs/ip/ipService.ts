// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@lbrlabs/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mycart = ovh.Order.getCart({
 *     ovhSubsidiary: "fr",
 *     description: "order ip block",
 * });
 * const ipblockCartProductPlan = mycart.then(mycart => ovh.Order.getCartProductPlan({
 *     cartId: mycart.id,
 *     priceCapacity: "renew",
 *     product: "ip",
 *     planCode: "ip-v4-s30-ripe",
 * }));
 * const ipblockIpService = new ovh.ip.IpService("ipblockIpService", {
 *     ovhSubsidiary: mycart.then(mycart => mycart.ovhSubsidiary),
 *     description: "my ip block",
 *     plan: {
 *         duration: ipblockCartProductPlan.then(ipblockCartProductPlan => ipblockCartProductPlan.selectedPrices?.[0]?.duration),
 *         planCode: ipblockCartProductPlan.then(ipblockCartProductPlan => ipblockCartProductPlan.planCode),
 *         pricingMode: ipblockCartProductPlan.then(ipblockCartProductPlan => ipblockCartProductPlan.selectedPrices?.[0]?.pricingMode),
 *         configurations: [{
 *             label: "country",
 *             value: "FR",
 *         }],
 *     },
 * });
 * ```
 */
export class IpService extends pulumi.CustomResource {
    /**
     * Get an existing IpService resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpServiceState, opts?: pulumi.CustomResourceOptions): IpService {
        return new IpService(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Ip/ipService:IpService';

    /**
     * Returns true if the given object is an instance of IpService.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpService {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpService.__pulumiType;
    }

    /**
     * can be terminated
     */
    public /*out*/ readonly canBeTerminated!: pulumi.Output<boolean>;
    /**
     * country
     */
    public /*out*/ readonly country!: pulumi.Output<string>;
    /**
     * Custom description on your ip.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * ip block
     */
    public /*out*/ readonly ip!: pulumi.Output<string>;
    /**
     * Details about an Order
     */
    public /*out*/ readonly orders!: pulumi.Output<outputs.Ip.IpServiceOrder[]>;
    /**
     * IP block organisation Id
     */
    public /*out*/ readonly organisationId!: pulumi.Output<string>;
    /**
     * OVHcloud Subsidiary
     */
    public readonly ovhSubsidiary!: pulumi.Output<string>;
    /**
     * Ovh payment mode
     *
     * @deprecated This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     */
    public readonly paymentMean!: pulumi.Output<string | undefined>;
    /**
     * Product Plan to order
     */
    public readonly plan!: pulumi.Output<outputs.Ip.IpServicePlan>;
    /**
     * Product Plan to order
     */
    public readonly planOptions!: pulumi.Output<outputs.Ip.IpServicePlanOption[] | undefined>;
    /**
     * Routage information
     */
    public /*out*/ readonly routedTos!: pulumi.Output<outputs.Ip.IpServiceRoutedTo[]>;
    /**
     * service name
     */
    public /*out*/ readonly serviceName!: pulumi.Output<string>;
    /**
     * Possible values for ip type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a IpService resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpServiceArgs | IpServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpServiceState | undefined;
            resourceInputs["canBeTerminated"] = state ? state.canBeTerminated : undefined;
            resourceInputs["country"] = state ? state.country : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["orders"] = state ? state.orders : undefined;
            resourceInputs["organisationId"] = state ? state.organisationId : undefined;
            resourceInputs["ovhSubsidiary"] = state ? state.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = state ? state.paymentMean : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["planOptions"] = state ? state.planOptions : undefined;
            resourceInputs["routedTos"] = state ? state.routedTos : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as IpServiceArgs | undefined;
            if ((!args || args.ovhSubsidiary === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ovhSubsidiary'");
            }
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ovhSubsidiary"] = args ? args.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = args ? args.paymentMean : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["planOptions"] = args ? args.planOptions : undefined;
            resourceInputs["canBeTerminated"] = undefined /*out*/;
            resourceInputs["country"] = undefined /*out*/;
            resourceInputs["ip"] = undefined /*out*/;
            resourceInputs["orders"] = undefined /*out*/;
            resourceInputs["organisationId"] = undefined /*out*/;
            resourceInputs["routedTos"] = undefined /*out*/;
            resourceInputs["serviceName"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpService.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpService resources.
 */
export interface IpServiceState {
    /**
     * can be terminated
     */
    canBeTerminated?: pulumi.Input<boolean>;
    /**
     * country
     */
    country?: pulumi.Input<string>;
    /**
     * Custom description on your ip.
     */
    description?: pulumi.Input<string>;
    /**
     * ip block
     */
    ip?: pulumi.Input<string>;
    /**
     * Details about an Order
     */
    orders?: pulumi.Input<pulumi.Input<inputs.Ip.IpServiceOrder>[]>;
    /**
     * IP block organisation Id
     */
    organisationId?: pulumi.Input<string>;
    /**
     * OVHcloud Subsidiary
     */
    ovhSubsidiary?: pulumi.Input<string>;
    /**
     * Ovh payment mode
     *
     * @deprecated This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     */
    paymentMean?: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan?: pulumi.Input<inputs.Ip.IpServicePlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.Ip.IpServicePlanOption>[]>;
    /**
     * Routage information
     */
    routedTos?: pulumi.Input<pulumi.Input<inputs.Ip.IpServiceRoutedTo>[]>;
    /**
     * service name
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Possible values for ip type
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpService resource.
 */
export interface IpServiceArgs {
    /**
     * Custom description on your ip.
     */
    description?: pulumi.Input<string>;
    /**
     * OVHcloud Subsidiary
     */
    ovhSubsidiary: pulumi.Input<string>;
    /**
     * Ovh payment mode
     *
     * @deprecated This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     */
    paymentMean?: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan: pulumi.Input<inputs.Ip.IpServicePlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.Ip.IpServicePlanOption>[]>;
}
