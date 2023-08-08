// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProject
{
    public static class GetUsers
    {
        /// <summary>
        /// Get the list of all users of a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var projectUsers = Ovh.CloudProject.GetUsers.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///     });
        /// 
        ///     var users = .Where(user =&gt; user.Description == "S3-User").Select(user =&gt; 
        ///     {
        ///         return  user.UserId;
        ///     });
        /// 
        ///     var s3UserId = users[0];
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["userId"] = s3UserId,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetUsersResult> InvokeAsync(GetUsersArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetUsersResult>("ovh:CloudProject/getUsers:getUsers", args ?? new GetUsersArgs(), options.WithDefaults());

        /// <summary>
        /// Get the list of all users of a public cloud project.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var projectUsers = Ovh.CloudProject.GetUsers.Invoke(new()
        ///     {
        ///         ServiceName = "XXX",
        ///     });
        /// 
        ///     var users = .Where(user =&gt; user.Description == "S3-User").Select(user =&gt; 
        ///     {
        ///         return  user.UserId;
        ///     });
        /// 
        ///     var s3UserId = users[0];
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["userId"] = s3UserId,
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetUsersResult> Invoke(GetUsersInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetUsersResult>("ovh:CloudProject/getUsers:getUsers", args ?? new GetUsersInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetUsersArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetUsersArgs()
        {
        }
        public static new GetUsersArgs Empty => new GetUsersArgs();
    }

    public sealed class GetUsersInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetUsersInvokeArgs()
        {
        }
        public static new GetUsersInvokeArgs Empty => new GetUsersInvokeArgs();
    }


    [OutputType]
    public sealed class GetUsersResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string ServiceName;
        /// <summary>
        /// The list of users of a public cloud project.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetUsersUserResult> Users;

        [OutputConstructor]
        private GetUsersResult(
            string id,

            string serviceName,

            ImmutableArray<Outputs.GetUsersUserResult> users)
        {
            Id = id;
            ServiceName = serviceName;
            Users = users;
        }
    }
}
