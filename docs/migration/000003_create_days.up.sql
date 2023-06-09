CREATE TABLE `days` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `day_key` VARCHAR(12) NOT NULL UNIQUE,
    `spot_key` VARCHAR(12) NOT NULL,
    `count` VARCHAR(255) NOT NULL,
    `status` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
