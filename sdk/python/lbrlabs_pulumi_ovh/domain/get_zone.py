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
    'GetZoneResult',
    'AwaitableGetZoneResult',
    'get_zone',
    'get_zone_output',
]

@pulumi.output_type
class GetZoneResult:
    """
    A collection of values returned by getZone.
    """
    def __init__(__self__, dnssec_supported=None, has_dns_anycast=None, id=None, last_update=None, name=None, name_servers=None, urn=None):
        if dnssec_supported and not isinstance(dnssec_supported, bool):
            raise TypeError("Expected argument 'dnssec_supported' to be a bool")
        pulumi.set(__self__, "dnssec_supported", dnssec_supported)
        if has_dns_anycast and not isinstance(has_dns_anycast, bool):
            raise TypeError("Expected argument 'has_dns_anycast' to be a bool")
        pulumi.set(__self__, "has_dns_anycast", has_dns_anycast)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_update and not isinstance(last_update, str):
            raise TypeError("Expected argument 'last_update' to be a str")
        pulumi.set(__self__, "last_update", last_update)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_servers and not isinstance(name_servers, list):
            raise TypeError("Expected argument 'name_servers' to be a list")
        pulumi.set(__self__, "name_servers", name_servers)
        if urn and not isinstance(urn, str):
            raise TypeError("Expected argument 'urn' to be a str")
        pulumi.set(__self__, "urn", urn)

    @property
    @pulumi.getter(name="dnssecSupported")
    def dnssec_supported(self) -> bool:
        """
        Is DNSSEC supported by this zone
        """
        return pulumi.get(self, "dnssec_supported")

    @property
    @pulumi.getter(name="hasDnsAnycast")
    def has_dns_anycast(self) -> bool:
        """
        hasDnsAnycast flag of the DNS zone
        """
        return pulumi.get(self, "has_dns_anycast")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastUpdate")
    def last_update(self) -> str:
        """
        Last update date of the DNS zone
        """
        return pulumi.get(self, "last_update")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameServers")
    def name_servers(self) -> Sequence[str]:
        """
        Name servers that host the DNS zone
        """
        return pulumi.get(self, "name_servers")

    @property
    @pulumi.getter
    def urn(self) -> str:
        """
        URN of the DNS zone
        """
        return pulumi.get(self, "urn")


class AwaitableGetZoneResult(GetZoneResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetZoneResult(
            dnssec_supported=self.dnssec_supported,
            has_dns_anycast=self.has_dns_anycast,
            id=self.id,
            last_update=self.last_update,
            name=self.name,
            name_servers=self.name_servers,
            urn=self.urn)


def get_zone(name: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetZoneResult:
    """
    Use this data source to retrieve information about a domain zone.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    rootzone = ovh.Domain.get_zone(name="mysite.ovh")
    ```


    :param str name: The name of the domain zone.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Domain/getZone:getZone', __args__, opts=opts, typ=GetZoneResult).value

    return AwaitableGetZoneResult(
        dnssec_supported=pulumi.get(__ret__, 'dnssec_supported'),
        has_dns_anycast=pulumi.get(__ret__, 'has_dns_anycast'),
        id=pulumi.get(__ret__, 'id'),
        last_update=pulumi.get(__ret__, 'last_update'),
        name=pulumi.get(__ret__, 'name'),
        name_servers=pulumi.get(__ret__, 'name_servers'),
        urn=pulumi.get(__ret__, 'urn'))


@_utilities.lift_output_func(get_zone)
def get_zone_output(name: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetZoneResult]:
    """
    Use this data source to retrieve information about a domain zone.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    rootzone = ovh.Domain.get_zone(name="mysite.ovh")
    ```


    :param str name: The name of the domain zone.
    """
    ...
