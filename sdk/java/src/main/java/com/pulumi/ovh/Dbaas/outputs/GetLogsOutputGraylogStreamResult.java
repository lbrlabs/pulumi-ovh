// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.ovh.Dbaas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLogsOutputGraylogStreamResult {
    private Boolean canAlert;
    /**
     * @return Cold storage compression method
     * 
     */
    private String coldStorageCompression;
    /**
     * @return ColdStorage content
     * 
     */
    private String coldStorageContent;
    /**
     * @return Is Cold storage enabled?
     * 
     */
    private Boolean coldStorageEnabled;
    /**
     * @return Notify on new Cold storage archive
     * 
     */
    private Boolean coldStorageNotifyEnabled;
    /**
     * @return Cold storage retention in year
     * 
     */
    private Integer coldStorageRetention;
    /**
     * @return ColdStorage destination
     * 
     */
    private String coldStorageTarget;
    /**
     * @return Stream creation
     * 
     */
    private String createdAt;
    /**
     * @return Stream description
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Enable ES indexing
     * 
     */
    private Boolean indexingEnabled;
    /**
     * @return Maximum indexing size (in GB)
     * 
     */
    private Integer indexingMaxSize;
    /**
     * @return If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
     * 
     */
    private Boolean indexingNotifyEnabled;
    /**
     * @return Indicates if you are allowed to edit entry
     * 
     */
    private Boolean isEditable;
    /**
     * @return Indicates if you are allowed to share entry
     * 
     */
    private Boolean isShareable;
    /**
     * @return Number of alert condition
     * 
     */
    private Integer nbAlertCondition;
    /**
     * @return Number of coldstored archives
     * 
     */
    private Integer nbArchive;
    /**
     * @return Parent stream ID
     * 
     */
    private String parentStreamId;
    /**
     * @return If set, pause indexing when maximum size is reach
     * 
     */
    private Boolean pauseIndexingOnMaxSize;
    /**
     * @return Retention ID
     * 
     */
    private String retentionId;
    private String serviceName;
    /**
     * @return Stream ID
     * 
     */
    private String streamId;
    private String title;
    /**
     * @return Stream last update
     * 
     */
    private String updatedAt;
    /**
     * @return Enable Websocket
     * 
     */
    private Boolean webSocketEnabled;

    private GetLogsOutputGraylogStreamResult() {}
    public Boolean canAlert() {
        return this.canAlert;
    }
    /**
     * @return Cold storage compression method
     * 
     */
    public String coldStorageCompression() {
        return this.coldStorageCompression;
    }
    /**
     * @return ColdStorage content
     * 
     */
    public String coldStorageContent() {
        return this.coldStorageContent;
    }
    /**
     * @return Is Cold storage enabled?
     * 
     */
    public Boolean coldStorageEnabled() {
        return this.coldStorageEnabled;
    }
    /**
     * @return Notify on new Cold storage archive
     * 
     */
    public Boolean coldStorageNotifyEnabled() {
        return this.coldStorageNotifyEnabled;
    }
    /**
     * @return Cold storage retention in year
     * 
     */
    public Integer coldStorageRetention() {
        return this.coldStorageRetention;
    }
    /**
     * @return ColdStorage destination
     * 
     */
    public String coldStorageTarget() {
        return this.coldStorageTarget;
    }
    /**
     * @return Stream creation
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return Stream description
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Enable ES indexing
     * 
     */
    public Boolean indexingEnabled() {
        return this.indexingEnabled;
    }
    /**
     * @return Maximum indexing size (in GB)
     * 
     */
    public Integer indexingMaxSize() {
        return this.indexingMaxSize;
    }
    /**
     * @return If set, notify when size is near 80, 90 or 100 % of the maximum configured setting
     * 
     */
    public Boolean indexingNotifyEnabled() {
        return this.indexingNotifyEnabled;
    }
    /**
     * @return Indicates if you are allowed to edit entry
     * 
     */
    public Boolean isEditable() {
        return this.isEditable;
    }
    /**
     * @return Indicates if you are allowed to share entry
     * 
     */
    public Boolean isShareable() {
        return this.isShareable;
    }
    /**
     * @return Number of alert condition
     * 
     */
    public Integer nbAlertCondition() {
        return this.nbAlertCondition;
    }
    /**
     * @return Number of coldstored archives
     * 
     */
    public Integer nbArchive() {
        return this.nbArchive;
    }
    /**
     * @return Parent stream ID
     * 
     */
    public String parentStreamId() {
        return this.parentStreamId;
    }
    /**
     * @return If set, pause indexing when maximum size is reach
     * 
     */
    public Boolean pauseIndexingOnMaxSize() {
        return this.pauseIndexingOnMaxSize;
    }
    /**
     * @return Retention ID
     * 
     */
    public String retentionId() {
        return this.retentionId;
    }
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Stream ID
     * 
     */
    public String streamId() {
        return this.streamId;
    }
    public String title() {
        return this.title;
    }
    /**
     * @return Stream last update
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return Enable Websocket
     * 
     */
    public Boolean webSocketEnabled() {
        return this.webSocketEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLogsOutputGraylogStreamResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean canAlert;
        private String coldStorageCompression;
        private String coldStorageContent;
        private Boolean coldStorageEnabled;
        private Boolean coldStorageNotifyEnabled;
        private Integer coldStorageRetention;
        private String coldStorageTarget;
        private String createdAt;
        private String description;
        private String id;
        private Boolean indexingEnabled;
        private Integer indexingMaxSize;
        private Boolean indexingNotifyEnabled;
        private Boolean isEditable;
        private Boolean isShareable;
        private Integer nbAlertCondition;
        private Integer nbArchive;
        private String parentStreamId;
        private Boolean pauseIndexingOnMaxSize;
        private String retentionId;
        private String serviceName;
        private String streamId;
        private String title;
        private String updatedAt;
        private Boolean webSocketEnabled;
        public Builder() {}
        public Builder(GetLogsOutputGraylogStreamResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.canAlert = defaults.canAlert;
    	      this.coldStorageCompression = defaults.coldStorageCompression;
    	      this.coldStorageContent = defaults.coldStorageContent;
    	      this.coldStorageEnabled = defaults.coldStorageEnabled;
    	      this.coldStorageNotifyEnabled = defaults.coldStorageNotifyEnabled;
    	      this.coldStorageRetention = defaults.coldStorageRetention;
    	      this.coldStorageTarget = defaults.coldStorageTarget;
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.indexingEnabled = defaults.indexingEnabled;
    	      this.indexingMaxSize = defaults.indexingMaxSize;
    	      this.indexingNotifyEnabled = defaults.indexingNotifyEnabled;
    	      this.isEditable = defaults.isEditable;
    	      this.isShareable = defaults.isShareable;
    	      this.nbAlertCondition = defaults.nbAlertCondition;
    	      this.nbArchive = defaults.nbArchive;
    	      this.parentStreamId = defaults.parentStreamId;
    	      this.pauseIndexingOnMaxSize = defaults.pauseIndexingOnMaxSize;
    	      this.retentionId = defaults.retentionId;
    	      this.serviceName = defaults.serviceName;
    	      this.streamId = defaults.streamId;
    	      this.title = defaults.title;
    	      this.updatedAt = defaults.updatedAt;
    	      this.webSocketEnabled = defaults.webSocketEnabled;
        }

        @CustomType.Setter
        public Builder canAlert(Boolean canAlert) {
            this.canAlert = Objects.requireNonNull(canAlert);
            return this;
        }
        @CustomType.Setter
        public Builder coldStorageCompression(String coldStorageCompression) {
            this.coldStorageCompression = Objects.requireNonNull(coldStorageCompression);
            return this;
        }
        @CustomType.Setter
        public Builder coldStorageContent(String coldStorageContent) {
            this.coldStorageContent = Objects.requireNonNull(coldStorageContent);
            return this;
        }
        @CustomType.Setter
        public Builder coldStorageEnabled(Boolean coldStorageEnabled) {
            this.coldStorageEnabled = Objects.requireNonNull(coldStorageEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder coldStorageNotifyEnabled(Boolean coldStorageNotifyEnabled) {
            this.coldStorageNotifyEnabled = Objects.requireNonNull(coldStorageNotifyEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder coldStorageRetention(Integer coldStorageRetention) {
            this.coldStorageRetention = Objects.requireNonNull(coldStorageRetention);
            return this;
        }
        @CustomType.Setter
        public Builder coldStorageTarget(String coldStorageTarget) {
            this.coldStorageTarget = Objects.requireNonNull(coldStorageTarget);
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            this.createdAt = Objects.requireNonNull(createdAt);
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder indexingEnabled(Boolean indexingEnabled) {
            this.indexingEnabled = Objects.requireNonNull(indexingEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder indexingMaxSize(Integer indexingMaxSize) {
            this.indexingMaxSize = Objects.requireNonNull(indexingMaxSize);
            return this;
        }
        @CustomType.Setter
        public Builder indexingNotifyEnabled(Boolean indexingNotifyEnabled) {
            this.indexingNotifyEnabled = Objects.requireNonNull(indexingNotifyEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder isEditable(Boolean isEditable) {
            this.isEditable = Objects.requireNonNull(isEditable);
            return this;
        }
        @CustomType.Setter
        public Builder isShareable(Boolean isShareable) {
            this.isShareable = Objects.requireNonNull(isShareable);
            return this;
        }
        @CustomType.Setter
        public Builder nbAlertCondition(Integer nbAlertCondition) {
            this.nbAlertCondition = Objects.requireNonNull(nbAlertCondition);
            return this;
        }
        @CustomType.Setter
        public Builder nbArchive(Integer nbArchive) {
            this.nbArchive = Objects.requireNonNull(nbArchive);
            return this;
        }
        @CustomType.Setter
        public Builder parentStreamId(String parentStreamId) {
            this.parentStreamId = Objects.requireNonNull(parentStreamId);
            return this;
        }
        @CustomType.Setter
        public Builder pauseIndexingOnMaxSize(Boolean pauseIndexingOnMaxSize) {
            this.pauseIndexingOnMaxSize = Objects.requireNonNull(pauseIndexingOnMaxSize);
            return this;
        }
        @CustomType.Setter
        public Builder retentionId(String retentionId) {
            this.retentionId = Objects.requireNonNull(retentionId);
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            this.serviceName = Objects.requireNonNull(serviceName);
            return this;
        }
        @CustomType.Setter
        public Builder streamId(String streamId) {
            this.streamId = Objects.requireNonNull(streamId);
            return this;
        }
        @CustomType.Setter
        public Builder title(String title) {
            this.title = Objects.requireNonNull(title);
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            this.updatedAt = Objects.requireNonNull(updatedAt);
            return this;
        }
        @CustomType.Setter
        public Builder webSocketEnabled(Boolean webSocketEnabled) {
            this.webSocketEnabled = Objects.requireNonNull(webSocketEnabled);
            return this;
        }
        public GetLogsOutputGraylogStreamResult build() {
            final var o = new GetLogsOutputGraylogStreamResult();
            o.canAlert = canAlert;
            o.coldStorageCompression = coldStorageCompression;
            o.coldStorageContent = coldStorageContent;
            o.coldStorageEnabled = coldStorageEnabled;
            o.coldStorageNotifyEnabled = coldStorageNotifyEnabled;
            o.coldStorageRetention = coldStorageRetention;
            o.coldStorageTarget = coldStorageTarget;
            o.createdAt = createdAt;
            o.description = description;
            o.id = id;
            o.indexingEnabled = indexingEnabled;
            o.indexingMaxSize = indexingMaxSize;
            o.indexingNotifyEnabled = indexingNotifyEnabled;
            o.isEditable = isEditable;
            o.isShareable = isShareable;
            o.nbAlertCondition = nbAlertCondition;
            o.nbArchive = nbArchive;
            o.parentStreamId = parentStreamId;
            o.pauseIndexingOnMaxSize = pauseIndexingOnMaxSize;
            o.retentionId = retentionId;
            o.serviceName = serviceName;
            o.streamId = streamId;
            o.title = title;
            o.updatedAt = updatedAt;
            o.webSocketEnabled = webSocketEnabled;
            return o;
        }
    }
}
