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

    public sealed class IpLoadBalancingHttpRouteActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// HTTP status code for "redirect" and "reject" actions
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Farm ID for "farm" action type or URL template for "redirect" action. You may use ${uri}, ${protocol}, ${host}, ${port} and ${path} variables in redirect target
        /// </summary>
        [Input("target")]
        public Input<string>? Target { get; set; }

        /// <summary>
        /// Action to trigger if all the rules of this route matches
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IpLoadBalancingHttpRouteActionArgs()
        {
        }
        public static new IpLoadBalancingHttpRouteActionArgs Empty => new IpLoadBalancingHttpRouteActionArgs();
    }
}
