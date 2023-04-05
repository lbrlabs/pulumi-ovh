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

    public sealed class GetKubeCustomizationApiserverAdmissionpluginInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("disableds", required: true)]
        private InputList<string>? _disableds;

        /// <summary>
        /// Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.
        /// </summary>
        public InputList<string> Disableds
        {
            get => _disableds ?? (_disableds = new InputList<string>());
            set => _disableds = value;
        }

        [Input("enableds", required: true)]
        private InputList<string>? _enableds;

        /// <summary>
        /// Array of admission plugins enabled, default is ["NodeRestriction","AlwaysPulImages"] and only these admission plugins can be enabled at this time.
        /// </summary>
        public InputList<string> Enableds
        {
            get => _enableds ?? (_enableds = new InputList<string>());
            set => _enableds = value;
        }

        public GetKubeCustomizationApiserverAdmissionpluginInputArgs()
        {
        }
        public static new GetKubeCustomizationApiserverAdmissionpluginInputArgs Empty => new GetKubeCustomizationApiserverAdmissionpluginInputArgs();
    }
}