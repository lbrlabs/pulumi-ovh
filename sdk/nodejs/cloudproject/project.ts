// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Orders a public cloud project.
 *
 * ## Important
 *
 * This resource is in beta state. Use with caution.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@lbrlabs/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const mycart = ovh.Order.getCart({
 *     ovhSubsidiary: "fr",
 *     description: "my cloud order cart",
 * });
 * const cloudCartProductPlan = mycart.then(mycart => ovh.Order.getCartProductPlan({
 *     cartId: mycart.id,
 *     priceCapacity: "renew",
 *     product: "cloud",
 *     planCode: "project.2018",
 * }));
 * const cloudProject = new ovh.cloudproject.Project("cloudProject", {
 *     ovhSubsidiary: mycart.then(mycart => mycart.ovhSubsidiary),
 *     description: "my cloud project",
 *     paymentMean: "fidelity",
 *     plan: {
 *         duration: cloudCartProductPlan.then(cloudCartProductPlan => cloudCartProductPlan.selectedPrices?[0]?.duration),
 *         planCode: cloudCartProductPlan.then(cloudCartProductPlan => cloudCartProductPlan.planCode),
 *         pricingMode: cloudCartProductPlan.then(cloudCartProductPlan => cloudCartProductPlan.selectedPrices?[0]?.pricingMode),
 *     },
 * });
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    /**
     * project access
     */
    public /*out*/ readonly access!: pulumi.Output<string>;
    /**
     * A description associated with the user.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Details about an Order
     */
    public /*out*/ readonly orders!: pulumi.Output<outputs.CloudProject.ProjectOrder[]>;
    /**
     * OVHcloud Subsidiary
     */
    public readonly ovhSubsidiary!: pulumi.Output<string>;
    /**
     * OVHcloud payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
     */
    public readonly paymentMean!: pulumi.Output<string>;
    /**
     * Product Plan to order
     */
    public readonly plan!: pulumi.Output<outputs.CloudProject.ProjectPlan>;
    /**
     * Product Plan to order
     */
    public readonly planOptions!: pulumi.Output<outputs.CloudProject.ProjectPlanOption[] | undefined>;
    /**
     * openstack project id
     */
    public /*out*/ readonly projectId!: pulumi.Output<string>;
    /**
     * openstack project name
     */
    public /*out*/ readonly projectName!: pulumi.Output<string>;
    /**
     * project status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["access"] = state ? state.access : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["orders"] = state ? state.orders : undefined;
            resourceInputs["ovhSubsidiary"] = state ? state.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = state ? state.paymentMean : undefined;
            resourceInputs["plan"] = state ? state.plan : undefined;
            resourceInputs["planOptions"] = state ? state.planOptions : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.ovhSubsidiary === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ovhSubsidiary'");
            }
            if ((!args || args.paymentMean === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paymentMean'");
            }
            if ((!args || args.plan === undefined) && !opts.urn) {
                throw new Error("Missing required property 'plan'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ovhSubsidiary"] = args ? args.ovhSubsidiary : undefined;
            resourceInputs["paymentMean"] = args ? args.paymentMean : undefined;
            resourceInputs["plan"] = args ? args.plan : undefined;
            resourceInputs["planOptions"] = args ? args.planOptions : undefined;
            resourceInputs["access"] = undefined /*out*/;
            resourceInputs["orders"] = undefined /*out*/;
            resourceInputs["projectId"] = undefined /*out*/;
            resourceInputs["projectName"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    /**
     * project access
     */
    access?: pulumi.Input<string>;
    /**
     * A description associated with the user.
     */
    description?: pulumi.Input<string>;
    /**
     * Details about an Order
     */
    orders?: pulumi.Input<pulumi.Input<inputs.CloudProject.ProjectOrder>[]>;
    /**
     * OVHcloud Subsidiary
     */
    ovhSubsidiary?: pulumi.Input<string>;
    /**
     * OVHcloud payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
     */
    paymentMean?: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan?: pulumi.Input<inputs.CloudProject.ProjectPlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.CloudProject.ProjectPlanOption>[]>;
    /**
     * openstack project id
     */
    projectId?: pulumi.Input<string>;
    /**
     * openstack project name
     */
    projectName?: pulumi.Input<string>;
    /**
     * project status
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    /**
     * A description associated with the user.
     */
    description?: pulumi.Input<string>;
    /**
     * OVHcloud Subsidiary
     */
    ovhSubsidiary: pulumi.Input<string>;
    /**
     * OVHcloud payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
     */
    paymentMean: pulumi.Input<string>;
    /**
     * Product Plan to order
     */
    plan: pulumi.Input<inputs.CloudProject.ProjectPlan>;
    /**
     * Product Plan to order
     */
    planOptions?: pulumi.Input<pulumi.Input<inputs.CloudProject.ProjectPlanOption>[]>;
}
