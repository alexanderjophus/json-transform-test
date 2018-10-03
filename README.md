# json-transform-test

This is a simple repo with a couple of parts for the purpose of testing json transformation tools/frameworks

Frameworks being tested;

| Language  | Framework | pachyderm implementation |
| - | - | - |
| Go  | [qntfy/kazaam](https://github.com/qntfy/kazaam)  | :construction: |
| Python  | [Onyo/jsonbender](https://github.com/Onyo/jsonbender)  | :x: |
| Javascript  | [json-transforms](https://github.com/ColinEberhardt/json-transforms) | :x: |
| JVM | [bazaarvoice/jolt](https://github.com/bazaarvoice/jolt) | :x: |

The file based solution will be using pachyderm, this allows us a few benefits such as data provenance.
This does require us to build a docker image for each of these frameworks.

## Services

The following are helper services

### feed-gen

This is the feed generator, it produces bad data. Sometimes integers, sometimes strings. Sometimes one format, sometimes another. Sometimes nothing at all.

TODO:

- Make the output of this deterministic (or at least expose an endpoint that is)
- Add a "correct" format response somewhere (possibly output to a pachyderm repo)
- Add a `doc` top level, make the API worse (move `startDate` into a time struct with end and expected duration?)

### feed-gen-pachydem-wrapper

This service is a simple wrapper that makes requests to feed-gen, then places those files into a pachyderm `feed-raw` repo

## Usage

Each services README should have a usage section. Order does not matter, `feed-gen-pachyderm-wrapper` depends on `feed-gen`.
