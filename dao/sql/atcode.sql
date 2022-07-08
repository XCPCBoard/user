CREATE DATABASE IF NOT EXISTS xcpc_board_mysql;
use xcpc_board_mysql;
CREATE TABLE IF NOT EXISTS `atcoder`
(
      `submission_id`    VARCHAR(64)    NOT NULL  DEFAULT '',
      `user_id`          VARCHAR(64)    NOT NULL  DEFAULT '',
      `contest_id`       VARCHAR(64)    NOT NULL  DEFAULT '',
      `problem_id`       VARCHAR(64)    NOT NULL  DEFAULT '',
      `score`            INT            NOT NULL  default  0,
      primary key (`submission_id`),
      index (`user_id`)
) ENGINE = innodb DEFAULT CHARSET = utf8;
