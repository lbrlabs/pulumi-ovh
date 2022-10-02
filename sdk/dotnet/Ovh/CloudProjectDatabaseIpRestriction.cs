// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh
{
    /// <summary>
    /// Apply IP restrictions to an OVHcloud Managed Database cluster.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Ovh = Lbrlabs.PulumiPackage.Ovh;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var db = Ovh.GetCloudProjectDatabase.Invoke(new()
    ///     {
    ///         ServiceName = "XXXX",
    ///         Engine = "YYYY",
    ///         ClusterId = "ZZZZ",
    ///     });
    /// 
    ///     var iprestriction = new Ovh.CloudProjectDatabaseIpRestriction("iprestriction", new()
    ///     {
    ///         ServiceName = ovh_cloud_project_database.Db.Service_name,
    ///         Engine = ovh_cloud_project_database.Db.Engine,
    ///         ClusterId = ovh_cloud_project_database.Db.Cluster_id,
    ///         Ip = "178.97.6.0/24",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OVHcloud Managed database cluster IP restrictions can be imported using the `service_name`, `engine`, `cluster_id` and the `ip`, separated by "/" E.g.,
    /// 
    /// ```sh
    ///  $ pulumi import ovh:index/cloudProjectDatabaseIpRestriction:CloudProjectDatabaseIpRestriction my_ip_restriction &lt;service_name&gt;/&lt;engine&gt;/&lt;cluster_id&gt;/178.97.6.0/24
    /// ```
    /// </summary>
    [OvhResourceType("ovh:index/cloudProjectDatabaseIpRestriction:CloudProjectDatabaseIpRestriction")]
    public partial class CloudProjectDatabaseIpRestriction : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Description of the IP restriction.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Output("engine")]
        public Output<string> Engine { get; private set; } = null!;

        /// <summary>
        /// Authorized IP.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Current status of the IP restriction.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a CloudProjectDatabaseIpRestriction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CloudProjectDatabaseIpRestriction(string name, CloudProjectDatabaseIpRestrictionArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/cloudProjectDatabaseIpRestriction:CloudProjectDatabaseIpRestriction", name, args ?? new CloudProjectDatabaseIpRestrictionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CloudProjectDatabaseIpRestriction(string name, Input<string> id, CloudProjectDatabaseIpRestrictionState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/cloudProjectDatabaseIpRestriction:CloudProjectDatabaseIpRestriction", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CloudProjectDatabaseIpRestriction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CloudProjectDatabaseIpRestriction Get(string name, Input<string> id, CloudProjectDatabaseIpRestrictionState? state = null, CustomResourceOptions? options = null)
        {
            return new CloudProjectDatabaseIpRestriction(name, id, state, options);
        }
    }

    public sealed class CloudProjectDatabaseIpRestrictionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Description of the IP restriction.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// Authorized IP.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public CloudProjectDatabaseIpRestrictionArgs()
        {
        }
        public static new CloudProjectDatabaseIpRestrictionArgs Empty => new CloudProjectDatabaseIpRestrictionArgs();
    }

    public sealed class CloudProjectDatabaseIpRestrictionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Description of the IP restriction.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit.
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine")]
        public Input<string>? Engine { get; set; }

        /// <summary>
        /// Authorized IP.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Current status of the IP restriction.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public CloudProjectDatabaseIpRestrictionState()
        {
        }
        public static new CloudProjectDatabaseIpRestrictionState Empty => new CloudProjectDatabaseIpRestrictionState();
    }
}