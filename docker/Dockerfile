FROM alpine:3.19.1 AS build
ARG VERSION=5.02.5184

WORKDIR /root

RUN apk add --no-cache cmake gcc g++ libsodium-dev make libc-dev readline-dev openssl-dev ncurses-dev zlib-dev gnu-libiconv git && \
    git clone --branch ${VERSION} --depth 1 https://github.com/SoftEtherVPN/SoftEtherVPN.git && \
    cd /root/SoftEtherVPN && \
    git submodule init && git submodule update && \
    ./configure && make -C build vpnserver&& \
    mkdir /vpnserver && mv /root/SoftEtherVPN/build/vpnserver /vpnserver/ && \
    mv /root/SoftEtherVPN/build/hamcore.se2 /vpnserver/
    echo 127.0.0.1 > /vpnserver/adminip.txt
COPY genconfig /vpnserver/
COPY start.sh /vpnserver/
RUN chmod 755 /vpnserver/genconfig /vpnserver/start.sh

# normal
FROM alpine:3.19.1
ENV USRNAME="alice"
ENV PASSWORD=""

COPY --from=build /vpnserver /usr/vpnserver

RUN apk add --no-cache readline

EXPOSE 8080

CMD "/usr/vpnserver/start.sh"
