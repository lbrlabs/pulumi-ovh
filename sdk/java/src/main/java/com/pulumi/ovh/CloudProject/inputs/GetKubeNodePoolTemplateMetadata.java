// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetKubeNodePoolTemplateMetadata extends com.pulumi.resources.InvokeArgs {

    public static final GetKubeNodePoolTemplateMetadata Empty = new GetKubeNodePoolTemplateMetadata();

    @Import(name="annotations")
    private @Nullable Map<String,String> annotations;

    public Optional<Map<String,String>> annotations() {
        return Optional.ofNullable(this.annotations);
    }

    @Import(name="finalizers")
    private @Nullable List<String> finalizers;

    public Optional<List<String>> finalizers() {
        return Optional.ofNullable(this.finalizers);
    }

    @Import(name="labels")
    private @Nullable Map<String,String> labels;

    public Optional<Map<String,String>> labels() {
        return Optional.ofNullable(this.labels);
    }

    private GetKubeNodePoolTemplateMetadata() {}

    private GetKubeNodePoolTemplateMetadata(GetKubeNodePoolTemplateMetadata $) {
        this.annotations = $.annotations;
        this.finalizers = $.finalizers;
        this.labels = $.labels;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubeNodePoolTemplateMetadata defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubeNodePoolTemplateMetadata $;

        public Builder() {
            $ = new GetKubeNodePoolTemplateMetadata();
        }

        public Builder(GetKubeNodePoolTemplateMetadata defaults) {
            $ = new GetKubeNodePoolTemplateMetadata(Objects.requireNonNull(defaults));
        }

        public Builder annotations(@Nullable Map<String,String> annotations) {
            $.annotations = annotations;
            return this;
        }

        public Builder finalizers(@Nullable List<String> finalizers) {
            $.finalizers = finalizers;
            return this;
        }

        public Builder finalizers(String... finalizers) {
            return finalizers(List.of(finalizers));
        }

        public Builder labels(@Nullable Map<String,String> labels) {
            $.labels = labels;
            return this;
        }

        public GetKubeNodePoolTemplateMetadata build() {
            return $;
        }
    }

}
