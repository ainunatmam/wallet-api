-- +goose Up
CREATE TABLE transactions (
    `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
    `trx_id` VARCHAR(255) NOT NULL UNIQUE,
    `wallet_id` BIGINT NOT NULL,
    `transfered_wallet_id` BIGINT NULL,
    `transaction_type` VARCHAR(50) NOT NULL,
    `amount` BIGINT NOT NULL,
    `status` VARCHAR(100) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE transactions;
