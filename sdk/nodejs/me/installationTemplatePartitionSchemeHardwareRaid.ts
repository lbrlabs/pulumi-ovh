// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this resource to create a hardware raid group in the partition scheme of a custom installation template available for dedicated servers.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@lbrlabs/pulumi-ovh";
 *
 * const mytemplate = new ovh.me.InstallationTemplate("mytemplate", {
 *     baseTemplateName: "centos7_64",
 *     templateName: "mytemplate",
 *     defaultLanguage: "fr",
 * });
 * const scheme = new ovh.me.InstallationTemplatePartitionScheme("scheme", {
 *     templateName: mytemplate.templateName,
 *     priority: 1,
 * });
 * const group1 = new ovh.me.InstallationTemplatePartitionSchemeHardwareRaid("group1", {
 *     templateName: scheme.templateName,
 *     schemeName: scheme.name,
 *     disks: [
 *         "[c1:d1,c1:d2,c1:d3]",
 *         "[c1:d10,c1:d20,c1:d30]",
 *     ],
 *     mode: "raid50",
 *     step: 1,
 * });
 * ```
 *
 * ## Import
 *
 * The resource can be imported using the `template_name`, `scheme_name`, `name` of the cluster, separated by "/" E.g., bash
 *
 * ```sh
 *  $ pulumi import ovh:Me/installationTemplatePartitionSchemeHardwareRaid:InstallationTemplatePartitionSchemeHardwareRaid group1 template_name/scheme_name/name
 * ```
 */
export class InstallationTemplatePartitionSchemeHardwareRaid extends pulumi.CustomResource {
    /**
     * Get an existing InstallationTemplatePartitionSchemeHardwareRaid resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstallationTemplatePartitionSchemeHardwareRaidState, opts?: pulumi.CustomResourceOptions): InstallationTemplatePartitionSchemeHardwareRaid {
        return new InstallationTemplatePartitionSchemeHardwareRaid(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:Me/installationTemplatePartitionSchemeHardwareRaid:InstallationTemplatePartitionSchemeHardwareRaid';

    /**
     * Returns true if the given object is an instance of InstallationTemplatePartitionSchemeHardwareRaid.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstallationTemplatePartitionSchemeHardwareRaid {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstallationTemplatePartitionSchemeHardwareRaid.__pulumiType;
    }

    /**
     * Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
     */
    public readonly disks!: pulumi.Output<string[]>;
    /**
     * RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
     */
    public readonly mode!: pulumi.Output<string>;
    /**
     * Hardware RAID name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The partition scheme name.
     */
    public readonly schemeName!: pulumi.Output<string>;
    /**
     * Specifies the creation order of the hardware RAID.
     */
    public readonly step!: pulumi.Output<number>;
    /**
     * The template name of the partition scheme.
     */
    public readonly templateName!: pulumi.Output<string>;

    /**
     * Create a InstallationTemplatePartitionSchemeHardwareRaid resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstallationTemplatePartitionSchemeHardwareRaidArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstallationTemplatePartitionSchemeHardwareRaidArgs | InstallationTemplatePartitionSchemeHardwareRaidState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstallationTemplatePartitionSchemeHardwareRaidState | undefined;
            resourceInputs["disks"] = state ? state.disks : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["schemeName"] = state ? state.schemeName : undefined;
            resourceInputs["step"] = state ? state.step : undefined;
            resourceInputs["templateName"] = state ? state.templateName : undefined;
        } else {
            const args = argsOrState as InstallationTemplatePartitionSchemeHardwareRaidArgs | undefined;
            if ((!args || args.disks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'disks'");
            }
            if ((!args || args.mode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mode'");
            }
            if ((!args || args.schemeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'schemeName'");
            }
            if ((!args || args.step === undefined) && !opts.urn) {
                throw new Error("Missing required property 'step'");
            }
            if ((!args || args.templateName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'templateName'");
            }
            resourceInputs["disks"] = args ? args.disks : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["schemeName"] = args ? args.schemeName : undefined;
            resourceInputs["step"] = args ? args.step : undefined;
            resourceInputs["templateName"] = args ? args.templateName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstallationTemplatePartitionSchemeHardwareRaid.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstallationTemplatePartitionSchemeHardwareRaid resources.
 */
export interface InstallationTemplatePartitionSchemeHardwareRaidState {
    /**
     * Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
     */
    disks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
     */
    mode?: pulumi.Input<string>;
    /**
     * Hardware RAID name.
     */
    name?: pulumi.Input<string>;
    /**
     * The partition scheme name.
     */
    schemeName?: pulumi.Input<string>;
    /**
     * Specifies the creation order of the hardware RAID.
     */
    step?: pulumi.Input<number>;
    /**
     * The template name of the partition scheme.
     */
    templateName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstallationTemplatePartitionSchemeHardwareRaid resource.
 */
export interface InstallationTemplatePartitionSchemeHardwareRaidArgs {
    /**
     * Disk List. Syntax is cX:dY for disks and [cX:dY,cX:dY] for groups. With X and Y resp. the controller id and the disk id.
     */
    disks: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * RAID mode (raid0, raid1, raid10, raid5, raid50, raid6, raid60).
     */
    mode: pulumi.Input<string>;
    /**
     * Hardware RAID name.
     */
    name?: pulumi.Input<string>;
    /**
     * The partition scheme name.
     */
    schemeName: pulumi.Input<string>;
    /**
     * Specifies the creation order of the hardware RAID.
     */
    step: pulumi.Input<number>;
    /**
     * The template name of the partition scheme.
     */
    templateName: pulumi.Input<string>;
}
