use xcpc_board_mysql;
DROP TABLE IF EXISTS `codeforces`;
CREATE TABLE `codeforces` (
                        `Id` INT(10) NOT NULL AUTO_INCREMENT COMMENT 'cf主键ID',
                        `user_id` VARCHAR(40) COMMENT 'cf账号ID',
                        `contest_id` int(10) comment '比赛号',
                        `problem_index` varchar(5) comment '题目编号',
                        `problem_name` varchar(70) comment '题目名称',
                        `rating` int(10) comment  '难度',
                        PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='集体表格汇总';
alter table codeforces add unique key `name_problem` (`user_id`,`contest_id`,`problem_index`);