create database blockchain;
use blockchain;
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `firm_id` varchar(64) NOT NULL COMMENT '企业id',
  `address` char(42) NOT NULL COMMENT '钱包地址',
  `priv_key` char(64) NOT NULL COMMENT '私钥',
  `passphrase` varchar(64) NOT NULL DEFAULT '' COMMENT '钱包密码',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_firm` (`firm_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `cmd_procedure`;
CREATE TABLE `cmd_procedure` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `transaction_id` int(11) NOT NULL,
  `start_nonce` bigint(20) DEFAULT NULL,
  `tx_hashes` JSON DEFAULT NULL,
  `state` enum('unprocess', 'processing', 'processed') NOT NULL DEFAULT 'unprocess',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


DROP TABLE IF EXISTS `transactions`;
CREATE TABLE `transactions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `deal_id` varchar(32) NOT NULL COMMENT '业务交易ID',
  `input` JSON,
  `output` JSON,
  `state` enum('payment', 'discount', 'splitToken', 'mintToken') NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;