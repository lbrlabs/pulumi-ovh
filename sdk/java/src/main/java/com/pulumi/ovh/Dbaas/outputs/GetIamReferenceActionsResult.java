// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dbaas.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.ovh.Dbaas.outputs.GetIamReferenceActionsAction;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetIamReferenceActionsResult {
    /**
     * @return List of actions
     * 
     */
    private List<GetIamReferenceActionsAction> actions;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String type;

    private GetIamReferenceActionsResult() {}
    /**
     * @return List of actions
     * 
     */
    public List<GetIamReferenceActionsAction> actions() {
        return this.actions;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIamReferenceActionsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetIamReferenceActionsAction> actions;
        private String id;
        private String type;
        public Builder() {}
        public Builder(GetIamReferenceActionsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actions = defaults.actions;
    	      this.id = defaults.id;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder actions(List<GetIamReferenceActionsAction> actions) {
            this.actions = Objects.requireNonNull(actions);
            return this;
        }
        public Builder actions(GetIamReferenceActionsAction... actions) {
            return actions(List.of(actions));
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public GetIamReferenceActionsResult build() {
            final var o = new GetIamReferenceActionsResult();
            o.actions = actions;
            o.id = id;
            o.type = type;
            return o;
        }
    }
}
