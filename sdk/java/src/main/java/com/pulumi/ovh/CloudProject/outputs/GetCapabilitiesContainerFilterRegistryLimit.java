// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class GetCapabilitiesContainerFilterRegistryLimit {
    /**
     * @return Docker image storage limits in bytes
     * 
     */
    private Integer imageStorage;
    /**
     * @return Parallel requests on Docker image API (/v2 Docker registry API)
     * 
     */
    private Integer parallelRequest;

    private GetCapabilitiesContainerFilterRegistryLimit() {}
    /**
     * @return Docker image storage limits in bytes
     * 
     */
    public Integer imageStorage() {
        return this.imageStorage;
    }
    /**
     * @return Parallel requests on Docker image API (/v2 Docker registry API)
     * 
     */
    public Integer parallelRequest() {
        return this.parallelRequest;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCapabilitiesContainerFilterRegistryLimit defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer imageStorage;
        private Integer parallelRequest;
        public Builder() {}
        public Builder(GetCapabilitiesContainerFilterRegistryLimit defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.imageStorage = defaults.imageStorage;
    	      this.parallelRequest = defaults.parallelRequest;
        }

        @CustomType.Setter
        public Builder imageStorage(Integer imageStorage) {
            this.imageStorage = Objects.requireNonNull(imageStorage);
            return this;
        }
        @CustomType.Setter
        public Builder parallelRequest(Integer parallelRequest) {
            this.parallelRequest = Objects.requireNonNull(parallelRequest);
            return this;
        }
        public GetCapabilitiesContainerFilterRegistryLimit build() {
            final var o = new GetCapabilitiesContainerFilterRegistryLimit();
            o.imageStorage = imageStorage;
            o.parallelRequest = parallelRequest;
            return o;
        }
    }
}
