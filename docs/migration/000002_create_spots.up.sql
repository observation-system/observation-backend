CREATE TABLE `spots` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `spot_key` VARCHAR(12) NOT NULL UNIQUE,
    `user_key` VARCHAR(12) NOT NULL UNIQUE,
    `name` VARCHAR(255) NOT NULL,
    `url` VARCHAR(255) NOT NULL,
    `address` VARCHAR(255) NOT NULL,
    `latitude` VARCHAR(255) NOT NULL,
    `longitude` VARCHAR(255) NOT NULL,
    `count` VARCHAR(255) NOT NULL,
    `count_day` VARCHAR(255) NOT NULL,
    `status` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
