// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get the managed database of a public cloud project.
 *
 * ## Example Usage
 *
 * To get information of a database cluster service:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const db = ovh.CloudProjectDatabase.getDatabase({
 *     serviceName: "XXXXXX",
 *     engine: "YYYY",
 *     id: "ZZZZ",
 * });
 * export const clusterId = db.then(db => db.id);
 * ```
 */
export function getDatabase(args: GetDatabaseArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("ovh:CloudProjectDatabase/getDatabase:getDatabase", {
        "engine": args.engine,
        "id": args.id,
        "serviceName": args.serviceName,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabase.
 */
export interface GetDatabaseArgs {
    /**
     * The database engine you want to get information. To get a full list of available engine visit:
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     */
    engine: string;
    /**
     * Cluster ID
     */
    id: string;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: string;
}

/**
 * A collection of values returned by getDatabase.
 */
export interface GetDatabaseResult {
    /**
     * Time on which backups start every day.
     */
    readonly backupTime: string;
    /**
     * Date of the creation of the cluster.
     */
    readonly createdAt: string;
    /**
     * Small description of the database service.
     */
    readonly description: string;
    /**
     * The disk size (in GB) of the database service.
     */
    readonly diskSize: number;
    /**
     * The disk type of the database service.
     */
    readonly diskType: string;
    /**
     * List of all endpoints objects of the service.
     */
    readonly endpoints: outputs.CloudProjectDatabase.GetDatabaseEndpoint[];
    /**
     * See Argument Reference above.
     */
    readonly engine: string;
    /**
     * A valid OVHcloud public cloud database flavor name in which the nodes will be started.
     */
    readonly flavor: string;
    /**
     * See Argument Reference above.
     */
    readonly id: string;
    /**
     * Defines whether the REST API is enabled on a kafka cluster.
     */
    readonly kafkaRestApi: boolean;
    /**
     * Time on which maintenances can start every day.
     */
    readonly maintenanceTime: string;
    /**
     * Type of network of the cluster.
     */
    readonly networkType: string;
    /**
     * List of nodes object.
     */
    readonly nodes: outputs.CloudProjectDatabase.GetDatabaseNode[];
    readonly opensearchAclsEnabled: boolean;
    /**
     * Plan of the cluster.
     */
    readonly plan: string;
    /**
     * See Argument Reference above.
     */
    readonly serviceName: string;
    /**
     * Current status of the cluster.
     */
    readonly status: string;
    /**
     * The version of the engine in which the service should be deployed
     */
    readonly version: string;
}
/**
 * Use this data source to get the managed database of a public cloud project.
 *
 * ## Example Usage
 *
 * To get information of a database cluster service:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@pulumi/ovh";
 *
 * const db = ovh.CloudProjectDatabase.getDatabase({
 *     serviceName: "XXXXXX",
 *     engine: "YYYY",
 *     id: "ZZZZ",
 * });
 * export const clusterId = db.then(db => db.id);
 * ```
 */
export function getDatabaseOutput(args: GetDatabaseOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseResult> {
    return pulumi.output(args).apply((a: any) => getDatabase(a, opts))
}

/**
 * A collection of arguments for invoking getDatabase.
 */
export interface GetDatabaseOutputArgs {
    /**
     * The database engine you want to get information. To get a full list of available engine visit:
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     */
    engine: pulumi.Input<string>;
    /**
     * Cluster ID
     */
    id: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
