## Setting up the environment

To set up the repository, run:

```sh
$ ./scripts/bootstrap
$ ./scripts/build
```

This will install all the required dependencies and build the Provider binary into the root directory.

You can also [install go 1.22+ manually](https://go.dev/doc/install).

## Public API types

`proto/` is a copy of the Serval public API protos from ServalHQ/mono (the files reachable from
`svflow/publicapi/service.proto` and `svauth/public/service.proto`), and `internal/gen/proto` is
generated from it. Don't edit either by hand. To refresh both, install
[`buf`](https://buf.build/docs/installation) and `protoc-gen-go` v1.36.11, then run:

```sh
$ MONO_DIR=/path/to/mono ./scripts/sync-protos
```

`proto/MONO_COMMIT` records the mono commit the copy came from.

## Running the Provider locally

You can build the provider locally and have your `.tf` files refer to the local build instead of the one in the Hashicorp registry.

First, build the provider binary:

```sh
$ ./scripts/build
```

Then edit (or create) your `~/.terraformrc` to look something like this:

```hcl
  provider_installation {
    dev_overrides {
      "ServalHQ/serval" = "/local/path/to/this/repo"
    }
    direct {}
  }
```

## Running tests

To execute the schema and unit tests, run:

```sh
$ ./scripts/test
```

Note that this does not run [acceptance tests](https://developer.hashicorp.com/terraform/plugin/framework/acctests) by default, because
those tests interact with real resources in the cloud and could incur fees.

To enable the running of acceptance tests, use the `TF_ACC` environment variable:

```sh
$ TF_ACC=1 ./scripts/test
```

## Formatting

This library uses the standard gofmt code formatter:

```sh
$ ./scripts/format
```

## Running Tests

To run schema tests,
