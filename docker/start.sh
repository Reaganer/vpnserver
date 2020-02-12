#!/bin/sh
cd /usr/vpnserver
if [ ! -f "/usr/vpnserver/vpn_server.config" ]; then
    /usr/vpnserver/genconfig
fi
/usr/vpnserver/vpnserver execsvc