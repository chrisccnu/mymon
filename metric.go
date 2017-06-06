package main

import (
	"fmt"
	"os"
	"time"
	"strings"
)

const (
	TIME_OUT = 30

	ORIGIN   = "GAUGE"
	DELTA_PS = "COUNTER"
	DELTA    = ""
)

// COUNTER: Speed per second
// GAUGE: Original, DEFAULT
var DataType = map[string]string{
	"innodb_buffer_pool_reads":         DELTA_PS,
	"innodb_buffer_pool_read_requests": DELTA_PS,
	"innodb_compress_time":             DELTA_PS,
	"innodb_data_fsyncs":               DELTA_PS,
	"innodb_data_read":                 DELTA_PS,
	"innodb_data_reads":                DELTA_PS,
	"innodb_data_writes":               DELTA_PS,
	"innodb_data_written":              DELTA_PS,
	"innodb_last_checkpoint_at":        DELTA_PS,
	"innodb_log_flushed_up_to":         DELTA_PS,
	"innodb_log_sequence_number":       DELTA_PS,
	"innodb_mutex_os_waits":            DELTA_PS,
	"innodb_mutex_spin_rounds":         DELTA_PS,
	"innodb_mutex_spin_waits":          DELTA_PS,
	"innodb_pages_flushed_up_to":       DELTA_PS,
	"innodb_rows_deleted":              DELTA_PS,
	"innodb_rows_inserted":             DELTA_PS,
	"innodb_rows_locked":               DELTA_PS,
	"innodb_rows_modified":             DELTA_PS,
	"innodb_rows_read":                 DELTA_PS,
	"innodb_rows_updated":              DELTA_PS,
	"innodb_row_lock_time":             DELTA_PS,
	"innodb_row_lock_waits":            DELTA_PS,
	"innodb_uncompress_time":           DELTA_PS,

	"binlog_event_count": DELTA_PS,
	"binlog_number":      DELTA_PS,
	"slave_count":        DELTA_PS,

	"com_admin_commands":        DELTA_PS,
	"com_assign_to_keycache":    DELTA_PS,
	"com_alter_db":              DELTA_PS,
	"com_alter_db_upgrade":      DELTA_PS,
	"com_alter_event":           DELTA_PS,
	"com_alter_function":        DELTA_PS,
	"com_alter_procedure":       DELTA_PS,
	"com_alter_server":          DELTA_PS,
	"com_alter_table":           DELTA_PS,
	"com_alter_tablespace":      DELTA_PS,
	"com_analyze":               DELTA_PS,
	"com_begin":                 DELTA_PS,
	"com_binlog":                DELTA_PS,
	"com_call_procedure":        DELTA_PS,
	"com_change_db":             DELTA_PS,
	"com_change_master":         DELTA_PS,
	"com_check":                 DELTA_PS,
	"com_checksum":              DELTA_PS,
	"com_commit":                DELTA_PS,
	"com_create_db":             DELTA_PS,
	"com_create_event":          DELTA_PS,
	"com_create_function":       DELTA_PS,
	"com_create_index":          DELTA_PS,
	"com_create_procedure":      DELTA_PS,
	"com_create_server":         DELTA_PS,
	"com_create_table":          DELTA_PS,
	"com_create_trigger":        DELTA_PS,
	"com_create_udf":            DELTA_PS,
	"com_create_user":           DELTA_PS,
	"com_create_view":           DELTA_PS,
	"com_dealloc_sql":           DELTA_PS,
	"com_delete":                DELTA_PS,
	"com_delete_multi":          DELTA_PS,
	"com_do":                    DELTA_PS,
	"com_drop_db":               DELTA_PS,
	"com_drop_event":            DELTA_PS,
	"com_drop_function":         DELTA_PS,
	"com_drop_index":            DELTA_PS,
	"com_drop_procedure":        DELTA_PS,
	"com_drop_server":           DELTA_PS,
	"com_drop_table":            DELTA_PS,
	"com_drop_trigger":          DELTA_PS,
	"com_drop_user":             DELTA_PS,
	"com_drop_view":             DELTA_PS,
	"com_empty_query":           DELTA_PS,
	"com_execute_sql":           DELTA_PS,
	"com_flush":                 DELTA_PS,
	"com_grant":                 DELTA_PS,
	"com_ha_close":              DELTA_PS,
	"com_ha_open":               DELTA_PS,
	"com_ha_read":               DELTA_PS,
	"com_help":                  DELTA_PS,
	"com_insert":                DELTA_PS,
	"com_insert_select":         DELTA_PS,
	"com_install_plugin":        DELTA_PS,
	"com_kill":                  DELTA_PS,
	"com_load":                  DELTA_PS,
	"com_lock_tables":           DELTA_PS,
	"com_optimize":              DELTA_PS,
	"com_preload_keys":          DELTA_PS,
	"com_prepare_sql":           DELTA_PS,
	"com_purge":                 DELTA_PS,
	"com_purge_before_date":     DELTA_PS,
	"com_release_savepoint":     DELTA_PS,
	"com_rename_table":          DELTA_PS,
	"com_rename_user":           DELTA_PS,
	"com_repair":                DELTA_PS,
	"com_replace":               DELTA_PS,
	"com_replace_select":        DELTA_PS,
	"com_reset":                 DELTA_PS,
	"com_resignal":              DELTA_PS,
	"com_revoke":                DELTA_PS,
	"com_revoke_all":            DELTA_PS,
	"com_rollback":              DELTA_PS,
	"com_rollback_to_savepoint": DELTA_PS,
	"com_savepoint":             DELTA_PS,
	"com_select":                DELTA_PS,
	"com_set_option":            DELTA_PS,
	"com_signal":                DELTA_PS,
	"com_show_authors":          DELTA_PS,
	"com_show_binlog_events":    DELTA_PS,
	"com_show_binlogs":          DELTA_PS,
	"com_show_charsets":         DELTA_PS,
	"com_show_collations":       DELTA_PS,
	"com_show_contributors":     DELTA_PS,
	"com_show_create_db":        DELTA_PS,
	"com_show_create_event":     DELTA_PS,
	"com_show_create_func":      DELTA_PS,
	"com_show_create_proc":      DELTA_PS,
	"com_show_create_table":     DELTA_PS,
	"com_show_create_trigger":   DELTA_PS,
	"com_show_databases":        DELTA_PS,
	"com_show_engine_logs":      DELTA_PS,
	"com_show_engine_mutex":     DELTA_PS,
	"com_show_engine_status":    DELTA_PS,
	"com_show_events":           DELTA_PS,
	"com_show_errors":           DELTA_PS,
	"com_show_fields":           DELTA_PS,
	"com_show_function_status":  DELTA_PS,
	"com_show_grants":           DELTA_PS,
	"com_show_keys":             DELTA_PS,
	"com_show_master_status":    DELTA_PS,
	"com_show_open_tables":      DELTA_PS,
	"com_show_plugins":          DELTA_PS,
	"com_show_privileges":       DELTA_PS,
	"com_show_procedure_status": DELTA_PS,
	"com_show_processlist":      DELTA_PS,
	"com_show_profile":          DELTA_PS,
	"com_show_profiles":         DELTA_PS,
	"com_show_relaylog_events":  DELTA_PS,
	"com_show_slave_hosts":      DELTA_PS,
	"com_show_slave_status":     DELTA_PS,
	"com_show_status":           DELTA_PS,
	"com_show_storage_engines":  DELTA_PS,
	"com_show_table_status":     DELTA_PS,
	"com_show_tables":           DELTA_PS,
	"com_show_triggers":         DELTA_PS,
	"com_show_variables":        DELTA_PS,
	"com_show_warnings":         DELTA_PS,
	"com_slave_start":           DELTA_PS,
	"com_slave_stop":            DELTA_PS,
	"com_stmt_close":            DELTA_PS,
	"com_stmt_execute":          DELTA_PS,
	"com_stmt_fetch":            DELTA_PS,
	"com_stmt_prepare":          DELTA_PS,
	"com_stmt_reprepare":        DELTA_PS,
	"com_stmt_reset":            DELTA_PS,
	"com_stmt_send_long_data":   DELTA_PS,
	"com_truncate":              DELTA_PS,
	"com_uninstall_plugin":      DELTA_PS,
	"com_unlock_tables":         DELTA_PS,
	"com_update":                DELTA_PS,
	"com_update_multi":          DELTA_PS,
	"com_xa_commit":             DELTA_PS,
	"com_xa_end":                DELTA_PS,
	"com_xa_prepare":            DELTA_PS,
	"com_xa_recover":            DELTA_PS,
	"com_xa_rollback":           DELTA_PS,
	"com_xa_start":              DELTA_PS,

	"aborted_clients":            DELTA_PS,
	"aborted_connects":           DELTA_PS,
	"access_denied_errors":       DELTA_PS,
	"binlog_bytes_written":       DELTA_PS,
	"binlog_cache_disk_use":      DELTA_PS,
	"binlog_cache_use":           DELTA_PS,
	"binlog_stmt_cache_disk_use": DELTA_PS,
	"binlog_stmt_cache_use":      DELTA_PS,
	"bytes_received":             DELTA_PS,
	"bytes_sent":                 DELTA_PS,
	"connections":                DELTA_PS,
	"created_tmp_disk_tables":    DELTA_PS,
	"created_tmp_files":          DELTA_PS,
	"created_tmp_tables":         DELTA_PS,
	"handler_delete":             DELTA_PS,
	"handler_read_first":         DELTA_PS,
	"handler_read_key":           DELTA_PS,
	"handler_read_last":          DELTA_PS,
	"handler_read_next":          DELTA_PS,
	"handler_read_prev":          DELTA_PS,
	"handler_read_rnd":           DELTA_PS,
	"handler_read_rnd_next":      DELTA_PS,
	"handler_update":             DELTA_PS,
	"handler_write":              DELTA_PS,
	"opened_files":               DELTA_PS,
	"opened_tables":              DELTA_PS,
	"opened_table_definitions":   DELTA_PS,
	"qcache_hits":                DELTA_PS,
	"qcache_inserts":             DELTA_PS,
	"qcache_lowmem_prunes":       DELTA_PS,
	"qcache_not_cached":          DELTA_PS,
	"queries":                    DELTA_PS,
	"questions":                  DELTA_PS,
	"select_full_join":           DELTA_PS,
	"select_full_range_join":     DELTA_PS,
	"select_range_check":         DELTA_PS,
	"select_scan":                DELTA_PS,
	"slow_queries":               DELTA_PS,
	"sort_merge_passes":          DELTA_PS,
	"sort_range":                 DELTA_PS,
	"sort_rows":                  DELTA_PS,
	"sort_scan":                  DELTA_PS,
	"table_locks_immediate":      DELTA_PS,
	"table_locks_waited":         DELTA_PS,
	"threads_created":            DELTA_PS,
}

type MysqlIns struct {
	Host string
	Port int
	Tag  string
}

func dataType(key_ string) string {
	if v, ok := DataType[key_]; ok {
		return v
	}
	return ORIGIN
}

type MetaData struct {
	Metric      string      `json:"metric"`      //key
	Endpoint    string      `json:"endpoint"`    //hostname
	Value       interface{} `json:"value"`       // number or string
	CounterType string      `json:"counterType"` // GAUGE  原值   COUNTER 差值(ps)
	Tags        string      `json:"tags"`        // port=3306,k=v
	Timestamp   int64       `json:"timestamp"`
	Step        int64       `json:"step"`
}

func (m *MetaData) String() string {
	s := fmt.Sprintf("MetaData Metric:%s Endpoint:%s Value:%v CounterType:%s Tags:%s Timestamp:%d Step:%d",
		m.Metric, m.Endpoint, m.Value, m.CounterType, m.Tags, m.Timestamp, m.Step)
	return s
}

func NewMetric(name string) *MetaData {
	return &MetaData{
		Metric:      name,
		Endpoint:    hostname(),
		CounterType: dataType(strings.ToLower(name)),
		Tags:        fmt.Sprintf("port=%d", cfg.Port),
		Timestamp:   time.Now().Unix(),
		Step:        60,
	}
}

func hostname() string {
	host := cfg.Endpoint
	if host != "" {
		return host
	}
	host, err := os.Hostname()
	if err != nil {
		host = cfg.Host
	}
	return host
}

func (m *MetaData) SetValue(v interface{}) {
	m.Value = v
}
