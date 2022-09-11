// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * Creates a container registry associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 * import * as pulumi_ovh from "@lbrlabs/pulumi-ovh";
 *
 * const regcap = ovh.getCloudProjectCapabilitiesContainerFilter({
 *     serviceName: "XXXXXX",
 *     planName: "SMALL",
 *     region: "GRA",
 * });
 * const reg = new ovh.CloudProjectContainerRegistry("reg", {
 *     serviceName: regcap.then(regcap => regcap.serviceName),
 *     planId: regcap.then(regcap => regcap.id),
 *     region: regcap.then(regcap => regcap.region),
 * });
 * ```
 */
export class CloudProjectContainerRegistry extends pulumi.CustomResource {
    /**
     * Get an existing CloudProjectContainerRegistry resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CloudProjectContainerRegistryState, opts?: pulumi.CustomResourceOptions): CloudProjectContainerRegistry {
        return new CloudProjectContainerRegistry(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:index/cloudProjectContainerRegistry:CloudProjectContainerRegistry';

    /**
     * Returns true if the given object is an instance of CloudProjectContainerRegistry.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CloudProjectContainerRegistry {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CloudProjectContainerRegistry.__pulumiType;
    }

    /**
     * Plan creation date
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Registry name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Plan ID of the registry
     */
    public readonly planId!: pulumi.Output<string>;
    /**
     * Plan of the registry
     */
    public /*out*/ readonly plans!: pulumi.Output<outputs.CloudProjectContainerRegistryPlan[]>;
    /**
     * Project ID of your registry
     */
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    /**
     * Region of the registry
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Current size of the registry (bytes)
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * Registry status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Registry last update date
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * Access url of the registry
     */
    public /*out*/ readonly url!: pulumi.Output<string>;
    /**
     * Version of your registry
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a CloudProjectContainerRegistry resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CloudProjectContainerRegistryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CloudProjectContainerRegistryArgs | CloudProjectContainerRegistryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CloudProjectContainerRegistryState | undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["planId"] = state ? state.planId : undefined;
            resourceInputs["plans"] = state ? state.plans : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as CloudProjectContainerRegistryArgs | undefined;
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["planId"] = args ? args.planId : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["plans"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["url"] = undefined /*out*/;
            resourceInputs["version"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CloudProjectContainerRegistry.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CloudProjectContainerRegistry resources.
 */
export interface CloudProjectContainerRegistryState {
    /**
     * Plan creation date
     */
    createdAt?: pulumi.Input<string>;
    /**
     * Registry name
     */
    name?: pulumi.Input<string>;
    /**
     * Plan ID of the registry
     */
    planId?: pulumi.Input<string>;
    /**
     * Plan of the registry
     */
    plans?: pulumi.Input<pulumi.Input<inputs.CloudProjectContainerRegistryPlan>[]>;
    /**
     * Project ID of your registry
     */
    projectId?: pulumi.Input<string>;
    /**
     * Region of the registry
     */
    region?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Current size of the registry (bytes)
     */
    size?: pulumi.Input<number>;
    /**
     * Registry status
     */
    status?: pulumi.Input<string>;
    /**
     * Registry last update date
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * Access url of the registry
     */
    url?: pulumi.Input<string>;
    /**
     * Version of your registry
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CloudProjectContainerRegistry resource.
 */
export interface CloudProjectContainerRegistryArgs {
    /**
     * Registry name
     */
    name?: pulumi.Input<string>;
    /**
     * Plan ID of the registry
     */
    planId?: pulumi.Input<string>;
    /**
     * Region of the registry
     */
    region: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
