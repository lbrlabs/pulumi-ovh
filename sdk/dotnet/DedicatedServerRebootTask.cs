// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Ovh
{
    [OvhResourceType("ovh:index/dedicatedServerRebootTask:DedicatedServerRebootTask")]
    public partial class DedicatedServerRebootTask : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Details of this task. (should be `Reboot asked`)
        /// </summary>
        [Output("comment")]
        public Output<string> Comment { get; private set; } = null!;

        /// <summary>
        /// Completion date in RFC3339 format.
        /// </summary>
        [Output("doneDate")]
        public Output<string> DoneDate { get; private set; } = null!;

        /// <summary>
        /// Function name (should be `hardReboot`).
        /// </summary>
        [Output("function")]
        public Output<string> Function { get; private set; } = null!;

        /// <summary>
        /// List of values traccked to trigger reboot, used also to form implicit dependencies
        /// </summary>
        [Output("keepers")]
        public Output<ImmutableArray<string>> Keepers { get; private set; } = null!;

        /// <summary>
        /// Last update in RFC3339 format.
        /// </summary>
        [Output("lastUpdate")]
        public Output<string> LastUpdate { get; private set; } = null!;

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Task creation date in RFC3339 format.
        /// </summary>
        [Output("startDate")]
        public Output<string> StartDate { get; private set; } = null!;

        /// <summary>
        /// Task status (should be `done`)
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a DedicatedServerRebootTask resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DedicatedServerRebootTask(string name, DedicatedServerRebootTaskArgs args, CustomResourceOptions? options = null)
            : base("ovh:index/dedicatedServerRebootTask:DedicatedServerRebootTask", name, args ?? new DedicatedServerRebootTaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DedicatedServerRebootTask(string name, Input<string> id, DedicatedServerRebootTaskState? state = null, CustomResourceOptions? options = null)
            : base("ovh:index/dedicatedServerRebootTask:DedicatedServerRebootTask", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DedicatedServerRebootTask resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DedicatedServerRebootTask Get(string name, Input<string> id, DedicatedServerRebootTaskState? state = null, CustomResourceOptions? options = null)
        {
            return new DedicatedServerRebootTask(name, id, state, options);
        }
    }

    public sealed class DedicatedServerRebootTaskArgs : global::Pulumi.ResourceArgs
    {
        [Input("keepers", required: true)]
        private InputList<string>? _keepers;

        /// <summary>
        /// List of values traccked to trigger reboot, used also to form implicit dependencies
        /// </summary>
        public InputList<string> Keepers
        {
            get => _keepers ?? (_keepers = new InputList<string>());
            set => _keepers = value;
        }

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public DedicatedServerRebootTaskArgs()
        {
        }
        public static new DedicatedServerRebootTaskArgs Empty => new DedicatedServerRebootTaskArgs();
    }

    public sealed class DedicatedServerRebootTaskState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details of this task. (should be `Reboot asked`)
        /// </summary>
        [Input("comment")]
        public Input<string>? Comment { get; set; }

        /// <summary>
        /// Completion date in RFC3339 format.
        /// </summary>
        [Input("doneDate")]
        public Input<string>? DoneDate { get; set; }

        /// <summary>
        /// Function name (should be `hardReboot`).
        /// </summary>
        [Input("function")]
        public Input<string>? Function { get; set; }

        [Input("keepers")]
        private InputList<string>? _keepers;

        /// <summary>
        /// List of values traccked to trigger reboot, used also to form implicit dependencies
        /// </summary>
        public InputList<string> Keepers
        {
            get => _keepers ?? (_keepers = new InputList<string>());
            set => _keepers = value;
        }

        /// <summary>
        /// Last update in RFC3339 format.
        /// </summary>
        [Input("lastUpdate")]
        public Input<string>? LastUpdate { get; set; }

        /// <summary>
        /// The service_name of your dedicated server.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Task creation date in RFC3339 format.
        /// </summary>
        [Input("startDate")]
        public Input<string>? StartDate { get; set; }

        /// <summary>
        /// Task status (should be `done`)
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DedicatedServerRebootTaskState()
        {
        }
        public static new DedicatedServerRebootTaskState Empty => new DedicatedServerRebootTaskState();
    }
}
