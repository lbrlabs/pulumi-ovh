// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Iam
{
    public static class GetReferenceResourceType
    {
        /// <summary>
        /// Use this data source to list all the IAM resource types.
        /// 
        /// ## Important
        /// 
        /// &gt; Using this resource requires that the account is enrolled in the OVHcloud [IAM beta](https://labs.ovhcloud.com/en/iam/) 
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
        ///     var types = Ovh.Iam.GetReferenceResourceType.Invoke();
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetReferenceResourceTypeResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetReferenceResourceTypeResult>("ovh:Iam/getReferenceResourceType:getReferenceResourceType", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetReferenceResourceTypeResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// List of resource types
        /// </summary>
        public readonly ImmutableArray<string> Types;

        [OutputConstructor]
        private GetReferenceResourceTypeResult(
            string id,

            ImmutableArray<string> types)
        {
            Id = id;
            Types = types;
        }
    }
}
