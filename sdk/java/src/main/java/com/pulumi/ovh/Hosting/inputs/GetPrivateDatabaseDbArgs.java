// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetPrivateDatabaseDbArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetPrivateDatabaseDbArgs Empty = new GetPrivateDatabaseDbArgs();

    /**
     * Database name
     * 
     */
    @Import(name="databaseName", required=true)
    private Output<String> databaseName;

    /**
     * @return Database name
     * 
     */
    public Output<String> databaseName() {
        return this.databaseName;
    }

    /**
     * The internal name of your private database
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The internal name of your private database
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private GetPrivateDatabaseDbArgs() {}

    private GetPrivateDatabaseDbArgs(GetPrivateDatabaseDbArgs $) {
        this.databaseName = $.databaseName;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetPrivateDatabaseDbArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetPrivateDatabaseDbArgs $;

        public Builder() {
            $ = new GetPrivateDatabaseDbArgs();
        }

        public Builder(GetPrivateDatabaseDbArgs defaults) {
            $ = new GetPrivateDatabaseDbArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param databaseName Database name
         * 
         * @return builder
         * 
         */
        public Builder databaseName(Output<String> databaseName) {
            $.databaseName = databaseName;
            return this;
        }

        /**
         * @param databaseName Database name
         * 
         * @return builder
         * 
         */
        public Builder databaseName(String databaseName) {
            return databaseName(Output.of(databaseName));
        }

        /**
         * @param serviceName The internal name of your private database
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your private database
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public GetPrivateDatabaseDbArgs build() {
            $.databaseName = Objects.requireNonNull($.databaseName, "expected parameter 'databaseName' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
