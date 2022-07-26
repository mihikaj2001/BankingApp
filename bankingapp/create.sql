
CREATE TABLE banks(
    bank_id bigserial PRIMARY KEY, 
    bank_code text UNIQUE, 
    bank_name text, 
    bank_address text, 
    created_at timestamp NOT NULL DEFAULT NOW(), 
    updated_at timestamp NOT NULL DEFAULT NOW() 
);

CREATE TABLE branches(
    branch_id bigserial PRIMARY KEY, 
    ifsc_code varchar(11) UNIQUE NOT NULL, 
    fk_bank_id bigint REFERENCES banks(bank_id), 
    branch_name text, 
    branch_addr text, 
    created_at timestamp NOT NULL DEFAULT NOW(),  
    updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE customers(
    customer_id bigserial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    pan_number text NOT NULL,
    aadhar_number numeric(12) NOT NULL,
    dob date NOT NULL,
    email text NOT NULL,
    contact_number numeric(10) NOT NULL,
    addr text,
    gender text,
    occupation text,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    fk_bank_id bigint REFERENCES banks(bank_id)
);

CREATE TABLE accounts(
    account_id bigserial PRIMARY KEY,
    fk_customer_id bigint REFERENCES customers(customer_id),
    account_number numeric(10) UNIQUE NOT NULL,
    is_active boolean,
    account_type text NOT NULL,
    current_balance bigint NOT NULL,
    fk_ifsc_code varchar(11) REFERENCES branches(ifsc_code),
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);


CREATE TABLE loans(
    loan_id bigserial PRIMARY Key,
    fk_customer_id bigint REFERENCES customers(customer_id),
    loan_amount bigint,
    loan_term int,
    loan_interest int,
    total_interest int,
    installments int,
    monthly_amount int,
    monthly_interest int,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE transactions(
    transaction_id bigserial PRIMARY KEY,
    fk_account_id bigint REFERENCES accounts(account_id),
    credited_amount int,
    debited_amount int,
    running_balance int,
    other_party_ifsc varchar(11) NOT NULL,
    other_party_account_number numeric(10) NOT NULL,
    other_party_bank_name text NOT NULL,
    other_party_branch_name text NOT NULL,
    transaction_at timestamp NOT NULL DEFAULT NOW()
);






ALTER TABLE accounts 
ADD CONSTRAINT fk_customer_id
FOREIGN KEY (fk_customer_id) 
REFERENCES customers (customer_id) ON DELETE CASCADE;


  ALTER TABLE accounts 
  DROP CONSTRAINT
  accounts_fk_customer_id_fkey;

  ALTER TABLE loans 
  DROP CONSTRAINT
  loans_fk_customer_id_fkey;

ALTER TABLE loans 
ADD CONSTRAINT fk_customer_id
FOREIGN KEY (fk_customer_id) 
REFERENCES customers (customer_id) ON DELETE CASCADE;

ALTER TABLE transactions 
  DROP CONSTRAINT
  transactions_fk_account_id_fkey;

ALTER TABLE transactions ADD CONSTRAINT fk_account_id
FOREIGN KEY (fk_account_id) 
REFERENCES accounts (account_id) ON DELETE CASCADE;



  

