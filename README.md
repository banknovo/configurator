# Configurator
Configurator is a tool for fetching configuration from AWS Parameters Store and convert and export in required format. This library is heavily inspired from [chamber](https://github.com/segmentio/chamber), which was evaluated, but it did not meet all requirements, specifically hierarchical configuration and configuration that is common between multiple services/applications.

## Concepts
This library makes heavy use of the fact that AWS Parameters Store supports hierarchical configuration. Thus, it translates it back to configuration which preserves this hierarchy.

There are 3 modes of export supported:
1. Flat: Creates a flat representation of parameters in Parameters Store, it ignores the nesting in parameters store and treats key names as-is.
2. Hierarchical: It creates a hierarchical representation based on the nested keys in parameters store.
3. Blueprint-Based: This mode is NOT RECOMMENDED, flat/hierarchical mode should be used. This is for manual mapping to an existing document (json mostly) where parameters store values are to be replaced in a pre-existing document. This was created for backward compatibility with json based configuration in existing apps, so no heavy changes would be required.

## Installation
If you have a functional go environment, you can install with:

```shell script
$ go get github.com/banknovo/configurator
```
To pull the docker image run:
```shell script
$ docker pull banknovo/configurator
```
## Usage
`configurator` uses AWS SDK to call Parameters Store API. It does not take care of authentication and assumes that relevant credentials are available in the environment where the CLI runs.
The easiest way is to export the following variables:
```shell script
$ export AWS_SECRET_ACCESS_KEY=
$ export AWS_ACCESS_KEY_ID=
$ export AWS_DEFAULT_REGION=
```

### Export
This command exports the Parameters into a specified format. Currently, only exporting as JSON and YAML is supported.
```shell script
$ ./configurator export --paths /Path1,/Path2
```
The paths are the secrets defined in Parameters Store, and all properties starting with the path name are fetched.
The parameter `excludePrefix (-x)` is used to remove those number of prefixes from the final export.

For example, if the property is `/Key1/Key2/Key3`, then `excludePrefix=1` will export property names as `Key2/Key3`.

If you're using `mode=blueprint`, an additional parameter is required which is path of blueprint json file:
```shell script
$ ./configurator export --paths /Path1,/Path2 -m blueprint -f json -b <path-to-blueprint-json> -o <output-file>
```
If `-o <output-file>` is omitted, it prints to standard output by default.

Parameter `-i` or `--indent` generates indented/pretty output. 

### Env
This command is used to print the export linux commands on stdout. Because of this, only `flat` mode is supported with `Env`.
```shell script
$ ./configurator env --paths /Path1,/Path2
```