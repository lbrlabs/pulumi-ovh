// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPaymentmeanCreditCardResult {
    /**
     * @return a boolean which tells if the retrieved credit card
     * is marked as the default payment mean
     * 
     */
    private Boolean default_;
    /**
     * @return the description attribute of the credit card
     * 
     */
    private String description;
    private @Nullable String descriptionRegexp;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return the state attribute of the credit card
     * 
     */
    private String state;
    private @Nullable List<String> states;
    private @Nullable Boolean useDefault;
    private @Nullable Boolean useLastToExpire;

    private GetPaymentmeanCreditCardResult() {}
    /**
     * @return a boolean which tells if the retrieved credit card
     * is marked as the default payment mean
     * 
     */
    public Boolean default_() {
        return this.default_;
    }
    /**
     * @return the description attribute of the credit card
     * 
     */
    public String description() {
        return this.description;
    }
    public Optional<String> descriptionRegexp() {
        return Optional.ofNullable(this.descriptionRegexp);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return the state attribute of the credit card
     * 
     */
    public String state() {
        return this.state;
    }
    public List<String> states() {
        return this.states == null ? List.of() : this.states;
    }
    public Optional<Boolean> useDefault() {
        return Optional.ofNullable(this.useDefault);
    }
    public Optional<Boolean> useLastToExpire() {
        return Optional.ofNullable(this.useLastToExpire);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPaymentmeanCreditCardResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean default_;
        private String description;
        private @Nullable String descriptionRegexp;
        private String id;
        private String state;
        private @Nullable List<String> states;
        private @Nullable Boolean useDefault;
        private @Nullable Boolean useLastToExpire;
        public Builder() {}
        public Builder(GetPaymentmeanCreditCardResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.default_ = defaults.default_;
    	      this.description = defaults.description;
    	      this.descriptionRegexp = defaults.descriptionRegexp;
    	      this.id = defaults.id;
    	      this.state = defaults.state;
    	      this.states = defaults.states;
    	      this.useDefault = defaults.useDefault;
    	      this.useLastToExpire = defaults.useLastToExpire;
        }

        @CustomType.Setter("default")
        public Builder default_(Boolean default_) {
            this.default_ = Objects.requireNonNull(default_);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder descriptionRegexp(@Nullable String descriptionRegexp) {
            this.descriptionRegexp = descriptionRegexp;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder states(@Nullable List<String> states) {
            this.states = states;
            return this;
        }
        public Builder states(String... states) {
            return states(List.of(states));
        }
        @CustomType.Setter
        public Builder useDefault(@Nullable Boolean useDefault) {
            this.useDefault = useDefault;
            return this;
        }
        @CustomType.Setter
        public Builder useLastToExpire(@Nullable Boolean useLastToExpire) {
            this.useLastToExpire = useLastToExpire;
            return this;
        }
        public GetPaymentmeanCreditCardResult build() {
            final var o = new GetPaymentmeanCreditCardResult();
            o.default_ = default_;
            o.description = description;
            o.descriptionRegexp = descriptionRegexp;
            o.id = id;
            o.state = state;
            o.states = states;
            o.useDefault = useDefault;
            o.useLastToExpire = useLastToExpire;
            return o;
        }
    }
}