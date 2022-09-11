// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Ovh.Outputs
{

    [OutputType]
    public sealed class GetCloudProjectContainerRegistryUsersResultResult
    {
        /// <summary>
        /// User email
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// User ID
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// User name
        /// </summary>
        public readonly string User;

        [OutputConstructor]
        private GetCloudProjectContainerRegistryUsersResultResult(
            string email,

            string id,

            string user)
        {
            Email = email;
            Id = id;
            User = user;
        }
    }
}
