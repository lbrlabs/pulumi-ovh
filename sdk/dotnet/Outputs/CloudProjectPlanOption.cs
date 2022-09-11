// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Lbrlabs_Pulumi.Ovh.Outputs
{

    [OutputType]
    public sealed class CloudProjectPlanOption
    {
        /// <summary>
        /// Catalog name
        /// </summary>
        public readonly string? CatalogName;
        /// <summary>
        /// Representation of a configuration item for personalizing product
        /// </summary>
        public readonly ImmutableArray<Outputs.CloudProjectPlanOptionConfiguration> Configurations;
        /// <summary>
        /// duration
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// Plan code
        /// </summary>
        public readonly string PlanCode;
        /// <summary>
        /// Pricing model identifier
        /// </summary>
        public readonly string PricingMode;

        [OutputConstructor]
        private CloudProjectPlanOption(
            string? catalogName,

            ImmutableArray<Outputs.CloudProjectPlanOptionConfiguration> configurations,

            string duration,

            string planCode,

            string pricingMode)
        {
            CatalogName = catalogName;
            Configurations = configurations;
            Duration = duration;
            PlanCode = planCode;
            PricingMode = pricingMode;
        }
    }
}
