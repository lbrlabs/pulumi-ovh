// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Ovh.Inputs
{

    public sealed class CloudProjectKubePrivateNetworkConfigurationArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultVrackGateway", required: true)]
        public Input<string> DefaultVrackGateway { get; set; } = null!;

        [Input("privateNetworkRoutingAsDefault", required: true)]
        public Input<bool> PrivateNetworkRoutingAsDefault { get; set; } = null!;

        public CloudProjectKubePrivateNetworkConfigurationArgs()
        {
        }
        public static new CloudProjectKubePrivateNetworkConfigurationArgs Empty => new CloudProjectKubePrivateNetworkConfigurationArgs();
    }
}
