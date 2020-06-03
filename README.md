# vpnserver
Docker image for SoftEther VPN server. Small and ready to use.

This docker image is based on alpine base image. Total size is very small (about 14MB).

## Features

- Can be configured as a multi-hop vpn system.
- Can add more than one vpn user account.


## Contents

* `src` folder contains `genconfig` tool's source code.
* `docker` folder contains `Dockerfile` and it's dependents. 

## Config & Run

When first running, `genconfig` will config the vpnserver automatically.
For security reason, it will randomly generate server and virtual hub admin password.

The vpnserver is listening on port **8080** inside container.
You only need to set vpn user account and/or next hop vpn account (depends on whether you need multi-hop vpn system) in environment variable.

### **Single node (Standard) mode**

```shell
docker run -d -e UserList='user1:pass1;user2:pass2...' -p xxxx:8080 reaganer/vpnserver:latest
```

If `UserList` is not set or empty, `genconfig` will create an account with randomly generated username and password and print it on the `STDOUT`.

### **Relay node (Mutli-Hop vpn) mode**

The last hop node needs to be configured as **Single node mode**. Other nodes need to be configured as relay node mode.

You can use an env-file to configure all the environment variables at once.

```shell
$ cat node.env
RelayNode
NextNodeHost=xxx.xxx.xxx.xxx
NextNodePort=xxxxx
NextNodeUserName=username
NextNodePassword=password
UserList=user1:pass1;user2:pass2...

$ docker run -d --env-file ./node.env -p xxxx:8080 reaganer/vpnserver:latest
```

