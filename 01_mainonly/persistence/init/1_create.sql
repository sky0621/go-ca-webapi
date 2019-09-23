CREATE TABLE IF NOT EXISTS `item` (
  `id` varchar(64) NOT NULL,
  `name` varchar(256) NOT NULL,
  `price` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;
