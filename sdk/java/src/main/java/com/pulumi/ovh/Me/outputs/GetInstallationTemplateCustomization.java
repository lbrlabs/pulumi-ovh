// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetInstallationTemplateCustomization {
    /**
     * @return (DEPRECATED) Template change log details.
     * 
     * @deprecated
     * field is not used anymore
     * 
     */
    @Deprecated /* field is not used anymore */
    private String changeLog;
    /**
     * @return Set up the server using the provided hostname instead of the default hostname.
     * 
     */
    private String customHostname;
    /**
     * @return Indicate the URL where your postinstall customisation script is located.
     * 
     */
    private String postInstallationScriptLink;
    /**
     * @return indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is &#39;loh1Xee7eo OK OK OK UGh8Ang1Gu&#39;.
     * 
     */
    private String postInstallationScriptReturn;
    /**
     * @return (DEPRECATED) Rating.
     * 
     * @deprecated
     * field is not used anymore
     * 
     */
    @Deprecated /* field is not used anymore */
    private Integer rating;
    /**
     * @return Name of the ssh key that should be installed. Password login will be disabled.
     * 
     */
    private String sshKeyName;
    /**
     * @return Use the distribution&#39;s native kernel instead of the recommended OVHcloud Kernel.
     * 
     */
    private Boolean useDistributionKernel;

    private GetInstallationTemplateCustomization() {}
    /**
     * @return (DEPRECATED) Template change log details.
     * 
     * @deprecated
     * field is not used anymore
     * 
     */
    @Deprecated /* field is not used anymore */
    public String changeLog() {
        return this.changeLog;
    }
    /**
     * @return Set up the server using the provided hostname instead of the default hostname.
     * 
     */
    public String customHostname() {
        return this.customHostname;
    }
    /**
     * @return Indicate the URL where your postinstall customisation script is located.
     * 
     */
    public String postInstallationScriptLink() {
        return this.postInstallationScriptLink;
    }
    /**
     * @return indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is &#39;loh1Xee7eo OK OK OK UGh8Ang1Gu&#39;.
     * 
     */
    public String postInstallationScriptReturn() {
        return this.postInstallationScriptReturn;
    }
    /**
     * @return (DEPRECATED) Rating.
     * 
     * @deprecated
     * field is not used anymore
     * 
     */
    @Deprecated /* field is not used anymore */
    public Integer rating() {
        return this.rating;
    }
    /**
     * @return Name of the ssh key that should be installed. Password login will be disabled.
     * 
     */
    public String sshKeyName() {
        return this.sshKeyName;
    }
    /**
     * @return Use the distribution&#39;s native kernel instead of the recommended OVHcloud Kernel.
     * 
     */
    public Boolean useDistributionKernel() {
        return this.useDistributionKernel;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstallationTemplateCustomization defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String changeLog;
        private String customHostname;
        private String postInstallationScriptLink;
        private String postInstallationScriptReturn;
        private Integer rating;
        private String sshKeyName;
        private Boolean useDistributionKernel;
        public Builder() {}
        public Builder(GetInstallationTemplateCustomization defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.changeLog = defaults.changeLog;
    	      this.customHostname = defaults.customHostname;
    	      this.postInstallationScriptLink = defaults.postInstallationScriptLink;
    	      this.postInstallationScriptReturn = defaults.postInstallationScriptReturn;
    	      this.rating = defaults.rating;
    	      this.sshKeyName = defaults.sshKeyName;
    	      this.useDistributionKernel = defaults.useDistributionKernel;
        }

        @CustomType.Setter
        public Builder changeLog(String changeLog) {
            this.changeLog = Objects.requireNonNull(changeLog);
            return this;
        }
        @CustomType.Setter
        public Builder customHostname(String customHostname) {
            this.customHostname = Objects.requireNonNull(customHostname);
            return this;
        }
        @CustomType.Setter
        public Builder postInstallationScriptLink(String postInstallationScriptLink) {
            this.postInstallationScriptLink = Objects.requireNonNull(postInstallationScriptLink);
            return this;
        }
        @CustomType.Setter
        public Builder postInstallationScriptReturn(String postInstallationScriptReturn) {
            this.postInstallationScriptReturn = Objects.requireNonNull(postInstallationScriptReturn);
            return this;
        }
        @CustomType.Setter
        public Builder rating(Integer rating) {
            this.rating = Objects.requireNonNull(rating);
            return this;
        }
        @CustomType.Setter
        public Builder sshKeyName(String sshKeyName) {
            this.sshKeyName = Objects.requireNonNull(sshKeyName);
            return this;
        }
        @CustomType.Setter
        public Builder useDistributionKernel(Boolean useDistributionKernel) {
            this.useDistributionKernel = Objects.requireNonNull(useDistributionKernel);
            return this;
        }
        public GetInstallationTemplateCustomization build() {
            final var o = new GetInstallationTemplateCustomization();
            o.changeLog = changeLog;
            o.customHostname = customHostname;
            o.postInstallationScriptLink = postInstallationScriptLink;
            o.postInstallationScriptReturn = postInstallationScriptReturn;
            o.rating = rating;
            o.sshKeyName = sshKeyName;
            o.useDistributionKernel = useDistributionKernel;
            return o;
        }
    }
}
