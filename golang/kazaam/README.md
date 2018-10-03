# kazaam implementation

This directory is for the kazaam json transformation experiment

[qntfy/kazaam](https://github.com/qntfy/kazaam)

## Quick Start (pachyderm)

### pre-requisites

- pachyderm running locally
- populated `feed-raw` repo (possible through the [feed-gen-pachyderm-wrapper](../../feed-gen-pachyderm-wrapper/README.md))

### steps

- `make docker-build`
- `make docker-push`
- `make pipeline`
