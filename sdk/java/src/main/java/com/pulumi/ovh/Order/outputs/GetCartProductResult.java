// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Order.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.Order.outputs.GetCartProductResultPrice;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCartProductResult {
    /**
     * @return Product offer identifier
     * 
     */
    private String planCode;
    /**
     * @return Prices of the product offer
     * 
     */
    private List<GetCartProductResultPrice> prices;
    /**
     * @return Name of the product
     * 
     */
    private String productName;
    /**
     * @return Product type
     * 
     */
    private String productType;

    private GetCartProductResult() {}
    /**
     * @return Product offer identifier
     * 
     */
    public String planCode() {
        return this.planCode;
    }
    /**
     * @return Prices of the product offer
     * 
     */
    public List<GetCartProductResultPrice> prices() {
        return this.prices;
    }
    /**
     * @return Name of the product
     * 
     */
    public String productName() {
        return this.productName;
    }
    /**
     * @return Product type
     * 
     */
    public String productType() {
        return this.productType;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCartProductResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String planCode;
        private List<GetCartProductResultPrice> prices;
        private String productName;
        private String productType;
        public Builder() {}
        public Builder(GetCartProductResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.planCode = defaults.planCode;
    	      this.prices = defaults.prices;
    	      this.productName = defaults.productName;
    	      this.productType = defaults.productType;
        }

        @CustomType.Setter
        public Builder planCode(String planCode) {
            this.planCode = Objects.requireNonNull(planCode);
            return this;
        }
        @CustomType.Setter
        public Builder prices(List<GetCartProductResultPrice> prices) {
            this.prices = Objects.requireNonNull(prices);
            return this;
        }
        public Builder prices(GetCartProductResultPrice... prices) {
            return prices(List.of(prices));
        }
        @CustomType.Setter
        public Builder productName(String productName) {
            this.productName = Objects.requireNonNull(productName);
            return this;
        }
        @CustomType.Setter
        public Builder productType(String productType) {
            this.productType = Objects.requireNonNull(productType);
            return this;
        }
        public GetCartProductResult build() {
            final var o = new GetCartProductResult();
            o.planCode = planCode;
            o.prices = prices;
            o.productName = productName;
            o.productType = productType;
            return o;
        }
    }
}
