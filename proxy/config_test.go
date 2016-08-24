package main

const file = `
[upsql-proxy]
#标识
proxy-domain = <service_id>
proxy-name = <unit_id>
#默认数据库节点
default-datanode = default
#使用场景列表
#读写分离
need-rw-split = false
#分库分表
need-shard = false
#autocommit = true
#charset-names = gbk
#proxy客户端连接地址
proxy-address = <proxy_ip_addr>:<proxy_data_port>
#线程数
event-threads-count = <cpu_num>
#基路径,其他路径相关参数的基路径,相对路径都以该基路径计算
#必须用户手动配置,且配置为绝对路径
basedir = /usr/local/upproxy
#拓扑结构和虚拟数据库配置,相对路径
topology-config = /DBAASCNF/topology.json
connect-time-out = 60
read-time-out = 30
write-time-out = 60
max-connections = 5000
max-packet-num-in-shard = 100000
max-shard-num =8

need-daemon=true
need-keepalive=true
[log]
log-dir = /DBAASLOG
log-file = proxy.log 
log-file-size = 10240
log-max-days = 7
log-file-num = 100
log-level = warning
log-use = file
[adm-cli]
#adm-svr对外服务地址
adm-svr-address = <swm_ip_addr>:<swm_port>
adm-cli-event-threads-count = 2
adm-cli-address = <proxy_ip_addr>:<proxy_adm_port>

[supervise]
supervise-read-time-out = 5
supervise-write-time-out = 5
supervise-event-threads-count =2
supervise-address = <proxy_ip_addr>:<proxy_adm_port>
`
