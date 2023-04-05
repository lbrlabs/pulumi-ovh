// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.ovh.CloudProject.S3CredentialArgs;
import com.pulumi.ovh.CloudProject.inputs.S3CredentialState;
import com.pulumi.ovh.Utilities;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Creates an S3 Credential for a user in a public cloud project.
 * 
 * ## Import
 * 
 * OVHcloud User S3 Credentials can be imported using the `service_name`, `user_id` and `access_key_id` of the credential, separated by &#34;/&#34; E.g., bash
 * 
 * ```sh
 *  $ pulumi import ovh:CloudProject/s3Credential:S3Credential s3_credential service_name/user_id/access_key_id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProject/s3Credential:S3Credential")
public class S3Credential extends com.pulumi.resources.CustomResource {
    /**
     * the Access Key ID
     * 
     */
    @Export(name="accessKeyId", refs={String.class}, tree="[0]")
    private Output<String> accessKeyId;

    /**
     * @return the Access Key ID
     * 
     */
    public Output<String> accessKeyId() {
        return this.accessKeyId;
    }
    @Export(name="internalUserId", refs={String.class}, tree="[0]")
    private Output<String> internalUserId;

    public Output<String> internalUserId() {
        return this.internalUserId;
    }
    /**
     * (Sensitive) the Secret Access Key
     * 
     */
    @Export(name="secretAccessKey", refs={String.class}, tree="[0]")
    private Output<String> secretAccessKey;

    /**
     * @return (Sensitive) the Secret Access Key
     * 
     */
    public Output<String> secretAccessKey() {
        return this.secretAccessKey;
    }
    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * The ID of a public cloud project&#39;s user.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output<String> userId;

    /**
     * @return The ID of a public cloud project&#39;s user.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public S3Credential(String name) {
        this(name, S3CredentialArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public S3Credential(String name, S3CredentialArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public S3Credential(String name, S3CredentialArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/s3Credential:S3Credential", name, args == null ? S3CredentialArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private S3Credential(String name, Output<String> id, @Nullable S3CredentialState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/s3Credential:S3Credential", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secretAccessKey"
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
    public static S3Credential get(String name, Output<String> id, @Nullable S3CredentialState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new S3Credential(name, id, state, options);
    }
}
