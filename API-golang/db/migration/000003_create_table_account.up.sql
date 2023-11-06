 CREATE TABLE account(
    account_id bigserial PRIMARY KEY,
    fk_customer_id bigint,
    account_number numeric(10) UNIQUE NOT NULL,
    is_active boolean,
    account_type text NOT NULL,
    current_balance bigint NOT NULL,
    account_agency text
 );  

 