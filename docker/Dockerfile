FROM alpine:3.11 AS build
ARG VERSION=v4.32-9731

WORKDIR /root

RUN apk add --no-cache gcc make libc-dev readline-dev openssl-dev ncurses-dev zlib-dev gnu-libiconv wget && \
    wget -O vpnserver.tar.gz https://github.com/SoftEtherVPN/SoftEtherVPN_Stable/releases/download/${VERSION}-beta/softether-src-${VERSION}-beta.tar.gz && \
    tar xzf vpnserver.tar.gz && \
    cd /root/${VERSION} && \
    ./configure && make && \
    mv /root/${VERSION}/bin/vpnserver / && \
    echo 127.0.0.1 > /vpnserver/adminip.txt
COPY genconfig /vpnserver/
COPY start.sh /vpnserver/
RUN chmod 755 /vpnserver/genconfig /vpnserver/start.sh

# normal
FROM alpine:3.11
ENV USRNAME="steve"
ENV PASSWORD="Zqpm06275#"

COPY --from=build /vpnserver /usr/vpnserver

RUN apk add --no-cache readline

EXPOSE 8080

CMD "/usr/vpnserver/start.sh"
