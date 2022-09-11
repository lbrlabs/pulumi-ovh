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
    public sealed class CloudProjectDatabaseNode
    {
        /// <summary>
        /// Private network id in which the node should be deployed.
        /// </summary>
        public readonly string? NetworkId;
        /// <summary>
        /// Public cloud region in which the node should be deployed.
        /// Ex: "GRA'.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Private subnet ID in which the node is.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private CloudProjectDatabaseNode(
            string? networkId,

            string region,

            string? subnetId)
        {
            NetworkId = networkId;
            Region = region;
            SubnetId = subnetId;
        }
    }
}
