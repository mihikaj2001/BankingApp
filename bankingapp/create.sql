

CREATE TYPE gender_enum AS ENUM ('Male', 'Female', 'Other');
CREATE TYPE account_type_enum AS ENUM ('Loan', 'Savings', 'Current');
CREATE TYPE transaction_type_enum AS ENUM ('Withdrawal', 'Deposit');


CREATE TABLE banks(
    id bigserial PRIMARY KEY, 
    code text UNIQUE, 
    name text, 
    addr text, 
    created_at timestamp NOT NULL DEFAULT NOW(), 
    updated_at timestamp NOT NULL DEFAULT NOW() 
);

CREATE TABLE branches(
    id bigserial PRIMARY KEY, 
    ifsc_code varchar(11) UNIQUE NOT NULL, 
    fk_bank_id bigint REFERENCES banks(id), 
    name text, 
    addr text, 
    created_at timestamp NOT NULL DEFAULT NOW(),  
    updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE customers(
    id bigserial PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    pan_number text NOT NULL,
    aadhar_number numeric(12) NOT NULL,
    dob date NOT NULL,
    email text NOT NULL,
    contact_number numeric(10) NOT NULL,
    addr text,
    gender gender_enum,
    occupation text,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    fk_bank_id bigint REFERENCES banks(id)
);

CREATE TABLE accounts(
    id bigserial PRIMARY KEY,
    fk_customer_id bigint REFERENCES customers(id),
    fk_branch_id varchar(11) REFERENCES branches(id),
    account_number numeric(10) UNIQUE NOT NULL,
    is_active boolean,
    type account_type_enum text NOT NULL,
    current_balance bigint NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);


CREATE TABLE loans(
    id bigserial PRIMARY KEY,
    fk_customer_id bigint REFERENCES customers(id),
    fk_branch_id varchar(11) REFERENCES branches(id),
    amount bigint,
    term int,
    interest_percent int,
    total_interest int,
    installments int,
    monthly_amount int,
    monthly_interest int,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE transactions(
    id bigserial PRIMARY KEY,
    fk_account_id bigint REFERENCES accounts(id),
    fk_branch_id bigint REFERENCES branches(id)
    amount int,
    type transaction_type_enum,
    running_balance int,
    other_party_account_id numeric(10),
    transaction_at timestamp NOT NULL DEFAULT NOW()
);






-- ALTER TABLE accounts 
-- ADD CONSTRAINT fk_customer_id
-- FOREIGN KEY (fk_customer_id) 
-- REFERENCES customers (customer_id) ON DELETE CASCADE;


--   ALTER TABLE accounts 
--   DROP CONSTRAINT
--   accounts_fk_customer_id_fkey;

--   ALTER TABLE loans 
--   DROP CONSTRAINT
--   loans_fk_customer_id_fkey;

-- ALTER TABLE loans 
-- ADD CONSTRAINT fk_customer_id
-- FOREIGN KEY (fk_customer_id) 
-- REFERENCES customers (customer_id) ON DELETE CASCADE;

-- ALTER TABLE transactions 
--   DROP CONSTRAINT
--   transactions_fk_account_id_fkey;

-- ALTER TABLE transactions ADD CONSTRAINT fk_account_id
-- FOREIGN KEY (fk_account_id) 
-- REFERENCES accounts (account_id) ON DELETE CASCADE;



  

