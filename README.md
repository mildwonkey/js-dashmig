# dashboards in jsonschema 

*NOTE*: lots of libraries to choose from: https://json-schema.org/implementations.html

## Get a json schema for dashboards
I started with a fun and hacky cue -> openAPI -> jsonschema transformation (openAPI is a _subset_ of jsonschema); the "real" dashboard schema will likely need some massaging.

hacky prereqs:
* cue installed
* openapi-transformer-toolkit (npm install)

1. Started with the dashboard_kind.cue from grafana/grafana
1. Removed everything but the "spec" ( -> #Dashboard)
1. `cue export --out openapi dashboard_kind.cue > dashboard_kind.openapi`
1. `openapi-transformer-toolkit oas2json -i  dashboard_kind.openapi -o ./schemas`

## what do

* gotypes from jsonschemas: https://github.com/omissis/go-jsonschema
`gojsonschema -p dashboard schemas/Dashboard.json -o internal/kinds/dashboard.go`

* read a json file & validate against jsonschema
* migrate an older version to latest & export
* panel cfg & options - get schema from external source 

Validation:
Started with https://github.com/santhosh-tekuri/jsonschema
- validates schemas against meta-schema
- full support of remote references

## stretch goals! 
* json schema store webservice dealie