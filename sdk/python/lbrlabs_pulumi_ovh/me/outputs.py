# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'InstallationTemplateCustomization',
    'GetInstallationTemplateCustomizationResult',
    'GetInstallationTemplatePartitionSchemeResult',
    'GetInstallationTemplatePartitionSchemeHardwareRaidResult',
    'GetInstallationTemplatePartitionSchemePartitionResult',
    'GetMeCurrencyResult',
]

@pulumi.output_type
class InstallationTemplateCustomization(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "changeLog":
            suggest = "change_log"
        elif key == "customHostname":
            suggest = "custom_hostname"
        elif key == "postInstallationScriptLink":
            suggest = "post_installation_script_link"
        elif key == "postInstallationScriptReturn":
            suggest = "post_installation_script_return"
        elif key == "sshKeyName":
            suggest = "ssh_key_name"
        elif key == "useDistributionKernel":
            suggest = "use_distribution_kernel"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in InstallationTemplateCustomization. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        InstallationTemplateCustomization.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        InstallationTemplateCustomization.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 change_log: Optional[str] = None,
                 custom_hostname: Optional[str] = None,
                 post_installation_script_link: Optional[str] = None,
                 post_installation_script_return: Optional[str] = None,
                 rating: Optional[int] = None,
                 ssh_key_name: Optional[str] = None,
                 use_distribution_kernel: Optional[bool] = None):
        if change_log is not None:
            pulumi.set(__self__, "change_log", change_log)
        if custom_hostname is not None:
            pulumi.set(__self__, "custom_hostname", custom_hostname)
        if post_installation_script_link is not None:
            pulumi.set(__self__, "post_installation_script_link", post_installation_script_link)
        if post_installation_script_return is not None:
            pulumi.set(__self__, "post_installation_script_return", post_installation_script_return)
        if rating is not None:
            pulumi.set(__self__, "rating", rating)
        if ssh_key_name is not None:
            pulumi.set(__self__, "ssh_key_name", ssh_key_name)
        if use_distribution_kernel is not None:
            pulumi.set(__self__, "use_distribution_kernel", use_distribution_kernel)

    @property
    @pulumi.getter(name="changeLog")
    def change_log(self) -> Optional[str]:
        return pulumi.get(self, "change_log")

    @property
    @pulumi.getter(name="customHostname")
    def custom_hostname(self) -> Optional[str]:
        return pulumi.get(self, "custom_hostname")

    @property
    @pulumi.getter(name="postInstallationScriptLink")
    def post_installation_script_link(self) -> Optional[str]:
        return pulumi.get(self, "post_installation_script_link")

    @property
    @pulumi.getter(name="postInstallationScriptReturn")
    def post_installation_script_return(self) -> Optional[str]:
        return pulumi.get(self, "post_installation_script_return")

    @property
    @pulumi.getter
    def rating(self) -> Optional[int]:
        return pulumi.get(self, "rating")

    @property
    @pulumi.getter(name="sshKeyName")
    def ssh_key_name(self) -> Optional[str]:
        return pulumi.get(self, "ssh_key_name")

    @property
    @pulumi.getter(name="useDistributionKernel")
    def use_distribution_kernel(self) -> Optional[bool]:
        return pulumi.get(self, "use_distribution_kernel")


@pulumi.output_type
class GetInstallationTemplateCustomizationResult(dict):
    def __init__(__self__, *,
                 change_log: str,
                 custom_hostname: str,
                 post_installation_script_link: str,
                 post_installation_script_return: str,
                 rating: int,
                 ssh_key_name: str,
                 use_distribution_kernel: bool):
        pulumi.set(__self__, "change_log", change_log)
        pulumi.set(__self__, "custom_hostname", custom_hostname)
        pulumi.set(__self__, "post_installation_script_link", post_installation_script_link)
        pulumi.set(__self__, "post_installation_script_return", post_installation_script_return)
        pulumi.set(__self__, "rating", rating)
        pulumi.set(__self__, "ssh_key_name", ssh_key_name)
        pulumi.set(__self__, "use_distribution_kernel", use_distribution_kernel)

    @property
    @pulumi.getter(name="changeLog")
    def change_log(self) -> str:
        return pulumi.get(self, "change_log")

    @property
    @pulumi.getter(name="customHostname")
    def custom_hostname(self) -> str:
        return pulumi.get(self, "custom_hostname")

    @property
    @pulumi.getter(name="postInstallationScriptLink")
    def post_installation_script_link(self) -> str:
        return pulumi.get(self, "post_installation_script_link")

    @property
    @pulumi.getter(name="postInstallationScriptReturn")
    def post_installation_script_return(self) -> str:
        return pulumi.get(self, "post_installation_script_return")

    @property
    @pulumi.getter
    def rating(self) -> int:
        return pulumi.get(self, "rating")

    @property
    @pulumi.getter(name="sshKeyName")
    def ssh_key_name(self) -> str:
        return pulumi.get(self, "ssh_key_name")

    @property
    @pulumi.getter(name="useDistributionKernel")
    def use_distribution_kernel(self) -> bool:
        return pulumi.get(self, "use_distribution_kernel")


@pulumi.output_type
class GetInstallationTemplatePartitionSchemeResult(dict):
    def __init__(__self__, *,
                 hardware_raids: Sequence['outputs.GetInstallationTemplatePartitionSchemeHardwareRaidResult'],
                 name: str,
                 partitions: Sequence['outputs.GetInstallationTemplatePartitionSchemePartitionResult'],
                 priority: int):
        pulumi.set(__self__, "hardware_raids", hardware_raids)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "partitions", partitions)
        pulumi.set(__self__, "priority", priority)

    @property
    @pulumi.getter(name="hardwareRaids")
    def hardware_raids(self) -> Sequence['outputs.GetInstallationTemplatePartitionSchemeHardwareRaidResult']:
        return pulumi.get(self, "hardware_raids")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def partitions(self) -> Sequence['outputs.GetInstallationTemplatePartitionSchemePartitionResult']:
        return pulumi.get(self, "partitions")

    @property
    @pulumi.getter
    def priority(self) -> int:
        return pulumi.get(self, "priority")


@pulumi.output_type
class GetInstallationTemplatePartitionSchemeHardwareRaidResult(dict):
    def __init__(__self__, *,
                 disks: Sequence[str],
                 mode: str,
                 name: str,
                 step: int):
        pulumi.set(__self__, "disks", disks)
        pulumi.set(__self__, "mode", mode)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "step", step)

    @property
    @pulumi.getter
    def disks(self) -> Sequence[str]:
        return pulumi.get(self, "disks")

    @property
    @pulumi.getter
    def mode(self) -> str:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def step(self) -> int:
        return pulumi.get(self, "step")


@pulumi.output_type
class GetInstallationTemplatePartitionSchemePartitionResult(dict):
    def __init__(__self__, *,
                 filesystem: str,
                 mountpoint: str,
                 order: int,
                 raid: str,
                 size: int,
                 type: str,
                 volume_name: str):
        pulumi.set(__self__, "filesystem", filesystem)
        pulumi.set(__self__, "mountpoint", mountpoint)
        pulumi.set(__self__, "order", order)
        pulumi.set(__self__, "raid", raid)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "volume_name", volume_name)

    @property
    @pulumi.getter
    def filesystem(self) -> str:
        return pulumi.get(self, "filesystem")

    @property
    @pulumi.getter
    def mountpoint(self) -> str:
        return pulumi.get(self, "mountpoint")

    @property
    @pulumi.getter
    def order(self) -> int:
        return pulumi.get(self, "order")

    @property
    @pulumi.getter
    def raid(self) -> str:
        return pulumi.get(self, "raid")

    @property
    @pulumi.getter
    def size(self) -> int:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="volumeName")
    def volume_name(self) -> str:
        return pulumi.get(self, "volume_name")


@pulumi.output_type
class GetMeCurrencyResult(dict):
    def __init__(__self__, *,
                 code: str,
                 symbol: str):
        pulumi.set(__self__, "code", code)
        pulumi.set(__self__, "symbol", symbol)

    @property
    @pulumi.getter
    def code(self) -> str:
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def symbol(self) -> str:
        return pulumi.get(self, "symbol")


