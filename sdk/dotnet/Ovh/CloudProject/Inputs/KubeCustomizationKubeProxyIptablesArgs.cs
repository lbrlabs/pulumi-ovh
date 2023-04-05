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

    public sealed class KubeCustomizationKubeProxyIptablesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Minimum period that IPVS rules are refreshed in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration (e.g. `PT60S`).
        /// </summary>
        [Input("minSyncPeriod")]
        public Input<string>? MinSyncPeriod { get; set; }

        /// <summary>
        /// Minimum period that IPVS rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format (e.g. `PT60S`).
        /// </summary>
        [Input("syncPeriod")]
        public Input<string>? SyncPeriod { get; set; }

        public KubeCustomizationKubeProxyIptablesArgs()
        {
        }
        public static new KubeCustomizationKubeProxyIptablesArgs Empty => new KubeCustomizationKubeProxyIptablesArgs();
    }
}
