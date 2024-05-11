FROM alpine:3.19.1 AS build
ARG VERSION=v4.43-9799

WORKDIR /root

RUN apk add --no-cache gcc make libc-dev readline-dev openssl-dev ncurses-dev zlib-dev gnu-libiconv wget && \
    wget -O vpnserver.tar.gz https://www.softether-download.com/files/softether/v4.43-9799-beta-2023.08.31-tree/Linux/SoftEther_VPN_Server/64bit_-_Intel_x64_or_AMD64/softether-vpnserver-v4.43-9799-beta-2023.08.31-linux-x64-64bit.tar.gz && \
    tar xzf vpnserver.tar.gz && \
    cd /root/${VERSION} && \
    ./configure && make && \
    mv /root/${VERSION}/bin/vpnserver / && \
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
