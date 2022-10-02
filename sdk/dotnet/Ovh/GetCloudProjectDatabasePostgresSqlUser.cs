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
    public static class GetCloudProjectDatabasePostgresSqlUser
    {
        /// <summary>
        /// Use this data source to get information about a user of a postgresql cluster associated with a public cloud project.
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
        ///     var pguser = Ovh.GetCloudProjectDatabasePostgresSqlUser.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///         Name = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["pguserRoles"] = pguser.Apply(getCloudProjectDatabasePostgresSqlUserResult =&gt; getCloudProjectDatabasePostgresSqlUserResult.Roles),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCloudProjectDatabasePostgresSqlUserResult> InvokeAsync(GetCloudProjectDatabasePostgresSqlUserArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetCloudProjectDatabasePostgresSqlUserResult>("ovh:index/getCloudProjectDatabasePostgresSqlUser:getCloudProjectDatabasePostgresSqlUser", args ?? new GetCloudProjectDatabasePostgresSqlUserArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get information about a user of a postgresql cluster associated with a public cloud project.
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
        ///     var pguser = Ovh.GetCloudProjectDatabasePostgresSqlUser.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///         ClusterId = "YYY",
        ///         Name = "ZZZ",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["pguserRoles"] = pguser.Apply(getCloudProjectDatabasePostgresSqlUserResult =&gt; getCloudProjectDatabasePostgresSqlUserResult.Roles),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCloudProjectDatabasePostgresSqlUserResult> Invoke(GetCloudProjectDatabasePostgresSqlUserInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetCloudProjectDatabasePostgresSqlUserResult>("ovh:index/getCloudProjectDatabasePostgresSqlUser:getCloudProjectDatabasePostgresSqlUser", args ?? new GetCloudProjectDatabasePostgresSqlUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCloudProjectDatabasePostgresSqlUserArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public string ClusterId { get; set; } = null!;

        /// <summary>
        /// Name of the user.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCloudProjectDatabasePostgresSqlUserArgs()
        {
        }
        public static new GetCloudProjectDatabasePostgresSqlUserArgs Empty => new GetCloudProjectDatabasePostgresSqlUserArgs();
    }

    public sealed class GetCloudProjectDatabasePostgresSqlUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cluster ID
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// Name of the user.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCloudProjectDatabasePostgresSqlUserInvokeArgs()
        {
        }
        public static new GetCloudProjectDatabasePostgresSqlUserInvokeArgs Empty => new GetCloudProjectDatabasePostgresSqlUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetCloudProjectDatabasePostgresSqlUserResult
    {
        public readonly string ClusterId;
        /// <summary>
        /// Date of the creation of the user.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name of the user.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Roles the user belongs to.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        public readonly string ServiceName;
        /// <summary>
        /// Current status of the user.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private GetCloudProjectDatabasePostgresSqlUserResult(
            string clusterId,

            string createdAt,

            string id,

            string name,

            ImmutableArray<string> roles,

            string serviceName,

            string status)
        {
            ClusterId = clusterId;
            CreatedAt = createdAt;
            Id = id;
            Name = name;
            Roles = roles;
            ServiceName = serviceName;
            Status = status;
        }
    }
}