CREATE TABLE transaction(
    transaction_id bigserial PRIMARY KEY,
    fk_account_id bigint,
    credited_amount int,
    debited_amount int,
    running_balance int,
    credit_origin text,
    other_party_account_number numeric(10) NOT NULL,
    other_party_bank_name text NOT NULL
);