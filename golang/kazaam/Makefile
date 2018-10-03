all: build

build: 
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kazaam .

docker-build: build ## some steps
	docker build -t trelore/kazaam -f Dockerfile .

docker-push: 
	docker push trelore/kazaam:latest

pipeline:
	pachctl create-pipeline -f kazaam.json
