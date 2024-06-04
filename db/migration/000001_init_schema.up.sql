CREATE TABLE Accounts (
    id SERIAL PRIMARY KEY,
    "owner" VARCHAR(350) NOT NULL,
    balance BIGINT NOT NULL,
    currency VARCHAR(150) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE Entries (
    id SERIAL PRIMARY KEY,
    account_id INTEGER REFERENCES Accounts(id) NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE Transfers (
    id SERIAL PRIMARY KEY , 
    from_account_id INTEGER REFERENCES Accounts(id) NOT NULL,
    to_account_id INTEGER REFERENCES Accounts(id) NOT NULL,
    amount BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX ON Accounts ("owner");
CREATE INDEX ON Entries ("account_id");
CREATE INDEX ON Transfers ("from_account_id");
CREATE INDEX ON Transfers ("to_account_id");
CREATE INDEX ON Transfers ("from_account_id","to_account_id");

