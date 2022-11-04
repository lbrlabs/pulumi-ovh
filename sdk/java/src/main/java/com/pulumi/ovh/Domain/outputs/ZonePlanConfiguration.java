// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Domain.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ZonePlanConfiguration {
    /**
     * @return Identifier of the resource
     * 
     */
    private String label;
    /**
     * @return Path to the resource in API.OVH.COM
     * 
     */
    private String value;

    private ZonePlanConfiguration() {}
    /**
     * @return Identifier of the resource
     * 
     */
    public String label() {
        return this.label;
    }
    /**
     * @return Path to the resource in API.OVH.COM
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ZonePlanConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String label;
        private String value;
        public Builder() {}
        public Builder(ZonePlanConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.label = defaults.label;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder label(String label) {
            this.label = Objects.requireNonNull(label);
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }
        public ZonePlanConfiguration build() {
            final var o = new ZonePlanConfiguration();
            o.label = label;
            o.value = value;
            return o;
        }
    }
}