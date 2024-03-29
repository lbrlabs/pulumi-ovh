// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KubeNodePoolTemplateSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final KubeNodePoolTemplateSpecArgs Empty = new KubeNodePoolTemplateSpecArgs();

    /**
     * Taints to apply to each node
     * 
     */
    @Import(name="taints")
    private @Nullable Output<List<Map<String,Object>>> taints;

    /**
     * @return Taints to apply to each node
     * 
     */
    public Optional<Output<List<Map<String,Object>>>> taints() {
        return Optional.ofNullable(this.taints);
    }

    /**
     * If true, set nodes as un-schedulable
     * 
     */
    @Import(name="unschedulable")
    private @Nullable Output<Boolean> unschedulable;

    /**
     * @return If true, set nodes as un-schedulable
     * 
     */
    public Optional<Output<Boolean>> unschedulable() {
        return Optional.ofNullable(this.unschedulable);
    }

    private KubeNodePoolTemplateSpecArgs() {}

    private KubeNodePoolTemplateSpecArgs(KubeNodePoolTemplateSpecArgs $) {
        this.taints = $.taints;
        this.unschedulable = $.unschedulable;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KubeNodePoolTemplateSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KubeNodePoolTemplateSpecArgs $;

        public Builder() {
            $ = new KubeNodePoolTemplateSpecArgs();
        }

        public Builder(KubeNodePoolTemplateSpecArgs defaults) {
            $ = new KubeNodePoolTemplateSpecArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param taints Taints to apply to each node
         * 
         * @return builder
         * 
         */
        public Builder taints(@Nullable Output<List<Map<String,Object>>> taints) {
            $.taints = taints;
            return this;
        }

        /**
         * @param taints Taints to apply to each node
         * 
         * @return builder
         * 
         */
        public Builder taints(List<Map<String,Object>> taints) {
            return taints(Output.of(taints));
        }

        /**
         * @param taints Taints to apply to each node
         * 
         * @return builder
         * 
         */
        public Builder taints(Map<String,Object>... taints) {
            return taints(List.of(taints));
        }

        /**
         * @param unschedulable If true, set nodes as un-schedulable
         * 
         * @return builder
         * 
         */
        public Builder unschedulable(@Nullable Output<Boolean> unschedulable) {
            $.unschedulable = unschedulable;
            return this;
        }

        /**
         * @param unschedulable If true, set nodes as un-schedulable
         * 
         * @return builder
         * 
         */
        public Builder unschedulable(Boolean unschedulable) {
            return unschedulable(Output.of(unschedulable));
        }

        public KubeNodePoolTemplateSpecArgs build() {
            return $;
        }
    }

}
