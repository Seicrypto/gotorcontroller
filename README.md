# gotorcontroller

now support: [torcontroller](https://github.com/Seicrypto/torcontroller) v1.0-1

With this package, you can easily control your backend application to use different IP addresses through Tor's service. This enables users to leverage multiple IPs, making it ideal for tasks such as web scraping and other applications that require a large number of IP addresses.

## QuickStart

![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)

Please make sure your docker image base on debian.
Such as bullseye, bookworm:

* golang:bullseye
* golang:bookworm
* so on...

Step1. Download and Install package

```dockerfile
# dockerfile
From golang:bullseye
WORKDIR /app

# Make sure there's wget in your docker image.
RUN apt-get update

# Intel / AMD cpu:
RUN wget https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_amd64.deb
RUN apt-get install -y /app/torcontroller_1.0-1_amd64.deb
# Be careful about your download and install path.

# ARM cpu:
# RUN wget https://github.com/Seicrypto/torcontroller/releases/download/v1.0-1/torcontroller_1.0-1_arm64.deb
# RUN apt-get install -y /app/torcontroller_1.0-1_arm64.deb
```

Step2. Set up your AUTHENTICATE password

Terminal (Recommend)

```bash
# Terminal
# This sample using 9050 as socketPort, 9051 as controlPort on machine.
docker run -it -p 9050:9050 -p 9051:9051 --name <your-container-name> <docker-image>

# Check torcontroller installed properly.
docker exec <your-container-name> torcontroller --version
# torcontroller version X.X-X

# On terminal to control docker.
docker exec -i <your-container-name> torcontroller --resetpassword
# torcontroller info ...
# Enter old TOR password:
# (torcontroller set as default password)
```

step3. get gotorcontroller lib

```bash
# Terminal
go get -u github.com/Seicrypto/gotorcontroller
# go ... gotorcontroller v1.0.0
```

## Reference

This go lib basic on [torcontroller](https://github.com/Seicrypto/torcontroller), if you're looking for devolep with different languages or more feature. Please read it.
