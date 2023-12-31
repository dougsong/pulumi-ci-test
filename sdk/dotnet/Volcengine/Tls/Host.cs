// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace ZhenRan.PulumiPackage.Volcengine.Tls
{
    [VolcengineResourceType("volcengine:tls/host:Host")]
    public partial class Host : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of host group.
        /// </summary>
        [Output("hostGroupId")]
        public Output<string> HostGroupId { get; private set; } = null!;

        /// <summary>
        /// The ip address.
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;


        /// <summary>
        /// Create a Host resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Host(string name, HostArgs args, CustomResourceOptions? options = null)
            : base("volcengine:tls/host:Host", name, args ?? new HostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Host(string name, Input<string> id, HostState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:tls/host:Host", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Host resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Host Get(string name, Input<string> id, HostState? state = null, CustomResourceOptions? options = null)
        {
            return new Host(name, id, state, options);
        }
    }

    public sealed class HostArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of host group.
        /// </summary>
        [Input("hostGroupId", required: true)]
        public Input<string> HostGroupId { get; set; } = null!;

        /// <summary>
        /// The ip address.
        /// </summary>
        [Input("ip", required: true)]
        public Input<string> Ip { get; set; } = null!;

        public HostArgs()
        {
        }
        public static new HostArgs Empty => new HostArgs();
    }

    public sealed class HostState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of host group.
        /// </summary>
        [Input("hostGroupId")]
        public Input<string>? HostGroupId { get; set; }

        /// <summary>
        /// The ip address.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public HostState()
        {
        }
        public static new HostState Empty => new HostState();
    }
}
