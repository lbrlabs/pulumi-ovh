// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dedicated.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class ServerNetworkingInterface {
    /**
     * @return List of mac addresses to bind together.
     * 
     */
    private List<String> macs;
    /**
     * @return Type of bonding to create.
     * 
     */
    private String type;

    private ServerNetworkingInterface() {}
    /**
     * @return List of mac addresses to bind together.
     * 
     */
    public List<String> macs() {
        return this.macs;
    }
    /**
     * @return Type of bonding to create.
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ServerNetworkingInterface defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> macs;
        private String type;
        public Builder() {}
        public Builder(ServerNetworkingInterface defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.macs = defaults.macs;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder macs(List<String> macs) {
            this.macs = Objects.requireNonNull(macs);
            return this;
        }
        public Builder macs(String... macs) {
            return macs(List.of(macs));
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public ServerNetworkingInterface build() {
            final var o = new ServerNetworkingInterface();
            o.macs = macs;
            o.type = type;
            return o;
        }
    }
}
