// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class KubeCustomizationApiserverAdmissionplugins {
    private @Nullable List<String> disableds;
    private @Nullable List<String> enableds;

    private KubeCustomizationApiserverAdmissionplugins() {}
    public List<String> disableds() {
        return this.disableds == null ? List.of() : this.disableds;
    }
    public List<String> enableds() {
        return this.enableds == null ? List.of() : this.enableds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubeCustomizationApiserverAdmissionplugins defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> disableds;
        private @Nullable List<String> enableds;
        public Builder() {}
        public Builder(KubeCustomizationApiserverAdmissionplugins defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.disableds = defaults.disableds;
    	      this.enableds = defaults.enableds;
        }

        @CustomType.Setter
        public Builder disableds(@Nullable List<String> disableds) {
            this.disableds = disableds;
            return this;
        }
        public Builder disableds(String... disableds) {
            return disableds(List.of(disableds));
        }
        @CustomType.Setter
        public Builder enableds(@Nullable List<String> enableds) {
            this.enableds = enableds;
            return this;
        }
        public Builder enableds(String... enableds) {
            return enableds(List.of(enableds));
        }
        public KubeCustomizationApiserverAdmissionplugins build() {
            final var o = new KubeCustomizationApiserverAdmissionplugins();
            o.disableds = disableds;
            o.enableds = enableds;
            return o;
        }
    }
}