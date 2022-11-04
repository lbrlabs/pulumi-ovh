// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetDatabaseInstancesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatabaseInstancesPlainArgs Empty = new GetDatabaseInstancesPlainArgs();

    /**
     * Cluster ID
     * 
     */
    @Import(name="clusterId", required=true)
    private String clusterId;

    /**
     * @return Cluster ID
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }

    /**
     * The engine of the database cluster you want to list databases. To get a full list of available engine visit:
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * Available engines:
     * * `mysql`
     * * `postgresql`
     * 
     */
    @Import(name="engine", required=true)
    private String engine;

    /**
     * @return The engine of the database cluster you want to list databases. To get a full list of available engine visit:
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * Available engines:
     * * `mysql`
     * * `postgresql`
     * 
     */
    public String engine() {
        return this.engine;
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetDatabaseInstancesPlainArgs() {}

    private GetDatabaseInstancesPlainArgs(GetDatabaseInstancesPlainArgs $) {
        this.clusterId = $.clusterId;
        this.engine = $.engine;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatabaseInstancesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatabaseInstancesPlainArgs $;

        public Builder() {
            $ = new GetDatabaseInstancesPlainArgs();
        }

        public Builder(GetDatabaseInstancesPlainArgs defaults) {
            $ = new GetDatabaseInstancesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param engine The engine of the database cluster you want to list databases. To get a full list of available engine visit:
         * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * Available engines:
         * * `mysql`
         * * `postgresql`
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetDatabaseInstancesPlainArgs build() {
            $.clusterId = Objects.requireNonNull($.clusterId, "expected parameter 'clusterId' to be non-null");
            $.engine = Objects.requireNonNull($.engine, "expected parameter 'engine' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            return $;
        }
    }

}