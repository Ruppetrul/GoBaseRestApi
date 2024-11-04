CREATE TABLE IF NOT EXISTS coingecko (
    id            VARCHAR(20) NOT NULL UNIQUE,
    symbol        VARCHAR(100) NOT NULL UNIQUE,
    name          VARCHAR(100) NOT NULL UNIQUE,
    current_price NUMERIC(20, 8) NOT NULL
);

CREATE TABLE IF NOT EXISTS general_html (
    html text NULL
);