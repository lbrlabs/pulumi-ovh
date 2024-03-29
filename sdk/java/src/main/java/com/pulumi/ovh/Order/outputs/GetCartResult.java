// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Order.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetCartResult {
    private @Nullable Boolean assign;
    /**
     * @return Cart identifier
     * 
     */
    private String cartId;
    private @Nullable String description;
    private String expire;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Items of your cart
     * 
     */
    private List<Integer> items;
    private String ovhSubsidiary;
    /**
     * @return Indicates if the cart has already been validated
     * 
     */
    private Boolean readOnly;

    private GetCartResult() {}
    public Optional<Boolean> assign() {
        return Optional.ofNullable(this.assign);
    }
    /**
     * @return Cart identifier
     * 
     */
    public String cartId() {
        return this.cartId;
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public String expire() {
        return this.expire;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Items of your cart
     * 
     */
    public List<Integer> items() {
        return this.items;
    }
    public String ovhSubsidiary() {
        return this.ovhSubsidiary;
    }
    /**
     * @return Indicates if the cart has already been validated
     * 
     */
    public Boolean readOnly() {
        return this.readOnly;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCartResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean assign;
        private String cartId;
        private @Nullable String description;
        private String expire;
        private String id;
        private List<Integer> items;
        private String ovhSubsidiary;
        private Boolean readOnly;
        public Builder() {}
        public Builder(GetCartResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.assign = defaults.assign;
    	      this.cartId = defaults.cartId;
    	      this.description = defaults.description;
    	      this.expire = defaults.expire;
    	      this.id = defaults.id;
    	      this.items = defaults.items;
    	      this.ovhSubsidiary = defaults.ovhSubsidiary;
    	      this.readOnly = defaults.readOnly;
        }

        @CustomType.Setter
        public Builder assign(@Nullable Boolean assign) {
            this.assign = assign;
            return this;
        }
        @CustomType.Setter
        public Builder cartId(String cartId) {
            this.cartId = Objects.requireNonNull(cartId);
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder expire(String expire) {
            this.expire = Objects.requireNonNull(expire);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder items(List<Integer> items) {
            this.items = Objects.requireNonNull(items);
            return this;
        }
        public Builder items(Integer... items) {
            return items(List.of(items));
        }
        @CustomType.Setter
        public Builder ovhSubsidiary(String ovhSubsidiary) {
            this.ovhSubsidiary = Objects.requireNonNull(ovhSubsidiary);
            return this;
        }
        @CustomType.Setter
        public Builder readOnly(Boolean readOnly) {
            this.readOnly = Objects.requireNonNull(readOnly);
            return this;
        }
        public GetCartResult build() {
            final var o = new GetCartResult();
            o.assign = assign;
            o.cartId = cartId;
            o.description = description;
            o.expire = expire;
            o.id = id;
            o.items = items;
            o.ovhSubsidiary = ovhSubsidiary;
            o.readOnly = readOnly;
            return o;
        }
    }
}
