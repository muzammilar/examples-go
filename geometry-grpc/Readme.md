# gRPC Example

A basic golang gRPC example (with server certificates).

```sh
## Step 1: (Can be skipped)
# Build the protobufs (and generate the *.pb.go files).
# The second command will run "make protos" in the container as well as recreate go module/update dependencies
docker-compose up --detach --build protobuilder
docker-compose run --rm protobuilder

## Step 2:
# Build the base image.
# Since docker-compose currently doesn't support copying from local images (between multiple Dockerfiles),
# we use the host Operating system to copy the files as needed.
docker-compose up --detach --build basebuilder
docker-compose run --rm basebuilder
# To access the container itself, use any of the following command(s)
docker-compose run --rm basebuilder bash
docker-compose run --rm basebuilder sh
docker-compose run --rm basebuilder /bin/sh


# Run the grpc servers and clients
docker-compose up --build --detach


# Shutdown everything (and remove volumes)
docker-compose rm --force --stop -v
```

## Examples:
Some good gRPC examples are available here:
* https://github.com/grpc/grpc-web
