// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetUserS3CredentialArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserS3CredentialArgs Empty = new GetUserS3CredentialArgs();

    /**
     * the Access Key ID
     * 
     */
    @Import(name="accessKeyId", required=true)
    private Output<String> accessKeyId;

    /**
     * @return the Access Key ID
     * 
     */
    public Output<String> accessKeyId() {
        return this.accessKeyId;
    }

    /**
     * The ID of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
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
    @Import(name="userId", required=true)
    private Output<String> userId;

    /**
     * @return The ID of a public cloud project&#39;s user.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    private GetUserS3CredentialArgs() {}

    private GetUserS3CredentialArgs(GetUserS3CredentialArgs $) {
        this.accessKeyId = $.accessKeyId;
        this.serviceName = $.serviceName;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserS3CredentialArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserS3CredentialArgs $;

        public Builder() {
            $ = new GetUserS3CredentialArgs();
        }

        public Builder(GetUserS3CredentialArgs defaults) {
            $ = new GetUserS3CredentialArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessKeyId the Access Key ID
         * 
         * @return builder
         * 
         */
        public Builder accessKeyId(Output<String> accessKeyId) {
            $.accessKeyId = accessKeyId;
            return this;
        }

        /**
         * @param accessKeyId the Access Key ID
         * 
         * @return builder
         * 
         */
        public Builder accessKeyId(String accessKeyId) {
            return accessKeyId(Output.of(accessKeyId));
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param userId The ID of a public cloud project&#39;s user.
         * 
         * @return builder
         * 
         */
        public Builder userId(Output<String> userId) {
            $.userId = userId;
            return this;
        }

        /**
         * @param userId The ID of a public cloud project&#39;s user.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            return userId(Output.of(userId));
        }

        public GetUserS3CredentialArgs build() {
            $.accessKeyId = Objects.requireNonNull($.accessKeyId, "expected parameter 'accessKeyId' to be non-null");
            $.serviceName = Objects.requireNonNull($.serviceName, "expected parameter 'serviceName' to be non-null");
            $.userId = Objects.requireNonNull($.userId, "expected parameter 'userId' to be non-null");
            return $;
        }
    }

}