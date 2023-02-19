# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['DatabaseInstanceArgs', 'DatabaseInstance']

@pulumi.input_type
class DatabaseInstanceArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 engine: pulumi.Input[str],
                 service_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabaseInstance resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
               Available engines:
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        :param pulumi.Input[str] name: Name of the database.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "engine", engine)
        pulumi.set(__self__, "service_name", service_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Input[str]:
        """
        The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        Available engines:
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: pulumi.Input[str]):
        pulumi.set(self, "engine", value)

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
        Name of the database.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _DatabaseInstanceState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 default: Optional[pulumi.Input[bool]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering DatabaseInstance resources.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[bool] default: Defines if the database has been created by default.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
               Available engines:
        :param pulumi.Input[str] name: Name of the database.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if engine is not None:
            pulumi.set(__self__, "engine", engine)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines if the database has been created by default.
        """
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default", value)

    @property
    @pulumi.getter
    def engine(self) -> Optional[pulumi.Input[str]]:
        """
        The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        Available engines:
        """
        return pulumi.get(self, "engine")

    @engine.setter
    def engine(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "engine", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the database.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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


class DatabaseInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a database for a database cluster associated with a public cloud project.

        With this resource you can create a database for the following database engine:

          * `mysql`
          * `postgresql`

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh
        import pulumi_ovh as ovh

        db = ovh.CloudProjectDatabase.get_database(service_name="XXXX",
            engine="YYYY",
            id="ZZZZ")
        database = ovh.cloud_project_database.DatabaseInstance("database",
            service_name=db.service_name,
            engine=db.engine,
            cluster_id=db.id)
        ```

        ## Import

        OVHcloud Managed database clusters databases can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the database, separated by "/" E.g., bash

        ```sh
         $ pulumi import ovh:CloudProjectDatabase/databaseInstance:DatabaseInstance my_database service_name/engine/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
               Available engines:
        :param pulumi.Input[str] name: Name of the database.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a database for a database cluster associated with a public cloud project.

        With this resource you can create a database for the following database engine:

          * `mysql`
          * `postgresql`

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh
        import pulumi_ovh as ovh

        db = ovh.CloudProjectDatabase.get_database(service_name="XXXX",
            engine="YYYY",
            id="ZZZZ")
        database = ovh.cloud_project_database.DatabaseInstance("database",
            service_name=db.service_name,
            engine=db.engine,
            cluster_id=db.id)
        ```

        ## Import

        OVHcloud Managed database clusters databases can be imported using the `service_name`, `engine`, `cluster_id` and `id` of the database, separated by "/" E.g., bash

        ```sh
         $ pulumi import ovh:CloudProjectDatabase/databaseInstance:DatabaseInstance my_database service_name/engine/cluster_id/id
        ```

        :param str resource_name: The name of the resource.
        :param DatabaseInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 engine: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseInstanceArgs.__new__(DatabaseInstanceArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if engine is None and not opts.urn:
                raise TypeError("Missing required property 'engine'")
            __props__.__dict__["engine"] = engine
            __props__.__dict__["name"] = name
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["default"] = None
        super(DatabaseInstance, __self__).__init__(
            'ovh:CloudProjectDatabase/databaseInstance:DatabaseInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            default: Optional[pulumi.Input[bool]] = None,
            engine: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None) -> 'DatabaseInstance':
        """
        Get an existing DatabaseInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: Cluster ID.
        :param pulumi.Input[bool] default: Defines if the database has been created by default.
        :param pulumi.Input[str] engine: The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
               Available engines:
        :param pulumi.Input[str] name: Name of the database.
        :param pulumi.Input[str] service_name: The id of the public cloud project. If omitted,
               the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DatabaseInstanceState.__new__(_DatabaseInstanceState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["default"] = default
        __props__.__dict__["engine"] = engine
        __props__.__dict__["name"] = name
        __props__.__dict__["service_name"] = service_name
        return DatabaseInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        Cluster ID.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def default(self) -> pulumi.Output[bool]:
        """
        Defines if the database has been created by default.
        """
        return pulumi.get(self, "default")

    @property
    @pulumi.getter
    def engine(self) -> pulumi.Output[str]:
        """
        The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
        Available engines:
        """
        return pulumi.get(self, "engine")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the database.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The id of the public cloud project. If omitted,
        the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        """
        return pulumi.get(self, "service_name")

