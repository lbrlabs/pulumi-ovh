# Ovh Resource Provider

The Ovh Resource Provider lets you manage cloud resources for [Ovh](http://ovh.com) resources.

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @lbrlabs/ovh
```

or `yarn`:

```bash
yarn add @lbrlabs/ovh
```

### Python

To use from Python, install using `pip`:

```bash
pip install lbrlabs_ovh
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/lbrlabs/pulumi-ovh/sdk/go/...
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package lbrlabs.Ovh
```

## Configuration

The following configuration points are available for the `Ovh` provider:

- `ovh:endpoint` (environment: `OVH_ENDPOINT`) - the Ovh endpoint, such `ovh-us`
- `ovh:applicationKey` (environment: `OVH_APPLICATION_KEY`) - the Ovh application key
- `ovh:applicationSecret` (environment: `OVH_APPLICATION_SECRET`) - the Ovh application secret
- `ovh:consumerKey` (environment: `OVH_CONSUMER_KEY`) - the Ovh consumer key

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/ovh/api-docs/).

