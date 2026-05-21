package version

import (
	"runtime"

	"k8soperation/cmd/migrate/migration"

	"gorm.io/gorm"
)

func init() {
	_, fileName, _, _ := runtime.Caller(0)
	migration.RegisterNamed(migration.VersionFromFilename(fileName), "migrate", func(tx *gorm.DB) error {
		// ============================================================
		// 在这里编写本次数据库变更逻辑。
		//
		// 注意事项：
		// 1. 当前函数已经运行在事务中，不需要再手动开启 tx.Transaction。
		// 2. 返回 error 时，本次迁移会回滚，并且不会写入 sys_migration。
		// 3. 执行成功后，框架会自动写入 sys_migration(version, apply_time)。
		// 4. 下方示例代码仅用于参考，使用时请删除不需要的示例。
		// ============================================================

		statements := []string{
			`CREATE TABLE IF NOT EXISTS k8s_cluster (
				id int unsigned NOT NULL AUTO_INCREMENT,
				cluster_name varchar(128) NOT NULL DEFAULT '' COMMENT '集群名',
				cluster_version varchar(64) NOT NULL DEFAULT '' COMMENT '集群版本',
				kube_config longtext COMMENT 'KubeConfig文本',
				status int NOT NULL DEFAULT 0 COMMENT '集群状态',
				created_at int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
				modified_at int unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
				deleted_at int unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
				is_del tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否删除,1表示删除,0表示未删除',
				PRIMARY KEY (id),
				UNIQUE KEY uniq_cluster_name (cluster_name),
				KEY idx_is_del (is_del)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='K8s集群信息';`,
			`CREATE TABLE IF NOT EXISTS event_item (
				namespace varchar(128) NOT NULL DEFAULT '' COMMENT '命名空间',
				kind varchar(64) NOT NULL DEFAULT '' COMMENT '资源类型',
				name varchar(255) NOT NULL DEFAULT '' COMMENT '资源名称',
				type varchar(32) NOT NULL DEFAULT '' COMMENT '事件类型',
				reason varchar(255) NOT NULL DEFAULT '' COMMENT '事件原因',
				message text COMMENT '事件详情',
				count int NOT NULL DEFAULT 0 COMMENT '事件次数',
				event_time datetime(3) NULL COMMENT '事件发生时间',
				source_component varchar(128) NOT NULL DEFAULT '' COMMENT '事件来源组件',
				source_instance varchar(128) NOT NULL DEFAULT '' COMMENT '事件来源实例',
				KEY idx_namespace_kind_name (namespace, kind, name),
				KEY idx_event_time (event_time)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Kubernetes事件';`,
			`CREATE TABLE IF NOT EXISTS login_session_info (
				username varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
				token varchar(512) NOT NULL DEFAULT '' COMMENT '登录Token',
				login_time datetime(3) NULL COMMENT '登录时间',
				KEY idx_username (username),
				KEY idx_token (token)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='登录会话信息';`,
			`CREATE TABLE IF NOT EXISTS user (
				id int unsigned NOT NULL AUTO_INCREMENT,
				username varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
				password varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
				created_at int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
				modified_at int unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
				deleted_at int unsigned NOT NULL DEFAULT 0 COMMENT '删除时间',
				is_del tinyint unsigned NOT NULL DEFAULT 0 COMMENT '是否删除,1表示删除,0表示未删除',
				PRIMARY KEY (id),
				UNIQUE KEY uniq_username (username),
				KEY idx_is_del (is_del)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户';`,
			`CREATE TABLE IF NOT EXISTS sys_migration (
				version varchar(64) NOT NULL COMMENT '迁移版本',
				apply_time int unsigned NOT NULL DEFAULT 0 COMMENT '执行时间',
				PRIMARY KEY (version)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='数据库迁移记录';`,
			`CREATE TABLE IF NOT EXISTS node_metric_item (
				name varchar(255) NOT NULL DEFAULT '' COMMENT '节点名称',
				timestamp datetime(3) NULL COMMENT '采集时间',
				window_seconds bigint NOT NULL DEFAULT 0 COMMENT '时间窗口大小',
				cpu_usage_milli bigint NOT NULL DEFAULT 0 COMMENT 'CPU使用量',
				mem_usage_bytes bigint NOT NULL DEFAULT 0 COMMENT '内存使用量',
				cpu_alloc_milli bigint NOT NULL DEFAULT 0 COMMENT 'CPU已分配量',
				mem_alloc_bytes bigint NOT NULL DEFAULT 0 COMMENT '内存已分配量',
				cpu_cap_milli bigint NOT NULL DEFAULT 0 COMMENT 'CPU容量',
				mem_cap_bytes bigint NOT NULL DEFAULT 0 COMMENT '内存容量',
				cpu_usage_percent double NOT NULL DEFAULT 0 COMMENT 'CPU使用率',
				mem_usage_percent double NOT NULL DEFAULT 0 COMMENT '内存使用率',
				KEY idx_name_timestamp (name, timestamp)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='节点资源指标';`,
		}

		for _, statement := range statements {
			if err := tx.Exec(statement).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
