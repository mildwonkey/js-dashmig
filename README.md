# dashboards in jsonschema 

## Overview 
This repository contains experiments converting our CUE-coded schemas into
jsonschema format (via openAPI spec) and modeling schema & data operations using
jsonschema. 

*NOTE*: lots of libraries to choose from: https://json-schema.org/implementations.html

## Create a json schema for dashboards
I started with a fun and hacky cue -> openAPI -> jsonschema transformation (openAPI is a _subset_ of jsonschema); the "real" dashboard schema will likely need some massaging.

### hacky prereqs:
* [cue](https://github.com/cue-lang/cue)    
* [openapi-transformer-toolkit](https://github.com/nearform/openapi-transformer-toolkit)
* [gojsonschema](https://github.com/omissis/go-jsonschema)

1. Start with the dashboard_kind.cue from grafana/grafana
1. Remove everything but the `"spec"` section & renamed to #Dashboard
1. openAPI spec: `cue export --out openapi dashboard_kind.cue > dashboard_kind.openapi`
1. jsonschema spec: `openapi-transformer-toolkit oas2json -i  dashboard_kind.openapi -o ./schemas`
1. gotypes: `gojsonschema -p dashboard schemas/Dashboard.json -o internal/kinds/dashboard.go` 

I am struggling with gojsonschema the library - the most basic usage isn't working for me, and the codebase hasn't been touched in years, so I switched to "github.com/santhosh-tekuri/jsonschema" (that had been my first choice for a library; `gojsonschema` was what i used to generate the go types)

## Doin' Stuff
  - [x] read a json file & validate against jsonschema
    - the go types generated w/`gojsonschema` include custom unmarshallers
    - one possible migration plan is to extend those unmarshallers to support older formats
  - [] migrate an older version to latest & export
  - [] meta schemas: kindsys (core, custom, composable)
  - [] meta schemas: panel cfg & options

### Validation:
Using https://github.com/santhosh-tekuri/jsonschema
- [x] validate against local schema
- [x] validate schemas against meta-schema (draft/2020-12)
- [] validate schemas against meta-schema (core kind)
- [] validate schemas with remote references (see stretch goals below)

## stretch goals!
- [] json schema store webservice dealie 
- [] individual plugin serving schema over gRPC 

### Appendix

Generated schema weirdnesses (probably due to Cue -> openAPI -> json weirdnesses)

* DashboardCursorSync enum should have been 0, 1, 2 but ended up w/ 0.0, 0.1 etc in  `enumValues_DashboardCursorSyncJson` (fixed manually for now)
* exclusiveMinimum and exclusiveMaximum were originally bools in the generated schema, but they should be int.