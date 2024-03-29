// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Me;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.Me.IpxeScriptArgs;
import com.pulumi.ovh.Me.inputs.IpxeScriptState;
import com.pulumi.ovh.Utilities;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Creates an IPXE Script.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Me.IpxeScript;
 * import com.pulumi.ovh.Me.IpxeScriptArgs;
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
 *         var script = new IpxeScript(&#34;script&#34;, IpxeScriptArgs.builder()        
 *             .script(Files.readString(Paths.get(String.format(&#34;%s/boot.ipxe&#34;, path.module()))))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="ovh:Me/ipxeScript:IpxeScript")
public class IpxeScript extends com.pulumi.resources.CustomResource {
    /**
     * For documentation purpose only. This attribute is not passed to the OVHcloud API as it cannot be retrieved back. Instead a fake description (&#39;$name auto description&#39;) is passed at creation time.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return For documentation purpose only. This attribute is not passed to the OVHcloud API as it cannot be retrieved back. Instead a fake description (&#39;$name auto description&#39;) is passed at creation time.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The name of the IPXE Script.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the IPXE Script.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The content of the script.
     * 
     */
    @Export(name="script", refs={String.class}, tree="[0]")
    private Output<String> script;

    /**
     * @return The content of the script.
     * 
     */
    public Output<String> script() {
        return this.script;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public IpxeScript(String name) {
        this(name, IpxeScriptArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public IpxeScript(String name, IpxeScriptArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public IpxeScript(String name, IpxeScriptArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/ipxeScript:IpxeScript", name, args == null ? IpxeScriptArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private IpxeScript(String name, Output<String> id, @Nullable IpxeScriptState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/ipxeScript:IpxeScript", name, state, makeResourceOptions(options, id));
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
    public static IpxeScript get(String name, Output<String> id, @Nullable IpxeScriptState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new IpxeScript(name, id, state, options);
    }
}
