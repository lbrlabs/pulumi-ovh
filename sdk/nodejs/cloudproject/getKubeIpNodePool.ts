// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Use this data source to get a OVHcloud Managed Kubernetes node pool.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const nodepool = ovh.CloudProject.getKubeIpNodePool({
 *     serviceName: "XXXXXX",
 *     kubeId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
 *     name: "xxxxxx",
 * });
 * export const maxNodes = nodepool.then(nodepool => nodepool.maxNodes);
 * ```
 */
export function getKubeIpNodePool(args: GetKubeIpNodePoolArgs, opts?: pulumi.InvokeOptions): Promise<GetKubeIpNodePoolResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("ovh:CloudProject/getKubeIpNodePool:getKubeIpNodePool", {
        "kubeId": args.kubeId,
        "maxNodes": args.maxNodes,
        "minNodes": args.minNodes,
        "name": args.name,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getKubeIpNodePool.
 */
export interface GetKubeIpNodePoolArgs {
    /**
     * The id of the managed kubernetes cluster.
     */
    kubeId: string;
    /**
     * maximum number of nodes allowed in the pool.
     * Setting `desiredNodes` over this value will raise an error.
     */
    maxNodes?: number;
    /**
     * minimum number of nodes allowed in the pool.
     * Setting `desiredNodes` under this value will raise an error.
     */
    minNodes?: number;
    /**
     * The name of the node pool.
     */
    name: string;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getKubeIpNodePool.
 */
export interface GetKubeIpNodePoolResult {
    /**
     * (Optional) should the pool use the anti-affinity feature. Default to `false`.
     */
    readonly antiAffinity: boolean;
    /**
     * (Optional) Enable auto-scaling for the pool. Default to `false`.
     */
    readonly autoscale: boolean;
    /**
     * Number of nodes which are actually ready in the pool
     */
    readonly availableNodes: number;
    /**
     * Creation date
     */
    readonly createdAt: string;
    /**
     * Number of nodes present in the pool
     */
    readonly currentNodes: number;
    /**
     * Number of nodes you desire in the pool
     */
    readonly desiredNodes: number;
    /**
     * Flavor name
     */
    readonly flavor: string;
    /**
     * a valid OVHcloud public cloud flavor ID in which the nodes will be started.
     * Ex: "b2-7". Changing this value recreates the resource.
     * You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
     */
    readonly flavorName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * See Argument Reference above.
     */
    readonly kubeId: string;
    /**
     * maximum number of nodes allowed in the pool.
     * Setting `desiredNodes` over this value will raise an error.
     */
    readonly maxNodes: number;
    /**
     * minimum number of nodes allowed in the pool.
     * Setting `desiredNodes` under this value will raise an error.
     */
    readonly minNodes: number;
    /**
     * (Optional) should the nodes be billed on a monthly basis. Default to `false`.
     */
    readonly monthlyBilled: boolean;
    /**
     * (Optional) The name of the nodepool.
     * Changing this value recreates the resource.
     * Warning: "_" char is not allowed!
     */
    readonly name: string;
    /**
     * Project id
     */
    readonly projectId: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
    /**
     * Status describing the state between number of nodes wanted and available ones
     */
    readonly sizeStatus: string;
    /**
     * Current status
     */
    readonly status: string;
    /**
     * Number of nodes with the latest version installed in the pool
     */
    readonly upToDateNodes: number;
    /**
     * Last update date
     */
    readonly updatedAt: string;
}

export function getKubeIpNodePoolOutput(args: GetKubeIpNodePoolOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKubeIpNodePoolResult> {
    return pulumi.output(args).apply(a => getKubeIpNodePool(a, opts))
}

/**
 * A collection of arguments for invoking getKubeIpNodePool.
 */
export interface GetKubeIpNodePoolOutputArgs {
    /**
     * The id of the managed kubernetes cluster.
     */
    kubeId: pulumi.Input<string>;
    /**
     * maximum number of nodes allowed in the pool.
     * Setting `desiredNodes` over this value will raise an error.
     */
    maxNodes?: pulumi.Input<number>;
    /**
     * minimum number of nodes allowed in the pool.
     * Setting `desiredNodes` under this value will raise an error.
     */
    minNodes?: pulumi.Input<number>;
    /**
     * The name of the node pool.
     */
    name: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
