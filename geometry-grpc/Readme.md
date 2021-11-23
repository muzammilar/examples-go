# gRPC Example

A basic golang gRPC example (with server certificates).

```sh
# Build the protobufs (and generate the *.pb.go files).
# The second command will run "make protos" in the container as well as recreate go module/update dependencies
docker-compose up --detach --build protobuilder
docker-compose run --rm protobuilder
# To access the container itself, use any of the following command(s)
docker-compose run --rm protobuilder bash
docker-compose run --rm protobuilder sh
docker-compose run --rm protobuilder /bin/sh

# Run the grpc servers and clients
docker-compose up --build


# Shutdown everything (and remove volumes)
docker-compose rm --force --stop -v
```

## Examples:
Some good gRPC examples are available here:
* https://github.com/grpc/grpc-web
