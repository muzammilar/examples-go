# examples-go
Example codes in golang for fun. All examples should have associated `Dockerfile` and `docker-compose.yml` files for experimenting and development.

The `ext` directory project that are imported as git submodules.

## Summary of the Projects

`benchmark-ip-firewall-updates`: An example on peformance benchmarks for increment integers atomically.

`chocolate-errors`: An example of passing custom errors in golang (using chocolates as references).

`clickhouse-multitable-bulk-ingest`: Bulk ingest example to clickhouse using `clickhouse-go`.

`clickhouse-struct-ingest-performance`:  Performance evaluation of `clickhouse-go`.

`ext/geomrpc`: An example of gRPC clients and servers, including both server-side and client-side streaming and gRPC metrics collection using Prometheus (including both connection stats and RPC stats).

`file-shreder`: An example to implement a `shred` function like the [shred](https://manpages.ubuntu.com/manpages/jammy/man1/shred.1.html) command line utility with some tests.

`guage-approximator`: A basic example of implement an average gauge metric over a given time interval. The example uses a circular ring buffer to store the last n-values. The example is simliar to Prometheus' *summary* metric.

`json-parser`: A basic JSON parser example that Unmarshals a JSON stream into different structs.

`kafka-trees`: A multi-topic example of sync/async producers (publishers) and a consumer group (subsribers) allowing horizontal scaling of kafka consumers. The example uses tree names as references.

`mockery-of-the-language`: An example to use mockery to generate golang interfances for uses in tests.

`struct-embedding`: A basic struct embedding example in Golang.

`sqlc-students`: A basic example of using sqlc to convert sql queries into golang structs.

`titan-prometheus`: A basic example of building a stats/metrics server for a running application using Prometheus.
