drop database ftp;
create database if not exists ftp;
use ftp;
CREATE TABLE if not exists  `user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(25) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `creator_id` bigint(20) NOT NULL DEFAULT '0',
  `user_type` tinyint(3) NOT NULL DEFAULT '0' COMMENT '1/admin',
  `available_flag` tinyint(3) NOT NULL DEFAULT '1',
  `email` varchar(25) NOT NULL DEFAULT '',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

CREATE TABLE if not exists `ip_list` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ip_type` tinyint(5) NOT NULL DEFAULT '0',
  `version` tinyint(5) NOT NULL DEFAULT '4',
  `creator_id` bigint(20) NOT NULL DEFAULT '0',
  `addr` varchar(25) NOT NULL,
  `available_flag` tinyint(3) NOT NULL DEFAULT '1',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `addr` (`addr`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- only for test. do not store password as plaintext
insert into `user`(id,user_name,password)value (1,'weisiqian','abc');
insert into `ip_list`(id,ip_type,version,addr)value (1,1,4,'127.0.0.1');