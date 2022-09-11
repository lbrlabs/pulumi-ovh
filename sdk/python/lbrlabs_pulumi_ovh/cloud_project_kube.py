# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['CloudProjectKubeArgs', 'CloudProjectKube']

@pulumi.input_type
class CloudProjectKubeArgs:
    def __init__(__self__, *,
                 region: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 private_network_configuration: Optional[pulumi.Input['CloudProjectKubePrivateNetworkConfigurationArgs']] = None,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 update_policy: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CloudProjectKube resource.
        :param pulumi.Input[str] region: a valid OVH public cloud region ID in which the kubernetes
               cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
               Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] name: The name of the kubernetes cluster.
        :param pulumi.Input['CloudProjectKubePrivateNetworkConfigurationArgs'] private_network_configuration: The private network configuration
               * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
               * private_network_routing_as_default - Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
        :param pulumi.Input[str] private_network_id: OpenStack private network ID to use.
               Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
        :param pulumi.Input[str] update_policy: Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]'.
        :param pulumi.Input[str] version: kubernetes version to use.
               Changing this value updates the resource. Defaults to latest available.
        """
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "service_name", service_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if private_network_configuration is not None:
            pulumi.set(__self__, "private_network_configuration", private_network_configuration)
        if private_network_id is not None:
            pulumi.set(__self__, "private_network_id", private_network_id)
        if update_policy is not None:
            pulumi.set(__self__, "update_policy", update_policy)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        a valid OVH public cloud region ID in which the kubernetes
        cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        Changing this value recreates the resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the kubernetes cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="privateNetworkConfiguration")
    def private_network_configuration(self) -> Optional[pulumi.Input['CloudProjectKubePrivateNetworkConfigurationArgs']]:
        """
        The private network configuration
        * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
        * private_network_routing_as_default - Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
        """
        return pulumi.get(self, "private_network_configuration")

    @private_network_configuration.setter
    def private_network_configuration(self, value: Optional[pulumi.Input['CloudProjectKubePrivateNetworkConfigurationArgs']]):
        pulumi.set(self, "private_network_configuration", value)

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        OpenStack private network ID to use.
        Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
        """
        return pulumi.get(self, "private_network_id")

    @private_network_id.setter
    def private_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_network_id", value)

    @property
    @pulumi.getter(name="updatePolicy")
    def update_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]'.
        """
        return pulumi.get(self, "update_policy")

    @update_policy.setter
    def update_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_policy", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        kubernetes version to use.
        Changing this value updates the resource. Defaults to latest available.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _CloudProjectKubeState:
    def __init__(__self__, *,
                 control_plane_is_up_to_date: Optional[pulumi.Input[bool]] = None,
                 is_up_to_date: Optional[pulumi.Input[bool]] = None,
                 kubeconfig: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 next_upgrade_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 nodes_url: Optional[pulumi.Input[str]] = None,
                 private_network_configuration: Optional[pulumi.Input['CloudProjectKubePrivateNetworkConfigurationArgs']] = None,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 update_policy: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CloudProjectKube resources.
        :param pulumi.Input[bool] control_plane_is_up_to_date: True if control-plane is up to date.
        :param pulumi.Input[bool] is_up_to_date: True if all nodes and control-plane are up to date.
        :param pulumi.Input[str] kubeconfig: The kubeconfig file. Use this file to connect to your kubernetes cluster.
        :param pulumi.Input[str] name: The name of the kubernetes cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] next_upgrade_versions: Kubernetes versions available for upgrade.
        :param pulumi.Input[str] nodes_url: Cluster nodes URL.
        :param pulumi.Input['CloudProjectKubePrivateNetworkConfigurationArgs'] private_network_configuration: The private network configuration
               * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
               * private_network_routing_as_default - Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
        :param pulumi.Input[str] private_network_id: OpenStack private network ID to use.
               Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
        :param pulumi.Input[str] region: a valid OVH public cloud region ID in which the kubernetes
               cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
               Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] status: Cluster status. Should be normally set to 'READY'.
        :param pulumi.Input[str] update_policy: Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]'.
        :param pulumi.Input[str] url: Management URL of your cluster.
        :param pulumi.Input[str] version: kubernetes version to use.
               Changing this value updates the resource. Defaults to latest available.
        """
        if control_plane_is_up_to_date is not None:
            pulumi.set(__self__, "control_plane_is_up_to_date", control_plane_is_up_to_date)
        if is_up_to_date is not None:
            pulumi.set(__self__, "is_up_to_date", is_up_to_date)
        if kubeconfig is not None:
            pulumi.set(__self__, "kubeconfig", kubeconfig)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if next_upgrade_versions is not None:
            pulumi.set(__self__, "next_upgrade_versions", next_upgrade_versions)
        if nodes_url is not None:
            pulumi.set(__self__, "nodes_url", nodes_url)
        if private_network_configuration is not None:
            pulumi.set(__self__, "private_network_configuration", private_network_configuration)
        if private_network_id is not None:
            pulumi.set(__self__, "private_network_id", private_network_id)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if update_policy is not None:
            pulumi.set(__self__, "update_policy", update_policy)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="controlPlaneIsUpToDate")
    def control_plane_is_up_to_date(self) -> Optional[pulumi.Input[bool]]:
        """
        True if control-plane is up to date.
        """
        return pulumi.get(self, "control_plane_is_up_to_date")

    @control_plane_is_up_to_date.setter
    def control_plane_is_up_to_date(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "control_plane_is_up_to_date", value)

    @property
    @pulumi.getter(name="isUpToDate")
    def is_up_to_date(self) -> Optional[pulumi.Input[bool]]:
        """
        True if all nodes and control-plane are up to date.
        """
        return pulumi.get(self, "is_up_to_date")

    @is_up_to_date.setter
    def is_up_to_date(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_up_to_date", value)

    @property
    @pulumi.getter
    def kubeconfig(self) -> Optional[pulumi.Input[str]]:
        """
        The kubeconfig file. Use this file to connect to your kubernetes cluster.
        """
        return pulumi.get(self, "kubeconfig")

    @kubeconfig.setter
    def kubeconfig(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubeconfig", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the kubernetes cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nextUpgradeVersions")
    def next_upgrade_versions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Kubernetes versions available for upgrade.
        """
        return pulumi.get(self, "next_upgrade_versions")

    @next_upgrade_versions.setter
    def next_upgrade_versions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "next_upgrade_versions", value)

    @property
    @pulumi.getter(name="nodesUrl")
    def nodes_url(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster nodes URL.
        """
        return pulumi.get(self, "nodes_url")

    @nodes_url.setter
    def nodes_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nodes_url", value)

    @property
    @pulumi.getter(name="privateNetworkConfiguration")
    def private_network_configuration(self) -> Optional[pulumi.Input['CloudProjectKubePrivateNetworkConfigurationArgs']]:
        """
        The private network configuration
        * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
        * private_network_routing_as_default - Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
        """
        return pulumi.get(self, "private_network_configuration")

    @private_network_configuration.setter
    def private_network_configuration(self, value: Optional[pulumi.Input['CloudProjectKubePrivateNetworkConfigurationArgs']]):
        pulumi.set(self, "private_network_configuration", value)

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        OpenStack private network ID to use.
        Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
        """
        return pulumi.get(self, "private_network_id")

    @private_network_id.setter
    def private_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_network_id", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        a valid OVH public cloud region ID in which the kubernetes
        cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        Changing this value recreates the resource.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster status. Should be normally set to 'READY'.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updatePolicy")
    def update_policy(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]'.
        """
        return pulumi.get(self, "update_policy")

    @update_policy.setter
    def update_policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_policy", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        Management URL of your cluster.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        kubernetes version to use.
        Changing this value updates the resource. Defaults to latest available.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class CloudProjectKube(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_network_configuration: Optional[pulumi.Input[pulumi.InputType['CloudProjectKubePrivateNetworkConfigurationArgs']]] = None,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 update_policy: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a OVH Managed Kubernetes Service cluster in a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh

        mykube = ovh.CloudProjectKube("mykube",
            service_name="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            region="GRA7",
            private_network_id=xxxxxxxx_xxxx_xxxx_xxxxx_xxxxxxxxxxxx,
            private_network_configuration=ovh.CloudProjectKubePrivateNetworkConfigurationArgs(
                default_vrack_gateway="10.4.0.1",
                private_network_routing_as_default=True,
            ),
            opts=pulumi.ResourceOptions(depends_on=[ovh_cloud_project_network_private["network1"]]))
        ```

        ## Import

        OVHcloud Managed Kubernetes Service clusters can be imported using the `serviceName` and the `id` of the cluster, separated by "/" E.g.,

        ```sh
         $ pulumi import ovh:index/cloudProjectKube:CloudProjectKube my_kube_cluster a6678gggjh76hggjh7f59/a123bc45-a1b2-34c5-678d-678ghg7676ebc
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the kubernetes cluster.
        :param pulumi.Input[pulumi.InputType['CloudProjectKubePrivateNetworkConfigurationArgs']] private_network_configuration: The private network configuration
               * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
               * private_network_routing_as_default - Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
        :param pulumi.Input[str] private_network_id: OpenStack private network ID to use.
               Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
        :param pulumi.Input[str] region: a valid OVH public cloud region ID in which the kubernetes
               cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
               Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] update_policy: Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]'.
        :param pulumi.Input[str] version: kubernetes version to use.
               Changing this value updates the resource. Defaults to latest available.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CloudProjectKubeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a OVH Managed Kubernetes Service cluster in a public cloud project.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh

        mykube = ovh.CloudProjectKube("mykube",
            service_name="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
            region="GRA7",
            private_network_id=xxxxxxxx_xxxx_xxxx_xxxxx_xxxxxxxxxxxx,
            private_network_configuration=ovh.CloudProjectKubePrivateNetworkConfigurationArgs(
                default_vrack_gateway="10.4.0.1",
                private_network_routing_as_default=True,
            ),
            opts=pulumi.ResourceOptions(depends_on=[ovh_cloud_project_network_private["network1"]]))
        ```

        ## Import

        OVHcloud Managed Kubernetes Service clusters can be imported using the `serviceName` and the `id` of the cluster, separated by "/" E.g.,

        ```sh
         $ pulumi import ovh:index/cloudProjectKube:CloudProjectKube my_kube_cluster a6678gggjh76hggjh7f59/a123bc45-a1b2-34c5-678d-678ghg7676ebc
        ```

        :param str resource_name: The name of the resource.
        :param CloudProjectKubeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CloudProjectKubeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 private_network_configuration: Optional[pulumi.Input[pulumi.InputType['CloudProjectKubePrivateNetworkConfigurationArgs']]] = None,
                 private_network_id: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 update_policy: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CloudProjectKubeArgs.__new__(CloudProjectKubeArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["private_network_configuration"] = private_network_configuration
            __props__.__dict__["private_network_id"] = private_network_id
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["update_policy"] = update_policy
            __props__.__dict__["version"] = version
            __props__.__dict__["control_plane_is_up_to_date"] = None
            __props__.__dict__["is_up_to_date"] = None
            __props__.__dict__["kubeconfig"] = None
            __props__.__dict__["next_upgrade_versions"] = None
            __props__.__dict__["nodes_url"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["url"] = None
        super(CloudProjectKube, __self__).__init__(
            'ovh:index/cloudProjectKube:CloudProjectKube',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            control_plane_is_up_to_date: Optional[pulumi.Input[bool]] = None,
            is_up_to_date: Optional[pulumi.Input[bool]] = None,
            kubeconfig: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            next_upgrade_versions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            nodes_url: Optional[pulumi.Input[str]] = None,
            private_network_configuration: Optional[pulumi.Input[pulumi.InputType['CloudProjectKubePrivateNetworkConfigurationArgs']]] = None,
            private_network_id: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            update_policy: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'CloudProjectKube':
        """
        Get an existing CloudProjectKube resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] control_plane_is_up_to_date: True if control-plane is up to date.
        :param pulumi.Input[bool] is_up_to_date: True if all nodes and control-plane are up to date.
        :param pulumi.Input[str] kubeconfig: The kubeconfig file. Use this file to connect to your kubernetes cluster.
        :param pulumi.Input[str] name: The name of the kubernetes cluster.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] next_upgrade_versions: Kubernetes versions available for upgrade.
        :param pulumi.Input[str] nodes_url: Cluster nodes URL.
        :param pulumi.Input[pulumi.InputType['CloudProjectKubePrivateNetworkConfigurationArgs']] private_network_configuration: The private network configuration
               * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
               * private_network_routing_as_default - Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
        :param pulumi.Input[str] private_network_id: OpenStack private network ID to use.
               Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
        :param pulumi.Input[str] region: a valid OVH public cloud region ID in which the kubernetes
               cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
               Changing this value recreates the resource.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] status: Cluster status. Should be normally set to 'READY'.
        :param pulumi.Input[str] update_policy: Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]'.
        :param pulumi.Input[str] url: Management URL of your cluster.
        :param pulumi.Input[str] version: kubernetes version to use.
               Changing this value updates the resource. Defaults to latest available.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CloudProjectKubeState.__new__(_CloudProjectKubeState)

        __props__.__dict__["control_plane_is_up_to_date"] = control_plane_is_up_to_date
        __props__.__dict__["is_up_to_date"] = is_up_to_date
        __props__.__dict__["kubeconfig"] = kubeconfig
        __props__.__dict__["name"] = name
        __props__.__dict__["next_upgrade_versions"] = next_upgrade_versions
        __props__.__dict__["nodes_url"] = nodes_url
        __props__.__dict__["private_network_configuration"] = private_network_configuration
        __props__.__dict__["private_network_id"] = private_network_id
        __props__.__dict__["region"] = region
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["status"] = status
        __props__.__dict__["update_policy"] = update_policy
        __props__.__dict__["url"] = url
        __props__.__dict__["version"] = version
        return CloudProjectKube(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="controlPlaneIsUpToDate")
    def control_plane_is_up_to_date(self) -> pulumi.Output[bool]:
        """
        True if control-plane is up to date.
        """
        return pulumi.get(self, "control_plane_is_up_to_date")

    @property
    @pulumi.getter(name="isUpToDate")
    def is_up_to_date(self) -> pulumi.Output[bool]:
        """
        True if all nodes and control-plane are up to date.
        """
        return pulumi.get(self, "is_up_to_date")

    @property
    @pulumi.getter
    def kubeconfig(self) -> pulumi.Output[str]:
        """
        The kubeconfig file. Use this file to connect to your kubernetes cluster.
        """
        return pulumi.get(self, "kubeconfig")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the kubernetes cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nextUpgradeVersions")
    def next_upgrade_versions(self) -> pulumi.Output[Sequence[str]]:
        """
        Kubernetes versions available for upgrade.
        """
        return pulumi.get(self, "next_upgrade_versions")

    @property
    @pulumi.getter(name="nodesUrl")
    def nodes_url(self) -> pulumi.Output[str]:
        """
        Cluster nodes URL.
        """
        return pulumi.get(self, "nodes_url")

    @property
    @pulumi.getter(name="privateNetworkConfiguration")
    def private_network_configuration(self) -> pulumi.Output[Optional['outputs.CloudProjectKubePrivateNetworkConfiguration']]:
        """
        The private network configuration
        * default_vrack_gateway - If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
        * private_network_routing_as_default - Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
        """
        return pulumi.get(self, "private_network_configuration")

    @property
    @pulumi.getter(name="privateNetworkId")
    def private_network_id(self) -> pulumi.Output[Optional[str]]:
        """
        OpenStack private network ID to use.
        Changing this value delete the resource(including ETCD user data). Defaults - not use private network.
        """
        return pulumi.get(self, "private_network_id")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        a valid OVH public cloud region ID in which the kubernetes
        cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.
        Changing this value recreates the resource.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Cluster status. Should be normally set to 'READY'.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updatePolicy")
    def update_policy(self) -> pulumi.Output[str]:
        """
        Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]'.
        """
        return pulumi.get(self, "update_policy")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        Management URL of your cluster.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        kubernetes version to use.
        Changing this value updates the resource. Defaults to latest available.
        """
        return pulumi.get(self, "version")
