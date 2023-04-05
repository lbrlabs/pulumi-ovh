// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProject.Inputs
{

    public sealed class GetKubeCustomizationKubeProxyIpvsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Minimum period that IPVS rules are refreshed in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        [Input("minSyncPeriod")]
        public string? MinSyncPeriod { get; set; }

        /// <summary>
        /// IPVS scheduler.
        /// </summary>
        [Input("scheduler")]
        public string? Scheduler { get; set; }

        /// <summary>
        /// Minimum period that IPVS rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format.
        /// </summary>
        [Input("syncPeriod")]
        public string? SyncPeriod { get; set; }

        /// <summary>
        /// Timeout value used for IPVS TCP sessions after receiving a FIN in RFC3339 duration.
        /// </summary>
        [Input("tcpFinTimeout")]
        public string? TcpFinTimeout { get; set; }

        /// <summary>
        /// Timeout value used for idle IPVS TCP sessions in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        [Input("tcpTimeout")]
        public string? TcpTimeout { get; set; }

        /// <summary>
        /// timeout value used for IPVS UDP packets in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        [Input("udpTimeout")]
        public string? UdpTimeout { get; set; }

        public GetKubeCustomizationKubeProxyIpvsArgs()
        {
        }
        public static new GetKubeCustomizationKubeProxyIpvsArgs Empty => new GetKubeCustomizationKubeProxyIpvsArgs();
    }
}
