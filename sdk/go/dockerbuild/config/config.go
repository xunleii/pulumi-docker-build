// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-docker-build/sdk/go/dockerbuild/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// The build daemon's address.
func GetHost(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "docker-build:host")
	if err == nil {
		return v
	}
	var value string
	if d := internal.GetEnvOrDefault("", nil, "DOCKER_HOST"); d != nil {
		value = d.(string)
	}
	return value
}
func GetRegistries(ctx *pulumi.Context) string {
	return config.Get(ctx, "docker-build:registries")
}
