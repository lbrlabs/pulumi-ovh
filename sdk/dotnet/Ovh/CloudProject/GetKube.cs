// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProject
{
    public static class GetKube
    {
        /// <summary>
        /// Use this data source to get a OVHcloud Managed Kubernetes Service cluster.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKubeCluster = Ovh.CloudProject.GetKube.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["version"] = myKubeCluster.Apply(getKubeResult =&gt; getKubeResult.Version),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetKubeResult> InvokeAsync(GetKubeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubeResult>("ovh:CloudProject/getKube:getKube", args ?? new GetKubeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a OVHcloud Managed Kubernetes Service cluster.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKubeCluster = Ovh.CloudProject.GetKube.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["version"] = myKubeCluster.Apply(getKubeResult =&gt; getKubeResult.Version),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetKubeResult> Invoke(GetKubeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubeResult>("ovh:CloudProject/getKube:getKube", args ?? new GetKubeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Customer customization object
        /// * apiserver - Kubernetes API server customization
        /// * admissionplugins - Kubernetes API server admission plugins customization
        /// * enabled - Array of admission plugins enabled, default is ["NodeRestriction","AlwaysPulImages"] and only these admission plugins can be enabled at this time.
        /// * disabled - Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.
        /// </summary>
        [Input("customization")]
        public Inputs.GetKubeCustomizationArgs? Customization { get; set; }

        /// <summary>
        /// The id of the managed kubernetes cluster.
        /// </summary>
        [Input("kubeId", required: true)]
        public string KubeId { get; set; } = null!;

        /// <summary>
        /// The name of the managed kubernetes cluster.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The OVHcloud public cloud region ID of the managed kubernetes cluster.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
        /// </summary>
        [Input("updatePolicy")]
        public string? UpdatePolicy { get; set; }

        /// <summary>
        /// Kubernetes version of the managed kubernetes cluster.
        /// </summary>
        [Input("version")]
        public string? Version { get; set; }

        public GetKubeArgs()
        {
        }
        public static new GetKubeArgs Empty => new GetKubeArgs();
    }

    public sealed class GetKubeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Customer customization object
        /// * apiserver - Kubernetes API server customization
        /// * admissionplugins - Kubernetes API server admission plugins customization
        /// * enabled - Array of admission plugins enabled, default is ["NodeRestriction","AlwaysPulImages"] and only these admission plugins can be enabled at this time.
        /// * disabled - Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.
        /// </summary>
        [Input("customization")]
        public Input<Inputs.GetKubeCustomizationInputArgs>? Customization { get; set; }

        /// <summary>
        /// The id of the managed kubernetes cluster.
        /// </summary>
        [Input("kubeId", required: true)]
        public Input<string> KubeId { get; set; } = null!;

        /// <summary>
        /// The name of the managed kubernetes cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The OVHcloud public cloud region ID of the managed kubernetes cluster.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
        /// </summary>
        [Input("updatePolicy")]
        public Input<string>? UpdatePolicy { get; set; }

        /// <summary>
        /// Kubernetes version of the managed kubernetes cluster.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public GetKubeInvokeArgs()
        {
        }
        public static new GetKubeInvokeArgs Empty => new GetKubeInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubeResult
    {
        /// <summary>
        /// True if control-plane is up to date.
        /// </summary>
        public readonly bool ControlPlaneIsUpToDate;
        /// <summary>
        /// Customer customization object
        /// * apiserver - Kubernetes API server customization
        /// * admissionplugins - Kubernetes API server admission plugins customization
        /// * enabled - Array of admission plugins enabled, default is ["NodeRestriction","AlwaysPulImages"] and only these admission plugins can be enabled at this time.
        /// * disabled - Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.
        /// </summary>
        public readonly Outputs.GetKubeCustomizationResult Customization;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// True if all nodes and control-plane are up to date.
        /// </summary>
        public readonly bool IsUpToDate;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string KubeId;
        /// <summary>
        /// The name of the managed kubernetes cluster.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Kubernetes versions available for upgrade.
        /// </summary>
        public readonly ImmutableArray<string> NextUpgradeVersions;
        /// <summary>
        /// Cluster nodes URL.
        /// </summary>
        public readonly string NodesUrl;
        /// <summary>
        /// OpenStack private network (or vrack) ID to use.
        /// </summary>
        public readonly string PrivateNetworkId;
        /// <summary>
        /// The OVHcloud public cloud region ID of the managed kubernetes cluster.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Cluster status. Should be normally set to 'READY'.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
        /// </summary>
        public readonly string? UpdatePolicy;
        /// <summary>
        /// Management URL of your cluster.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Kubernetes version of the managed kubernetes cluster.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetKubeResult(
            bool controlPlaneIsUpToDate,

            Outputs.GetKubeCustomizationResult customization,

            string id,

            bool isUpToDate,

            string kubeId,

            string? name,

            ImmutableArray<string> nextUpgradeVersions,

            string nodesUrl,

            string privateNetworkId,

            string? region,

            string serviceName,

            string status,

            string? updatePolicy,

            string url,

            string? version)
        {
            ControlPlaneIsUpToDate = controlPlaneIsUpToDate;
            Customization = customization;
            Id = id;
            IsUpToDate = isUpToDate;
            KubeId = kubeId;
            Name = name;
            NextUpgradeVersions = nextUpgradeVersions;
            NodesUrl = nodesUrl;
            PrivateNetworkId = privateNetworkId;
            Region = region;
            ServiceName = serviceName;
            Status = status;
            UpdatePolicy = updatePolicy;
            Url = url;
            Version = version;
        }
    }
}
