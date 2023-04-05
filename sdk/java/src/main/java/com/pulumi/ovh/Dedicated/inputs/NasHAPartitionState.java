// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dedicated.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NasHAPartitionState extends com.pulumi.resources.ResourceArgs {

    public static final NasHAPartitionState Empty = new NasHAPartitionState();

    /**
     * Percentage of partition space used in %
     * 
     */
    @Import(name="capacity")
    private @Nullable Output<Integer> capacity;

    /**
     * @return Percentage of partition space used in %
     * 
     */
    public Optional<Output<Integer>> capacity() {
        return Optional.ofNullable(this.capacity);
    }

    /**
     * A brief description of the partition
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A brief description of the partition
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * name of the partition
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return name of the partition
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * one of &#34;NFS&#34;, &#34;CIFS&#34; or &#34;NFS_CIFS&#34;
     * 
     */
    @Import(name="protocol")
    private @Nullable Output<String> protocol;

    /**
     * @return one of &#34;NFS&#34;, &#34;CIFS&#34; or &#34;NFS_CIFS&#34;
     * 
     */
    public Optional<Output<String>> protocol() {
        return Optional.ofNullable(this.protocol);
    }

    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * size of the partition in GB
     * 
     */
    @Import(name="size")
    private @Nullable Output<Integer> size;

    /**
     * @return size of the partition in GB
     * 
     */
    public Optional<Output<Integer>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * Percentage of partition space used by snapshots in %
     * 
     */
    @Import(name="usedBySnapshots")
    private @Nullable Output<Integer> usedBySnapshots;

    /**
     * @return Percentage of partition space used by snapshots in %
     * 
     */
    public Optional<Output<Integer>> usedBySnapshots() {
        return Optional.ofNullable(this.usedBySnapshots);
    }

    private NasHAPartitionState() {}

    private NasHAPartitionState(NasHAPartitionState $) {
        this.capacity = $.capacity;
        this.description = $.description;
        this.name = $.name;
        this.protocol = $.protocol;
        this.serviceName = $.serviceName;
        this.size = $.size;
        this.usedBySnapshots = $.usedBySnapshots;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NasHAPartitionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NasHAPartitionState $;

        public Builder() {
            $ = new NasHAPartitionState();
        }

        public Builder(NasHAPartitionState defaults) {
            $ = new NasHAPartitionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param capacity Percentage of partition space used in %
         * 
         * @return builder
         * 
         */
        public Builder capacity(@Nullable Output<Integer> capacity) {
            $.capacity = capacity;
            return this;
        }

        /**
         * @param capacity Percentage of partition space used in %
         * 
         * @return builder
         * 
         */
        public Builder capacity(Integer capacity) {
            return capacity(Output.of(capacity));
        }

        /**
         * @param description A brief description of the partition
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A brief description of the partition
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param name name of the partition
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name name of the partition
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param protocol one of &#34;NFS&#34;, &#34;CIFS&#34; or &#34;NFS_CIFS&#34;
         * 
         * @return builder
         * 
         */
        public Builder protocol(@Nullable Output<String> protocol) {
            $.protocol = protocol;
            return this;
        }

        /**
         * @param protocol one of &#34;NFS&#34;, &#34;CIFS&#34; or &#34;NFS_CIFS&#34;
         * 
         * @return builder
         * 
         */
        public Builder protocol(String protocol) {
            return protocol(Output.of(protocol));
        }

        /**
         * @param serviceName The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param size size of the partition in GB
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size size of the partition in GB
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param usedBySnapshots Percentage of partition space used by snapshots in %
         * 
         * @return builder
         * 
         */
        public Builder usedBySnapshots(@Nullable Output<Integer> usedBySnapshots) {
            $.usedBySnapshots = usedBySnapshots;
            return this;
        }

        /**
         * @param usedBySnapshots Percentage of partition space used by snapshots in %
         * 
         * @return builder
         * 
         */
        public Builder usedBySnapshots(Integer usedBySnapshots) {
            return usedBySnapshots(Output.of(usedBySnapshots));
        }

        public NasHAPartitionState build() {
            return $;
        }
    }

}
