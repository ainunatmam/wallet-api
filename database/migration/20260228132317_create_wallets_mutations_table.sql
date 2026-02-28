-- +goose Up
CREATE TABLE wallets_mutations (
    id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    wallet_id BIGINT NOT NULL,
    mutation_type VARCHAR(50) NOT NULL,
    amount BIGINT NOT NULL,
    balance_before BIGINT NOT NULL,
    balance_after BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE wallets_mutations;
