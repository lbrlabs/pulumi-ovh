// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetMeCurrency {
    /**
     * @return Currency code used by this account (e.g EUR, USD, ...)
     * 
     */
    private String code;
    /**
     * @return Currency symbol used by this account (e.g €, $, ...)
     * 
     */
    private String symbol;

    private GetMeCurrency() {}
    /**
     * @return Currency code used by this account (e.g EUR, USD, ...)
     * 
     */
    public String code() {
        return this.code;
    }
    /**
     * @return Currency symbol used by this account (e.g €, $, ...)
     * 
     */
    public String symbol() {
        return this.symbol;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMeCurrency defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String code;
        private String symbol;
        public Builder() {}
        public Builder(GetMeCurrency defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.code = defaults.code;
    	      this.symbol = defaults.symbol;
        }

        @CustomType.Setter
        public Builder code(String code) {
            this.code = Objects.requireNonNull(code);
            return this;
        }
        @CustomType.Setter
        public Builder symbol(String symbol) {
            this.symbol = Objects.requireNonNull(symbol);
            return this;
        }
        public GetMeCurrency build() {
            final var o = new GetMeCurrency();
            o.code = code;
            o.symbol = symbol;
            return o;
        }
    }
}
