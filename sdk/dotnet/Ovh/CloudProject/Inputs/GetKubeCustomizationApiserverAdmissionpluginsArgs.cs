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

    public sealed class GetKubeCustomizationApiserverAdmissionpluginsInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("disableds", required: true)]
        private InputList<string>? _disableds;
        public InputList<string> Disableds
        {
            get => _disableds ?? (_disableds = new InputList<string>());
            set => _disableds = value;
        }

        [Input("enableds", required: true)]
        private InputList<string>? _enableds;
        public InputList<string> Enableds
        {
            get => _enableds ?? (_enableds = new InputList<string>());
            set => _enableds = value;
        }

        public GetKubeCustomizationApiserverAdmissionpluginsInputArgs()
        {
        }
        public static new GetKubeCustomizationApiserverAdmissionpluginsInputArgs Empty => new GetKubeCustomizationApiserverAdmissionpluginsInputArgs();
    }
}
