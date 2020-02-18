# kubernetes-iot
The software stack for developing lightweight IoT platform on kubernetes. It includes:
* Elasticsearch as storage service
* Mosquitto to accept JSON data via MQTT protocol
* Kafka used to pass data from Mosquitto brokers to Elasticsearch(based on confluentinc kafka)
* Kafka connector(based on confluentinc kafka)
* zookeeper
* jdk8(zulu8.36)

bootstrap is a programs implemented with go. It can generate all k8s yaml files which are used to deploy all services above. There is also a example file named 'values.yaml' as input of bootstrap.
