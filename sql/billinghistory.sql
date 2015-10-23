CREATE TABLE `billinghistory` (
	`order_id` VARCHAR(128) NOT NULL,
	`platform_type` SMALLINT(4) NOT NULL,
	`account_id` VARCHAR(80) NOT NULL DEFAULT '',
	`extra_order_id` VARCHAR(128) NOT NULL DEFAULT '',
	`uid` INT(32) NOT NULL,
	`store_item_id` VARCHAR(32) NOT NULL,
	`cost` INT(10) NOT NULL DEFAULT '0',
	`status` SMALLINT(4) NOT NULL DEFAULT '0',
	`insert_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY (`order_id`, `platform_type`)
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB;
