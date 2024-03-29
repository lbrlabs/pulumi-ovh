// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.CloudProject.outputs.GetKubeCustomization;
import com.pulumi.ovh.CloudProject.outputs.GetKubeCustomizationApiserver;
import com.pulumi.ovh.CloudProject.outputs.GetKubeCustomizationKubeProxy;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetKubeResult {
    /**
     * @return True if control-plane is up-to-date.
     * 
     */
    private Boolean controlPlaneIsUpToDate;
    /**
     * @return Kubernetes API server customization
     * 
     */
    private List<GetKubeCustomizationApiserver> customizationApiservers;
    /**
     * @return Kubernetes kube-proxy customization
     * 
     */
    private @Nullable GetKubeCustomizationKubeProxy customizationKubeProxy;
    /**
     * @return **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    private List<GetKubeCustomization> customizations;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return True if all nodes and control-plane are up-to-date.
     * 
     */
    private Boolean isUpToDate;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String kubeId;
    /**
     * @return Selected mode for kube-proxy.
     * 
     */
    private @Nullable String kubeProxyMode;
    /**
     * @return The name of the managed kubernetes cluster.
     * 
     */
    private @Nullable String name;
    /**
     * @return Kubernetes versions available for upgrade.
     * 
     */
    private List<String> nextUpgradeVersions;
    /**
     * @return Cluster nodes URL.
     * 
     */
    private String nodesUrl;
    /**
     * @return OpenStack private network (or vrack) ID to use.
     * 
     */
    private String privateNetworkId;
    /**
     * @return The OVHcloud public cloud region ID of the managed kubernetes cluster.
     * 
     */
    private @Nullable String region;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;
    /**
     * @return Cluster status. Should be normally set to &#39;READY&#39;.
     * 
     */
    private String status;
    /**
     * @return Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]&#39;.
     * 
     */
    private @Nullable String updatePolicy;
    /**
     * @return Management URL of your cluster.
     * 
     */
    private String url;
    /**
     * @return Kubernetes version of the managed kubernetes cluster.
     * 
     */
    private @Nullable String version;

    private GetKubeResult() {}
    /**
     * @return True if control-plane is up-to-date.
     * 
     */
    public Boolean controlPlaneIsUpToDate() {
        return this.controlPlaneIsUpToDate;
    }
    /**
     * @return Kubernetes API server customization
     * 
     */
    public List<GetKubeCustomizationApiserver> customizationApiservers() {
        return this.customizationApiservers;
    }
    /**
     * @return Kubernetes kube-proxy customization
     * 
     */
    public Optional<GetKubeCustomizationKubeProxy> customizationKubeProxy() {
        return Optional.ofNullable(this.customizationKubeProxy);
    }
    /**
     * @return **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    public List<GetKubeCustomization> customizations() {
        return this.customizations;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return True if all nodes and control-plane are up-to-date.
     * 
     */
    public Boolean isUpToDate() {
        return this.isUpToDate;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String kubeId() {
        return this.kubeId;
    }
    /**
     * @return Selected mode for kube-proxy.
     * 
     */
    public Optional<String> kubeProxyMode() {
        return Optional.ofNullable(this.kubeProxyMode);
    }
    /**
     * @return The name of the managed kubernetes cluster.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Kubernetes versions available for upgrade.
     * 
     */
    public List<String> nextUpgradeVersions() {
        return this.nextUpgradeVersions;
    }
    /**
     * @return Cluster nodes URL.
     * 
     */
    public String nodesUrl() {
        return this.nodesUrl;
    }
    /**
     * @return OpenStack private network (or vrack) ID to use.
     * 
     */
    public String privateNetworkId() {
        return this.privateNetworkId;
    }
    /**
     * @return The OVHcloud public cloud region ID of the managed kubernetes cluster.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Cluster status. Should be normally set to &#39;READY&#39;.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]&#39;.
     * 
     */
    public Optional<String> updatePolicy() {
        return Optional.ofNullable(this.updatePolicy);
    }
    /**
     * @return Management URL of your cluster.
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return Kubernetes version of the managed kubernetes cluster.
     * 
     */
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean controlPlaneIsUpToDate;
        private List<GetKubeCustomizationApiserver> customizationApiservers;
        private @Nullable GetKubeCustomizationKubeProxy customizationKubeProxy;
        private List<GetKubeCustomization> customizations;
        private String id;
        private Boolean isUpToDate;
        private String kubeId;
        private @Nullable String kubeProxyMode;
        private @Nullable String name;
        private List<String> nextUpgradeVersions;
        private String nodesUrl;
        private String privateNetworkId;
        private @Nullable String region;
        private String serviceName;
        private String status;
        private @Nullable String updatePolicy;
        private String url;
        private @Nullable String version;
        public Builder() {}
        public Builder(GetKubeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.controlPlaneIsUpToDate = defaults.controlPlaneIsUpToDate;
    	      this.customizationApiservers = defaults.customizationApiservers;
    	      this.customizationKubeProxy = defaults.customizationKubeProxy;
    	      this.customizations = defaults.customizations;
    	      this.id = defaults.id;
    	      this.isUpToDate = defaults.isUpToDate;
    	      this.kubeId = defaults.kubeId;
    	      this.kubeProxyMode = defaults.kubeProxyMode;
    	      this.name = defaults.name;
    	      this.nextUpgradeVersions = defaults.nextUpgradeVersions;
    	      this.nodesUrl = defaults.nodesUrl;
    	      this.privateNetworkId = defaults.privateNetworkId;
    	      this.region = defaults.region;
    	      this.serviceName = defaults.serviceName;
    	      this.status = defaults.status;
    	      this.updatePolicy = defaults.updatePolicy;
    	      this.url = defaults.url;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder controlPlaneIsUpToDate(Boolean controlPlaneIsUpToDate) {
            this.controlPlaneIsUpToDate = Objects.requireNonNull(controlPlaneIsUpToDate);
            return this;
        }
        @CustomType.Setter
        public Builder customizationApiservers(List<GetKubeCustomizationApiserver> customizationApiservers) {
            this.customizationApiservers = Objects.requireNonNull(customizationApiservers);
            return this;
        }
        public Builder customizationApiservers(GetKubeCustomizationApiserver... customizationApiservers) {
            return customizationApiservers(List.of(customizationApiservers));
        }
        @CustomType.Setter
        public Builder customizationKubeProxy(@Nullable GetKubeCustomizationKubeProxy customizationKubeProxy) {
            this.customizationKubeProxy = customizationKubeProxy;
            return this;
        }
        @CustomType.Setter
        public Builder customizations(List<GetKubeCustomization> customizations) {
            this.customizations = Objects.requireNonNull(customizations);
            return this;
        }
        public Builder customizations(GetKubeCustomization... customizations) {
            return customizations(List.of(customizations));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder isUpToDate(Boolean isUpToDate) {
            this.isUpToDate = Objects.requireNonNull(isUpToDate);
            return this;
        }
        @CustomType.Setter
        public Builder kubeId(String kubeId) {
            this.kubeId = Objects.requireNonNull(kubeId);
            return this;
        }
        @CustomType.Setter
        public Builder kubeProxyMode(@Nullable String kubeProxyMode) {
            this.kubeProxyMode = kubeProxyMode;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder nextUpgradeVersions(List<String> nextUpgradeVersions) {
            this.nextUpgradeVersions = Objects.requireNonNull(nextUpgradeVersions);
            return this;
        }
        public Builder nextUpgradeVersions(String... nextUpgradeVersions) {
            return nextUpgradeVersions(List.of(nextUpgradeVersions));
        }
        @CustomType.Setter
        public Builder nodesUrl(String nodesUrl) {
            this.nodesUrl = Objects.requireNonNull(nodesUrl);
            return this;
        }
        @CustomType.Setter
        public Builder privateNetworkId(String privateNetworkId) {
            this.privateNetworkId = Objects.requireNonNull(privateNetworkId);
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            this.status = Objects.requireNonNull(status);
            return this;
        }
        @CustomType.Setter
        public Builder updatePolicy(@Nullable String updatePolicy) {
            this.updatePolicy = updatePolicy;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        @CustomType.Setter
        public Builder version(@Nullable String version) {
            this.version = version;
            return this;
        }
        public GetKubeResult build() {
            final var o = new GetKubeResult();
            o.controlPlaneIsUpToDate = controlPlaneIsUpToDate;
            o.customizationApiservers = customizationApiservers;
            o.customizationKubeProxy = customizationKubeProxy;
            o.customizations = customizations;
            o.id = id;
            o.isUpToDate = isUpToDate;
            o.kubeId = kubeId;
            o.kubeProxyMode = kubeProxyMode;
            o.name = name;
            o.nextUpgradeVersions = nextUpgradeVersions;
            o.nodesUrl = nodesUrl;
            o.privateNetworkId = privateNetworkId;
            o.region = region;
            o.serviceName = serviceName;
            o.status = status;
            o.updatePolicy = updatePolicy;
            o.url = url;
            o.version = version;
            return o;
        }
    }
}
