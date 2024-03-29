// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.Hosting.PrivateDatabaseAllowlistArgs;
import com.pulumi.ovh.Hosting.inputs.PrivateDatabaseAllowlistState;
import com.pulumi.ovh.Utilities;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Create a new IP whitelist on your private cloud database instance.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Hosting.PrivateDatabaseAllowlist;
 * import com.pulumi.ovh.Hosting.PrivateDatabaseAllowlistArgs;
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
 *         var ip = new PrivateDatabaseAllowlist(&#34;ip&#34;, PrivateDatabaseAllowlistArgs.builder()        
 *             .ip(&#34;1.2.3.4&#34;)
 *             .service(true)
 *             .serviceName(&#34;XXXXXX&#34;)
 *             .sftp(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * OVHcloud database whitelist can be imported using the `service_name` and the `ip`, separated by &#34;/&#34; E.g.,
 * 
 * ```sh
 *  $ pulumi import ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist ip service_name/ip
 * ```
 * 
 */
@ResourceType(type="ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist")
public class PrivateDatabaseAllowlist extends com.pulumi.resources.CustomResource {
    /**
     * The whitelisted IP in your instance.
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return The whitelisted IP in your instance.
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * Custom name for your Whitelisted IP.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Custom name for your Whitelisted IP.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Authorize this IP to access service port. Values can be `true` or `false`
     * 
     */
    @Export(name="service", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> service;

    /**
     * @return Authorize this IP to access service port. Values can be `true` or `false`
     * 
     */
    public Output<Boolean> service() {
        return this.service;
    }
    /**
     * The internal name of your private database.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your private database.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Authorize this IP to access SFTP port. Values can be `true` or `false`
     * 
     */
    @Export(name="sftp", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> sftp;

    /**
     * @return Authorize this IP to access SFTP port. Values can be `true` or `false`
     * 
     */
    public Output<Boolean> sftp() {
        return this.sftp;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivateDatabaseAllowlist(String name) {
        this(name, PrivateDatabaseAllowlistArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivateDatabaseAllowlist(String name, PrivateDatabaseAllowlistArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivateDatabaseAllowlist(String name, PrivateDatabaseAllowlistArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist", name, args == null ? PrivateDatabaseAllowlistArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PrivateDatabaseAllowlist(String name, Output<String> id, @Nullable PrivateDatabaseAllowlistState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Hosting/privateDatabaseAllowlist:PrivateDatabaseAllowlist", name, state, makeResourceOptions(options, id));
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
    public static PrivateDatabaseAllowlist get(String name, Output<String> id, @Nullable PrivateDatabaseAllowlistState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivateDatabaseAllowlist(name, id, state, options);
    }
}
