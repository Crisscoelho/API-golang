CREATE TABLE customer(
    customer_id bigserial PRIMARY KEY,
    name text NOT NULL,
    fk_current_balance bigint,
    phone_number text NOT NULL,
    email text NOT NULL,
    fk_bank_id bigint,
    fk_account_id bigint,
    account_agency text
);


