// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Ovh
{
    /// <summary>
    /// Orders an IP load balancing.
    /// 
    /// ## Important
    /// 
    /// This resource orders an OVH product for a long period of time and may generate heavy costs !
    /// Use with caution.
    /// 
    /// __NOTE__ 1: the "default-payment-mean" will scan your registered bank accounts, credit card and paypal payment means to find your default payment mean.
    /// 
    /// __NOTE__ 2: this resource is in beta state. Use with caution.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Ovh = Lbrlabs_Pulumi.Ovh;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mycart = Ovh.GetOrderCart.Invoke(new()
    ///     {
    ///         OvhSubsidiary = "fr",
    ///         Description = "mycart",
    ///     });
    /// 
    ///     var iplb = Ovh.GetOrderCartProductPlan.Invoke(new()
    ///     {
    ///         CartId = mycart.Apply(getOrderCartResult =&gt; getOrderCartResult.Id),
    ///         PriceCapacity = "renew",
    ///         Product = "ipLoadbalancing",
    ///         PlanCode = "iplb-lb1",
    ///     });
    /// 
    ///     var bhs = Ovh.GetOrderCartProductOptionsPlan.Invoke(new()
    ///     {
    ///         CartId = iplb.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.CartId),
    ///         PriceCapacity = iplb.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.PriceCapacity),
    ///         Product = iplb.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.Product),
    ///         PlanCode = iplb.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.PlanCode),
    ///         OptionsPlanCode = "iplb-zone-lb1-rbx",
    ///     });
    /// 
    ///     var iplb_lb1 = new Ovh.IpLoadBalancing("iplb-lb1", new()
    ///     {
    ///         OvhSubsidiary = mycart.Apply(getOrderCartResult =&gt; getOrderCartResult.OvhSubsidiary),
    ///         DisplayName = "my ip loadbalancing",
    ///         PaymentMean = "ovh-account",
    ///         Plan = new Ovh.Inputs.IpLoadBalancingPlanArgs
    ///         {
    ///             Duration = iplb.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.SelectedPrices[0]?.Duration),
    ///             PlanCode = iplb.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.PlanCode),
    ///             PricingMode = iplb.Apply(getOrderCartProductPlanResult =&gt; getOrderCartProductPlanResult.SelectedPrices[0]?.PricingMode),
    ///         },
    ///         PlanOptions = new[]
    ///         {
    ///             new Ovh.Inputs.IpLoadBalancingPlanOptionArgs
    ///             {
    ///                 Duration = bhs.Apply(getOrderCartProductOptionsPlanResult =&gt; getOrderCartProductOptionsPlanResult.SelectedPrices[0]?.Duration),
    ///                 PlanCode = bhs.Apply(getOrderCartProductOptionsPlanResult =&gt; getOrderCartProductOptionsPlanResult.PlanCode),
    ///                 PricingMode = bhs.Apply(getOrderCartProductOptionsPlanResult =&gt; getOrderCartProductOptionsPlanResult.SelectedPrices[0]?.PricingMode),
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:index/ipLoadBalancing:IpLoadBalancing")]
    public partial class IpLoadBalancing : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Your IP load balancing
        /// </summary>
        [Output("ipLoadbalancing")]
        public Output<string> IpLoadbalancing { get; private set; } = null!;

        /// <summary>
        /// The IPV4 associated to your IP load balancing
        /// </summary>
        [Output("ipv4")]
        public Output<string> Ipv4 { get; private set; } = null!;

        /// <summary>
        /// The IPV6 associated to your IP load balancing. DEPRECATED.
        /// </summary>
        [Output("ipv6")]
        public Output<string> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// The metrics token associated with your IP load balancing
        /// </summary>
        [Output("metricsToken")]
        public Output<string> MetricsToken { get; private set; } = null!;

        /// <summary>
        /// The offer of your IP load balancing
        /// </summary>
        [Output("offer")]
        public Output<string> Offer { get; private set; } = null!;

        /// <summary>
        /// Available additional zone for your Load Balancer
        /// </summary>
        [Output("orderableZones")]
        public Output<ImmutableArray<Outputs.IpLoadBalancingOrderableZone>> OrderableZones { get; private set; } = null!;

        /// <summary>
        /// Details about an Order
        /// </summary>
        [Output("orders")]
        public Output<ImmutableArray<Outputs.IpLoadBalancingOrder>> Orders { get; private set; } = null!;

        /// <summary>
        /// Ovh Subsidiary
        /// </summary>
        [Output("ovhSubsidiary")]
        public Output<string> OvhSubsidiary { get; private set; } = null!;

        /// <summary>
        /// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
        /// </summary>
        [Output("paymentMean")]
        public Output<string> PaymentMean { get; private set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Output("plan")]
        public Output<Outputs.IpLoadBalancingPlan> Plan { get; private set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Output("planOptions")]
        public Output<ImmutableArray<Outputs.IpLoadBalancingPlanOption>> PlanOptions { get; private set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
        /// </summary>
        [Output("sslConfiguration")]
        public Output<string> SslConfiguration { get; private set; } = null!;

        /// <summary>
        /// Current state of your IP
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Vrack eligibility
        /// </summary>
        [Output("vrackEligibility")]
        public Output<bool> VrackEligibility { get; private set; } = null!;

        /// <summary>
        /// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
        /// </summary>
        [Output("vrackName")]
        public Output<string> VrackName { get; private set; } = null!;

        /// <summary>
        /// Location where your service is
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a IpLoadBalancing resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpLoadBalancing(string name, IpLoadBalancingArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/ipLoadBalancing:IpLoadBalancing", name, args ?? new IpLoadBalancingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpLoadBalancing(string name, Input<string> id, IpLoadBalancingState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/ipLoadBalancing:IpLoadBalancing", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpLoadBalancing resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpLoadBalancing Get(string name, Input<string> id, IpLoadBalancingState? state = null, CustomResourceOptions? options = null)
        {
            return new IpLoadBalancing(name, id, state, options);
        }
    }

    public sealed class IpLoadBalancingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Ovh Subsidiary
        /// </summary>
        [Input("ovhSubsidiary", required: true)]
        public Input<string> OvhSubsidiary { get; set; } = null!;

        /// <summary>
        /// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
        /// </summary>
        [Input("paymentMean", required: true)]
        public Input<string> PaymentMean { get; set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Input("plan", required: true)]
        public Input<Inputs.IpLoadBalancingPlanArgs> Plan { get; set; } = null!;

        [Input("planOptions")]
        private InputList<Inputs.IpLoadBalancingPlanOptionArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.IpLoadBalancingPlanOptionArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.IpLoadBalancingPlanOptionArgs>());
            set => _planOptions = value;
        }

        /// <summary>
        /// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
        /// </summary>
        [Input("sslConfiguration")]
        public Input<string>? SslConfiguration { get; set; }

        public IpLoadBalancingArgs()
        {
        }
        public static new IpLoadBalancingArgs Empty => new IpLoadBalancingArgs();
    }

    public sealed class IpLoadBalancingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Your IP load balancing
        /// </summary>
        [Input("ipLoadbalancing")]
        public Input<string>? IpLoadbalancing { get; set; }

        /// <summary>
        /// The IPV4 associated to your IP load balancing
        /// </summary>
        [Input("ipv4")]
        public Input<string>? Ipv4 { get; set; }

        /// <summary>
        /// The IPV6 associated to your IP load balancing. DEPRECATED.
        /// </summary>
        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        /// <summary>
        /// The metrics token associated with your IP load balancing
        /// </summary>
        [Input("metricsToken")]
        public Input<string>? MetricsToken { get; set; }

        /// <summary>
        /// The offer of your IP load balancing
        /// </summary>
        [Input("offer")]
        public Input<string>? Offer { get; set; }

        [Input("orderableZones")]
        private InputList<Inputs.IpLoadBalancingOrderableZoneGetArgs>? _orderableZones;

        /// <summary>
        /// Available additional zone for your Load Balancer
        /// </summary>
        public InputList<Inputs.IpLoadBalancingOrderableZoneGetArgs> OrderableZones
        {
            get => _orderableZones ?? (_orderableZones = new InputList<Inputs.IpLoadBalancingOrderableZoneGetArgs>());
            set => _orderableZones = value;
        }

        [Input("orders")]
        private InputList<Inputs.IpLoadBalancingOrderGetArgs>? _orders;

        /// <summary>
        /// Details about an Order
        /// </summary>
        public InputList<Inputs.IpLoadBalancingOrderGetArgs> Orders
        {
            get => _orders ?? (_orders = new InputList<Inputs.IpLoadBalancingOrderGetArgs>());
            set => _orders = value;
        }

        /// <summary>
        /// Ovh Subsidiary
        /// </summary>
        [Input("ovhSubsidiary")]
        public Input<string>? OvhSubsidiary { get; set; }

        /// <summary>
        /// Ovh payment mode (One of "default-payment-mean", "fidelity", "ovh-account")
        /// </summary>
        [Input("paymentMean")]
        public Input<string>? PaymentMean { get; set; }

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Input("plan")]
        public Input<Inputs.IpLoadBalancingPlanGetArgs>? Plan { get; set; }

        [Input("planOptions")]
        private InputList<Inputs.IpLoadBalancingPlanOptionGetArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.IpLoadBalancingPlanOptionGetArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.IpLoadBalancingPlanOptionGetArgs>());
            set => _planOptions = value;
        }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
        /// </summary>
        [Input("sslConfiguration")]
        public Input<string>? SslConfiguration { get; set; }

        /// <summary>
        /// Current state of your IP
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Vrack eligibility
        /// </summary>
        [Input("vrackEligibility")]
        public Input<bool>? VrackEligibility { get; set; }

        /// <summary>
        /// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
        /// </summary>
        [Input("vrackName")]
        public Input<string>? VrackName { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// Location where your service is
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public IpLoadBalancingState()
        {
        }
        public static new IpLoadBalancingState Empty => new IpLoadBalancingState();
    }
}
