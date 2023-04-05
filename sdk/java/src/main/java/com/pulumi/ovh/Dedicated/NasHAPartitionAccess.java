// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dedicated;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.Dedicated.NasHAPartitionAccessArgs;
import com.pulumi.ovh.Dedicated.inputs.NasHAPartitionAccessState;
import com.pulumi.ovh.Utilities;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource for managing access rights to partitions on HA-NAS services
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Dedicated.NasHAPartitionAccess;
 * import com.pulumi.ovh.Dedicated.NasHAPartitionAccessArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var my_partition = new NasHAPartitionAccess(&#34;my-partition&#34;, NasHAPartitionAccessArgs.builder()        
 *             .ip(&#34;123.123.123.123/32&#34;)
 *             .partitionName(&#34;my-partition&#34;)
 *             .serviceName(&#34;zpool-12345&#34;)
 *             .type(&#34;readwrite&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * HA-NAS partition access can be imported using the `{service_name}/{partition_name}/{ip}`, e.g.
 * 
 * ```sh
 *  $ pulumi import ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess my-partition zpool-12345/my-partition/123.123.123.123%2F32`
 * ```
 * 
 */
@ResourceType(type="ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess")
public class NasHAPartitionAccess extends com.pulumi.resources.CustomResource {
    /**
     * ip block in x.x.x.x/x format
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return ip block in x.x.x.x/x format
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * name of the partition
     * 
     */
    @Export(name="partitionName", refs={String.class}, tree="[0]")
    private Output<String> partitionName;

    /**
     * @return name of the partition
     * 
     */
    public Output<String> partitionName() {
        return this.partitionName;
    }
    /**
     * The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * one of &#34;readwrite&#34;, &#34;readonly&#34;
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return one of &#34;readwrite&#34;, &#34;readonly&#34;
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NasHAPartitionAccess(String name) {
        this(name, NasHAPartitionAccessArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NasHAPartitionAccess(String name, NasHAPartitionAccessArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NasHAPartitionAccess(String name, NasHAPartitionAccessArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess", name, args == null ? NasHAPartitionAccessArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NasHAPartitionAccess(String name, Output<String> id, @Nullable NasHAPartitionAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dedicated/nasHAPartitionAccess:NasHAPartitionAccess", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static NasHAPartitionAccess get(String name, Output<String> id, @Nullable NasHAPartitionAccessState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NasHAPartitionAccess(name, id, state, options);
    }
}