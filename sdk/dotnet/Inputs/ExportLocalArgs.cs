// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DockerBuild.Inputs
{

    public sealed class ExportLocalArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output path.
        /// </summary>
        [Input("dest", required: true)]
        public Input<string> Dest { get; set; } = null!;

        public ExportLocalArgs()
        {
        }
        public static new ExportLocalArgs Empty => new ExportLocalArgs();
    }
}
