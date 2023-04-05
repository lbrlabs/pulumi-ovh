// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Order
{
    public static class GetCartProductOptions
    {
        /// <summary>
        /// Use this data source to retrieve information of order cart product options.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mycart = Ovh.Order.GetCart.Invoke(new()
        ///     {
        ///         OvhSubsidiary = "fr",
        ///         Description = "my cart",
        ///     });
        /// 
        ///     var options = Ovh.Order.GetCartProductOptions.Invoke(new()
        ///     {
        ///         CartId = mycart.Apply(getCartResult =&gt; getCartResult.Id),
        ///         Product = "cloud",
        ///         PlanCode = "project",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCartProductOptionsResult> InvokeAsync(GetCartProductOptionsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCartProductOptionsResult>("ovh:Order/getCartProductOptions:getCartProductOptions", args ?? new GetCartProductOptionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information of order cart product options.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mycart = Ovh.Order.GetCart.Invoke(new()
        ///     {
        ///         OvhSubsidiary = "fr",
        ///         Description = "my cart",
        ///     });
        /// 
        ///     var options = Ovh.Order.GetCartProductOptions.Invoke(new()
        ///     {
        ///         CartId = mycart.Apply(getCartResult =&gt; getCartResult.Id),
        ///         Product = "cloud",
        ///         PlanCode = "project",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCartProductOptionsResult> Invoke(GetCartProductOptionsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCartProductOptionsResult>("ovh:Order/getCartProductOptions:getCartProductOptions", args ?? new GetCartProductOptionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCartProductOptionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cart identifier
        /// </summary>
        [Input("cartId", required: true)]
        public string CartId { get; set; } = null!;

        /// <summary>
        /// Catalog name
        /// </summary>
        [Input("catalogName")]
        public string? CatalogName { get; set; }

        /// <summary>
        /// Product offer identifier
        /// </summary>
        [Input("planCode", required: true)]
        public string PlanCode { get; set; } = null!;

        /// <summary>
        /// Product
        /// </summary>
        [Input("product", required: true)]
        public string Product { get; set; } = null!;

        public GetCartProductOptionsArgs()
        {
        }
        public static new GetCartProductOptionsArgs Empty => new GetCartProductOptionsArgs();
    }

    public sealed class GetCartProductOptionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Cart identifier
        /// </summary>
        [Input("cartId", required: true)]
        public Input<string> CartId { get; set; } = null!;

        /// <summary>
        /// Catalog name
        /// </summary>
        [Input("catalogName")]
        public Input<string>? CatalogName { get; set; }

        /// <summary>
        /// Product offer identifier
        /// </summary>
        [Input("planCode", required: true)]
        public Input<string> PlanCode { get; set; } = null!;

        /// <summary>
        /// Product
        /// </summary>
        [Input("product", required: true)]
        public Input<string> Product { get; set; } = null!;

        public GetCartProductOptionsInvokeArgs()
        {
        }
        public static new GetCartProductOptionsInvokeArgs Empty => new GetCartProductOptionsInvokeArgs();
    }


    [OutputType]
    public sealed class GetCartProductOptionsResult
    {
        public readonly string CartId;
        public readonly string? CatalogName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Product offer identifier
        /// </summary>
        public readonly string PlanCode;
        public readonly string Product;
        /// <summary>
        /// products results
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCartProductOptionsResultResult> Results;

        [OutputConstructor]
        private GetCartProductOptionsResult(
            string cartId,

            string? catalogName,

            string id,

            string planCode,

            string product,

            ImmutableArray<Outputs.GetCartProductOptionsResultResult> results)
        {
            CartId = cartId;
            CatalogName = catalogName;
            Id = id;
            PlanCode = planCode;
            Product = product;
            Results = results;
        }
    }
}
