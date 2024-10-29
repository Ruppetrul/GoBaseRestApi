CREATE TABLE IF NOT EXISTS general (
    symbol VARCHAR(20) NOT NULL UNIQUE,
    last_price NUMERIC(20, 8) NOT NULL,
    quote_volume NUMERIC(20, 8) NOT NULL,
    price_change_percent NUMERIC(5, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS binance (
    symbol VARCHAR(20) NOT NULL UNIQUE,
    last_price NUMERIC(20, 8) NOT NULL,
    quote_volume NUMERIC(20, 8) NOT NULL,
    price_change_percent NUMERIC(5, 2) NOT NULL
);