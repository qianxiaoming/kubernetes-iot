FROM eclipse-mosquitto:1.5

ENV IOT_PROTOCOL=MQTT \
    IOT_PROTOCOL_PORT=1883 \
    IOT_PROTOCOL_SSL_PORT=8883 \
    MOSQUITTO_VERSION=1.5.5

RUN set -x && \
    apk --no-cache add --virtual bash

COPY docker-entrypoint.sh /
ENTRYPOINT ["/docker-entrypoint.sh"]
