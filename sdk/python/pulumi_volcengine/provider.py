# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 access_key: pulumi.Input[str],
                 region: pulumi.Input[str],
                 secret_key: pulumi.Input[str],
                 customer_endpoints: Optional[pulumi.Input[str]] = None,
                 customer_headers: Optional[pulumi.Input[str]] = None,
                 disable_ssl: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 proxy_url: Optional[pulumi.Input[str]] = None,
                 session_token: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] access_key: The Access Key for Volcengine Provider
        :param pulumi.Input[str] region: The Region for Volcengine Provider
        :param pulumi.Input[str] secret_key: The Secret Key for Volcengine Provider
        :param pulumi.Input[str] customer_endpoints: CUSTOMER ENDPOINTS for Volcengine Provider
        :param pulumi.Input[str] customer_headers: CUSTOMER HEADERS for Volcengine Provider
        :param pulumi.Input[bool] disable_ssl: Disable SSL for Volcengine Provider
        :param pulumi.Input[str] endpoint: The Customer Endpoint for Volcengine Provider
        :param pulumi.Input[str] proxy_url: PROXY URL for Volcengine Provider
        :param pulumi.Input[str] session_token: The Session Token for Volcengine Provider
        """
        pulumi.set(__self__, "access_key", access_key)
        pulumi.set(__self__, "region", region)
        pulumi.set(__self__, "secret_key", secret_key)
        if customer_endpoints is not None:
            pulumi.set(__self__, "customer_endpoints", customer_endpoints)
        if customer_headers is not None:
            pulumi.set(__self__, "customer_headers", customer_headers)
        if disable_ssl is not None:
            pulumi.set(__self__, "disable_ssl", disable_ssl)
        if endpoint is not None:
            pulumi.set(__self__, "endpoint", endpoint)
        if proxy_url is not None:
            pulumi.set(__self__, "proxy_url", proxy_url)
        if session_token is not None:
            pulumi.set(__self__, "session_token", session_token)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Input[str]:
        """
        The Access Key for Volcengine Provider
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        """
        The Region for Volcengine Provider
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Input[str]:
        """
        The Secret Key for Volcengine Provider
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="customerEndpoints")
    def customer_endpoints(self) -> Optional[pulumi.Input[str]]:
        """
        CUSTOMER ENDPOINTS for Volcengine Provider
        """
        return pulumi.get(self, "customer_endpoints")

    @customer_endpoints.setter
    def customer_endpoints(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_endpoints", value)

    @property
    @pulumi.getter(name="customerHeaders")
    def customer_headers(self) -> Optional[pulumi.Input[str]]:
        """
        CUSTOMER HEADERS for Volcengine Provider
        """
        return pulumi.get(self, "customer_headers")

    @customer_headers.setter
    def customer_headers(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "customer_headers", value)

    @property
    @pulumi.getter(name="disableSsl")
    def disable_ssl(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable SSL for Volcengine Provider
        """
        return pulumi.get(self, "disable_ssl")

    @disable_ssl.setter
    def disable_ssl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_ssl", value)

    @property
    @pulumi.getter
    def endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        The Customer Endpoint for Volcengine Provider
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="proxyUrl")
    def proxy_url(self) -> Optional[pulumi.Input[str]]:
        """
        PROXY URL for Volcengine Provider
        """
        return pulumi.get(self, "proxy_url")

    @proxy_url.setter
    def proxy_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "proxy_url", value)

    @property
    @pulumi.getter(name="sessionToken")
    def session_token(self) -> Optional[pulumi.Input[str]]:
        """
        The Session Token for Volcengine Provider
        """
        return pulumi.get(self, "session_token")

    @session_token.setter
    def session_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "session_token", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 customer_endpoints: Optional[pulumi.Input[str]] = None,
                 customer_headers: Optional[pulumi.Input[str]] = None,
                 disable_ssl: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 proxy_url: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 session_token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the volcengine package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The Access Key for Volcengine Provider
        :param pulumi.Input[str] customer_endpoints: CUSTOMER ENDPOINTS for Volcengine Provider
        :param pulumi.Input[str] customer_headers: CUSTOMER HEADERS for Volcengine Provider
        :param pulumi.Input[bool] disable_ssl: Disable SSL for Volcengine Provider
        :param pulumi.Input[str] endpoint: The Customer Endpoint for Volcengine Provider
        :param pulumi.Input[str] proxy_url: PROXY URL for Volcengine Provider
        :param pulumi.Input[str] region: The Region for Volcengine Provider
        :param pulumi.Input[str] secret_key: The Secret Key for Volcengine Provider
        :param pulumi.Input[str] session_token: The Session Token for Volcengine Provider
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the volcengine package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 customer_endpoints: Optional[pulumi.Input[str]] = None,
                 customer_headers: Optional[pulumi.Input[str]] = None,
                 disable_ssl: Optional[pulumi.Input[bool]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 proxy_url: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 session_token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if access_key is None and not opts.urn:
                raise TypeError("Missing required property 'access_key'")
            __props__.__dict__["access_key"] = access_key
            __props__.__dict__["customer_endpoints"] = customer_endpoints
            __props__.__dict__["customer_headers"] = customer_headers
            __props__.__dict__["disable_ssl"] = pulumi.Output.from_input(disable_ssl).apply(pulumi.runtime.to_json) if disable_ssl is not None else None
            __props__.__dict__["endpoint"] = endpoint
            __props__.__dict__["proxy_url"] = proxy_url
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            if secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_key'")
            __props__.__dict__["secret_key"] = secret_key
            __props__.__dict__["session_token"] = session_token
        super(Provider, __self__).__init__(
            'volcengine',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[str]:
        """
        The Access Key for Volcengine Provider
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="customerEndpoints")
    def customer_endpoints(self) -> pulumi.Output[Optional[str]]:
        """
        CUSTOMER ENDPOINTS for Volcengine Provider
        """
        return pulumi.get(self, "customer_endpoints")

    @property
    @pulumi.getter(name="customerHeaders")
    def customer_headers(self) -> pulumi.Output[Optional[str]]:
        """
        CUSTOMER HEADERS for Volcengine Provider
        """
        return pulumi.get(self, "customer_headers")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        The Customer Endpoint for Volcengine Provider
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter(name="proxyUrl")
    def proxy_url(self) -> pulumi.Output[Optional[str]]:
        """
        PROXY URL for Volcengine Provider
        """
        return pulumi.get(self, "proxy_url")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The Region for Volcengine Provider
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        The Secret Key for Volcengine Provider
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="sessionToken")
    def session_token(self) -> pulumi.Output[Optional[str]]:
        """
        The Session Token for Volcengine Provider
        """
        return pulumi.get(self, "session_token")

