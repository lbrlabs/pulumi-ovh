// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to list the IAM action associated with a resource type.
 *
 * ## Important
 *
 * > Using this resource requires that the account is enrolled in the OVHcloud [IAM beta](https://labs.ovhcloud.com/en/iam/)
 */
export function getIamReferenceActions(args: GetIamReferenceActionsArgs, opts?: pulumi.InvokeOptions): Promise<GetIamReferenceActionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:Dbaas/getIamReferenceActions:getIamReferenceActions", {
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getIamReferenceActions.
 */
export interface GetIamReferenceActionsArgs {
    /**
     * Kind of resource we want the actions for
     */
    type: string;
}

/**
 * A collection of values returned by getIamReferenceActions.
 */
export interface GetIamReferenceActionsResult {
    /**
     * List of actions
     */
    readonly actions: outputs.Dbaas.GetIamReferenceActionsAction[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly type: string;
}
/**
 * Use this data source to list the IAM action associated with a resource type.
 *
 * ## Important
 *
 * > Using this resource requires that the account is enrolled in the OVHcloud [IAM beta](https://labs.ovhcloud.com/en/iam/)
 */
export function getIamReferenceActionsOutput(args: GetIamReferenceActionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetIamReferenceActionsResult> {
    return pulumi.output(args).apply((a: any) => getIamReferenceActions(a, opts))
}

/**
 * A collection of arguments for invoking getIamReferenceActions.
 */
export interface GetIamReferenceActionsOutputArgs {
    /**
     * Kind of resource we want the actions for
     */
    type: pulumi.Input<string>;
}