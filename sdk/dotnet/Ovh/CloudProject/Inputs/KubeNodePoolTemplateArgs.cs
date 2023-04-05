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

    public sealed class KubeNodePoolTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Metadata of each node in the pool
        /// </summary>
        [Input("metadata")]
        public Input<Inputs.KubeNodePoolTemplateMetadataArgs>? Metadata { get; set; }

        /// <summary>
        /// Spec of each node in the pool
        /// </summary>
        [Input("spec")]
        public Input<Inputs.KubeNodePoolTemplateSpecArgs>? Spec { get; set; }

        public KubeNodePoolTemplateArgs()
        {
        }
        public static new KubeNodePoolTemplateArgs Empty => new KubeNodePoolTemplateArgs();
    }
}
