# feed-gen-pachyderm-wrapper

This API wraps the node API and dumps the feed into a pachyderm repo.
This allows any number of pipelines to be generated simply ingesting the repo, making an easy comparison of json transformers.

## pre-requisites

Follow the pachyderm getting [started guide](https://pachyderm.readthedocs.io/en/stable/getting_started/local_installation.html)

Create the `feed-raw` repo

## Usage

```sh
go run main.go
```
