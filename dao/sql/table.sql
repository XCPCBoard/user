CREATE TABLE `lg` (
                      uid varchar(20) NOT NULL,
                      problem_number int(10) DEFAULT NULL,
                      ranting	 int(10) DEFAULT NULL,
                      simple_problem_number int(10) DEFAULT NULL,
                      base_problem_number int(10)	 DEFAULT NULL,
                      elevated_problem_number int(10) DEFAULT NULL,
                      hard_problem_number int(10)  DEFAULT NULL,
                      unKnow_problem_number int(10) DEFAULT NULL,
                      PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
