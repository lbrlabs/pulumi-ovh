// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Vrack.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class DedicatedServerState extends com.pulumi.resources.ResourceArgs {

    public static final DedicatedServerState Empty = new DedicatedServerState();

    /**
     * The id of the dedicated server.
     * 
     */
    @Import(name="serverId")
    private @Nullable Output<String> serverId;

    /**
     * @return The id of the dedicated server.
     * 
     */
    public Optional<Output<String>> serverId() {
        return Optional.ofNullable(this.serverId);
    }

    /**
     * The service name of the vrack. If omitted,
     * the `OVH_VRACK_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The service name of the vrack. If omitted,
     * the `OVH_VRACK_SERVICE` environment variable is used.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    private DedicatedServerState() {}

    private DedicatedServerState(DedicatedServerState $) {
        this.serverId = $.serverId;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DedicatedServerState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DedicatedServerState $;

        public Builder() {
            $ = new DedicatedServerState();
        }

        public Builder(DedicatedServerState defaults) {
            $ = new DedicatedServerState(Objects.requireNonNull(defaults));
        }

        /**
         * @param serverId The id of the dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serverId(@Nullable Output<String> serverId) {
            $.serverId = serverId;
            return this;
        }

        /**
         * @param serverId The id of the dedicated server.
         * 
         * @return builder
         * 
         */
        public Builder serverId(String serverId) {
            return serverId(Output.of(serverId));
        }

        /**
         * @param serviceName The service name of the vrack. If omitted,
         * the `OVH_VRACK_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The service name of the vrack. If omitted,
         * the `OVH_VRACK_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public DedicatedServerState build() {
            return $;
        }
    }

}