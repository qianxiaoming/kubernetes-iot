FROM alpine:3.8
MAINTAINER titangrm-v2

ENV ZULU_ARCH=zulu8.36.0.1-ca-jdk8.0.202-linux_musl_x64

ENV LANG=en_US.UTF-8   \
    LANGUAGE=en_US:en  \
    LC_ALL=en_US.UTF-8 \
    JAVA_HOME=/usr/lib/jvm/${ZULU_ARCH} \
    PATH=/usr/lib/jvm/${ZULU_ARCH}/bin:$PATH

RUN INSTALL_DIR=/usr/lib/jvm && \
	BIN_DIR=/usr/bin && \
	MAN_DIR=/usr/share/man/man1 && \
	wget -q https://cdn.azul.com/zulu/bin/${ZULU_ARCH}.tar.gz && \
	mkdir -p ${INSTALL_DIR} && \
	tar -xf ./${ZULU_ARCH}.tar.gz -C ${INSTALL_DIR} && rm -f ${ZULU_ARCH}.tar.gz && \
	cd ${BIN_DIR} && find ${INSTALL_DIR}/${ZULU_ARCH}/bin -type f -perm -a=x -exec ln -s {} . \; && \
	mkdir -p ${MAN_DIR} && \
	cd ${MAN_DIR} && find ${INSTALL_DIR}/${ZULU_ARCH}/man/man1 -type f -name "*.1" -exec ln -s {} . \;

