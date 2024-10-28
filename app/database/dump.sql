CREATE TABLE IF NOT EXISTS prices (
    name VARCHAR(255) NOT NULL UNIQUE,
    price NUMERIC(20, 8) NOT NULL
);

CREATE TABLE IF NOT EXISTS tickers (
    symbol VARCHAR(20) NOT NULL UNIQUE,
    last_price NUMERIC(20, 8) NOT NULL,
    price_change_percent NUMERIC(5, 2) NOT NULL
);