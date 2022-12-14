// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetM3dbNamespacesResult {
    /**
     * @return See Argument Reference above.
     * 
     */
    private String clusterId;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The list of namespaces ids of the M3DB cluster associated with the project.
     * 
     */
    private List<String> namespaceIds;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;

    private GetM3dbNamespacesResult() {}
    /**
     * @return See Argument Reference above.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The list of namespaces ids of the M3DB cluster associated with the project.
     * 
     */
    public List<String> namespaceIds() {
        return this.namespaceIds;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetM3dbNamespacesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String clusterId;
        private String id;
        private List<String> namespaceIds;
        private String serviceName;
        public Builder() {}
        public Builder(GetM3dbNamespacesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterId = defaults.clusterId;
    	      this.id = defaults.id;
    	      this.namespaceIds = defaults.namespaceIds;
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder clusterId(String clusterId) {
            this.clusterId = Objects.requireNonNull(clusterId);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder namespaceIds(List<String> namespaceIds) {
            this.namespaceIds = Objects.requireNonNull(namespaceIds);
            return this;
        }
        public Builder namespaceIds(String... namespaceIds) {
            return namespaceIds(List.of(namespaceIds));
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        public GetM3dbNamespacesResult build() {
            final var o = new GetM3dbNamespacesResult();
            o.clusterId = clusterId;
            o.id = id;
            o.namespaceIds = namespaceIds;
            o.serviceName = serviceName;
            return o;
        }
    }
}
