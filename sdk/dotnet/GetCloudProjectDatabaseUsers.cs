// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Ovh
{
    public static class GetCloudProjectDatabaseUsers
    {
        /// <summary>
        /// Use this data source to get the list of users of a database cluster associated with a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var users = Ovh.GetCloudProjectDatabaseUsers.Invoke(new()
        ///     {
        ///         ServiceName = "XXXX",
        ///         Engine = "YYYY",
        ///         ClusterId = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["userIds"] = users.Apply(getCloudProjectDatabaseUsersResult =&gt; getCloudProjectDatabaseUsersResult.UserIds),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudProjectDatabaseUsersResult> InvokeAsync(GetCloudProjectDatabaseUsersArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectDatabaseUsersResult>("ovh:index/getCloudProjectDatabaseUsers:getCloudProjectDatabaseUsers", args ?? new GetCloudProjectDatabaseUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the list of users of a database cluster associated with a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var users = Ovh.GetCloudProjectDatabaseUsers.Invoke(new()
        ///     {
        ///         ServiceName = "XXXX",
        ///         Engine = "YYYY",
        ///         ClusterId = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["userIds"] = users.Apply(getCloudProjectDatabaseUsersResult =&gt; getCloudProjectDatabaseUsersResult.UserIds),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudProjectDatabaseUsersResult> Invoke(GetCloudProjectDatabaseUsersInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProjectDatabaseUsersResult>("ovh:index/getCloudProjectDatabaseUsers:getCloudProjectDatabaseUsers", args ?? new GetCloudProjectDatabaseUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProjectDatabaseUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// The engine of the database cluster you want to list users. To get a full list of available engine visit.
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine", required: true)]
        public string Engine { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectDatabaseUsersArgs()
        {
        }
        public static new GetCloudProjectDatabaseUsersArgs Empty => new GetCloudProjectDatabaseUsersArgs();
    }

    public sealed class GetCloudProjectDatabaseUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The engine of the database cluster you want to list users. To get a full list of available engine visit.
        /// [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        /// </summary>
        [Input("engine", required: true)]
        public Input<string> Engine { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCloudProjectDatabaseUsersInvokeArgs()
        {
        }
        public static new GetCloudProjectDatabaseUsersInvokeArgs Empty => new GetCloudProjectDatabaseUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudProjectDatabaseUsersResult
    {
        public readonly string ClusterId;
        public readonly string Engine;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ServiceName;
        /// <summary>
        /// The list of users ids of the database cluster associated with the project.
        /// </summary>
        public readonly ImmutableArray<string> UserIds;

        [OutputConstructor]
        private GetCloudProjectDatabaseUsersResult(
            string clusterId,

            string engine,

            string id,

            string serviceName,

            ImmutableArray<string> userIds)
        {
            ClusterId = clusterId;
            Engine = engine;
            Id = id;
            ServiceName = serviceName;
            UserIds = userIds;
        }
    }
}
