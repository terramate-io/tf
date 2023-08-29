# Building from Source

## Prerequisites

1. Ensure you've installed the Go language version specified in [`.go-version`](https://github.com/terramate-io/tf/blob/main/.go-version).
2. Clone this repository to a location of your choice.

## Terraform Build Options

Terraform accepts certain options passed using `ldflags` at build time which control the behavior of the resulting binary.

### Dev Version Reporting

Terraform will include a `-dev` flag when reporting its own version (ex: 1.5.0-dev) unless `version.dev` is set to `no`:

```
go build -ldflags "-w -s -X 'github.com/terramate-io/tf/version.dev=no'" -o bin/ .
```

### Experimental Features

Experimental features of Terraform will be disabled unless `main.experimentsAllowed` is set to `yes`:

```
go build -ldflags "-w -s -X 'main.experimentsAllowed=yes'" -o bin/ .
```

In the official build process for Terraform, experiments are only allowed in alpha release builds. We recommend that third-party distributors follow that convention in order to reduce user confusion.

## Go Options

For the most part, the Terraform release process relies on the Go toolchain defaults for the target operating system and processor architecture.

### `CGO_ENABLED`

One exception is the `CGO_ENABLED` option, which is set explicitly when building Terraform binaries. For most platforms, we build with `CGO_ENABLED=0` in order to produce a statically linked binary. For MacOS/Darwin operating systems, we build with `CGO_ENABLED=1` to avoid a platform-specific issue with DNS resolution. 


