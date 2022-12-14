# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CloudProjectKubeOidcArgs', 'CloudProjectKubeOidc']

@pulumi.input_type
class CloudProjectKubeOidcArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[str],
                 issuer_url: pulumi.Input[str],
                 kube_id: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a CloudProjectKubeOidc resource.
        :param pulumi.Input[str] client_id: The OIDC client ID.
        :param pulumi.Input[str] issuer_url: The OIDC issuer url.
        :param pulumi.Input[str] kube_id: The ID of the managed kubernetes cluster.
        :param pulumi.Input[str] service_name: The ID of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "issuer_url", issuer_url)
        pulumi.set(__self__, "kube_id", kube_id)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        """
        The OIDC client ID.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="issuerUrl")
    def issuer_url(self) -> pulumi.Input[str]:
        """
        The OIDC issuer url.
        """
        return pulumi.get(self, "issuer_url")

    @issuer_url.setter
    def issuer_url(self, value: pulumi.Input[str]):
        pulumi.set(self, "issuer_url", value)

    @property
    @pulumi.getter(name="kubeId")
    def kube_id(self) -> pulumi.Input[str]:
        """
        The ID of the managed kubernetes cluster.
        """
        return pulumi.get(self, "kube_id")

    @kube_id.setter
    def kube_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "kube_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The ID of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _CloudProjectKubeOidcState:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 issuer_url: Optional[pulumi.Input[str]] = None,
                 kube_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CloudProjectKubeOidc resources.
        :param pulumi.Input[str] client_id: The OIDC client ID.
        :param pulumi.Input[str] issuer_url: The OIDC issuer url.
        :param pulumi.Input[str] kube_id: The ID of the managed kubernetes cluster.
        :param pulumi.Input[str] service_name: The ID of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if issuer_url is not None:
            pulumi.set(__self__, "issuer_url", issuer_url)
        if kube_id is not None:
            pulumi.set(__self__, "kube_id", kube_id)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The OIDC client ID.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="issuerUrl")
    def issuer_url(self) -> Optional[pulumi.Input[str]]:
        """
        The OIDC issuer url.
        """
        return pulumi.get(self, "issuer_url")

    @issuer_url.setter
    def issuer_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer_url", value)

    @property
    @pulumi.getter(name="kubeId")
    def kube_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the managed kubernetes cluster.
        """
        return pulumi.get(self, "kube_id")

    @kube_id.setter
    def kube_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kube_id", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)


class CloudProjectKubeOidc(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 issuer_url: Optional[pulumi.Input[str]] = None,
                 kube_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates an OIDC configuration in an OVHcloud Managed Kubernetes cluster.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_ovh as ovh

        my_oidc = ovh.CloudProjectKubeOidc("my-oidc",
            service_name=var["projectid"],
            kube_id=ovh_cloud_project_kube["mykube"]["id"],
            client_id="xxx",
            issuer_url="https://ovh.com")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The OIDC client ID.
        :param pulumi.Input[str] issuer_url: The OIDC issuer url.
        :param pulumi.Input[str] kube_id: The ID of the managed kubernetes cluster.
        :param pulumi.Input[str] service_name: The ID of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudProjectKubeOidcArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an OIDC configuration in an OVHcloud Managed Kubernetes cluster.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_ovh as ovh

        my_oidc = ovh.CloudProjectKubeOidc("my-oidc",
            service_name=var["projectid"],
            kube_id=ovh_cloud_project_kube["mykube"]["id"],
            client_id="xxx",
            issuer_url="https://ovh.com")
        ```

        :param str resource_name: The name of the resource.
        :param CloudProjectKubeOidcArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudProjectKubeOidcArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 issuer_url: Optional[pulumi.Input[str]] = None,
                 kube_id: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudProjectKubeOidcArgs.__new__(CloudProjectKubeOidcArgs)

            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            if issuer_url is None and not opts.urn:
                raise TypeError("Missing required property 'issuer_url'")
            __props__.__dict__["issuer_url"] = issuer_url
            if kube_id is None and not opts.urn:
                raise TypeError("Missing required property 'kube_id'")
            __props__.__dict__["kube_id"] = kube_id
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
        super(CloudProjectKubeOidc, __self__).__init__(
            'ovh:index/cloudProjectKubeOidc:CloudProjectKubeOidc',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            issuer_url: Optional[pulumi.Input[str]] = None,
            kube_id: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'CloudProjectKubeOidc':
        """
        Get an existing CloudProjectKubeOidc resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The OIDC client ID.
        :param pulumi.Input[str] issuer_url: The OIDC issuer url.
        :param pulumi.Input[str] kube_id: The ID of the managed kubernetes cluster.
        :param pulumi.Input[str] service_name: The ID of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudProjectKubeOidcState.__new__(_CloudProjectKubeOidcState)

        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["issuer_url"] = issuer_url
        __props__.__dict__["kube_id"] = kube_id
        __props__.__dict__["service_name"] = service_name
        return CloudProjectKubeOidc(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        The OIDC client ID.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="issuerUrl")
    def issuer_url(self) -> pulumi.Output[str]:
        """
        The OIDC issuer url.
        """
        return pulumi.get(self, "issuer_url")

    @property
    @pulumi.getter(name="kubeId")
    def kube_id(self) -> pulumi.Output[str]:
        """
        The ID of the managed kubernetes cluster.
        """
        return pulumi.get(self, "kube_id")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The ID of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

