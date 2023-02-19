# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetM3dbNamespaceResult',
    'AwaitableGetM3dbNamespaceResult',
    'get_m3db_namespace',
    'get_m3db_namespace_output',
]

@pulumi.output_type
class GetM3dbNamespaceResult:
    """
    A collection of values returned by getM3dbNamespace.
    """
    def __init__(__self__, cluster_id=None, id=None, name=None, resolution=None, retention_block_data_expiration_duration=None, retention_block_size_duration=None, retention_buffer_future_duration=None, retention_buffer_past_duration=None, retention_period_duration=None, service_name=None, snapshot_enabled=None, type=None, writes_to_commit_log_enabled=None):
        if cluster_id and not isinstance(cluster_id, str):
            raise TypeError("Expected argument 'cluster_id' to be a str")
        pulumi.set(__self__, "cluster_id", cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resolution and not isinstance(resolution, str):
            raise TypeError("Expected argument 'resolution' to be a str")
        pulumi.set(__self__, "resolution", resolution)
        if retention_block_data_expiration_duration and not isinstance(retention_block_data_expiration_duration, str):
            raise TypeError("Expected argument 'retention_block_data_expiration_duration' to be a str")
        pulumi.set(__self__, "retention_block_data_expiration_duration", retention_block_data_expiration_duration)
        if retention_block_size_duration and not isinstance(retention_block_size_duration, str):
            raise TypeError("Expected argument 'retention_block_size_duration' to be a str")
        pulumi.set(__self__, "retention_block_size_duration", retention_block_size_duration)
        if retention_buffer_future_duration and not isinstance(retention_buffer_future_duration, str):
            raise TypeError("Expected argument 'retention_buffer_future_duration' to be a str")
        pulumi.set(__self__, "retention_buffer_future_duration", retention_buffer_future_duration)
        if retention_buffer_past_duration and not isinstance(retention_buffer_past_duration, str):
            raise TypeError("Expected argument 'retention_buffer_past_duration' to be a str")
        pulumi.set(__self__, "retention_buffer_past_duration", retention_buffer_past_duration)
        if retention_period_duration and not isinstance(retention_period_duration, str):
            raise TypeError("Expected argument 'retention_period_duration' to be a str")
        pulumi.set(__self__, "retention_period_duration", retention_period_duration)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if snapshot_enabled and not isinstance(snapshot_enabled, bool):
            raise TypeError("Expected argument 'snapshot_enabled' to be a bool")
        pulumi.set(__self__, "snapshot_enabled", snapshot_enabled)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if writes_to_commit_log_enabled and not isinstance(writes_to_commit_log_enabled, bool):
            raise TypeError("Expected argument 'writes_to_commit_log_enabled' to be a bool")
        pulumi.set(__self__, "writes_to_commit_log_enabled", writes_to_commit_log_enabled)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def resolution(self) -> str:
        """
        Resolution for an aggregated namespace.
        """
        return pulumi.get(self, "resolution")

    @property
    @pulumi.getter(name="retentionBlockDataExpirationDuration")
    def retention_block_data_expiration_duration(self) -> str:
        """
        Controls how long we wait before expiring stale data.
        """
        return pulumi.get(self, "retention_block_data_expiration_duration")

    @property
    @pulumi.getter(name="retentionBlockSizeDuration")
    def retention_block_size_duration(self) -> str:
        """
        Controls how long to keep a block in memory before flushing to a fileset on disk.
        """
        return pulumi.get(self, "retention_block_size_duration")

    @property
    @pulumi.getter(name="retentionBufferFutureDuration")
    def retention_buffer_future_duration(self) -> str:
        """
        Controls how far into the future writes to the namespace will be accepted.
        """
        return pulumi.get(self, "retention_buffer_future_duration")

    @property
    @pulumi.getter(name="retentionBufferPastDuration")
    def retention_buffer_past_duration(self) -> str:
        """
        Controls how far into the past writes to the namespace will be accepted.
        """
        return pulumi.get(self, "retention_buffer_past_duration")

    @property
    @pulumi.getter(name="retentionPeriodDuration")
    def retention_period_duration(self) -> str:
        """
        Controls the duration of time that M3DB will retain data for the namespace.
        """
        return pulumi.get(self, "retention_period_duration")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        """
        See Argument Reference above.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter(name="snapshotEnabled")
    def snapshot_enabled(self) -> bool:
        """
        SDefines whether M3db will create snapshot files for this namespace.
        """
        return pulumi.get(self, "snapshot_enabled")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        Type of namespace.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="writesToCommitLogEnabled")
    def writes_to_commit_log_enabled(self) -> bool:
        """
        Defines whether M3DB will include writes to this namespace in the commit log.
        """
        return pulumi.get(self, "writes_to_commit_log_enabled")


class AwaitableGetM3dbNamespaceResult(GetM3dbNamespaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetM3dbNamespaceResult(
            cluster_id=self.cluster_id,
            id=self.id,
            name=self.name,
            resolution=self.resolution,
            retention_block_data_expiration_duration=self.retention_block_data_expiration_duration,
            retention_block_size_duration=self.retention_block_size_duration,
            retention_buffer_future_duration=self.retention_buffer_future_duration,
            retention_buffer_past_duration=self.retention_buffer_past_duration,
            retention_period_duration=self.retention_period_duration,
            service_name=self.service_name,
            snapshot_enabled=self.snapshot_enabled,
            type=self.type,
            writes_to_commit_log_enabled=self.writes_to_commit_log_enabled)


def get_m3db_namespace(cluster_id: Optional[str] = None,
                       name: Optional[str] = None,
                       service_name: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetM3dbNamespaceResult:
    """
    Use this data source to get information about a namespace of a M3DB cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    m3dbnamespace = ovh.CloudProject.get_m3db_namespace(service_name="XXX",
        cluster_id="YYY",
        name="ZZZ")
    pulumi.export("m3dbnamespaceType", m3dbnamespace.type)
    ```


    :param str cluster_id: Cluster ID
    :param str name: Name of the namespace.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    __args__ = dict()
    __args__['clusterId'] = cluster_id
    __args__['name'] = name
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:CloudProject/getM3dbNamespace:getM3dbNamespace', __args__, opts=opts, typ=GetM3dbNamespaceResult).value

    return AwaitableGetM3dbNamespaceResult(
        cluster_id=__ret__.cluster_id,
        id=__ret__.id,
        name=__ret__.name,
        resolution=__ret__.resolution,
        retention_block_data_expiration_duration=__ret__.retention_block_data_expiration_duration,
        retention_block_size_duration=__ret__.retention_block_size_duration,
        retention_buffer_future_duration=__ret__.retention_buffer_future_duration,
        retention_buffer_past_duration=__ret__.retention_buffer_past_duration,
        retention_period_duration=__ret__.retention_period_duration,
        service_name=__ret__.service_name,
        snapshot_enabled=__ret__.snapshot_enabled,
        type=__ret__.type,
        writes_to_commit_log_enabled=__ret__.writes_to_commit_log_enabled)


@_utilities.lift_output_func(get_m3db_namespace)
def get_m3db_namespace_output(cluster_id: Optional[pulumi.Input[str]] = None,
                              name: Optional[pulumi.Input[str]] = None,
                              service_name: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetM3dbNamespaceResult]:
    """
    Use this data source to get information about a namespace of a M3DB cluster associated with a public cloud project.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    m3dbnamespace = ovh.CloudProject.get_m3db_namespace(service_name="XXX",
        cluster_id="YYY",
        name="ZZZ")
    pulumi.export("m3dbnamespaceType", m3dbnamespace.type)
    ```


    :param str cluster_id: Cluster ID
    :param str name: Name of the namespace.
    :param str service_name: The id of the public cloud project. If omitted,
           the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
    """
    ...
