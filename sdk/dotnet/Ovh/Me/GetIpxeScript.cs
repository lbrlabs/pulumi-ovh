// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Me
{
    public static class GetIpxeScript
    {
        /// <summary>
        /// Use this data source to retrieve information about an IPXE Script.
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
        ///     var script = Ovh.Me.GetIpxeScript.Invoke(new()
        ///     {
        ///         Name = "myscript",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetIpxeScriptResult> InvokeAsync(GetIpxeScriptArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIpxeScriptResult>("ovh:Me/getIpxeScript:getIpxeScript", args ?? new GetIpxeScriptArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to retrieve information about an IPXE Script.
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
        ///     var script = Ovh.Me.GetIpxeScript.Invoke(new()
        ///     {
        ///         Name = "myscript",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetIpxeScriptResult> Invoke(GetIpxeScriptInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIpxeScriptResult>("ovh:Me/getIpxeScript:getIpxeScript", args ?? new GetIpxeScriptInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIpxeScriptArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the IPXE Script.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetIpxeScriptArgs()
        {
        }
        public static new GetIpxeScriptArgs Empty => new GetIpxeScriptArgs();
    }

    public sealed class GetIpxeScriptInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the IPXE Script.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetIpxeScriptInvokeArgs()
        {
        }
        public static new GetIpxeScriptInvokeArgs Empty => new GetIpxeScriptInvokeArgs();
    }


    [OutputType]
    public sealed class GetIpxeScriptResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The content of the script.
        /// </summary>
        public readonly string Script;

        [OutputConstructor]
        private GetIpxeScriptResult(
            string id,

            string name,

            string script)
        {
            Id = id;
            Name = name;
            Script = script;
        }
    }
}
