// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Order.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.Order.outputs.GetCartProductOptionsResult;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCartProductOptionsInvokeResult {
    private String cartId;
    private @Nullable String catalogName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Product offer identifier
     * 
     */
    private String planCode;
    private String product;
    /**
     * @return products results
     * 
     */
    private List<GetCartProductOptionsResult> results;

    private GetCartProductOptionsInvokeResult() {}
    public String cartId() {
        return this.cartId;
    }
    public Optional<String> catalogName() {
        return Optional.ofNullable(this.catalogName);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Product offer identifier
     * 
     */
    public String planCode() {
        return this.planCode;
    }
    public String product() {
        return this.product;
    }
    /**
     * @return products results
     * 
     */
    public List<GetCartProductOptionsResult> results() {
        return this.results;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCartProductOptionsInvokeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String cartId;
        private @Nullable String catalogName;
        private String id;
        private String planCode;
        private String product;
        private List<GetCartProductOptionsResult> results;
        public Builder() {}
        public Builder(GetCartProductOptionsInvokeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cartId = defaults.cartId;
    	      this.catalogName = defaults.catalogName;
    	      this.id = defaults.id;
    	      this.planCode = defaults.planCode;
    	      this.product = defaults.product;
    	      this.results = defaults.results;
        }

        @CustomType.Setter
        public Builder cartId(String cartId) {
            this.cartId = Objects.requireNonNull(cartId);
            return this;
        }
        @CustomType.Setter
        public Builder catalogName(@Nullable String catalogName) {
            this.catalogName = catalogName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder planCode(String planCode) {
            this.planCode = Objects.requireNonNull(planCode);
            return this;
        }
        @CustomType.Setter
        public Builder product(String product) {
            this.product = Objects.requireNonNull(product);
            return this;
        }
        @CustomType.Setter
        public Builder results(List<GetCartProductOptionsResult> results) {
            this.results = Objects.requireNonNull(results);
            return this;
        }
        public Builder results(GetCartProductOptionsResult... results) {
            return results(List.of(results));
        }
        public GetCartProductOptionsInvokeResult build() {
            final var o = new GetCartProductOptionsInvokeResult();
            o.cartId = cartId;
            o.catalogName = catalogName;
            o.id = id;
            o.planCode = planCode;
            o.product = product;
            o.results = results;
            return o;
        }
    }
}
