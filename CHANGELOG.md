## Unreleased

### Fixed

- Allow unknows value in the `dockerfile.inline` field during the preview. (https://github.com/pulumi/pulumi-docker-build/pull/89)
- Fixed the default value for `ACTIONS_CACHE_URL` when using GitHub action caching. (https://github.com/pulumi/pulumi-docker-build/pull/80)
- Fixed Java SDK publishing. (https://github.com/pulumi/pulumi-docker-build/pull/89)
- Fixed a panic that could occur when `context` was omitted. (https://github.com/pulumi/pulumi-docker-build/pull/83)

## 0.0.2 (2024-04-25)

### Fixed

- Upgraded pulumi-go-provider to fix a panic during cancellation.

## 0.0.1 (2024-04-23)

Initial release.
