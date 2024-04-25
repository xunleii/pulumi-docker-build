// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DockerBuild.Inputs
{

    public sealed class CacheFromS3Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to `$AWS_ACCESS_KEY_ID`.
        /// </summary>
        [Input("accessKeyId")]
        public Input<string>? AccessKeyId { get; set; }

        /// <summary>
        /// Prefix to prepend to blob filenames.
        /// </summary>
        [Input("blobsPrefix")]
        public Input<string>? BlobsPrefix { get; set; }

        /// <summary>
        /// Name of the S3 bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Endpoint of the S3 bucket.
        /// </summary>
        [Input("endpointUrl")]
        public Input<string>? EndpointUrl { get; set; }

        /// <summary>
        /// Prefix to prepend on manifest filenames.
        /// </summary>
        [Input("manifestsPrefix")]
        public Input<string>? ManifestsPrefix { get; set; }

        /// <summary>
        /// Name of the cache image.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The geographic location of the bucket. Defaults to `$AWS_REGION`.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("secretAccessKey")]
        private Input<string>? _secretAccessKey;

        /// <summary>
        /// Defaults to `$AWS_SECRET_ACCESS_KEY`.
        /// </summary>
        public Input<string>? SecretAccessKey
        {
            get => _secretAccessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretAccessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("sessionToken")]
        private Input<string>? _sessionToken;

        /// <summary>
        /// Defaults to `$AWS_SESSION_TOKEN`.
        /// </summary>
        public Input<string>? SessionToken
        {
            get => _sessionToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _sessionToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Uses `bucket` in the URL instead of hostname when `true`.
        /// </summary>
        [Input("usePathStyle")]
        public Input<bool>? UsePathStyle { get; set; }

        public CacheFromS3Args()
        {
            AccessKeyId = Utilities.GetEnv("AWS_ACCESS_KEY_ID") ?? "";
            Region = Utilities.GetEnv("AWS_REGION") ?? "";
            SecretAccessKey = Utilities.GetEnv("AWS_SECRET_ACCESS_KEY") ?? "";
            SessionToken = Utilities.GetEnv("AWS_SESSION_TOKEN") ?? "";
        }
        public static new CacheFromS3Args Empty => new CacheFromS3Args();
    }
}
