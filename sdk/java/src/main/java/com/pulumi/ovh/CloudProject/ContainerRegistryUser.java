// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.CloudProject.ContainerRegistryUserArgs;
import com.pulumi.ovh.CloudProject.inputs.ContainerRegistryUserState;
import com.pulumi.ovh.Utilities;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Creates a user for a container registry associated with a public cloud project.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProject.CloudProjectFunctions;
 * import com.pulumi.ovh.CloudProject.inputs.GetContainerRegistryArgs;
 * import com.pulumi.ovh.CloudProject.ContainerRegistryUser;
 * import com.pulumi.ovh.CloudProject.ContainerRegistryUserArgs;
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
 *         final var registry = CloudProjectFunctions.getContainerRegistry(GetContainerRegistryArgs.builder()
 *             .serviceName(&#34;XXXXXX&#34;)
 *             .registryId(&#34;yyyy&#34;)
 *             .build());
 * 
 *         var user = new ContainerRegistryUser(&#34;user&#34;, ContainerRegistryUserArgs.builder()        
 *             .serviceName(ovh_cloud_project_containerregistry.registry().service_name())
 *             .registryId(ovh_cloud_project_containerregistry.registry().id())
 *             .email(&#34;foo@bar.com&#34;)
 *             .login(&#34;foobar&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProject/containerRegistryUser:ContainerRegistryUser")
public class ContainerRegistryUser extends com.pulumi.resources.CustomResource {
    /**
     * User email
     * 
     */
    @Export(name="email", refs={String.class}, tree="[0]")
    private Output<String> email;

    /**
     * @return User email
     * 
     */
    public Output<String> email() {
        return this.email;
    }
    /**
     * Registry name
     * 
     */
    @Export(name="login", refs={String.class}, tree="[0]")
    private Output<String> login;

    /**
     * @return Registry name
     * 
     */
    public Output<String> login() {
        return this.login;
    }
    /**
     * (Sensitive) User password
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output<String> password;

    /**
     * @return (Sensitive) User password
     * 
     */
    public Output<String> password() {
        return this.password;
    }
    /**
     * Registry ID
     * 
     */
    @Export(name="registryId", refs={String.class}, tree="[0]")
    private Output<String> registryId;

    /**
     * @return Registry ID
     * 
     */
    public Output<String> registryId() {
        return this.registryId;
    }
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * User name
     * 
     */
    @Export(name="user", refs={String.class}, tree="[0]")
    private Output<String> user;

    /**
     * @return User name
     * 
     */
    public Output<String> user() {
        return this.user;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ContainerRegistryUser(String name) {
        this(name, ContainerRegistryUserArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ContainerRegistryUser(String name, ContainerRegistryUserArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ContainerRegistryUser(String name, ContainerRegistryUserArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/containerRegistryUser:ContainerRegistryUser", name, args == null ? ContainerRegistryUserArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ContainerRegistryUser(String name, Output<String> id, @Nullable ContainerRegistryUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/containerRegistryUser:ContainerRegistryUser", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static ContainerRegistryUser get(String name, Output<String> id, @Nullable ContainerRegistryUserState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ContainerRegistryUser(name, id, state, options);
    }
}
