// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Create a new IP whitelist on your private cloud database instance.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@lbrlabs/pulumi-ovh";
 *
 * const ip = new ovh.hosting.PrivateDatabaseAllowlist("ip", {
 *     ip: "1.2.3.4",
 *     service: true,
 *     serviceName: "XXXXXX",
 *     sftp: true,
 * });
 * ```
 *
 * ## Import
 *
 * OVHcloud database whitelist can be imported using the `service_name` and the `ip`, separated by "/" E.g.,
 *
 * ```sh
 *  $ pulumi import ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist ip service_name/ip
 * ```
 */
export class PrivateDatabaseAllowlist extends pulumi.CustomResource {
    /**
     * Get an existing PrivateDatabaseAllowlist resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PrivateDatabaseAllowlistState, opts?: pulumi.CustomResourceOptions): PrivateDatabaseAllowlist {
        return new PrivateDatabaseAllowlist(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist';

    /**
     * Returns true if the given object is an instance of PrivateDatabaseAllowlist.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PrivateDatabaseAllowlist {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PrivateDatabaseAllowlist.__pulumiType;
    }

    /**
     * The whitelisted IP in your instance.
     */
    public readonly ip!: pulumi.Output<string>;
    /**
     * Custom name for your Whitelisted IP.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Authorize this IP to access service port. Values can be `true` or `false`
     */
    public readonly service!: pulumi.Output<boolean>;
    /**
     * The internal name of your private database.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Authorize this IP to access SFTP port. Values can be `true` or `false`
     */
    public readonly sftp!: pulumi.Output<boolean>;

    /**
     * Create a PrivateDatabaseAllowlist resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PrivateDatabaseAllowlistArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PrivateDatabaseAllowlistArgs | PrivateDatabaseAllowlistState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PrivateDatabaseAllowlistState | undefined;
            resourceInputs["ip"] = state ? state.ip : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["sftp"] = state ? state.sftp : undefined;
        } else {
            const args = argsOrState as PrivateDatabaseAllowlistArgs | undefined;
            if ((!args || args.ip === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ip'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.sftp === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sftp'");
            }
            resourceInputs["ip"] = args ? args.ip : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["sftp"] = args ? args.sftp : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(PrivateDatabaseAllowlist.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering PrivateDatabaseAllowlist resources.
 */
export interface PrivateDatabaseAllowlistState {
    /**
     * The whitelisted IP in your instance.
     */
    ip?: pulumi.Input<string>;
    /**
     * Custom name for your Whitelisted IP.
     */
    name?: pulumi.Input<string>;
    /**
     * Authorize this IP to access service port. Values can be `true` or `false`
     */
    service?: pulumi.Input<boolean>;
    /**
     * The internal name of your private database.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Authorize this IP to access SFTP port. Values can be `true` or `false`
     */
    sftp?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a PrivateDatabaseAllowlist resource.
 */
export interface PrivateDatabaseAllowlistArgs {
    /**
     * The whitelisted IP in your instance.
     */
    ip: pulumi.Input<string>;
    /**
     * Custom name for your Whitelisted IP.
     */
    name?: pulumi.Input<string>;
    /**
     * Authorize this IP to access service port. Values can be `true` or `false`
     */
    service: pulumi.Input<boolean>;
    /**
     * The internal name of your private database.
     */
    serviceName: pulumi.Input<string>;
    /**
     * Authorize this IP to access SFTP port. Values can be `true` or `false`
     */
    sftp: pulumi.Input<boolean>;
}
