// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.CloudProject.inputs.KubeNodePoolTemplateMetadataArgs;
import com.pulumi.ovh.CloudProject.inputs.KubeNodePoolTemplateSpecArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeNodePoolTemplateArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubeNodePoolTemplateArgs Empty = new KubeNodePoolTemplateArgs();

    /**
     * Metadata of each node in the pool
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<KubeNodePoolTemplateMetadataArgs> metadata;

    /**
     * @return Metadata of each node in the pool
     * 
     */
    public Optional<Output<KubeNodePoolTemplateMetadataArgs>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * Spec of each node in the pool
     * 
     */
    @Import(name="spec")
    private @Nullable Output<KubeNodePoolTemplateSpecArgs> spec;

    /**
     * @return Spec of each node in the pool
     * 
     */
    public Optional<Output<KubeNodePoolTemplateSpecArgs>> spec() {
        return Optional.ofNullable(this.spec);
    }

    private KubeNodePoolTemplateArgs() {}

    private KubeNodePoolTemplateArgs(KubeNodePoolTemplateArgs $) {
        this.metadata = $.metadata;
        this.spec = $.spec;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeNodePoolTemplateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeNodePoolTemplateArgs $;

        public Builder() {
            $ = new KubeNodePoolTemplateArgs();
        }

        public Builder(KubeNodePoolTemplateArgs defaults) {
            $ = new KubeNodePoolTemplateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param metadata Metadata of each node in the pool
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<KubeNodePoolTemplateMetadataArgs> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata Metadata of each node in the pool
         * 
         * @return builder
         * 
         */
        public Builder metadata(KubeNodePoolTemplateMetadataArgs metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param spec Spec of each node in the pool
         * 
         * @return builder
         * 
         */
        public Builder spec(@Nullable Output<KubeNodePoolTemplateSpecArgs> spec) {
            $.spec = spec;
            return this;
        }

        /**
         * @param spec Spec of each node in the pool
         * 
         * @return builder
         * 
         */
        public Builder spec(KubeNodePoolTemplateSpecArgs spec) {
            return spec(Output.of(spec));
        }

        public KubeNodePoolTemplateArgs build() {
            return $;
        }
    }

}
