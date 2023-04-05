// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.IpLoadBalancing.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class HttpRouteRule {
    /**
     * @return Name of the field to match like &#34;protocol&#34; or &#34;host&#34; &#34;/ipLoadbalancing/{serviceName}/route/availableRules&#34; for a list of available rules
     * 
     */
    private @Nullable String field;
    /**
     * @return Matching operator. Not all operators are available for all fields. See &#34;availableRules&#34;
     * 
     */
    private @Nullable String match;
    /**
     * @return Invert the matching operator effect
     * 
     */
    private @Nullable Boolean negate;
    /**
     * @return Value to match against this match. Interpretation if this field depends on the match and field
     * 
     */
    private @Nullable String pattern;
    /**
     * @return Id of your rule
     * 
     */
    private @Nullable Integer ruleId;
    /**
     * @return Name of sub-field, if applicable. This may be a Cookie or Header name for instance
     * 
     */
    private @Nullable String subField;

    private HttpRouteRule() {}
    /**
     * @return Name of the field to match like &#34;protocol&#34; or &#34;host&#34; &#34;/ipLoadbalancing/{serviceName}/route/availableRules&#34; for a list of available rules
     * 
     */
    public Optional<String> field() {
        return Optional.ofNullable(this.field);
    }
    /**
     * @return Matching operator. Not all operators are available for all fields. See &#34;availableRules&#34;
     * 
     */
    public Optional<String> match() {
        return Optional.ofNullable(this.match);
    }
    /**
     * @return Invert the matching operator effect
     * 
     */
    public Optional<Boolean> negate() {
        return Optional.ofNullable(this.negate);
    }
    /**
     * @return Value to match against this match. Interpretation if this field depends on the match and field
     * 
     */
    public Optional<String> pattern() {
        return Optional.ofNullable(this.pattern);
    }
    /**
     * @return Id of your rule
     * 
     */
    public Optional<Integer> ruleId() {
        return Optional.ofNullable(this.ruleId);
    }
    /**
     * @return Name of sub-field, if applicable. This may be a Cookie or Header name for instance
     * 
     */
    public Optional<String> subField() {
        return Optional.ofNullable(this.subField);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(HttpRouteRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String field;
        private @Nullable String match;
        private @Nullable Boolean negate;
        private @Nullable String pattern;
        private @Nullable Integer ruleId;
        private @Nullable String subField;
        public Builder() {}
        public Builder(HttpRouteRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.field = defaults.field;
    	      this.match = defaults.match;
    	      this.negate = defaults.negate;
    	      this.pattern = defaults.pattern;
    	      this.ruleId = defaults.ruleId;
    	      this.subField = defaults.subField;
        }

        @CustomType.Setter
        public Builder field(@Nullable String field) {
            this.field = field;
            return this;
        }
        @CustomType.Setter
        public Builder match(@Nullable String match) {
            this.match = match;
            return this;
        }
        @CustomType.Setter
        public Builder negate(@Nullable Boolean negate) {
            this.negate = negate;
            return this;
        }
        @CustomType.Setter
        public Builder pattern(@Nullable String pattern) {
            this.pattern = pattern;
            return this;
        }
        @CustomType.Setter
        public Builder ruleId(@Nullable Integer ruleId) {
            this.ruleId = ruleId;
            return this;
        }
        @CustomType.Setter
        public Builder subField(@Nullable String subField) {
            this.subField = subField;
            return this;
        }
        public HttpRouteRule build() {
            final var o = new HttpRouteRule();
            o.field = field;
            o.match = match;
            o.negate = negate;
            o.pattern = pattern;
            o.ruleId = ruleId;
            o.subField = subField;
            return o;
        }
    }
}
