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
 * const rescue = ovh.Dedicated.getServerBoots({
 *     serviceName: "ns00000.ip-1-2-3.eu",
 *     bootType: "rescue",
 * });
 * const key = new ovh.me.SshKey("key", {
 *     keyName: "mykey",
 *     key: "ssh-ed25519 AAAAC3...",
 * });
 * const debian = new ovh.me.InstallationTemplate("debian", {
 *     baseTemplateName: "debian10_64",
 *     templateName: "mydebian10",
 *     defaultLanguage: "en",
 *     customization: {
 *         changeLog: "v1",
 *         customHostname: "mytest",
 *         sshKeyName: key.keyName,
 *     },
 * });
 * const serverInstall = new ovh.dedicated.ServiceInstallTask("serverInstall", {
 *     serviceName: "ns00000.ip-1-2-3.eu",
 *     templateName: debian.templateName,
 *     bootidOnDestroy: rescue.then(rescue => rescue.results?[0]),
 * });
 * ```
 */
export class ServiceInstallTask extends pulumi.CustomResource {
    /**
     * Get an existing ServiceInstallTask resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceInstallTaskState, opts?: pulumi.CustomResourceOptions): ServiceInstallTask {
        return new ServiceInstallTask(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Dedicated/serviceInstallTask:ServiceInstallTask';

    /**
     * Returns true if the given object is an instance of ServiceInstallTask.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceInstallTask {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceInstallTask.__pulumiType;
    }

    /**
     * If set, reboot the server on the specified boot id during destroy phase.
     */
    public readonly bootidOnDestroy!: pulumi.Output<number | undefined>;
    /**
     * Details of this task. (should be `Install asked`)
     */
    public /*out*/ readonly comment!: pulumi.Output<string>;
    /**
     * see `details` block below.
     */
    public readonly details!: pulumi.Output<outputs.Dedicated.ServiceInstallTaskDetails | undefined>;
    /**
     * Completion date in RFC3339 format.
     */
    public /*out*/ readonly doneDate!: pulumi.Output<string>;
    /**
     * Function name (should be `hardInstall`).
     */
    public /*out*/ readonly function!: pulumi.Output<string>;
    /**
     * Last update in RFC3339 format.
     */
    public /*out*/ readonly lastUpdate!: pulumi.Output<string>;
    /**
     * Partition scheme name.
     */
    public readonly partitionSchemeName!: pulumi.Output<string | undefined>;
    /**
     * The serviceName of your dedicated server.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Task creation date in RFC3339 format.
     */
    public /*out*/ readonly startDate!: pulumi.Output<string>;
    /**
     * Task status (should be `done`)
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Template name.
     */
    public readonly templateName!: pulumi.Output<string>;

    /**
     * Create a ServiceInstallTask resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceInstallTaskArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceInstallTaskArgs | ServiceInstallTaskState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceInstallTaskState | undefined;
            resourceInputs["bootidOnDestroy"] = state ? state.bootidOnDestroy : undefined;
            resourceInputs["comment"] = state ? state.comment : undefined;
            resourceInputs["details"] = state ? state.details : undefined;
            resourceInputs["doneDate"] = state ? state.doneDate : undefined;
            resourceInputs["function"] = state ? state.function : undefined;
            resourceInputs["lastUpdate"] = state ? state.lastUpdate : undefined;
            resourceInputs["partitionSchemeName"] = state ? state.partitionSchemeName : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["startDate"] = state ? state.startDate : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
        } else {
            const args = argsOrState as ServiceInstallTaskArgs | undefined;
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            resourceInputs["bootidOnDestroy"] = args ? args.bootidOnDestroy : undefined;
            resourceInputs["details"] = args ? args.details : undefined;
            resourceInputs["partitionSchemeName"] = args ? args.partitionSchemeName : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
            resourceInputs["comment"] = undefined /*out*/;
            resourceInputs["doneDate"] = undefined /*out*/;
            resourceInputs["function"] = undefined /*out*/;
            resourceInputs["lastUpdate"] = undefined /*out*/;
            resourceInputs["startDate"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceInstallTask.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceInstallTask resources.
 */
export interface ServiceInstallTaskState {
    /**
     * If set, reboot the server on the specified boot id during destroy phase.
     */
    bootidOnDestroy?: pulumi.Input<number>;
    /**
     * Details of this task. (should be `Install asked`)
     */
    comment?: pulumi.Input<string>;
    /**
     * see `details` block below.
     */
    details?: pulumi.Input<inputs.Dedicated.ServiceInstallTaskDetails>;
    /**
     * Completion date in RFC3339 format.
     */
    doneDate?: pulumi.Input<string>;
    /**
     * Function name (should be `hardInstall`).
     */
    function?: pulumi.Input<string>;
    /**
     * Last update in RFC3339 format.
     */
    lastUpdate?: pulumi.Input<string>;
    /**
     * Partition scheme name.
     */
    partitionSchemeName?: pulumi.Input<string>;
    /**
     * The serviceName of your dedicated server.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Task creation date in RFC3339 format.
     */
    startDate?: pulumi.Input<string>;
    /**
     * Task status (should be `done`)
     */
    status?: pulumi.Input<string>;
    /**
     * Template name.
     */
    templateName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceInstallTask resource.
 */
export interface ServiceInstallTaskArgs {
    /**
     * If set, reboot the server on the specified boot id during destroy phase.
     */
    bootidOnDestroy?: pulumi.Input<number>;
    /**
     * see `details` block below.
     */
    details?: pulumi.Input<inputs.Dedicated.ServiceInstallTaskDetails>;
    /**
     * Partition scheme name.
     */
    partitionSchemeName?: pulumi.Input<string>;
    /**
     * The serviceName of your dedicated server.
     */
    serviceName: pulumi.Input<string>;
    /**
     * Template name.
     */
    templateName: pulumi.Input<string>;
}
