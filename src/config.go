package main

const configTemplate = `
# Software Configuration File
# ---------------------------
# 
# You may edit this file when the VPN Server / Client / Bridge program is not running.
# 
# In prior to edit this file manually by your text editor,
# shutdown the VPN Server / Client / Bridge background service.
# Otherwise, all changes will be lost.
# 
declare root
{
	uint ConfigRevision 0
	bool IPsecMessageDisplayed true
	string Region US
	bool VgsMessageDisplayed true

	declare DDnsClient
	{
		bool Disabled true
		byte Key {{.DDnsClientKey}}
		string LocalHostname alpine
		string ProxyHostName $
		uint ProxyPort 0
		uint ProxyType 0
		string ProxyUsername $
	}
	declare IPsec
	{
		bool EtherIP_IPsec false
		string IPsec_Secret vpn
		string L2TP_DefaultHub DEFAULT
		bool L2TP_IPsec false
		bool L2TP_Raw false

		declare EtherIP_IDSettingsList
		{
		}
	}
	declare ListenerList
	{
		declare Listener0
		{
			bool DisableDos false
			bool Enabled true
			uint Port 8080
		}
	}
	declare LocalBridgeList
	{
		bool DoNotDisableOffloading false
	}
	declare ServerConfiguration
	{
		bool AcceptOnlyTls true
		uint64 AutoDeleteCheckDiskFreeSpaceMin 104857600
		uint AutoDeleteCheckIntervalSecs 300
		uint AutoSaveConfigSpan 300
		bool BackupConfigOnlyWhenModified true
		string CipherName AES128-SHA
		uint CurrentBuild 9731
		bool DisableCoreDumpOnUnix false
		bool DisableDeadLockCheck false
		bool DisableDosProction false
		bool DisableGetHostNameWhenAcceptTcp false
		bool DisableIntelAesAcceleration false
		bool DisableIPsecAggressiveMode false
		bool DisableIPv6Listener true
		bool DisableJsonRpcWebApi false
		bool DisableNatTraversal false
		bool DisableOpenVPNServer true
		bool DisableSessionReconnect false
		bool DisableSSTPServer false
		bool DontBackupConfig false
		bool EnableVpnAzure false
		bool EnableVpnOverDns false
		bool EnableVpnOverIcmp false
		byte HashedPassword {{.ServerHashedPassword}}
		string KeepConnectHost keepalive.softether.org
		uint KeepConnectInterval 50
		uint KeepConnectPort 80
		uint KeepConnectProtocol 1
		uint64 LoggerMaxLogSize 2
		uint MaxConcurrentDnsClientThreads 512
		uint MaxConnectionsPerIP 256
		uint MaxUnestablishedConnections 1000
		bool NoHighPriorityProcess false
		bool NoLinuxArpFilter false
		bool NoSendSignature false
		string OpenVPNDefaultClientOption dev-type$20tun,link-mtu$201500,tun-mtu$201500,cipher$20AES-128-CBC,auth$20SHA1,keysize$20128,key-method$202,tls-client
		string OpenVPN_UdpPortList 1194
		bool SaveDebugLog false
		uint ServerLogSwitchType 4
		uint ServerType 0
		bool StrictSyslogDatetimeFormat false
		bool Tls_Disable1_0 false
		bool Tls_Disable1_1 false
		bool Tls_Disable1_2 false
		bool UseKeepConnect true
		bool UseWebTimePage false
		bool UseWebUI false

		declare GlobalParams
		{
			uint FIFO_BUDGET 10240000
			uint HUB_ARP_SEND_INTERVAL 5000
			uint IP_TABLE_EXPIRE_TIME 60000
			uint IP_TABLE_EXPIRE_TIME_DHCP 300000
			uint MAC_TABLE_EXPIRE_TIME 600000
			uint MAX_BUFFERING_PACKET_SIZE 2560000
			uint MAX_HUB_LINKS 1024
			uint MAX_IP_TABLES 65536
			uint MAX_MAC_TABLES 65536
			uint MAX_SEND_SOCKET_QUEUE_NUM 128
			uint MAX_SEND_SOCKET_QUEUE_SIZE 2560000
			uint MAX_STORED_QUEUE_NUM 1024
			uint MEM_FIFO_REALLOC_MEM_SIZE 655360
			uint MIN_SEND_SOCKET_QUEUE_SIZE 320000
			uint QUEUE_BUDGET 2048
			uint SELECT_TIME 256
			uint SELECT_TIME_FOR_NAT 30
			uint STORM_CHECK_SPAN 500
			uint STORM_DISCARD_VALUE_END 1024
			uint STORM_DISCARD_VALUE_START 3
		}
		declare ServerTraffic
		{
			declare RecvTraffic
			{
				uint64 BroadcastBytes 0
				uint64 BroadcastCount 0
				uint64 UnicastBytes 0
				uint64 UnicastCount 0
			}
			declare SendTraffic
			{
				uint64 BroadcastBytes 0
				uint64 BroadcastCount 0
				uint64 UnicastBytes 0
				uint64 UnicastCount 0
			}
		}
		declare SyslogSettings
		{
			string HostName $
			uint Port 0
			uint SaveType 0
		}
	}
	declare VirtualHUB
	{
		declare DEFAULT
		{
			uint64 CreatedTime {{.CreatedTime}}
			byte HashedPassword {{.HubHashedPassword}}
			uint64 LastCommTime {{.CreatedTime}}
			uint64 LastLoginTime {{.CreatedTime}}
			uint NumLogin 0
			bool Online true
			bool RadiusConvertAllMsChapv2AuthRequestToEap false
			string RadiusRealm $
			uint RadiusRetryInterval 0
			uint RadiusServerPort 1812
			string RadiusSuffixFilter $
			bool RadiusUsePeapInsteadOfEap false
			byte SecurePassword {{.HubSecurePassword}}
			uint Type 0

			declare AccessList
			{
			}
			declare AdminOption
			{
				uint allow_hub_admin_change_option 0
				uint deny_bridge 0
				uint deny_change_user_password 0
				uint deny_empty_password 1
				uint deny_hub_admin_change_ext_option 0
				uint deny_qos 0
				uint deny_routing 0
				uint max_accesslists 0
				uint max_bitrates_download 0
				uint max_bitrates_upload 0
				uint max_groups 0
				uint max_multilogins_per_user 0
				uint max_sessions 0
				uint max_sessions_bridge 0
				uint max_sessions_client 0
				uint max_sessions_client_bridge_apply 0
				uint max_users 0
				uint no_access_list_include_file 0
				uint no_cascade 0
				uint no_change_access_control_list 0
				uint no_change_access_list 0
				uint no_change_admin_password 0
				uint no_change_cert_list 0
				uint no_change_crl_list 0
				uint no_change_groups 0
				uint no_change_log_config 0
				uint no_change_log_switch_type 0
				uint no_change_msg 0
				uint no_change_users 0
				uint no_delay_jitter_packet_loss 0
				uint no_delete_iptable 0
				uint no_delete_mactable 0
				uint no_disconnect_session 0
				uint no_enum_session 0
				uint no_offline 0
				uint no_online 0
				uint no_query_session 0
				uint no_read_log_file 0
				uint no_securenat 0
				uint no_securenat_enabledhcp 0
				uint no_securenat_enablenat 0
			}
			declare CascadeList
			{
			}
			declare LogSetting
			{
				uint PacketLogSwitchType 4
				uint PACKET_LOG_ARP 0
				uint PACKET_LOG_DHCP 1
				uint PACKET_LOG_ETHERNET 0
				uint PACKET_LOG_ICMP 0
				uint PACKET_LOG_IP 0
				uint PACKET_LOG_TCP 0
				uint PACKET_LOG_TCP_CONN 1
				uint PACKET_LOG_UDP 0
				bool SavePacketLog false
				bool SaveSecurityLog false
				uint SecurityLogSwitchType 4
			}
			declare Message
			{
			}
			declare Option
			{
				uint AccessListIncludeFileCacheLifetime 30
				uint AdjustTcpMssValue 0
				bool ApplyIPv4AccessListOnArpPacket false
				bool AssignVLanIdByRadiusAttribute false
				bool BroadcastLimiterStrictMode false
				uint BroadcastStormDetectionThreshold 0
				uint ClientMinimumRequiredBuild 0
				bool DenyAllRadiusLoginWithNoVlanAssign false
				uint DetectDormantSessionInterval 0
				bool DisableAdjustTcpMss false
				bool DisableCheckMacOnLocalBridge false
				bool DisableCorrectIpOffloadChecksum false
				bool DisableHttpParsing false
				bool DisableIPParsing false
				bool DisableIpRawModeSecureNAT false
				bool DisableKernelModeSecureNAT false
				bool DisableUdpAcceleration false
				bool DisableUdpFilterForLocalBridgeNic false
				bool DisableUserModeSecureNAT false
				bool DoNotSaveHeavySecurityLogs false
				bool DropArpInPrivacyFilterMode true
				bool DropBroadcastsInPrivacyFilterMode true
				bool FilterBPDU false
				bool FilterIPv4 false
				bool FilterIPv6 false
				bool FilterNonIP false
				bool FilterOSPF false
				bool FilterPPPoE false
				uint FloodingSendQueueBufferQuota 33554432
				bool ManageOnlyLocalUnicastIPv6 true
				bool ManageOnlyPrivateIP true
				uint MaxLoggedPacketsPerMinute 0
				uint MaxSession 0
				bool NoArpPolling false
				bool NoDhcpPacketLogOutsideHub true
				bool NoEnum false
				bool NoIpTable false
				bool NoIPv4PacketLog false
				bool NoIPv6AddrPolling false
				bool NoIPv6DefaultRouterInRAWhenIPv6 true
				bool NoIPv6PacketLog false
				bool NoLookBPDUBridgeId false
				bool NoMacAddressLog true
				bool NoManageVlanId false
				bool NoPhysicalIPOnPacketLog false
				bool NoSpinLockForPacketDelay false
				bool RemoveDefGwOnDhcpForLocalhost true
				uint RequiredClientId 0
				uint SecureNAT_MaxDnsSessionsPerIp 0
				uint SecureNAT_MaxIcmpSessionsPerIp 0
				uint SecureNAT_MaxTcpSessionsPerIp 0
				uint SecureNAT_MaxTcpSynSentPerIp 0
				uint SecureNAT_MaxUdpSessionsPerIp 0
				bool SecureNAT_RandomizeAssignIp false
				bool SuppressClientUpdateNotification false
				bool UseHubNameAsDhcpUserClassOption false
				bool UseHubNameAsRadiusNasId false
				string VlanTypeId 0x8100
				bool YieldAfterStorePacket false
			}
			declare SecureNAT
			{
				bool Disabled false
				bool SaveLog false

				declare VirtualDhcpServer
				{
					string DhcpDnsServerAddress 8.8.8.8
					string DhcpDnsServerAddress2 8.8.4.4
					string DhcpDomainName $
					bool DhcpEnabled true
					uint DhcpExpireTimeSpan 7200
					string DhcpGatewayAddress {{.Subnet}}.1
					string DhcpLeaseIPEnd {{.Subnet}}.20
					string DhcpLeaseIPStart {{.Subnet}}.10
					string DhcpPushRoutes $
					string DhcpSubnetMask 255.255.255.0
				}
				declare VirtualHost
				{
					string VirtualHostIp {{.Subnet}}.1
					string VirtualHostIpSubnetMask 255.255.255.0
					string VirtualHostMacAddress {{.VirtualHostMacAddress}}
				}
				declare VirtualRouter
				{
					bool NatEnabled true
					uint NatMtu 1500
					uint NatTcpTimeout 1800
					uint NatUdpTimeout 60
				}
			}
			declare SecurityAccountDatabase
			{
				declare CertList
				{
				}
				declare CrlList
				{
				}
				declare GroupList
				{
				}
				declare IPAccessControlList
				{
				}
				declare UserList
				{
					declare {{.UserName}}
					{
						byte AuthNtLmSecureHash {{.UserAuthNtLmSecureHash}}
						byte AuthPassword {{.UserAuthPassword}}
						uint AuthType 1
						uint64 CreatedTime {{.CreatedTime}}
						uint64 ExpireTime 0
						uint64 LastLoginTime 0
						string Note $
						uint NumLogin 0
						string RealName $
						uint64 UpdatedTime {{.CreatedTime}}

						declare Traffic
						{
							declare RecvTraffic
							{
								uint64 BroadcastBytes 0
								uint64 BroadcastCount 0
								uint64 UnicastBytes 0
								uint64 UnicastCount 0
							}
							declare SendTraffic
							{
								uint64 BroadcastBytes 0
								uint64 BroadcastCount 0
								uint64 UnicastBytes 0
								uint64 UnicastCount 0
							}
						}
					}
				}
			}
			declare Traffic
			{
				declare RecvTraffic
				{
					uint64 BroadcastBytes 0
					uint64 BroadcastCount 0
					uint64 UnicastBytes 0
					uint64 UnicastCount 0
				}
				declare SendTraffic
				{
					uint64 BroadcastBytes 0
					uint64 BroadcastCount 0
					uint64 UnicastBytes 0
					uint64 UnicastCount 0
				}
			}
		}
	}
	declare VirtualLayer3SwitchList
	{
	}
}
`
