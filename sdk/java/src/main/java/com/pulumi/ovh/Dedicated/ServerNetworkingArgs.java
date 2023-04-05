// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dedicated;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.ovh.Dedicated.inputs.ServerNetworkingInterfaceArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class ServerNetworkingArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerNetworkingArgs Empty = new ServerNetworkingArgs();

    /**
     * Interface or interfaces aggregation.
     * 
     */
    @Import(name="interfaces", required=true)
    private Output<List<ServerNetworkingInterfaceArgs>> interfaces;

    /**
     * @return Interface or interfaces aggregation.
     * 
     */
    public Output<List<ServerNetworkingInterfaceArgs>> interfaces() {
        return this.interfaces;
    }

    /**
     * The service_name of your dedicated server. The full list of available dedicated servers can be found using the `ovh.getServers` datasource.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The service_name of your dedicated server. The full list of available dedicated servers can be found using the `ovh.getServers` datasource.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private ServerNetworkingArgs() {}

    private ServerNetworkingArgs(ServerNetworkingArgs $) {
        this.interfaces = $.interfaces;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerNetworkingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerNetworkingArgs $;

        public Builder() {
            $ = new ServerNetworkingArgs();
        }

        public Builder(ServerNetworkingArgs defaults) {
            $ = new ServerNetworkingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param interfaces Interface or interfaces aggregation.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(Output<List<ServerNetworkingInterfaceArgs>> interfaces) {
            $.interfaces = interfaces;
            return this;
        }

        /**
         * @param interfaces Interface or interfaces aggregation.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(List<ServerNetworkingInterfaceArgs> interfaces) {
            return interfaces(Output.of(interfaces));
        }

        /**
         * @param interfaces Interface or interfaces aggregation.
         * 
         * @return builder
         * 
         */
        public Builder interfaces(ServerNetworkingInterfaceArgs... interfaces) {
            return interfaces(List.of(interfaces));
        }

        /**
         * @param serviceName The service_name of your dedicated server. The full list of available dedicated servers can be found using the `ovh.getServers` datasource.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The service_name of your dedicated server. The full list of available dedicated servers can be found using the `ovh.getServers` datasource.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public ServerNetworkingArgs build() {
            $.interfaces = Objects.requireNonNull($.interfaces, "expected parameter 'interfaces' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}
