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
    'GetPrivateDatabaseResult',
    'AwaitableGetPrivateDatabaseResult',
    'get_private_database',
    'get_private_database_output',
]

@pulumi.output_type
class GetPrivateDatabaseResult:
    """
    A collection of values returned by getPrivateDatabase.
    """
    def __init__(__self__, cpu=None, datacenter=None, display_name=None, hostname=None, hostname_ftp=None, id=None, infrastructure=None, offer=None, port=None, port_ftp=None, quota_size=None, quota_used=None, ram=None, server=None, service_name=None, state=None, type=None, version=None, version_label=None, version_number=None):
        if cpu and not isinstance(cpu, int):
            raise TypeError("Expected argument 'cpu' to be a int")
        pulumi.set(__self__, "cpu", cpu)
        if datacenter and not isinstance(datacenter, str):
            raise TypeError("Expected argument 'datacenter' to be a str")
        pulumi.set(__self__, "datacenter", datacenter)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if hostname and not isinstance(hostname, str):
            raise TypeError("Expected argument 'hostname' to be a str")
        pulumi.set(__self__, "hostname", hostname)
        if hostname_ftp and not isinstance(hostname_ftp, str):
            raise TypeError("Expected argument 'hostname_ftp' to be a str")
        pulumi.set(__self__, "hostname_ftp", hostname_ftp)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if infrastructure and not isinstance(infrastructure, str):
            raise TypeError("Expected argument 'infrastructure' to be a str")
        pulumi.set(__self__, "infrastructure", infrastructure)
        if offer and not isinstance(offer, str):
            raise TypeError("Expected argument 'offer' to be a str")
        pulumi.set(__self__, "offer", offer)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if port_ftp and not isinstance(port_ftp, int):
            raise TypeError("Expected argument 'port_ftp' to be a int")
        pulumi.set(__self__, "port_ftp", port_ftp)
        if quota_size and not isinstance(quota_size, int):
            raise TypeError("Expected argument 'quota_size' to be a int")
        pulumi.set(__self__, "quota_size", quota_size)
        if quota_used and not isinstance(quota_used, int):
            raise TypeError("Expected argument 'quota_used' to be a int")
        pulumi.set(__self__, "quota_used", quota_used)
        if ram and not isinstance(ram, int):
            raise TypeError("Expected argument 'ram' to be a int")
        pulumi.set(__self__, "ram", ram)
        if server and not isinstance(server, str):
            raise TypeError("Expected argument 'server' to be a str")
        pulumi.set(__self__, "server", server)
        if service_name and not isinstance(service_name, str):
            raise TypeError("Expected argument 'service_name' to be a str")
        pulumi.set(__self__, "service_name", service_name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)
        if version_label and not isinstance(version_label, str):
            raise TypeError("Expected argument 'version_label' to be a str")
        pulumi.set(__self__, "version_label", version_label)
        if version_number and not isinstance(version_number, float):
            raise TypeError("Expected argument 'version_number' to be a float")
        pulumi.set(__self__, "version_number", version_number)

    @property
    @pulumi.getter
    def cpu(self) -> int:
        """
        Number of CPU on your private database
        """
        return pulumi.get(self, "cpu")

    @property
    @pulumi.getter
    def datacenter(self) -> str:
        """
        Datacenter where this private database is located
        """
        return pulumi.get(self, "datacenter")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Name displayed in customer panel for your private database
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        Private database hostname
        """
        return pulumi.get(self, "hostname")

    @property
    @pulumi.getter(name="hostnameFtp")
    def hostname_ftp(self) -> str:
        """
        Private database FTP hostname
        """
        return pulumi.get(self, "hostname_ftp")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def infrastructure(self) -> str:
        """
        Infrastructure where service was stored
        """
        return pulumi.get(self, "infrastructure")

    @property
    @pulumi.getter
    def offer(self) -> str:
        """
        Type of the private database offer
        """
        return pulumi.get(self, "offer")

    @property
    @pulumi.getter
    def port(self) -> int:
        """
        Private database service port
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="portFtp")
    def port_ftp(self) -> int:
        """
        Private database FTP port
        """
        return pulumi.get(self, "port_ftp")

    @property
    @pulumi.getter(name="quotaSize")
    def quota_size(self) -> int:
        """
        Space allowed (in MB) on your private database
        """
        return pulumi.get(self, "quota_size")

    @property
    @pulumi.getter(name="quotaUsed")
    def quota_used(self) -> int:
        """
        Sapce used (in MB) on your private database
        """
        return pulumi.get(self, "quota_used")

    @property
    @pulumi.getter
    def ram(self) -> int:
        """
        Amount of ram (in MB) on your private database
        """
        return pulumi.get(self, "ram")

    @property
    @pulumi.getter
    def server(self) -> str:
        """
        Private database server name
        """
        return pulumi.get(self, "server")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> str:
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Private database state
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Private database available versions
        """
        return pulumi.get(self, "version")

    @property
    @pulumi.getter(name="versionLabel")
    def version_label(self) -> str:
        """
        Private database version label
        """
        return pulumi.get(self, "version_label")

    @property
    @pulumi.getter(name="versionNumber")
    def version_number(self) -> float:
        """
        Private database version number
        """
        return pulumi.get(self, "version_number")


class AwaitableGetPrivateDatabaseResult(GetPrivateDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPrivateDatabaseResult(
            cpu=self.cpu,
            datacenter=self.datacenter,
            display_name=self.display_name,
            hostname=self.hostname,
            hostname_ftp=self.hostname_ftp,
            id=self.id,
            infrastructure=self.infrastructure,
            offer=self.offer,
            port=self.port,
            port_ftp=self.port_ftp,
            quota_size=self.quota_size,
            quota_used=self.quota_used,
            ram=self.ram,
            server=self.server,
            service_name=self.service_name,
            state=self.state,
            type=self.type,
            version=self.version,
            version_label=self.version_label,
            version_number=self.version_number)


def get_private_database(service_name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPrivateDatabaseResult:
    """
    Use this data source to retrieve information about an hosting database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    database = ovh.Hosting.get_private_database(service_name="XXXXXX")
    ```


    :param str service_name: The internal name of your private database
    """
    __args__ = dict()
    __args__['serviceName'] = service_name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Hosting/getPrivateDatabase:getPrivateDatabase', __args__, opts=opts, typ=GetPrivateDatabaseResult).value

    return AwaitableGetPrivateDatabaseResult(
        cpu=__ret__.cpu,
        datacenter=__ret__.datacenter,
        display_name=__ret__.display_name,
        hostname=__ret__.hostname,
        hostname_ftp=__ret__.hostname_ftp,
        id=__ret__.id,
        infrastructure=__ret__.infrastructure,
        offer=__ret__.offer,
        port=__ret__.port,
        port_ftp=__ret__.port_ftp,
        quota_size=__ret__.quota_size,
        quota_used=__ret__.quota_used,
        ram=__ret__.ram,
        server=__ret__.server,
        service_name=__ret__.service_name,
        state=__ret__.state,
        type=__ret__.type,
        version=__ret__.version,
        version_label=__ret__.version_label,
        version_number=__ret__.version_number)


@_utilities.lift_output_func(get_private_database)
def get_private_database_output(service_name: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPrivateDatabaseResult]:
    """
    Use this data source to retrieve information about an hosting database.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    database = ovh.Hosting.get_private_database(service_name="XXXXXX")
    ```


    :param str service_name: The internal name of your private database
    """
    ...
