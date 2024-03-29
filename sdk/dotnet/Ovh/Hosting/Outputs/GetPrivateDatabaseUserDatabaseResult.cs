// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs.PulumiPackage.Ovh.Hosting.Outputs
{

    [OutputType]
    public sealed class GetPrivateDatabaseUserDatabaseResult
    {
        /// <summary>
        /// Database's name linked to this user
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// Grant of this user for this database
        /// </summary>
        public readonly string GrantType;

        [OutputConstructor]
        private GetPrivateDatabaseUserDatabaseResult(
            string databaseName,

            string grantType)
        {
            DatabaseName = databaseName;
            GrantType = grantType;
        }
    }
}
