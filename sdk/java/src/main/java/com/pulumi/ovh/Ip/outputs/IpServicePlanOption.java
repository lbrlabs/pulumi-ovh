// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Ip.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.Ip.outputs.IpServicePlanOptionConfiguration;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IpServicePlanOption {
    /**
     * @return Catalog name
     * 
     */
    private @Nullable String catalogName;
    /**
     * @return Representation of a configuration item for personalizing product
     * 
     */
    private @Nullable List<IpServicePlanOptionConfiguration> configurations;
    /**
     * @return duration
     * 
     */
    private String duration;
    /**
     * @return Plan code
     * 
     */
    private String planCode;
    /**
     * @return Pricing model identifier
     * 
     */
    private String pricingMode;

    private IpServicePlanOption() {}
    /**
     * @return Catalog name
     * 
     */
    public Optional<String> catalogName() {
        return Optional.ofNullable(this.catalogName);
    }
    /**
     * @return Representation of a configuration item for personalizing product
     * 
     */
    public List<IpServicePlanOptionConfiguration> configurations() {
        return this.configurations == null ? List.of() : this.configurations;
    }
    /**
     * @return duration
     * 
     */
    public String duration() {
        return this.duration;
    }
    /**
     * @return Plan code
     * 
     */
    public String planCode() {
        return this.planCode;
    }
    /**
     * @return Pricing model identifier
     * 
     */
    public String pricingMode() {
        return this.pricingMode;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IpServicePlanOption defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String catalogName;
        private @Nullable List<IpServicePlanOptionConfiguration> configurations;
        private String duration;
        private String planCode;
        private String pricingMode;
        public Builder() {}
        public Builder(IpServicePlanOption defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.catalogName = defaults.catalogName;
    	      this.configurations = defaults.configurations;
    	      this.duration = defaults.duration;
    	      this.planCode = defaults.planCode;
    	      this.pricingMode = defaults.pricingMode;
        }

        @CustomType.Setter
        public Builder catalogName(@Nullable String catalogName) {
            this.catalogName = catalogName;
            return this;
        }
        @CustomType.Setter
        public Builder configurations(@Nullable List<IpServicePlanOptionConfiguration> configurations) {
            this.configurations = configurations;
            return this;
        }
        public Builder configurations(IpServicePlanOptionConfiguration... configurations) {
            return configurations(List.of(configurations));
        }
        @CustomType.Setter
        public Builder duration(String duration) {
            this.duration = Objects.requireNonNull(duration);
            return this;
        }
        @CustomType.Setter
        public Builder planCode(String planCode) {
            this.planCode = Objects.requireNonNull(planCode);
            return this;
        }
        @CustomType.Setter
        public Builder pricingMode(String pricingMode) {
            this.pricingMode = Objects.requireNonNull(pricingMode);
            return this;
        }
        public IpServicePlanOption build() {
            final var o = new IpServicePlanOption();
            o.catalogName = catalogName;
            o.configurations = configurations;
            o.duration = duration;
            o.planCode = planCode;
            o.pricingMode = pricingMode;
            return o;
        }
    }
}
