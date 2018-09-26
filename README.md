# json-transform-test

This is a simple repo with a couple of parts for the purpose of testing json transformation tools/frameworks

Frameworks being tested;

| Language  | Framework | Complete? |
| - | - | - |
| Go  | [qntfy/kazaam](https://github.com/qntfy/kazaam)  | :x: |
| Python  | [Onyo/jsonbender](https://github.com/Onyo/jsonbender)  | :x: |
| Javascript  | [json-transforms](https://github.com/ColinEberhardt/json-transforms) | :x: |


## feed-gen

This is the feed generator, it produces bad data. Sometimes integers, sometimes strings. Sometimes one format, sometimes another. Sometimes nothing at all.

TODO:

- Make the output of this deterministic (or at least expose an endpoint that is)

