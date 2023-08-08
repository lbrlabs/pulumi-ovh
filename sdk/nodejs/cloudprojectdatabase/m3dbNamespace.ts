// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates a namespace for a M3DB cluster associated with a public cloud project.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@lbrlabs/pulumi-ovh";
 * import * as ovh from "@pulumi/ovh";
 *
 * const m3db = ovh.CloudProjectDatabase.getDatabase({
 *     serviceName: "XXX",
 *     engine: "m3db",
 *     id: "ZZZ",
 * });
 * const namespace = new ovh.cloudprojectdatabase.M3DbNamespace("namespace", {
 *     serviceName: m3db.then(m3db => m3db.serviceName),
 *     clusterId: m3db.then(m3db => m3db.id),
 *     resolution: "P2D",
 *     retentionPeriodDuration: "PT48H",
 * });
 * ```
 *
 * ## Import
 *
 * OVHcloud Managed M3DB clusters namespaces can be imported using the `service_name`, `cluster_id` and `id` of the namespace, separated by "/" E.g., bash
 *
 * ```sh
 *  $ pulumi import ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace my_namespace service_name/cluster_id/id
 * ```
 */
export class M3DbNamespace extends pulumi.CustomResource {
    /**
     * Get an existing M3DbNamespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: M3DbNamespaceState, opts?: pulumi.CustomResourceOptions): M3DbNamespace {
        return new M3DbNamespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProjectDatabase/m3DbNamespace:M3DbNamespace';

    /**
     * Returns true if the given object is an instance of M3DbNamespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is M3DbNamespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === M3DbNamespace.__pulumiType;
    }

    /**
     * Cluster ID.
     */
    public readonly clusterId!: pulumi.Output<string>;
    /**
     * Name of the namespace.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
     */
    public readonly resolution!: pulumi.Output<string>;
    /**
     * Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
     */
    public readonly retentionBlockDataExpirationDuration!: pulumi.Output<string | undefined>;
    /**
     * Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
     */
    public readonly retentionBlockSizeDuration!: pulumi.Output<string>;
    /**
     * Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     */
    public readonly retentionBufferFutureDuration!: pulumi.Output<string | undefined>;
    /**
     * Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     */
    public readonly retentionBufferPastDuration!: pulumi.Output<string | undefined>;
    /**
     * Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
     */
    public readonly retentionPeriodDuration!: pulumi.Output<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;
    /**
     * Defines whether M3DB will create snapshot files for this namespace.
     */
    public readonly snapshotEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Type of namespace.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Defines whether M3DB will include writes to this namespace in the commit log.
     */
    public readonly writesToCommitLogEnabled!: pulumi.Output<boolean | undefined>;

    /**
     * Create a M3DbNamespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: M3DbNamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: M3DbNamespaceArgs | M3DbNamespaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as M3DbNamespaceState | undefined;
            resourceInputs["clusterId"] = state ? state.clusterId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["resolution"] = state ? state.resolution : undefined;
            resourceInputs["retentionBlockDataExpirationDuration"] = state ? state.retentionBlockDataExpirationDuration : undefined;
            resourceInputs["retentionBlockSizeDuration"] = state ? state.retentionBlockSizeDuration : undefined;
            resourceInputs["retentionBufferFutureDuration"] = state ? state.retentionBufferFutureDuration : undefined;
            resourceInputs["retentionBufferPastDuration"] = state ? state.retentionBufferPastDuration : undefined;
            resourceInputs["retentionPeriodDuration"] = state ? state.retentionPeriodDuration : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
            resourceInputs["snapshotEnabled"] = state ? state.snapshotEnabled : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["writesToCommitLogEnabled"] = state ? state.writesToCommitLogEnabled : undefined;
        } else {
            const args = argsOrState as M3DbNamespaceArgs | undefined;
            if ((!args || args.clusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterId'");
            }
            if ((!args || args.resolution === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resolution'");
            }
            if ((!args || args.retentionPeriodDuration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'retentionPeriodDuration'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["clusterId"] = args ? args.clusterId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["resolution"] = args ? args.resolution : undefined;
            resourceInputs["retentionBlockDataExpirationDuration"] = args ? args.retentionBlockDataExpirationDuration : undefined;
            resourceInputs["retentionBlockSizeDuration"] = args ? args.retentionBlockSizeDuration : undefined;
            resourceInputs["retentionBufferFutureDuration"] = args ? args.retentionBufferFutureDuration : undefined;
            resourceInputs["retentionBufferPastDuration"] = args ? args.retentionBufferPastDuration : undefined;
            resourceInputs["retentionPeriodDuration"] = args ? args.retentionPeriodDuration : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["snapshotEnabled"] = args ? args.snapshotEnabled : undefined;
            resourceInputs["writesToCommitLogEnabled"] = args ? args.writesToCommitLogEnabled : undefined;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(M3DbNamespace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering M3DbNamespace resources.
 */
export interface M3DbNamespaceState {
    /**
     * Cluster ID.
     */
    clusterId?: pulumi.Input<string>;
    /**
     * Name of the namespace.
     */
    name?: pulumi.Input<string>;
    /**
     * Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
     */
    resolution?: pulumi.Input<string>;
    /**
     * Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
     */
    retentionBlockDataExpirationDuration?: pulumi.Input<string>;
    /**
     * Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
     */
    retentionBlockSizeDuration?: pulumi.Input<string>;
    /**
     * Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     */
    retentionBufferFutureDuration?: pulumi.Input<string>;
    /**
     * Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     */
    retentionBufferPastDuration?: pulumi.Input<string>;
    /**
     * Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
     */
    retentionPeriodDuration?: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * Defines whether M3DB will create snapshot files for this namespace.
     */
    snapshotEnabled?: pulumi.Input<boolean>;
    /**
     * Type of namespace.
     */
    type?: pulumi.Input<string>;
    /**
     * Defines whether M3DB will include writes to this namespace in the commit log.
     */
    writesToCommitLogEnabled?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a M3DbNamespace resource.
 */
export interface M3DbNamespaceArgs {
    /**
     * Cluster ID.
     */
    clusterId: pulumi.Input<string>;
    /**
     * Name of the namespace.
     */
    name?: pulumi.Input<string>;
    /**
     * Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
     */
    resolution: pulumi.Input<string>;
    /**
     * Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
     */
    retentionBlockDataExpirationDuration?: pulumi.Input<string>;
    /**
     * Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
     */
    retentionBlockSizeDuration?: pulumi.Input<string>;
    /**
     * Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     */
    retentionBufferFutureDuration?: pulumi.Input<string>;
    /**
     * Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     */
    retentionBufferPastDuration?: pulumi.Input<string>;
    /**
     * Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
     */
    retentionPeriodDuration: pulumi.Input<string>;
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
    /**
     * Defines whether M3DB will create snapshot files for this namespace.
     */
    snapshotEnabled?: pulumi.Input<boolean>;
    /**
     * Defines whether M3DB will include writes to this namespace in the commit log.
     */
    writesToCommitLogEnabled?: pulumi.Input<boolean>;
}
