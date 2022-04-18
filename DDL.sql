-- for sqlite
CREATE TABLE if not EXISTS ` users ` (
	` user_id ` integer,` user_name ` text,` birth_of_date ` text,` address ` text,` description ` text,` create_at ` datetime,
PRIMARY KEY ( ` user_id ` )
);

-- for mysql
CREATE TABLE IF NOT EXISTS `users` (
		`user_id` BIGINT ( 20 ) UNSIGNED NOT NULL AUTO_INCREMENT,
		`user_name` VARCHAR ( 50 ) NOT NULL COMMENT '用户名',
		`birth_of_date` VARCHAR ( 20 ) NOT NULL COMMENT '生日',
		`address` VARCHAR ( 100 ) NOT NULL COMMENT '地址',
		`description` VARCHAR ( 150 ) NOT NULL COMMENT '描述',
		`create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
		PRIMARY KEY ( `user_id` ) USING BTREE,
		UNIQUE KEY `key_user_name` ( `user_id` ) USING BTREE
	) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COMMENT = '用户信息表';