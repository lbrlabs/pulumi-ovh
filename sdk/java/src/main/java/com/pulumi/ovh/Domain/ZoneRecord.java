// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Domain;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.Domain.ZoneRecordArgs;
import com.pulumi.ovh.Domain.inputs.ZoneRecordState;
import com.pulumi.ovh.Utilities;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a OVHcloud domain zone record.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Domain.ZoneRecord;
 * import com.pulumi.ovh.Domain.ZoneRecordArgs;
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
 *         var test = new ZoneRecord(&#34;test&#34;, ZoneRecordArgs.builder()        
 *             .fieldtype(&#34;A&#34;)
 *             .subdomain(&#34;test&#34;)
 *             .target(&#34;0.0.0.0&#34;)
 *             .ttl(&#34;3600&#34;)
 *             .zone(&#34;testdemo.ovh&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * OVHcloud record can be imported using the `id` and the `zone`, egbash
 * 
 * ```sh
 *  $ pulumi import ovh:Domain/zoneRecord:ZoneRecord test 1234OVH_ID.zone.tld
 * ```
 * 
 */
@ResourceType(type="ovh:Domain/zoneRecord:ZoneRecord")
public class ZoneRecord extends com.pulumi.resources.CustomResource {
    /**
     * The type of the record
     * 
     */
    @Export(name="fieldtype", type=String.class, parameters={})
    private Output<String> fieldtype;

    /**
     * @return The type of the record
     * 
     */
    public Output<String> fieldtype() {
        return this.fieldtype;
    }
    /**
     * The name of the record
     * 
     */
    @Export(name="subdomain", type=String.class, parameters={})
    private Output</* @Nullable */ String> subdomain;

    /**
     * @return The name of the record
     * 
     */
    public Output<Optional<String>> subdomain() {
        return Codegen.optional(this.subdomain);
    }
    /**
     * The value of the record
     * 
     */
    @Export(name="target", type=String.class, parameters={})
    private Output<String> target;

    /**
     * @return The value of the record
     * 
     */
    public Output<String> target() {
        return this.target;
    }
    /**
     * The TTL of the record
     * 
     */
    @Export(name="ttl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> ttl;

    /**
     * @return The TTL of the record
     * 
     */
    public Output<Optional<Integer>> ttl() {
        return Codegen.optional(this.ttl);
    }
    /**
     * The domain to add the record to
     * 
     */
    @Export(name="zone", type=String.class, parameters={})
    private Output<String> zone;

    /**
     * @return The domain to add the record to
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ZoneRecord(String name) {
        this(name, ZoneRecordArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ZoneRecord(String name, ZoneRecordArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ZoneRecord(String name, ZoneRecordArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Domain/zoneRecord:ZoneRecord", name, args == null ? ZoneRecordArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ZoneRecord(String name, Output<String> id, @Nullable ZoneRecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Domain/zoneRecord:ZoneRecord", name, state, makeResourceOptions(options, id));
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
    public static ZoneRecord get(String name, Output<String> id, @Nullable ZoneRecordState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ZoneRecord(name, id, state, options);
    }
}
