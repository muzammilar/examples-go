# gRPC Example

A basic golang gRPC example (with server certificates).

```sh
# Build the protobufs (and generate the *.pb.go files)
docker-compose up --detach --build protobuilder
docker-compose run --rm --workdir="/code" protobuilder "make protos"

# Run the grpc servers and clients
docker-compose up --build


# Shutdown everything (and remove volumes)
docker-compose rm --force --stop -v
```

## Examples:
Some good gRPC examples are available here:
* https://github.com/grpc/grpc-web
