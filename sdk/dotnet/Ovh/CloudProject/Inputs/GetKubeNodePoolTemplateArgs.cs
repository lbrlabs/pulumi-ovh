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

    public sealed class GetKubeNodePoolTemplateInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("metadata")]
        public Input<Inputs.GetKubeNodePoolTemplateMetadataInputArgs>? Metadata { get; set; }

        [Input("spec")]
        public Input<Inputs.GetKubeNodePoolTemplateSpecInputArgs>? Spec { get; set; }

        public GetKubeNodePoolTemplateInputArgs()
        {
        }
        public static new GetKubeNodePoolTemplateInputArgs Empty => new GetKubeNodePoolTemplateInputArgs();
    }
}
