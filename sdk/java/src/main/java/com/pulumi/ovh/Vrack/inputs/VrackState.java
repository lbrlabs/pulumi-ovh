// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Vrack.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.Vrack.inputs.VrackOrderArgs;
import com.pulumi.ovh.Vrack.inputs.VrackPlanArgs;
import com.pulumi.ovh.Vrack.inputs.VrackPlanOptionArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class VrackState extends com.pulumi.resources.ResourceArgs {

    public static final VrackState Empty = new VrackState();

    /**
     * yourvrackdescription
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return yourvrackdescription
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * yourvrackname
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return yourvrackname
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Details about an Order
     * 
     */
    @Import(name="orders")
    private @Nullable Output<List<VrackOrderArgs>> orders;

    /**
     * @return Details about an Order
     * 
     */
    public Optional<Output<List<VrackOrderArgs>>> orders() {
        return Optional.ofNullable(this.orders);
    }

    /**
     * OVHcloud Subsidiary
     * 
     */
    @Import(name="ovhSubsidiary")
    private @Nullable Output<String> ovhSubsidiary;

    /**
     * @return OVHcloud Subsidiary
     * 
     */
    public Optional<Output<String>> ovhSubsidiary() {
        return Optional.ofNullable(this.ovhSubsidiary);
    }

    /**
     * OVHcloud payment mode (One of &#34;default-payment-mean&#34;, &#34;fidelity&#34;, &#34;ovh-account&#34;)
     * 
     */
    @Import(name="paymentMean")
    private @Nullable Output<String> paymentMean;

    /**
     * @return OVHcloud payment mode (One of &#34;default-payment-mean&#34;, &#34;fidelity&#34;, &#34;ovh-account&#34;)
     * 
     */
    public Optional<Output<String>> paymentMean() {
        return Optional.ofNullable(this.paymentMean);
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="plan")
    private @Nullable Output<VrackPlanArgs> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Optional<Output<VrackPlanArgs>> plan() {
        return Optional.ofNullable(this.plan);
    }

    /**
     * Product Plan to order
     * 
     */
    @Import(name="planOptions")
    private @Nullable Output<List<VrackPlanOptionArgs>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Optional<Output<List<VrackPlanOptionArgs>>> planOptions() {
        return Optional.ofNullable(this.planOptions);
    }

    /**
     * The internal name of your vrack
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The internal name of your vrack
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    private VrackState() {}

    private VrackState(VrackState $) {
        this.description = $.description;
        this.name = $.name;
        this.orders = $.orders;
        this.ovhSubsidiary = $.ovhSubsidiary;
        this.paymentMean = $.paymentMean;
        this.plan = $.plan;
        this.planOptions = $.planOptions;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(VrackState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private VrackState $;

        public Builder() {
            $ = new VrackState();
        }

        public Builder(VrackState defaults) {
            $ = new VrackState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description yourvrackdescription
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description yourvrackdescription
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name yourvrackname
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name yourvrackname
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param orders Details about an Order
         * 
         * @return builder
         * 
         */
        public Builder orders(@Nullable Output<List<VrackOrderArgs>> orders) {
            $.orders = orders;
            return this;
        }

        /**
         * @param orders Details about an Order
         * 
         * @return builder
         * 
         */
        public Builder orders(List<VrackOrderArgs> orders) {
            return orders(Output.of(orders));
        }

        /**
         * @param orders Details about an Order
         * 
         * @return builder
         * 
         */
        public Builder orders(VrackOrderArgs... orders) {
            return orders(List.of(orders));
        }

        /**
         * @param ovhSubsidiary OVHcloud Subsidiary
         * 
         * @return builder
         * 
         */
        public Builder ovhSubsidiary(@Nullable Output<String> ovhSubsidiary) {
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
        public Builder paymentMean(@Nullable Output<String> paymentMean) {
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
        public Builder plan(@Nullable Output<VrackPlanArgs> plan) {
            $.plan = plan;
            return this;
        }

        /**
         * @param plan Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder plan(VrackPlanArgs plan) {
            return plan(Output.of(plan));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(@Nullable Output<List<VrackPlanOptionArgs>> planOptions) {
            $.planOptions = planOptions;
            return this;
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(List<VrackPlanOptionArgs> planOptions) {
            return planOptions(Output.of(planOptions));
        }

        /**
         * @param planOptions Product Plan to order
         * 
         * @return builder
         * 
         */
        public Builder planOptions(VrackPlanOptionArgs... planOptions) {
            return planOptions(List.of(planOptions));
        }

        /**
         * @param serviceName The internal name of your vrack
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your vrack
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public VrackState build() {
            return $;
        }
    }

}