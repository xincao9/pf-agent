CREATE
DATABASE `pf_agent` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE
`pf_agent`;

DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `mobile`     char(14)  NOT NULL DEFAULT '',
    `expire`     timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `token`      char(128) NOT NULL DEFAULT '',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `real_name`  char(32)  NOT NULL DEFAULT '',
    `gender`     int       NOT NULL DEFAULT '0',
    `birthday`   timestamp NULL DEFAULT NULL,
    `address`    char(128)          DEFAULT NULL,
    `nickname`   char(32)           DEFAULT NULL,
    `level`      int       NOT NULL DEFAULT '1',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;