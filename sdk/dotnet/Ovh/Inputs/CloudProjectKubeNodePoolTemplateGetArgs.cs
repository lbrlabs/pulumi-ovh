// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Inputs
{

    public sealed class CloudProjectKubeNodePoolTemplateGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("metadata")]
        public Input<Inputs.CloudProjectKubeNodePoolTemplateMetadataGetArgs>? Metadata { get; set; }

        [Input("spec")]
        public Input<Inputs.CloudProjectKubeNodePoolTemplateSpecGetArgs>? Spec { get; set; }

        public CloudProjectKubeNodePoolTemplateGetArgs()
        {
        }
        public static new CloudProjectKubeNodePoolTemplateGetArgs Empty => new CloudProjectKubeNodePoolTemplateGetArgs();
    }
}