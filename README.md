# vpnserver
Docker image for SoftEther VPN server. Small and ready to use.

This docker image is based on alpine base image. Total size is very small (about 14MB).

## Contents

* `src` folder contains `genconfig` tool's source code.
* `docker` folder contains `Dockerfile` and it's dependents. 

## Config & Run

When first running, `genconfig` will config the vpnserver automatically.
For security reason, it will randomly generate server and virtual hub admin password.

The vpnserver is listening on port **8080** in container.
You only need to set vpn user name and password in environment variable `USRNAME` and `PASSWORD`.

```shell
docker run -d -e USRNAME='xxx' -e PASSWORD='xxx' -p xxxx:8080 vpnserver
```

If `USRNAME` is not set or empty, "**alice**" will be used. If `PASSWORD` is not set or empty, `genconfig` will randomly generate a password for you and print it on the `STDOUT`.

