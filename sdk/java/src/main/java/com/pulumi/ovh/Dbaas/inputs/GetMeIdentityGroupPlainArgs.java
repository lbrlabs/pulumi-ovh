// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dbaas.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetMeIdentityGroupPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMeIdentityGroupPlainArgs Empty = new GetMeIdentityGroupPlainArgs();

    /**
     * Group name.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Group name.
     * 
     */
    public String name() {
        return this.name;
    }

    private GetMeIdentityGroupPlainArgs() {}

    private GetMeIdentityGroupPlainArgs(GetMeIdentityGroupPlainArgs $) {
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMeIdentityGroupPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMeIdentityGroupPlainArgs $;

        public Builder() {
            $ = new GetMeIdentityGroupPlainArgs();
        }

        public Builder(GetMeIdentityGroupPlainArgs defaults) {
            $ = new GetMeIdentityGroupPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Group name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        public GetMeIdentityGroupPlainArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}