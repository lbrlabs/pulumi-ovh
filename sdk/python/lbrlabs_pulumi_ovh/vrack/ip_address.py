# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['IpAddressArgs', 'IpAddress']

@pulumi.input_type
class IpAddressArgs:
    def __init__(__self__, *,
                 block: pulumi.Input[str],
                 service_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a IpAddress resource.
        :param pulumi.Input[str] block: Your IP block.
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        pulumi.set(__self__, "block", block)
        pulumi.set(__self__, "service_name", service_name)

    @property
    @pulumi.getter
    def block(self) -> pulumi.Input[str]:
        """
        Your IP block.
        """
        return pulumi.get(self, "block")

    @block.setter
    def block(self, value: pulumi.Input[str]):
        pulumi.set(self, "block", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Input[str]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_name", value)


@pulumi.input_type
class _IpAddressState:
    def __init__(__self__, *,
                 block: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 zone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpAddress resources.
        :param pulumi.Input[str] block: Your IP block.
        :param pulumi.Input[str] gateway: Your gateway
        :param pulumi.Input[str] ip: Your IP block
        :param pulumi.Input[str] service_name: The internal name of your vrack
        :param pulumi.Input[str] zone: Where you want your block announced on the network
        """
        if block is not None:
            pulumi.set(__self__, "block", block)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if service_name is not None:
            pulumi.set(__self__, "service_name", service_name)
        if zone is not None:
            pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter
    def block(self) -> Optional[pulumi.Input[str]]:
        """
        Your IP block.
        """
        return pulumi.get(self, "block")

    @block.setter
    def block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "block", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        """
        Your gateway
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        Your IP block
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> Optional[pulumi.Input[str]]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

    @service_name.setter
    def service_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_name", value)

    @property
    @pulumi.getter
    def zone(self) -> Optional[pulumi.Input[str]]:
        """
        Where you want your block announced on the network
        """
        return pulumi.get(self, "zone")

    @zone.setter
    def zone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone", value)


class IpAddress(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Attach an IP block to a VRack.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh
        import pulumi_ovh as ovh

        mycart = ovh.Order.get_cart(ovh_subsidiary="fr",
            description="my cart")
        vrack_cart_product_plan = ovh.Order.get_cart_product_plan(cart_id=mycart.id,
            price_capacity="renew",
            product="vrack",
            plan_code="vrack")
        vrack_vrack = ovh.vrack.Vrack("vrackVrack",
            description=mycart.description,
            ovh_subsidiary=mycart.ovh_subsidiary,
            plan=ovh.vrack.VrackPlanArgs(
                duration=vrack_cart_product_plan.selected_prices[0].duration,
                plan_code=vrack_cart_product_plan.plan_code,
                pricing_mode=vrack_cart_product_plan.selected_prices[0].pricing_mode,
            ))
        ipblock_cart_product_plan = ovh.Order.get_cart_product_plan(cart_id=mycart.id,
            price_capacity="renew",
            product="ip",
            plan_code="ip-v4-s30-ripe")
        ipblock_ip_service = ovh.ip.IpService("ipblockIpService",
            ovh_subsidiary=mycart.ovh_subsidiary,
            description=mycart.description,
            plan=ovh.ip.IpServicePlanArgs(
                duration=ipblock_cart_product_plan.selected_prices[0].duration,
                plan_code=ipblock_cart_product_plan.plan_code,
                pricing_mode=ipblock_cart_product_plan.selected_prices[0].pricing_mode,
                configurations=[ovh.ip.IpServicePlanConfigurationArgs(
                    label="country",
                    value="FR",
                )],
            ))
        vrackblock = ovh.vrack.IpAddress("vrackblock",
            service_name=vrack_vrack.service_name,
            block=ipblock_ip_service.ip)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] block: Your IP block.
        :param pulumi.Input[str] service_name: The internal name of your vrack
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpAddressArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Attach an IP block to a VRack.

        ## Example Usage

        ```python
        import pulumi
        import lbrlabs_pulumi_ovh as ovh
        import pulumi_ovh as ovh

        mycart = ovh.Order.get_cart(ovh_subsidiary="fr",
            description="my cart")
        vrack_cart_product_plan = ovh.Order.get_cart_product_plan(cart_id=mycart.id,
            price_capacity="renew",
            product="vrack",
            plan_code="vrack")
        vrack_vrack = ovh.vrack.Vrack("vrackVrack",
            description=mycart.description,
            ovh_subsidiary=mycart.ovh_subsidiary,
            plan=ovh.vrack.VrackPlanArgs(
                duration=vrack_cart_product_plan.selected_prices[0].duration,
                plan_code=vrack_cart_product_plan.plan_code,
                pricing_mode=vrack_cart_product_plan.selected_prices[0].pricing_mode,
            ))
        ipblock_cart_product_plan = ovh.Order.get_cart_product_plan(cart_id=mycart.id,
            price_capacity="renew",
            product="ip",
            plan_code="ip-v4-s30-ripe")
        ipblock_ip_service = ovh.ip.IpService("ipblockIpService",
            ovh_subsidiary=mycart.ovh_subsidiary,
            description=mycart.description,
            plan=ovh.ip.IpServicePlanArgs(
                duration=ipblock_cart_product_plan.selected_prices[0].duration,
                plan_code=ipblock_cart_product_plan.plan_code,
                pricing_mode=ipblock_cart_product_plan.selected_prices[0].pricing_mode,
                configurations=[ovh.ip.IpServicePlanConfigurationArgs(
                    label="country",
                    value="FR",
                )],
            ))
        vrackblock = ovh.vrack.IpAddress("vrackblock",
            service_name=vrack_vrack.service_name,
            block=ipblock_ip_service.ip)
        ```

        :param str resource_name: The name of the resource.
        :param IpAddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpAddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block: Optional[pulumi.Input[str]] = None,
                 service_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpAddressArgs.__new__(IpAddressArgs)

            if block is None and not opts.urn:
                raise TypeError("Missing required property 'block'")
            __props__.__dict__["block"] = block
            if service_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_name'")
            __props__.__dict__["service_name"] = service_name
            __props__.__dict__["gateway"] = None
            __props__.__dict__["ip"] = None
            __props__.__dict__["zone"] = None
        super(IpAddress, __self__).__init__(
            'ovh:Vrack/ipAddress:IpAddress',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            block: Optional[pulumi.Input[str]] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            service_name: Optional[pulumi.Input[str]] = None,
            zone: Optional[pulumi.Input[str]] = None) -> 'IpAddress':
        """
        Get an existing IpAddress resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] block: Your IP block.
        :param pulumi.Input[str] gateway: Your gateway
        :param pulumi.Input[str] ip: Your IP block
        :param pulumi.Input[str] service_name: The internal name of your vrack
        :param pulumi.Input[str] zone: Where you want your block announced on the network
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpAddressState.__new__(_IpAddressState)

        __props__.__dict__["block"] = block
        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["ip"] = ip
        __props__.__dict__["service_name"] = service_name
        __props__.__dict__["zone"] = zone
        return IpAddress(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def block(self) -> pulumi.Output[str]:
        """
        Your IP block.
        """
        return pulumi.get(self, "block")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[str]:
        """
        Your gateway
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        Your IP block
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="serviceName")
    def service_name(self) -> pulumi.Output[str]:
        """
        The internal name of your vrack
        """
        return pulumi.get(self, "service_name")

    @property
    @pulumi.getter
    def zone(self) -> pulumi.Output[str]:
        """
        Where you want your block announced on the network
        """
        return pulumi.get(self, "zone")

