// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.CloudProject.inputs.GetKubeCustomization;
import com.pulumi.ovh.CloudProject.inputs.GetKubeCustomizationApiserver;
import com.pulumi.ovh.CloudProject.inputs.GetKubeCustomizationKubeProxy;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetKubePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetKubePlainArgs Empty = new GetKubePlainArgs();

    /**
     * Kubernetes API server customization
     * 
     */
    @Import(name="customizationApiservers")
    private @Nullable List<GetKubeCustomizationApiserver> customizationApiservers;

    /**
     * @return Kubernetes API server customization
     * 
     */
    public Optional<List<GetKubeCustomizationApiserver>> customizationApiservers() {
        return Optional.ofNullable(this.customizationApiservers);
    }

    /**
     * Kubernetes kube-proxy customization
     * 
     */
    @Import(name="customizationKubeProxy")
    private @Nullable GetKubeCustomizationKubeProxy customizationKubeProxy;

    /**
     * @return Kubernetes kube-proxy customization
     * 
     */
    public Optional<GetKubeCustomizationKubeProxy> customizationKubeProxy() {
        return Optional.ofNullable(this.customizationKubeProxy);
    }

    /**
     * **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    @Import(name="customizations")
    private @Nullable List<GetKubeCustomization> customizations;

    /**
     * @return **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
     * 
     * @deprecated
     * Use customization_apiserver instead
     * 
     */
    @Deprecated /* Use customization_apiserver instead */
    public Optional<List<GetKubeCustomization>> customizations() {
        return Optional.ofNullable(this.customizations);
    }

    /**
     * The id of the managed kubernetes cluster.
     * 
     */
    @Import(name="kubeId", required=true)
    private String kubeId;

    /**
     * @return The id of the managed kubernetes cluster.
     * 
     */
    public String kubeId() {
        return this.kubeId;
    }

    /**
     * Selected mode for kube-proxy.
     * 
     */
    @Import(name="kubeProxyMode")
    private @Nullable String kubeProxyMode;

    /**
     * @return Selected mode for kube-proxy.
     * 
     */
    public Optional<String> kubeProxyMode() {
        return Optional.ofNullable(this.kubeProxyMode);
    }

    /**
     * The name of the managed kubernetes cluster.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the managed kubernetes cluster.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The OVHcloud public cloud region ID of the managed kubernetes cluster.
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return The OVHcloud public cloud region ID of the managed kubernetes cluster.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    /**
     * Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]&#39;.
     * 
     */
    @Import(name="updatePolicy")
    private @Nullable String updatePolicy;

    /**
     * @return Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]&#39;.
     * 
     */
    public Optional<String> updatePolicy() {
        return Optional.ofNullable(this.updatePolicy);
    }

    /**
     * Kubernetes version of the managed kubernetes cluster.
     * 
     */
    @Import(name="version")
    private @Nullable String version;

    /**
     * @return Kubernetes version of the managed kubernetes cluster.
     * 
     */
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    private GetKubePlainArgs() {}

    private GetKubePlainArgs(GetKubePlainArgs $) {
        this.customizationApiservers = $.customizationApiservers;
        this.customizationKubeProxy = $.customizationKubeProxy;
        this.customizations = $.customizations;
        this.kubeId = $.kubeId;
        this.kubeProxyMode = $.kubeProxyMode;
        this.name = $.name;
        this.region = $.region;
        this.serviceName = $.serviceName;
        this.updatePolicy = $.updatePolicy;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubePlainArgs $;

        public Builder() {
            $ = new GetKubePlainArgs();
        }

        public Builder(GetKubePlainArgs defaults) {
            $ = new GetKubePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param customizationApiservers Kubernetes API server customization
         * 
         * @return builder
         * 
         */
        public Builder customizationApiservers(@Nullable List<GetKubeCustomizationApiserver> customizationApiservers) {
            $.customizationApiservers = customizationApiservers;
            return this;
        }

        /**
         * @param customizationApiservers Kubernetes API server customization
         * 
         * @return builder
         * 
         */
        public Builder customizationApiservers(GetKubeCustomizationApiserver... customizationApiservers) {
            return customizationApiservers(List.of(customizationApiservers));
        }

        /**
         * @param customizationKubeProxy Kubernetes kube-proxy customization
         * 
         * @return builder
         * 
         */
        public Builder customizationKubeProxy(@Nullable GetKubeCustomizationKubeProxy customizationKubeProxy) {
            $.customizationKubeProxy = customizationKubeProxy;
            return this;
        }

        /**
         * @param customizations **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
         * 
         * @return builder
         * 
         * @deprecated
         * Use customization_apiserver instead
         * 
         */
        @Deprecated /* Use customization_apiserver instead */
        public Builder customizations(@Nullable List<GetKubeCustomization> customizations) {
            $.customizations = customizations;
            return this;
        }

        /**
         * @param customizations **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
         * 
         * @return builder
         * 
         * @deprecated
         * Use customization_apiserver instead
         * 
         */
        @Deprecated /* Use customization_apiserver instead */
        public Builder customizations(GetKubeCustomization... customizations) {
            return customizations(List.of(customizations));
        }

        /**
         * @param kubeId The id of the managed kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder kubeId(String kubeId) {
            $.kubeId = kubeId;
            return this;
        }

        /**
         * @param kubeProxyMode Selected mode for kube-proxy.
         * 
         * @return builder
         * 
         */
        public Builder kubeProxyMode(@Nullable String kubeProxyMode) {
            $.kubeProxyMode = kubeProxyMode;
            return this;
        }

        /**
         * @param name The name of the managed kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param region The OVHcloud public cloud region ID of the managed kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param updatePolicy Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]&#39;.
         * 
         * @return builder
         * 
         */
        public Builder updatePolicy(@Nullable String updatePolicy) {
            $.updatePolicy = updatePolicy;
            return this;
        }

        /**
         * @param version Kubernetes version of the managed kubernetes cluster.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable String version) {
            $.version = version;
            return this;
        }

        public GetKubePlainArgs build() {
            $.kubeId = Objects.requireNonNull($.kubeId, "expected parameter 'kubeId' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
