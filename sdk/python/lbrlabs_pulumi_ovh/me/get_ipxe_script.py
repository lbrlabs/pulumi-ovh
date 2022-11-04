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
    'GetIpxeScriptResult',
    'AwaitableGetIpxeScriptResult',
    'get_ipxe_script',
    'get_ipxe_script_output',
]

@pulumi.output_type
class GetIpxeScriptResult:
    """
    A collection of values returned by getIpxeScript.
    """
    def __init__(__self__, id=None, name=None, script=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if script and not isinstance(script, str):
            raise TypeError("Expected argument 'script' to be a str")
        pulumi.set(__self__, "script", script)

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
    def script(self) -> str:
        """
        The content of the script.
        """
        return pulumi.get(self, "script")


class AwaitableGetIpxeScriptResult(GetIpxeScriptResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIpxeScriptResult(
            id=self.id,
            name=self.name,
            script=self.script)


def get_ipxe_script(name: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIpxeScriptResult:
    """
    Use this data source to retrieve information about an IPXE Script.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    script = ovh.Me.get_ipxe_script(name="myscript")
    ```


    :param str name: The name of the IPXE Script.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('ovh:Me/getIpxeScript:getIpxeScript', __args__, opts=opts, typ=GetIpxeScriptResult).value

    return AwaitableGetIpxeScriptResult(
        id=__ret__.id,
        name=__ret__.name,
        script=__ret__.script)


@_utilities.lift_output_func(get_ipxe_script)
def get_ipxe_script_output(name: Optional[pulumi.Input[str]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIpxeScriptResult]:
    """
    Use this data source to retrieve information about an IPXE Script.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_ovh as ovh

    script = ovh.Me.get_ipxe_script(name="myscript")
    ```


    :param str name: The name of the IPXE Script.
    """
    ...