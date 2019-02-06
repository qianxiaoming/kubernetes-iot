FROM confluentinc/cp-kafka-connect:5.1.0
MAINTAINER titangrm-v2

ADD confluentinc-kafka-connect-mqtt /usr/share/java/kafka-connect-mqtt

COPY start-connect.sh /etc/confluent/docker/

CMD ["/etc/confluent/docker/start-connect.sh"]
