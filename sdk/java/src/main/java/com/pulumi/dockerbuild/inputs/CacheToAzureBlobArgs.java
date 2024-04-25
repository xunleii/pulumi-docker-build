// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.dockerbuild.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.core.internal.Codegen;
import com.pulumi.dockerbuild.enums.CacheMode;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CacheToAzureBlobArgs extends com.pulumi.resources.ResourceArgs {

    public static final CacheToAzureBlobArgs Empty = new CacheToAzureBlobArgs();

    /**
     * Base URL of the storage account.
     * 
     */
    @Import(name="accountUrl")
    private @Nullable Output<String> accountUrl;

    /**
     * @return Base URL of the storage account.
     * 
     */
    public Optional<Output<String>> accountUrl() {
        return Optional.ofNullable(this.accountUrl);
    }

    /**
     * Ignore errors caused by failed cache exports.
     * 
     */
    @Import(name="ignoreError")
    private @Nullable Output<Boolean> ignoreError;

    /**
     * @return Ignore errors caused by failed cache exports.
     * 
     */
    public Optional<Output<Boolean>> ignoreError() {
        return Optional.ofNullable(this.ignoreError);
    }

    /**
     * The cache mode to use. Defaults to `min`.
     * 
     */
    @Import(name="mode")
    private @Nullable Output<CacheMode> mode;

    /**
     * @return The cache mode to use. Defaults to `min`.
     * 
     */
    public Optional<Output<CacheMode>> mode() {
        return Optional.ofNullable(this.mode);
    }

    /**
     * The name of the cache image.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the cache image.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Blob storage account key.
     * 
     */
    @Import(name="secretAccessKey")
    private @Nullable Output<String> secretAccessKey;

    /**
     * @return Blob storage account key.
     * 
     */
    public Optional<Output<String>> secretAccessKey() {
        return Optional.ofNullable(this.secretAccessKey);
    }

    private CacheToAzureBlobArgs() {}

    private CacheToAzureBlobArgs(CacheToAzureBlobArgs $) {
        this.accountUrl = $.accountUrl;
        this.ignoreError = $.ignoreError;
        this.mode = $.mode;
        this.name = $.name;
        this.secretAccessKey = $.secretAccessKey;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CacheToAzureBlobArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CacheToAzureBlobArgs $;

        public Builder() {
            $ = new CacheToAzureBlobArgs();
        }

        public Builder(CacheToAzureBlobArgs defaults) {
            $ = new CacheToAzureBlobArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountUrl Base URL of the storage account.
         * 
         * @return builder
         * 
         */
        public Builder accountUrl(@Nullable Output<String> accountUrl) {
            $.accountUrl = accountUrl;
            return this;
        }

        /**
         * @param accountUrl Base URL of the storage account.
         * 
         * @return builder
         * 
         */
        public Builder accountUrl(String accountUrl) {
            return accountUrl(Output.of(accountUrl));
        }

        /**
         * @param ignoreError Ignore errors caused by failed cache exports.
         * 
         * @return builder
         * 
         */
        public Builder ignoreError(@Nullable Output<Boolean> ignoreError) {
            $.ignoreError = ignoreError;
            return this;
        }

        /**
         * @param ignoreError Ignore errors caused by failed cache exports.
         * 
         * @return builder
         * 
         */
        public Builder ignoreError(Boolean ignoreError) {
            return ignoreError(Output.of(ignoreError));
        }

        /**
         * @param mode The cache mode to use. Defaults to `min`.
         * 
         * @return builder
         * 
         */
        public Builder mode(@Nullable Output<CacheMode> mode) {
            $.mode = mode;
            return this;
        }

        /**
         * @param mode The cache mode to use. Defaults to `min`.
         * 
         * @return builder
         * 
         */
        public Builder mode(CacheMode mode) {
            return mode(Output.of(mode));
        }

        /**
         * @param name The name of the cache image.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the cache image.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param secretAccessKey Blob storage account key.
         * 
         * @return builder
         * 
         */
        public Builder secretAccessKey(@Nullable Output<String> secretAccessKey) {
            $.secretAccessKey = secretAccessKey;
            return this;
        }

        /**
         * @param secretAccessKey Blob storage account key.
         * 
         * @return builder
         * 
         */
        public Builder secretAccessKey(String secretAccessKey) {
            return secretAccessKey(Output.of(secretAccessKey));
        }

        public CacheToAzureBlobArgs build() {
            $.ignoreError = Codegen.booleanProp("ignoreError").output().arg($.ignoreError).def(false).getNullable();
            $.mode = Codegen.objectProp("mode", CacheMode.class).output().arg($.mode).def(CacheMode.Min).getNullable();
            if ($.name == null) {
                throw new MissingRequiredPropertyException("CacheToAzureBlobArgs", "name");
            }
            return $;
        }
    }

}
