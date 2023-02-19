// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.Hosting.inputs.PrivateDatabasePlanArgs;
import com.pulumi.ovh.Hosting.inputs.PrivateDatabasePlanOptionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PrivateDatabaseArgs extends com.pulumi.resources.ResourceArgs {

    public static final PrivateDatabaseArgs Empty = new PrivateDatabaseArgs();

    /**
     * Name displayed in customer panel for your private database
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Name displayed in customer panel for your private database
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * OVHcloud Subsidiary
     * 
     */
    @Import(name="ovhSubsidiary", required=true)
    private Output<String> ovhSubsidiary;

    /**
     * @return OVHcloud Subsidiary
     * 
     */
    public Output<String> ovhSubsidiary() {
        return this.ovhSubsidiary;
    }

    /**
     * OVHcloud payment mode (One of &#34;default-payment-mean&#34;, &#34;fidelity&#34;, &#34;ovh-account&#34;)
     * 
     */
    @Import(name="paymentMean", required=true)
    private Output<String> paymentMean;

    /**
     * @return OVHcloud payment mode (One of &#34;default-payment-mean&#34;, &#34;fidelity&#34;, &#34;ovh-account&#34;)
     * 
     */
    public Output<String> paymentMean() {
        return this.paymentMean;
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="plan", required=true)
    private Output<PrivateDatabasePlanArgs> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<PrivateDatabasePlanArgs> plan() {
        return this.plan;
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="planOptions")
    private @Nullable Output<List<PrivateDatabasePlanOptionArgs>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Optional<Output<List<PrivateDatabasePlanOptionArgs>>> planOptions() {
        return Optional.ofNullable(this.planOptions);
    }

    /**
     * Service name
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return Service name
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    private PrivateDatabaseArgs() {}

    private PrivateDatabaseArgs(PrivateDatabaseArgs $) {
        this.displayName = $.displayName;
        this.ovhSubsidiary = $.ovhSubsidiary;
        this.paymentMean = $.paymentMean;
        this.plan = $.plan;
        this.planOptions = $.planOptions;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivateDatabaseArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivateDatabaseArgs $;

        public Builder() {
            $ = new PrivateDatabaseArgs();
        }

        public Builder(PrivateDatabaseArgs defaults) {
            $ = new PrivateDatabaseArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param displayName Name displayed in customer panel for your private database
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Name displayed in customer panel for your private database
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(Output<String> ovhSubsidiary) {
            $.ovhSubsidiary = ovhSubsidiary;
            return this;
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(String ovhSubsidiary) {
            return ovhSubsidiary(Output.of(ovhSubsidiary));
        }

        /**
         * @param paymentMean OVHcloud payment mode (One of &#34;default-payment-mean&#34;, &#34;fidelity&#34;, &#34;ovh-account&#34;)
         * 
         * @return builder
         * 
         */
        public Builder paymentMean(Output<String> paymentMean) {
            $.paymentMean = paymentMean;
            return this;
        }

        /**
         * @param paymentMean OVHcloud payment mode (One of &#34;default-payment-mean&#34;, &#34;fidelity&#34;, &#34;ovh-account&#34;)
         * 
         * @return builder
         * 
         */
        public Builder paymentMean(String paymentMean) {
            return paymentMean(Output.of(paymentMean));
        }

        /**
         * @param plan Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder plan(Output<PrivateDatabasePlanArgs> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder plan(PrivateDatabasePlanArgs plan) {
            return plan(Output.of(plan));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(@Nullable Output<List<PrivateDatabasePlanOptionArgs>> planOptions) {
            $.planOptions = planOptions;
            return this;
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(List<PrivateDatabasePlanOptionArgs> planOptions) {
            return planOptions(Output.of(planOptions));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(PrivateDatabasePlanOptionArgs... planOptions) {
            return planOptions(List.of(planOptions));
        }

        /**
         * @param serviceName Service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName Service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public PrivateDatabaseArgs build() {
            $.ovhSubsidiary = Objects.requireNonNull($.ovhSubsidiary, "expected parameter 'ovhSubsidiary' to be non-null");
            $.paymentMean = Objects.requireNonNull($.paymentMean, "expected parameter 'paymentMean' to be non-null");
            $.plan = Objects.requireNonNull($.plan, "expected parameter 'plan' to be non-null");
            return $;
        }
    }

}
