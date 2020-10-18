-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE `books` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(50) DEFAULT NULL,
    `author` varchar(80) DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS books;
