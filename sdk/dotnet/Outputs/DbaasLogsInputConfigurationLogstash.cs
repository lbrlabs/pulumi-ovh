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
    public sealed class DbaasLogsInputConfigurationLogstash
    {
        /// <summary>
        /// The filter section of logstash.conf
        /// </summary>
        public readonly string? FilterSection;
        /// <summary>
        /// The filter section of logstash.conf
        /// </summary>
        public readonly string InputSection;
        /// <summary>
        /// The list of customs Grok patterns
        /// </summary>
        public readonly string? PatternSection;

        [OutputConstructor]
        private DbaasLogsInputConfigurationLogstash(
            string? filterSection,

            string inputSection,

            string? patternSection)
        {
            FilterSection = filterSection;
            InputSection = inputSection;
            PatternSection = patternSection;
        }
    }
}
