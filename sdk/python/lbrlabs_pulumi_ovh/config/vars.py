# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('ovh')


class _ExportableConfig(types.ModuleType):
    @property
    def application_key(self) -> Optional[str]:
        """
        The OVH API Application Key.
        """
        return __config__.get('applicationKey') or _utilities.get_env('OVH_APPLICATION_KEY')

    @property
    def application_secret(self) -> Optional[str]:
        """
        The OVH API Application Secret.
        """
        return __config__.get('applicationSecret') or _utilities.get_env('OVH_APPLICATION_SECRET')

    @property
    def consumer_key(self) -> Optional[str]:
        """
        The OVH API Consumer key.
        """
        return __config__.get('consumerKey') or _utilities.get_env('OVH_CONSUMER_KEY')

    @property
    def endpoint(self) -> Optional[str]:
        """
        The OVH API endpoint to target (ex: "ovh-eu").
        """
        return __config__.get('endpoint') or _utilities.get_env('OVH_ENDPOINT')
