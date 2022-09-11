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

    public sealed class CloudProjectDatabaseEndpointGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of component the URI relates to.
        /// </summary>
        [Input("component")]
        public Input<string>? Component { get; set; }

        /// <summary>
        /// Domain of the cluster.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Path of the endpoint.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Connection port for the endpoint.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Scheme used to generate the URI.
        /// </summary>
        [Input("scheme")]
        public Input<string>? Scheme { get; set; }

        /// <summary>
        /// Defines whether the endpoint uses SSL.
        /// </summary>
        [Input("ssl")]
        public Input<bool>? Ssl { get; set; }

        /// <summary>
        /// SSL mode used to connect to the service if the SSL is enabled.
        /// </summary>
        [Input("sslMode")]
        public Input<string>? SslMode { get; set; }

        /// <summary>
        /// URI of the endpoint.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public CloudProjectDatabaseEndpointGetArgs()
        {
        }
        public static new CloudProjectDatabaseEndpointGetArgs Empty => new CloudProjectDatabaseEndpointGetArgs();
    }
}
