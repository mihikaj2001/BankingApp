

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
    fk_bank_customer_map bigint REFERENCES bank_customer_map(id)
);

CREATE TABLE bank_customer_map(
    id bigserial PRIMARY KEY,
    fk_customer_id BIGINT REFERENCES customers(id),
    fk_bank_id BIGINT REFERENCES banks(id)
);

CREATE TABLE accounts(
    id bigserial PRIMARY KEY,
    fk_customer_id bigint REFERENCES customers(id),
    fk_branch_id BIGINT REFERENCES branches(id),
    account_number numeric(10) UNIQUE NOT NULL,
    is_active boolean,
    account_type account_type_enum NOT NULL,
    current_balance bigint NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW()
);


CREATE TABLE loans(
    id bigserial PRIMARY KEY,
    fk_customer_id bigint REFERENCES customers(id),
    fk_branch_id BIGINT REFERENCES branches(id),
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
    fk_branch_id bigint REFERENCES branches(id),
    amount int,
    transaction_type transaction_type_enum,
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

-- ALTER TABLE customers ADD fk_bank_customer_map BIGINT REFERENCES bank_customer_map(id);
-- ALTER TABLE customers DROP COLUMN fk_bank_id;

  ALTER TABLE bank_customer_map 
  DROP CONSTRAINT
  bank_customer_map_fk_customer_id_fkey;
ALTER TABLE bank_customer_map 
ADD CONSTRAINT fk_customer_id
FOREIGN KEY (fk_customer_id) 
REFERENCES customers (id) ON DELETE CASCADE;


ALTER TABLE bank_customer_map 
  DROP CONSTRAINT
  bank_customer_map_fk_bank_id_fkey;

ALTER TABLE bank_customer_map ADD CONSTRAINT fk_cascade_bank
FOREIGN KEY (fk_bank_id) 
REFERENCES banks(id) ON DELETE CASCADE;


INSERT INTO banks(id, code, name, addr, created_at, updated_at) VALUES(1, '123456', 'HSBC', 'Hyderabad', NOW(), NOW());
 Insert into customers(first_name, last_name, pan_number, aadhar_number, dob, email, contact_number, addr, gender, occupation) values ('Mihika', 'Jain', 1010101010, 123412341234, '2001-04-03', 'mihikaj2001@gmail.com', 9700166414, '6-3-1100/4/5', 'Female', 'coding');
 INSERT INTO branches(id, ifsc_code, fk_bank_id, name, addr, created_at, updated_at) VALUES (1, '12345565432', 1, 'JHANDEWALAN', '91 springboard, jhandewalan', NOW(), NOW());

 INSERT INTO accounts(id, fk_customer_id, fk_branch_id, account_number, is_active, account_type, current_balance, created_at, updated_at) VALUES (1, 1, 1, 1234562345, true, 'Loan', 3000, NOW(), NOW());

 INSERT INTO accounts(id, fk_customer_id, fk_branch_id, account_number, is_active, account_type, current_balance, created_at, updated_at) VALUES (2, 1, 1, 1111122222, true, 'Savings', 993000, NOW(), NOW());
INSERT INTO bank_customer_map(fk_bank_id, fk_customer_id) VALUES (1, 2);

INSERT INTO transactions(id, fk_account_id, fk_branch_id, amount, transaction_type, running_balance, other_party_account_id, transaction_at) VALUES (1, 1, 1, 1000, 'Withdrawal', 30000, 1, NOW())