# Sarama Consumer Group

A basic example of a Kafka consumer group (and producers) using Shopify/sarama library with stats exposed to Prometheus.

```sh
# run containers
docker-compose up --build --detach

# Shutdown everything (and remove networks and local images). Networks are removed in this.
docker-compose down --volumes
# Use `docker-compose down --rmi all --volumes` with above to images as well
# Remove everything (and remove volumes). Networks are not removed here.
docker-compose rm --force --stop -v

```
## Kafka

Check the details about the Kafka docker image [here](https://github.com/wurstmeister/kafka-docker).


## Prometheus

Prometheus is available by default on port `9090`. Check the details about Prometheus docker [here](https://hub.docker.com/r/prom/prometheus).
