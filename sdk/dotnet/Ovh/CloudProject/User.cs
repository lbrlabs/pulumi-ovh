// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.CloudProject
{
    /// <summary>
    /// Creates a user in a public cloud project.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Lbrlabs.PulumiPackage.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var user1 = new Ovh.CloudProject.User("user1", new()
    ///     {
    ///         ServiceName = "XXX",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [OvhResourceType("ovh:CloudProject/user:User")]
    public partial class User : global::Pulumi.CustomResource
    {
        /// <summary>
        /// the date the user was created.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// A description associated with the user.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// a convenient map representing an openstack_rc file.
        /// Note: no password nor sensitive token is set in this map.
        /// </summary>
        [Output("openstackRc")]
        public Output<ImmutableDictionary<string, object>> OpenstackRc { get; private set; } = null!;

        /// <summary>
        /// (Sensitive) the password generated for the user. The password can
        /// be used with the Openstack API. This attribute is sensitive and will only be
        /// retrieve once during creation.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The name of a role. See `role_names`.
        /// </summary>
        [Output("roleName")]
        public Output<string?> RoleName { get; private set; } = null!;

        /// <summary>
        /// A list of role names. Values can be: 
        /// - administrator,
        /// - ai_training_operator
        /// - authentication
        /// - backup_operator
        /// - compute_operator
        /// - image_operator
        /// - infrastructure_supervisor
        /// - network_operator
        /// - network_security_operator
        /// - objectstore_operator
        /// - volume_operator
        /// </summary>
        [Output("roleNames")]
        public Output<ImmutableArray<string>> RoleNames { get; private set; } = null!;

        /// <summary>
        /// A list of roles associated with the user.
        /// </summary>
        [Output("roles")]
        public Output<ImmutableArray<Outputs.UserRole>> Roles { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// the status of the user. should be normally set to 'ok'.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// the username generated for the user. This username can be used with
        /// the Openstack API.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/user:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/user:User", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/lbrlabs",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, UserState? state = null, CustomResourceOptions? options = null)
        {
            return new User(name, id, state, options);
        }
    }

    public sealed class UserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description associated with the user.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("openstackRc")]
        private InputMap<object>? _openstackRc;

        /// <summary>
        /// a convenient map representing an openstack_rc file.
        /// Note: no password nor sensitive token is set in this map.
        /// </summary>
        public InputMap<object> OpenstackRc
        {
            get => _openstackRc ?? (_openstackRc = new InputMap<object>());
            set => _openstackRc = value;
        }

        /// <summary>
        /// The name of a role. See `role_names`.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        [Input("roleNames")]
        private InputList<string>? _roleNames;

        /// <summary>
        /// A list of role names. Values can be: 
        /// - administrator,
        /// - ai_training_operator
        /// - authentication
        /// - backup_operator
        /// - compute_operator
        /// - image_operator
        /// - infrastructure_supervisor
        /// - network_operator
        /// - network_security_operator
        /// - objectstore_operator
        /// - volume_operator
        /// </summary>
        public InputList<string> RoleNames
        {
            get => _roleNames ?? (_roleNames = new InputList<string>());
            set => _roleNames = value;
        }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public UserArgs()
        {
        }
        public static new UserArgs Empty => new UserArgs();
    }

    public sealed class UserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the date the user was created.
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        /// <summary>
        /// A description associated with the user.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("openstackRc")]
        private InputMap<object>? _openstackRc;

        /// <summary>
        /// a convenient map representing an openstack_rc file.
        /// Note: no password nor sensitive token is set in this map.
        /// </summary>
        public InputMap<object> OpenstackRc
        {
            get => _openstackRc ?? (_openstackRc = new InputMap<object>());
            set => _openstackRc = value;
        }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// (Sensitive) the password generated for the user. The password can
        /// be used with the Openstack API. This attribute is sensitive and will only be
        /// retrieve once during creation.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The name of a role. See `role_names`.
        /// </summary>
        [Input("roleName")]
        public Input<string>? RoleName { get; set; }

        [Input("roleNames")]
        private InputList<string>? _roleNames;

        /// <summary>
        /// A list of role names. Values can be: 
        /// - administrator,
        /// - ai_training_operator
        /// - authentication
        /// - backup_operator
        /// - compute_operator
        /// - image_operator
        /// - infrastructure_supervisor
        /// - network_operator
        /// - network_security_operator
        /// - objectstore_operator
        /// - volume_operator
        /// </summary>
        public InputList<string> RoleNames
        {
            get => _roleNames ?? (_roleNames = new InputList<string>());
            set => _roleNames = value;
        }

        [Input("roles")]
        private InputList<Inputs.UserRoleGetArgs>? _roles;

        /// <summary>
        /// A list of roles associated with the user.
        /// </summary>
        public InputList<Inputs.UserRoleGetArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.UserRoleGetArgs>());
            set => _roles = value;
        }

        /// <summary>
        /// The id of the public cloud project. If omitted,
        /// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// the status of the user. should be normally set to 'ok'.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// the username generated for the user. This username can be used with
        /// the Openstack API.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public UserState()
        {
        }
        public static new UserState Empty => new UserState();
    }
}
