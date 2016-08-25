package parser

import (
	"bytes"
	"testing"
)

const content = `
###
[mysqld]
max_connections=5000
slow_query_log=0
innodb_flush_method=O_DIRECT
relay_log_purge=on
innodb_buffer_pool_instances=8
innodb_log_files_in_group=7
innodb_log_group_home_dir=/DBAASLOG/RED
innodb_max_dirty_pages_pct=30
innodb_doublewrite=1
net_write_timeout=60
explicit_defaults_for_timestamp=true
innodb_purge_threads=8
loose_rpl_semi_sync_slave_enabled=1
rpl_semi_sync_master_enabled=on
max_user_connections=0
binlog_cache_size=1M
open_files_limit=10240
innodb_buffer_pool_size=805306368
innodb_read_io_threads=4
innodb_stats_sample_pages=1
datadir=/DBAASDAT
innodb_file_per_table=1
log_error=upsql.err
binlog_row_image=minimal
max_binlog_size=1G
long_query_time=1
innodb_rollback_on_timeout=on
log_slave_updates=on
relay_log_info_repository=TABLE
skip_external_locking=ON
sort_buffer_size=2M
tmpdir=/DBAASDAT
innodb_write_io_threads=4
plugin_load=rpl_semi_sync_master=semisync_master.so;rpl_semi_sync_slave=semisync_slave.so;upsql_auth=upsql_auth.so
max_relay_log_size=1G
interactive_timeout=31536000
net_read_timeout=30
master_info_repository=TABLE
sync_binlog=1
expire_logs_days=0
innodb_flush_log_at_trx_commit=1
plugin_dir=/usr/local/mysql/lib/plugin
rpl_semi_sync_slave_enabled=on
innodb_thread_concurrency=16
# bind_address= wrong_bind_address
BIND_Address=  192.168.20. 102
    BIND_Address=     192.168.20.102   
BIND_Add_ress=  192.168.x2 0. 102
b ind_address   =192.168.20.102
join_buffer_size=128K
user=upsql
key_buffer_size=160M
slow_query_log_file=/DBAASLOG/slow-query.log
log_queries_not_using_indexes=0
innodb_open_files=1024
innodb_stats_on_metadata=OFF
innodb_support_xa=1
slave_parallel_workers=5
auto_increment_offset=1
rpl_semi_sync_master_trace_level=32
max_connect_errors=50000
innodb_log_file_size=128M
binlog_checksum=CRC32
master_verify_checksum=ON
relay_log_recovery=on
socket=/DBAASDAT/upsql.sock
server_id=30004
max_allowed_packet=16M
binlog_format=row
innodb_lock_wait_timeout=60
slave_sql_verify_checksum=ON
auto_increment_increment=1
gtid_mode=on
character_set_server=utf8
log_bin=/DBAASLOG/BIN/ec6935a4_aaa_01-binlog
innodb_log_buffer_size=128M
innodb_checksums=1
innodb_io_capacity=500
innodb_purge_batch_size=500
innodb_stats_persistent_sample_pages=10
enforce_gtid_consistency=on
rpl_semi_sync_master_timeout=10000
slave_net_timeout=10
relay_log=/DBAASLOG/REL/ec6935a4_aaa_01-relay
lower_case_table_names=1
connect_timeout=60
wait_timeout=31536000
innodb_data_file_path=ibdata1:12M:autoextend
replicate_ignore_db=dbaas_check
rpl_semi_sync_master_wait_no_slave=on
rpl_semi_sync_slave_trace_level=32
#port=  wrong_port
 port =  30004#
adminport=  wrong_port
     	_port =  30004
 port* =  30004
log_bin_trust_function_creators=ON
optimizer_switch='mrr=on,mrr_cost_based=off'
loose_rpl_semi_sync_master_enabled=1

[mysqldump]
max_allowed_packet=16M

[myisamchk]
key_buffer_size=20M
sort_buffer_size=2M
`

func TestGetSectionBody(t *testing.T) {
	source := []byte(content)

	body := GetSectionBody(source, []byte("  Mysqld     "))
	if len(body) == 0 {
		t.Error("Unexpected,got nil")
	} else {
		t.Logf("%s", body)
	}

	body = GetSectionBody(source, []byte("   myisamchK   "))
	if len(body) == 0 {
		t.Error("Unexpected,got nil")
	} else {
		t.Logf("%s", body)
	}

	body = GetSectionBody(source, []byte("   MYSQLdump  "))
	if len(body) == 0 {
		t.Error("Unexpected,got nil")
	} else {
		t.Logf("%s", body)
	}

	body = GetSectionBody(source, []byte("   "))
	if !bytes.Equal(body, source) {
		t.Errorf("Unexpected,got:'%s'", body)
	}

	body = GetSectionBody(source, []byte("   afaufoaofsqlDump  "))
	if len(body) > 0 {
		t.Errorf("Unexpected,want nil but got '%s'", body)
	}

	body = GetSectionBody(source, []byte("mys qld   "))
	if len(body) > 0 {
		t.Errorf("Unexpected,want nil but got '%s'", body)
	}
}

func TestGetString(t *testing.T) {
	source := []byte(content)

	value := GetString(source, []byte(" key_buffer_size   "), func(s string) bool {
		if s == "20M" {
			return true
		}
		return false
	})
	if value == "" {
		t.Error("Unexpected")
	}

	body := GetSectionBody(source, []byte("   myisamchK   "))
	if len(body) == 0 {
		t.Error("Unexpected,got nil")
	} else {
		t.Logf("%s", body)
	}

	value = GetString(body, []byte(" soRt_bufFer_size   "), func(s string) bool {
		if s == "2M" {
			return true
		}
		return false
	})
	if value == "" {
		t.Error("Unexpected")
	}

	value = GetString(body, []byte(" soRt_bufFer_size   "), func(s string) bool {
		return false
	})
	if value != "" {
		t.Error("Unexpected")
	}

	value = GetString(body, []byte(" soRt_bufFer_size   "), func(s string) bool {
		return true
	})
	if value == "" {
		t.Error("Unexpected")
	}
}
