DROP TABLE IF EXISTS `sillyhat_mother`.`unknown_table`;
CREATE TABLE `sillyhat_mother`.`unknown_table`
(
  `id`                 varchar(32)  NOT NULL,
  `content`            text         NOT NULL,
  `table_type`         int                   DEFAULT NULL,
  `status`             int                   DEFAULT NULL,
  `created_time`       timestamp(3) NOT NULL DEFAULT current_timestamp(3),
  `last_modified_time` timestamp(3) NOT NULL DEFAULT current_timestamp(3) ON UPDATE current_timestamp(3),
  PRIMARY KEY (`id`),
  INDEX (`status`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;