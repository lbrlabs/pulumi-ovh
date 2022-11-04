// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Domain.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetZoneResult {
    /**
     * @return Is DNSSEC supported by this zone
     * 
     */
    private Boolean dnssecSupported;
    /**
     * @return hasDnsAnycast flag of the DNS zone
     * 
     */
    private Boolean hasDnsAnycast;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Last update date of the DNS zone
     * 
     */
    private String lastUpdate;
    private String name;
    /**
     * @return Name servers that host the DNS zone
     * 
     */
    private List<String> nameServers;

    private GetZoneResult() {}
    /**
     * @return Is DNSSEC supported by this zone
     * 
     */
    public Boolean dnssecSupported() {
        return this.dnssecSupported;
    }
    /**
     * @return hasDnsAnycast flag of the DNS zone
     * 
     */
    public Boolean hasDnsAnycast() {
        return this.hasDnsAnycast;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Last update date of the DNS zone
     * 
     */
    public String lastUpdate() {
        return this.lastUpdate;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return Name servers that host the DNS zone
     * 
     */
    public List<String> nameServers() {
        return this.nameServers;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetZoneResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean dnssecSupported;
        private Boolean hasDnsAnycast;
        private String id;
        private String lastUpdate;
        private String name;
        private List<String> nameServers;
        public Builder() {}
        public Builder(GetZoneResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.dnssecSupported = defaults.dnssecSupported;
    	      this.hasDnsAnycast = defaults.hasDnsAnycast;
    	      this.id = defaults.id;
    	      this.lastUpdate = defaults.lastUpdate;
    	      this.name = defaults.name;
    	      this.nameServers = defaults.nameServers;
        }

        @CustomType.Setter
        public Builder dnssecSupported(Boolean dnssecSupported) {
            this.dnssecSupported = Objects.requireNonNull(dnssecSupported);
            return this;
        }
        @CustomType.Setter
        public Builder hasDnsAnycast(Boolean hasDnsAnycast) {
            this.hasDnsAnycast = Objects.requireNonNull(hasDnsAnycast);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder lastUpdate(String lastUpdate) {
            this.lastUpdate = Objects.requireNonNull(lastUpdate);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder nameServers(List<String> nameServers) {
            this.nameServers = Objects.requireNonNull(nameServers);
            return this;
        }
        public Builder nameServers(String... nameServers) {
            return nameServers(List.of(nameServers));
        }
        public GetZoneResult build() {
            final var o = new GetZoneResult();
            o.dnssecSupported = dnssecSupported;
            o.hasDnsAnycast = hasDnsAnycast;
            o.id = id;
            o.lastUpdate = lastUpdate;
            o.name = name;
            o.nameServers = nameServers;
            return o;
        }
    }
}