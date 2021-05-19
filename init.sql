CREATE DATABASE IF NOT EXISTS amazon;
USE amazon;

CREATE TABLE IF NOT EXISTS `amazon_products` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `title` varchar(255) UNIQUE NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `description` varchar(8000) DEFAULT NULL,
  `price` varchar(255) DEFAULT NULL,
  `total_reviews` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`,`title`),
  UNIQUE KEY `image_url` (`image_url`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;