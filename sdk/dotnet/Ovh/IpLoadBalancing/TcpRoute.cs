// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.IpLoadBalancing
{
    /// <summary>
    /// Manage TCP route for a loadbalancer service
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Ovh = Lbrlabs.PulumiPackage.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var tcpreject = new Ovh.IpLoadBalancing.TcpRoute("tcpreject", new()
    ///     {
    ///         Action = new Ovh.IpLoadBalancing.Inputs.TcpRouteActionArgs
    ///         {
    ///             Type = "reject",
    ///         },
    ///         ServiceName = "loadbalancer-xxxxxxxxxxxxxxxxxx",
    ///         Weight = 1,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:IpLoadBalancing/tcpRoute:TcpRoute")]
    public partial class TcpRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Action triggered when all rules match
        /// </summary>
        [Output("action")]
        public Output<Outputs.TcpRouteAction> Action { get; private set; } = null!;

        /// <summary>
        /// Human readable name for your route, this field is for you
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Route traffic for this frontend
        /// </summary>
        [Output("frontendId")]
        public Output<int> FrontendId { get; private set; } = null!;

        /// <summary>
        /// List of rules to match to trigger action
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.TcpRouteRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Route status. Routes in "ok" state are ready to operate
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
        /// </summary>
        [Output("weight")]
        public Output<int> Weight { get; private set; } = null!;


        /// <summary>
        /// Create a TcpRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TcpRoute(string name, TcpRouteArgs args, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/tcpRoute:TcpRoute", name, args ?? new TcpRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TcpRoute(string name, Input<string> id, TcpRouteState? state = null, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/tcpRoute:TcpRoute", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TcpRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TcpRoute Get(string name, Input<string> id, TcpRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new TcpRoute(name, id, state, options);
        }
    }

    public sealed class TcpRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action triggered when all rules match
        /// </summary>
        [Input("action", required: true)]
        public Input<Inputs.TcpRouteActionArgs> Action { get; set; } = null!;

        /// <summary>
        /// Human readable name for your route, this field is for you
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Route traffic for this frontend
        /// </summary>
        [Input("frontendId")]
        public Input<int>? FrontendId { get; set; }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public TcpRouteArgs()
        {
        }
        public static new TcpRouteArgs Empty => new TcpRouteArgs();
    }

    public sealed class TcpRouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action triggered when all rules match
        /// </summary>
        [Input("action")]
        public Input<Inputs.TcpRouteActionGetArgs>? Action { get; set; }

        /// <summary>
        /// Human readable name for your route, this field is for you
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Route traffic for this frontend
        /// </summary>
        [Input("frontendId")]
        public Input<int>? FrontendId { get; set; }

        [Input("rules")]
        private InputList<Inputs.TcpRouteRuleGetArgs>? _rules;

        /// <summary>
        /// List of rules to match to trigger action
        /// </summary>
        public InputList<Inputs.TcpRouteRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.TcpRouteRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Route status. Routes in "ok" state are ready to operate
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public TcpRouteState()
        {
        }
        public static new TcpRouteState Empty => new TcpRouteState();
    }
}