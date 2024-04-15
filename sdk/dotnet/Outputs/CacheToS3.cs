// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DockerBuild.Outputs
{

    [OutputType]
    public sealed class CacheToS3
    {
        /// <summary>
        /// Defaults to `$AWS_ACCESS_KEY_ID`.
        /// </summary>
        public readonly string? AccessKeyId;
        /// <summary>
        /// Prefix to prepend to blob filenames.
        /// </summary>
        public readonly string? BlobsPrefix;
        /// <summary>
        /// Name of the S3 bucket.
        /// </summary>
        public readonly string Bucket;
        /// <summary>
        /// Endpoint of the S3 bucket.
        /// </summary>
        public readonly string? EndpointUrl;
        /// <summary>
        /// Ignore errors caused by failed cache exports.
        /// </summary>
        public readonly bool? IgnoreError;
        /// <summary>
        /// Prefix to prepend on manifest filenames.
        /// </summary>
        public readonly string? ManifestsPrefix;
        /// <summary>
        /// The cache mode to use. Defaults to `min`.
        /// </summary>
        public readonly Pulumi.DockerBuild.CacheMode? Mode;
        /// <summary>
        /// Name of the cache image.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The geographic location of the bucket. Defaults to `$AWS_REGION`.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Defaults to `$AWS_SECRET_ACCESS_KEY`.
        /// </summary>
        public readonly string? SecretAccessKey;
        /// <summary>
        /// Defaults to `$AWS_SESSION_TOKEN`.
        /// </summary>
        public readonly string? SessionToken;
        /// <summary>
        /// Uses `bucket` in the URL instead of hostname when `true`.
        /// </summary>
        public readonly bool? UsePathStyle;

        [OutputConstructor]
        private CacheToS3(
            string? accessKeyId,

            string? blobsPrefix,

            string bucket,

            string? endpointUrl,

            bool? ignoreError,

            string? manifestsPrefix,

            Pulumi.DockerBuild.CacheMode? mode,

            string? name,

            string region,

            string? secretAccessKey,

            string? sessionToken,

            bool? usePathStyle)
        {
            AccessKeyId = accessKeyId;
            BlobsPrefix = blobsPrefix;
            Bucket = bucket;
            EndpointUrl = endpointUrl;
            IgnoreError = ignoreError;
            ManifestsPrefix = manifestsPrefix;
            Mode = mode;
            Name = name;
            Region = region;
            SecretAccessKey = secretAccessKey;
            SessionToken = sessionToken;
            UsePathStyle = usePathStyle;
        }
    }
}
