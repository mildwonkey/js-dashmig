# dashboards in jsonschema 

*NOTE*: lots of libraries to choose from: https://json-schema.org/implementations.html

## Create a json schema for dashboards
I started with a fun and hacky cue -> openAPI -> jsonschema transformation (openAPI is a _subset_ of jsonschema); the "real" dashboard schema will likely need some massaging.

### hacky prereqs:
* cue installed
* openapi-transformer-toolkit (npm install)

1. Started with the dashboard_kind.cue from grafana/grafana
1. Removed everything but the "spec" (renamed to -> #Dashboard)
1. openAPI spec: `cue export --out openapi dashboard_kind.cue > dashboard_kind.openapi`
1. jsonschema spec: `openapi-transformer-toolkit oas2json -i  dashboard_kind.openapi -o ./schemas`
1. gotypes: `gojsonschema -p dashboard schemas/Dashboard.json -o internal/kinds/dashboard.go` [gojsonschema](https://github.com/omissis/go-jsonschema)

## Doin' Stuff

* read a json file & validate against jsonschema
- the go types generated w/`gojsonschema` include custom unmarshallers

* migrate an older version to latest & export
* metaschema: kindsys (core, custom, composable)
* meta schemas: panel cfg & options

### Validation:
Look at with https://github.com/santhosh-tekuri/jsonschema
- validates schemas against meta-schema
- full support of remote references

## stretch goals!
* json schema store webservice dealie 
* individual plugin serves schema (gRPC?) 