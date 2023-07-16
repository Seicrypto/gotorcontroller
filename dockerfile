# This dockerfile built for test gotorcontroller pakage.
FROM golang:bullseye
WORKDIR /app

RUN apt-get update
# ARM cpu:
RUN wget -P /usr/src/ https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_arm64.deb
RUN apt-get install -y /usr/src/torcontroller_1.0-1_arm64.deb

# Clone sources from github,
# which means no need to rebuild container again.
RUN git clone -b dev https://github.com/Seicrypto/gocontroller.git /app/
# Couldn't just command this,
# maybe build a bash script if needed.
# CMD [ "git fetch" ]