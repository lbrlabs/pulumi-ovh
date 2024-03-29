// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dbaas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetIamPolicyPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetIamPolicyPlainArgs Empty = new GetIamPolicyPlainArgs();

    /**
     * List of actions allowed by the policy.
     * 
     */
    @Import(name="allows")
    private @Nullable List<String> allows;

    /**
     * @return List of actions allowed by the policy.
     * 
     */
    public Optional<List<String>> allows() {
        return Optional.ofNullable(this.allows);
    }

    /**
     * Group description.
     * 
     */
    @Import(name="description")
    private @Nullable String description;

    /**
     * @return Group description.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * List of actions.
     * 
     */
    @Import(name="excepts")
    private @Nullable List<String> excepts;

    /**
     * @return List of actions.
     * 
     */
    public Optional<List<String>> excepts() {
        return Optional.ofNullable(this.excepts);
    }

    /**
     * UUID of the policy.
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return UUID of the policy.
     * 
     */
    public String id() {
        return this.id;
    }

    private GetIamPolicyPlainArgs() {}

    private GetIamPolicyPlainArgs(GetIamPolicyPlainArgs $) {
        this.allows = $.allows;
        this.description = $.description;
        this.excepts = $.excepts;
        this.id = $.id;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetIamPolicyPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetIamPolicyPlainArgs $;

        public Builder() {
            $ = new GetIamPolicyPlainArgs();
        }

        public Builder(GetIamPolicyPlainArgs defaults) {
            $ = new GetIamPolicyPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allows List of actions allowed by the policy.
         * 
         * @return builder
         * 
         */
        public Builder allows(@Nullable List<String> allows) {
            $.allows = allows;
            return this;
        }

        /**
         * @param allows List of actions allowed by the policy.
         * 
         * @return builder
         * 
         */
        public Builder allows(String... allows) {
            return allows(List.of(allows));
        }

        /**
         * @param description Group description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable String description) {
            $.description = description;
            return this;
        }

        /**
         * @param excepts List of actions.
         * 
         * @return builder
         * 
         */
        public Builder excepts(@Nullable List<String> excepts) {
            $.excepts = excepts;
            return this;
        }

        /**
         * @param excepts List of actions.
         * 
         * @return builder
         * 
         */
        public Builder excepts(String... excepts) {
            return excepts(List.of(excepts));
        }

        /**
         * @param id UUID of the policy.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
            return this;
        }

        public GetIamPolicyPlainArgs build() {
            $.id = Objects.requireNonNull($.id, "expected parameter 'id' to be non-null");
            return $;
        }
    }

}
