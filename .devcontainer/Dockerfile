# This dockerfile built for test gotorcontroller pakage.
FROM golang:bullseye
WORKDIR /workspaces/gotorcontroller/examples

RUN apt-get update
# Install torcontroller with Intel/AMD cpu:
RUN wget -P /usr/src/ https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_amd64.deb
RUN apt-get install -y /usr/src/torcontroller_1.0-1_amd64.deb
