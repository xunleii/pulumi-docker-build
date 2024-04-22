// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dockerbuild

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-docker-build/sdk/go/dockerbuild/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A wrapper around `docker buildx imagetools create` to create an index
// (or manifest list) referencing one or more existing images.
//
// In most cases you do not need an `Index` to build a multi-platform
// image -- specifying multiple platforms on the `Image` will handle this
// for you automatically.
//
// However, as of April 2024, building multi-platform images _with
// caching_ will only export a cache for one platform at a time (see [this
// discussion](https://github.com/docker/buildx/discussions/1382) for more
// details).
//
// Therefore this resource can be helpful if you are building
// multi-platform images with caching: each platform can be built and
// cached separately, and an `Index` can join them all together. An
// example of this is shown below.
//
// This resource creates an OCI image index or a Docker manifest list
// depending on the media types of the source images.
//
// ## Example Usage
// ### Multi-platform registry caching
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-docker-build/sdk/go/dockerbuild"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			amd64, err := dockerbuild.NewImage(ctx, "amd64", &dockerbuild.ImageArgs{
//				CacheFrom: dockerbuild.CacheFromArray{
//					&dockerbuild.CacheFromArgs{
//						Registry: &dockerbuild.CacheFromRegistryArgs{
//							Ref: pulumi.String("docker.io/pulumi/pulumi:cache-amd64"),
//						},
//					},
//				},
//				CacheTo: dockerbuild.CacheToArray{
//					&dockerbuild.CacheToArgs{
//						Registry: &dockerbuild.CacheToRegistryArgs{
//							Mode: dockerbuild.CacheModeMax,
//							Ref:  pulumi.String("docker.io/pulumi/pulumi:cache-amd64"),
//						},
//					},
//				},
//				Context: &dockerbuild.BuildContextArgs{
//					Location: pulumi.String("app"),
//				},
//				Platforms: docker - build.PlatformArray{
//					dockerbuild.Platform_Linux_amd64,
//				},
//				Tags: pulumi.StringArray{
//					pulumi.String("docker.io/pulumi/pulumi:3.107.0-amd64"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			arm64, err := dockerbuild.NewImage(ctx, "arm64", &dockerbuild.ImageArgs{
//				CacheFrom: dockerbuild.CacheFromArray{
//					&dockerbuild.CacheFromArgs{
//						Registry: &dockerbuild.CacheFromRegistryArgs{
//							Ref: pulumi.String("docker.io/pulumi/pulumi:cache-arm64"),
//						},
//					},
//				},
//				CacheTo: dockerbuild.CacheToArray{
//					&dockerbuild.CacheToArgs{
//						Registry: &dockerbuild.CacheToRegistryArgs{
//							Mode: dockerbuild.CacheModeMax,
//							Ref:  pulumi.String("docker.io/pulumi/pulumi:cache-arm64"),
//						},
//					},
//				},
//				Context: &dockerbuild.BuildContextArgs{
//					Location: pulumi.String("app"),
//				},
//				Platforms: docker - build.PlatformArray{
//					dockerbuild.Platform_Linux_arm64,
//				},
//				Tags: pulumi.StringArray{
//					pulumi.String("docker.io/pulumi/pulumi:3.107.0-arm64"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			index, err := dockerbuild.NewIndex(ctx, "index", &dockerbuild.IndexArgs{
//				Sources: pulumi.StringArray{
//					amd64.Ref,
//					arm64.Ref,
//				},
//				Tag: pulumi.String("docker.io/pulumi/pulumi:3.107.0"),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("ref", index.Ref)
//			return nil
//		})
//	}
//
// ```
type Index struct {
	pulumi.CustomResourceState

	// If true, push the index to the target registry.
	//
	// Defaults to `true`.
	Push pulumix.Output[*bool] `pulumi:"push"`
	// The pushed tag with digest.
	//
	// Identical to the tag if the index was not pushed.
	Ref pulumix.Output[string] `pulumi:"ref"`
	// Authentication for the registry where the tagged index will be pushed.
	//
	// Credentials can also be included with the provider's configuration.
	Registry pulumix.GPtrOutput[Registry, RegistryOutput] `pulumi:"registry"`
	// Existing images to include in the index.
	Sources pulumix.ArrayOutput[string] `pulumi:"sources"`
	// The tag to apply to the index.
	Tag pulumix.Output[string] `pulumi:"tag"`
}

// NewIndex registers a new resource with the given unique name, arguments, and options.
func NewIndex(ctx *pulumi.Context,
	name string, args *IndexArgs, opts ...pulumi.ResourceOption) (*Index, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	if args.Tag == nil {
		return nil, errors.New("invalid value for required argument 'Tag'")
	}
	if args.Push == nil {
		args.Push = pulumix.Ptr(true)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Index
	err := ctx.RegisterResource("docker-build:index:Index", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIndex gets an existing Index resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIndex(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IndexState, opts ...pulumi.ResourceOption) (*Index, error) {
	var resource Index
	err := ctx.ReadResource("docker-build:index:Index", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Index resources.
type indexState struct {
}

type IndexState struct {
}

func (IndexState) ElementType() reflect.Type {
	return reflect.TypeOf((*indexState)(nil)).Elem()
}

type indexArgs struct {
	// If true, push the index to the target registry.
	//
	// Defaults to `true`.
	Push *bool `pulumi:"push"`
	// Authentication for the registry where the tagged index will be pushed.
	//
	// Credentials can also be included with the provider's configuration.
	Registry *Registry `pulumi:"registry"`
	// Existing images to include in the index.
	Sources []string `pulumi:"sources"`
	// The tag to apply to the index.
	Tag string `pulumi:"tag"`
}

// The set of arguments for constructing a Index resource.
type IndexArgs struct {
	// If true, push the index to the target registry.
	//
	// Defaults to `true`.
	Push pulumix.Input[*bool]
	// Authentication for the registry where the tagged index will be pushed.
	//
	// Credentials can also be included with the provider's configuration.
	Registry pulumix.Input[*RegistryArgs]
	// Existing images to include in the index.
	Sources pulumix.Input[[]string]
	// The tag to apply to the index.
	Tag pulumix.Input[string]
}

func (IndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*indexArgs)(nil)).Elem()
}

type IndexOutput struct{ *pulumi.OutputState }

func (IndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Index)(nil)).Elem()
}

func (o IndexOutput) ToIndexOutput() IndexOutput {
	return o
}

func (o IndexOutput) ToIndexOutputWithContext(ctx context.Context) IndexOutput {
	return o
}

func (o IndexOutput) ToOutput(ctx context.Context) pulumix.Output[Index] {
	return pulumix.Output[Index]{
		OutputState: o.OutputState,
	}
}

// If true, push the index to the target registry.
//
// Defaults to `true`.
func (o IndexOutput) Push() pulumix.Output[*bool] {
	value := pulumix.Apply[Index](o, func(v Index) pulumix.Output[*bool] { return v.Push })
	return pulumix.Flatten[*bool, pulumix.Output[*bool]](value)
}

// The pushed tag with digest.
//
// Identical to the tag if the index was not pushed.
func (o IndexOutput) Ref() pulumix.Output[string] {
	value := pulumix.Apply[Index](o, func(v Index) pulumix.Output[string] { return v.Ref })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

// Authentication for the registry where the tagged index will be pushed.
//
// Credentials can also be included with the provider's configuration.
func (o IndexOutput) Registry() pulumix.GPtrOutput[Registry, RegistryOutput] {
	value := pulumix.Apply[Index](o, func(v Index) pulumix.GPtrOutput[Registry, RegistryOutput] { return v.Registry })
	unwrapped := pulumix.Flatten[*Registry, pulumix.GPtrOutput[Registry, RegistryOutput]](value)
	return pulumix.GPtrOutput[Registry, RegistryOutput]{OutputState: unwrapped.OutputState}
}

// Existing images to include in the index.
func (o IndexOutput) Sources() pulumix.ArrayOutput[string] {
	value := pulumix.Apply[Index](o, func(v Index) pulumix.ArrayOutput[string] { return v.Sources })
	unwrapped := pulumix.Flatten[[]string, pulumix.ArrayOutput[string]](value)
	return pulumix.ArrayOutput[string]{OutputState: unwrapped.OutputState}
}

// The tag to apply to the index.
func (o IndexOutput) Tag() pulumix.Output[string] {
	value := pulumix.Apply[Index](o, func(v Index) pulumix.Output[string] { return v.Tag })
	return pulumix.Flatten[string, pulumix.Output[string]](value)
}

func init() {
	pulumi.RegisterOutputType(IndexOutput{})
}
