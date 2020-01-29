CREATE TABLE `daily` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '本日知识点',
  `author` varchar(30) NOT NULL DEFAULT '' COMMENT '作者',
  `day` varchar(30) NOT NULL DEFAULT '' COMMENT '日期，年月日，用横杠分割，比如2019-03-01',
  `content` text NOT NULL COMMENT 'json存储的内容，{[title,link,comment]}',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='每日知识点';