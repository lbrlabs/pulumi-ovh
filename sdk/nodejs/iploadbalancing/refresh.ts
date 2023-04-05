// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Applies changes from other `ovh_iploadbalancing_*` resources to the production configuration of loadbalancers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@lbrlabs/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const lb = ovh.IpLoadBalancing.getIpLoadBalancing({
 *     serviceName: "ip-1.2.3.4",
 *     state: "ok",
 * });
 * const farmname = new ovh.iploadbalancing.TcpFarm("farmname", {
 *     port: 8080,
 *     serviceName: lb.then(lb => lb.id),
 *     zone: "all",
 * });
 * const backend = new ovh.iploadbalancing.TcpFarmServer("backend", {
 *     address: "4.5.6.7",
 *     backup: true,
 *     displayName: "mybackend",
 *     farmId: farmname.id,
 *     port: 80,
 *     probe: true,
 *     proxyProtocolVersion: "v2",
 *     serviceName: lb.then(lb => lb.id),
 *     ssl: false,
 *     status: "active",
 *     weight: 2,
 * });
 * const mylb = new ovh.iploadbalancing.Refresh("mylb", {
 *     keepers: [[backend].map(__item => __item.address)],
 *     serviceName: lb.then(lb => lb.id),
 * });
 * ```
 */
export class Refresh extends pulumi.CustomResource {
    /**
     * Get an existing Refresh resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RefreshState, opts?: pulumi.CustomResourceOptions): Refresh {
        return new Refresh(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:IpLoadBalancing/refresh:Refresh';

    /**
     * Returns true if the given object is an instance of Refresh.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Refresh {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Refresh.__pulumiType;
    }

    /**
     * List of values tracked to trigger refresh, used also to form implicit dependencies
     */
    public readonly keepers!: pulumi.Output<string[]>;
    /**
     * The internal name of your IP load balancing
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a Refresh resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RefreshArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RefreshArgs | RefreshState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RefreshState | undefined;
            resourceInputs["keepers"] = state ? state.keepers : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as RefreshArgs | undefined;
            if ((!args || args.keepers === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keepers'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["keepers"] = args ? args.keepers : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Refresh.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Refresh resources.
 */
export interface RefreshState {
    /**
     * List of values tracked to trigger refresh, used also to form implicit dependencies
     */
    keepers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Refresh resource.
 */
export interface RefreshArgs {
    /**
     * List of values tracked to trigger refresh, used also to form implicit dependencies
     */
    keepers: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The internal name of your IP load balancing
     */
    serviceName: pulumi.Input<string>;
}
