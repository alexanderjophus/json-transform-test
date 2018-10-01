# json-transform-test

This is a simple repo with a couple of parts for the purpose of testing json transformation tools/frameworks

Frameworks being tested;

| Language  | Framework | HTTP Complete | File based complete |
| - | - | - |
| Go  | [qntfy/kazaam](https://github.com/qntfy/kazaam)  | :x: | :x: |
| Python  | [Onyo/jsonbender](https://github.com/Onyo/jsonbender)  | :x: | :x: |
| Javascript  | [json-transforms](https://github.com/ColinEberhardt/json-transforms) | :x: | :x: |
| JVM | [bazaarvoice/jolt](https://github.com/bazaarvoice/jolt) | :x: | :x: |

The file based solution will be using pachyderm, this allows us a few benefits such as data provenance.
This does require us to build a docker image for each of these frameworks.

## feed-gen

This is the feed generator, it produces bad data. Sometimes integers, sometimes strings. Sometimes one format, sometimes another. Sometimes nothing at all.

TODO:

- Make the output of this deterministic (or at least expose an endpoint that is)
- add a "correct" format response somewhere

### Feed Gen Usage

```sh
DEBUG=feed-gen:* npm start
```

## feed-gen-pachydem-wrapper

This service is a simple wrapper that makes requests to feed-gen, then places those files into a pachyderm `feed-raw` repo

### Pachyderm Wrapper pre-requisites

Follow the pachyderm getting [started guide](https://pachyderm.readthedocs.io/en/stable/getting_started/local_installation.html)

Create the `feed-raw` repo

### Pachyderm Wrapper Usage

```sh
go run main.go
```
