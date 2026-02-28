-- +goose Up
CREATE TABLE currencies (
    `id` BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `code` VARCHAR(3) NOT NULL,
    `name` VARCHAR(100) NOT NULL,
    `precision` INT NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP NULL
);

INSERT INTO currencies (`code`, `name`, `precision`) VALUES
('IDR', 'Indonesian Rupiah', 0),
('USD', 'United States Dollar', 2),
('SGD', 'Singapore Dollar', 2),
('EUR', 'Euro', 2),
('JPY', 'Japanese Yen', 0);

-- +goose Down
DROP TABLE currencies;
