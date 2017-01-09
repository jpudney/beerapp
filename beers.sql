# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.17)
# Database: beers
# Generation Time: 2017-01-09 19:18:51 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table beers
# ------------------------------------------------------------

DROP TABLE IF EXISTS `beers`;

CREATE TABLE `beers` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(150) NOT NULL DEFAULT '',
  `brewery` varchar(150) NOT NULL,
  `abv` float(4,1) NOT NULL,
  `short_description` text NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `beers` WRITE;
/*!40000 ALTER TABLE `beers` DISABLE KEYS */;

INSERT INTO `beers` (`id`, `name`, `brewery`, `abv`, `short_description`, `created`)
VALUES
	(1,'Best Bitter','T&R Theakston Ltd',3.8,'Theakston Best Bitter is the leading session ale within the Theakston portfolio and has been for time immemorial. It is quite possible that when Robert Theakston founded the brewery in 1827 the range of ales would have been limited to just two or three of which almost certainly, one would have been a bitter beer. Consequently it would be reasonable to argue that Theakston Best Bitter is one of the longest established session ales in Yorkshire.','2017-01-09 18:55:43'),
	(2,'London Pride','Fuller Smith & Turner PLC',4.1,'Brewed under the watchful eye of our Griffin since the 1950s, London Pride is unmistakably London\'s beer. With its well-rounded flavour and rich history, everything about this authentic, characterful beer binds it to our capital city and the people who love it.','2017-01-09 18:57:37');

/*!40000 ALTER TABLE `beers` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table reviews
# ------------------------------------------------------------

DROP TABLE IF EXISTS `reviews`;

CREATE TABLE `reviews` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `beer_id` int(11) NOT NULL,
  `first_name` varchar(150) NOT NULL DEFAULT '',
  `last_name` varchar(150) NOT NULL DEFAULT '',
  `score` int(2) NOT NULL,
  `text` text NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `reviews` WRITE;
/*!40000 ALTER TABLE `reviews` DISABLE KEYS */;

INSERT INTO `reviews` (`id`, `beer_id`, `first_name`, `last_name`, `score`, `text`, `created`)
VALUES
	(1,2,'Bob','Thornton',4,'Incredible beer, copper in colour.','2017-01-09 19:00:59'),
	(2,2,'Ted','Newton',1,'Not the nicest beer.','2017-01-09 12:30:12');

/*!40000 ALTER TABLE `reviews` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
