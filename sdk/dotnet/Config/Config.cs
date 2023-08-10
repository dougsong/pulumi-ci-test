// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace ZhenRan.PulumiPackage.Volcengine
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly global::Pulumi.Config __config = new global::Pulumi.Config("volcengine");

        private static readonly __Value<string?> _accessKey = new __Value<string?>(() => __config.Get("accessKey"));
        /// <summary>
        /// The Access Key for Volcengine Provider
        /// </summary>
        public static string? AccessKey
        {
            get => _accessKey.Get();
            set => _accessKey.Set(value);
        }

        private static readonly __Value<string?> _customerEndpoints = new __Value<string?>(() => __config.Get("customerEndpoints"));
        /// <summary>
        /// CUSTOMER ENDPOINTS for Volcengine Provider
        /// </summary>
        public static string? CustomerEndpoints
        {
            get => _customerEndpoints.Get();
            set => _customerEndpoints.Set(value);
        }

        private static readonly __Value<string?> _customerHeaders = new __Value<string?>(() => __config.Get("customerHeaders"));
        /// <summary>
        /// CUSTOMER HEADERS for Volcengine Provider
        /// </summary>
        public static string? CustomerHeaders
        {
            get => _customerHeaders.Get();
            set => _customerHeaders.Set(value);
        }

        private static readonly __Value<bool?> _disableSsl = new __Value<bool?>(() => __config.GetBoolean("disableSsl"));
        /// <summary>
        /// Disable SSL for Volcengine Provider
        /// </summary>
        public static bool? DisableSsl
        {
            get => _disableSsl.Get();
            set => _disableSsl.Set(value);
        }

        private static readonly __Value<string?> _endpoint = new __Value<string?>(() => __config.Get("endpoint"));
        /// <summary>
        /// The Customer Endpoint for Volcengine Provider
        /// </summary>
        public static string? Endpoint
        {
            get => _endpoint.Get();
            set => _endpoint.Set(value);
        }

        private static readonly __Value<string?> _proxyUrl = new __Value<string?>(() => __config.Get("proxyUrl"));
        /// <summary>
        /// PROXY URL for Volcengine Provider
        /// </summary>
        public static string? ProxyUrl
        {
            get => _proxyUrl.Get();
            set => _proxyUrl.Set(value);
        }

        private static readonly __Value<string?> _region = new __Value<string?>(() => __config.Get("region"));
        /// <summary>
        /// The Region for Volcengine Provider
        /// </summary>
        public static string? Region
        {
            get => _region.Get();
            set => _region.Set(value);
        }

        private static readonly __Value<string?> _secretKey = new __Value<string?>(() => __config.Get("secretKey"));
        /// <summary>
        /// The Secret Key for Volcengine Provider
        /// </summary>
        public static string? SecretKey
        {
            get => _secretKey.Get();
            set => _secretKey.Set(value);
        }

        private static readonly __Value<string?> _sessionToken = new __Value<string?>(() => __config.Get("sessionToken"));
        /// <summary>
        /// The Session Token for Volcengine Provider
        /// </summary>
        public static string? SessionToken
        {
            get => _sessionToken.Get();
            set => _sessionToken.Set(value);
        }

    }
}
