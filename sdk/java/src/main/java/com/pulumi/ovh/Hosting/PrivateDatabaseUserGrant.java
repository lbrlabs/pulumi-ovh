// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Hosting;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.Hosting.PrivateDatabaseUserGrantArgs;
import com.pulumi.ovh.Hosting.inputs.PrivateDatabaseUserGrantState;
import com.pulumi.ovh.Utilities;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Add grant on a database in your private cloud database instance.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Hosting.PrivateDatabaseUserGrant;
 * import com.pulumi.ovh.Hosting.PrivateDatabaseUserGrantArgs;
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
 *         var userGrant = new PrivateDatabaseUserGrant(&#34;userGrant&#34;, PrivateDatabaseUserGrantArgs.builder()        
 *             .databaseName(&#34;ovhcloud&#34;)
 *             .grant(&#34;admin&#34;)
 *             .serviceName(&#34;XXXXXX&#34;)
 *             .userName(&#34;terraform&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * OVHcloud database user&#39;s grant can be imported using the `service_name`, the `user_name`, the `database_name` and the `grant`, separated by &#34;/&#34; E.g.,
 * 
 * ```sh
 *  $ pulumi import ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant user service_name/user_name/database_name/grant
 * ```
 * 
 */
@ResourceType(type="ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant")
public class PrivateDatabaseUserGrant extends com.pulumi.resources.CustomResource {
    /**
     * Database name where add grant.
     * 
     */
    @Export(name="databaseName", refs={String.class}, tree="[0]")
    private Output<String> databaseName;

    /**
     * @return Database name where add grant.
     * 
     */
    public Output<String> databaseName() {
        return this.databaseName;
    }
    /**
     * Database name where add grant. Values can be:
     * - admin
     * - none
     * - ro
     * - rw
     * 
     */
    @Export(name="grant", refs={String.class}, tree="[0]")
    private Output<String> grant;

    /**
     * @return Database name where add grant. Values can be:
     * - admin
     * - none
     * - ro
     * - rw
     * 
     */
    public Output<String> grant() {
        return this.grant;
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
     * User name used to connect on your databases.
     * 
     */
    @Export(name="userName", refs={String.class}, tree="[0]")
    private Output<String> userName;

    /**
     * @return User name used to connect on your databases.
     * 
     */
    public Output<String> userName() {
        return this.userName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivateDatabaseUserGrant(String name) {
        this(name, PrivateDatabaseUserGrantArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivateDatabaseUserGrant(String name, PrivateDatabaseUserGrantArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivateDatabaseUserGrant(String name, PrivateDatabaseUserGrantArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant", name, args == null ? PrivateDatabaseUserGrantArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PrivateDatabaseUserGrant(String name, Output<String> id, @Nullable PrivateDatabaseUserGrantState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Hosting/privateDatabaseUserGrant:PrivateDatabaseUserGrant", name, state, makeResourceOptions(options, id));
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
    public static PrivateDatabaseUserGrant get(String name, Output<String> id, @Nullable PrivateDatabaseUserGrantState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivateDatabaseUserGrant(name, id, state, options);
    }
}