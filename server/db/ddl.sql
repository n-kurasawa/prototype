CREATE DATABASE prototype;

USE prototype;

CREATE TABLE IF NOT EXISTS `users` (
   `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
   `name` varchar(128) NOT NULL,
   `age` int(10) unsigned NOT NULL,
   
   `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
   `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   PRIMARY KEY (`id`),
   KEY idx_users_name (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
