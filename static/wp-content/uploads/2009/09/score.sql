DROP TABLE IF EXISTS `score`;
CREATE TABLE `score` (
  `score_id` int(11) NOT NULL auto_increment,
  `student_name` varchar(32) character set utf8 default NULL,
  `score` int(10) unsigned default NULL,
  PRIMARY KEY  (`score_id`)
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=latin1;

INSERT INTO `score` VALUES (1,'Wallace',95),(2,'Gromit',97),(3,'Shaun',85),(4,'McGraw',92),(5,'Preston',92);

