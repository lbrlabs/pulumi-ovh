// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dbaas.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetLogsClusterArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLogsClusterArgs Empty = new GetLogsClusterArgs();

    /**
     * The service name. It&#39;s the ID of your Logs Data Platform instance.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The service name. It&#39;s the ID of your Logs Data Platform instance.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private GetLogsClusterArgs() {}

    private GetLogsClusterArgs(GetLogsClusterArgs $) {
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLogsClusterArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLogsClusterArgs $;

        public Builder() {
            $ = new GetLogsClusterArgs();
        }

        public Builder(GetLogsClusterArgs defaults) {
            $ = new GetLogsClusterArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param serviceName The service name. It&#39;s the ID of your Logs Data Platform instance.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The service name. It&#39;s the ID of your Logs Data Platform instance.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public GetLogsClusterArgs build() {
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
