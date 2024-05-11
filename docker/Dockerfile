FROM alpine:3.19.1 AS build
ARG VERSION=v4.43-9799-beta

WORKDIR /root

RUN apk add --no-cache gcc make libc-dev readline-dev openssl-dev ncurses-dev zlib-dev gnu-libiconv git && \
    git clone --branch ${VERSION} --depth 1 https://github.com/SoftEtherVPN/SoftEtherVPN_Stable.git && \
    cd /root/SoftEtherVPN_Stable && \
    ./configure && make && \
    mv /root/SoftEtherVPN_Stable/bin/vpnserver / && \
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
