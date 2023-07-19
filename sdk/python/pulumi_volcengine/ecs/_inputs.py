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
    'InstanceDataVolumeArgs',
    'InstanceGpuDeviceArgs',
    'InstanceSecondaryNetworkInterfaceArgs',
    'InstanceTagArgs',
]

@pulumi.input_type
class InstanceDataVolumeArgs:
    def __init__(__self__, *,
                 size: pulumi.Input[int],
                 volume_type: pulumi.Input[str],
                 delete_with_instance: Optional[pulumi.Input[bool]] = None):
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "volume_type", volume_type)
        if delete_with_instance is not None:
            pulumi.set(__self__, "delete_with_instance", delete_with_instance)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[int]:
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[int]):
        pulumi.set(self, "size", value)

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "volume_type")

    @volume_type.setter
    def volume_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "volume_type", value)

    @property
    @pulumi.getter(name="deleteWithInstance")
    def delete_with_instance(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "delete_with_instance")

    @delete_with_instance.setter
    def delete_with_instance(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_with_instance", value)


@pulumi.input_type
class InstanceGpuDeviceArgs:
    def __init__(__self__, *,
                 count: Optional[pulumi.Input[int]] = None,
                 encrypted_memory_size: Optional[pulumi.Input[int]] = None,
                 memory_size: Optional[pulumi.Input[int]] = None,
                 product_name: Optional[pulumi.Input[str]] = None):
        if count is not None:
            pulumi.set(__self__, "count", count)
        if encrypted_memory_size is not None:
            pulumi.set(__self__, "encrypted_memory_size", encrypted_memory_size)
        if memory_size is not None:
            pulumi.set(__self__, "memory_size", memory_size)
        if product_name is not None:
            pulumi.set(__self__, "product_name", product_name)

    @property
    @pulumi.getter
    def count(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "count")

    @count.setter
    def count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "count", value)

    @property
    @pulumi.getter(name="encryptedMemorySize")
    def encrypted_memory_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "encrypted_memory_size")

    @encrypted_memory_size.setter
    def encrypted_memory_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "encrypted_memory_size", value)

    @property
    @pulumi.getter(name="memorySize")
    def memory_size(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "memory_size")

    @memory_size.setter
    def memory_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "memory_size", value)

    @property
    @pulumi.getter(name="productName")
    def product_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "product_name")

    @product_name.setter
    def product_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "product_name", value)


@pulumi.input_type
class InstanceSecondaryNetworkInterfaceArgs:
    def __init__(__self__, *,
                 security_group_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subnet_id: pulumi.Input[str],
                 primary_ip_address: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if primary_ip_address is not None:
            pulumi.set(__self__, "primary_ip_address", primary_ip_address)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="primaryIpAddress")
    def primary_ip_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "primary_ip_address")

    @primary_ip_address.setter
    def primary_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_ip_address", value)


@pulumi.input_type
class InstanceTagArgs:
    def __init__(__self__, *,
                 key: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def key(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: pulumi.Input[str]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


