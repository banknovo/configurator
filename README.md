# Configurator
Configurator is a tool for fetching configuration from AWS Parameters Store and convert and export in required format. This library is heavily inspired from [chamber](https://github.com/segmentio/chamber), which was evaluated, but it did not meet all requirements, specifically hierarchical configuration and configuration that is common between multiple services/applications.

## Concepts
This library makes heavy use of the fact that AWS Parameters Store supports hierarchical configuration. Thus, it translates it back to configuration which preserves this hierarchy.

There are 3 modes of export supported:
1. Flat: Creates a flat representation of parameters in Parameters Store, it ignores the nesting in parameters store and treats key names as-is.
2. Hierarchical: It creates a hierarchical representation based on the nested keys in parameters store.
3. Blueprint-Based: This mode is NOT RECOMMENDED, flat/hierarchical mode should be used. This is for manual mapping to an existing document (json mostly) where parameters store values are to be replaced in a pre-existing document. This was created for backward compatibility with json based configuration in existing apps, so no heavy changes would be required.

## Usage

Currently, only exporting as JSON is supported.

```
$ ./configurator export -a <app> -e <environment> [-t <common-config1> -t <common-config2>...] -m <mode> -f json -o <output-file>
```
It assumes that the keys in Parameters Store are defined like the following:
- `/environment/app/...`
- `/environment/common-config-1/...`
- `/environment/common-config-2/...`

If you're using `mode=blueprint`, an additional parameter is required which is path of blueprint json file:
```
$ ./configurator export -a <app> -e <environment> -t <common-config1> -t <common-config2> -m <mode> -b <path-to-blueprint-json> -f json -o <output-file>
```
If `-o <output-file>` is omitted, it prints to standard output by default.