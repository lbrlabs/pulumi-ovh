// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Me
{
    public static class GetInstallationTemplates
    {
        /// <summary>
        /// Use this data source to get the list of custom installation templates available for dedicated servers.
        /// </summary>
        public static Task<GetInstallationTemplatesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetInstallationTemplatesResult>("ovh:Me/getInstallationTemplates:getInstallationTemplates", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetInstallationTemplatesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The list of custom installation templates IDs available for dedicated servers.
        /// </summary>
        public readonly ImmutableArray<string> Results;

        [OutputConstructor]
        private GetInstallationTemplatesResult(
            string id,

            ImmutableArray<string> results)
        {
            Id = id;
            Results = results;
        }
    }
}