CREATE TABLE IF NOT EXISTS coingecko (
    id                           VARCHAR(20) NOT NULL UNIQUE,
    symbol                       VARCHAR(100) NOT NULL UNIQUE,
    name                         VARCHAR(100) NOT NULL UNIQUE,
    current_price                NUMERIC(28, 8) NOT NULL,
    market_cap                   NUMERIC(28) NOT NULL,
    price_change_percentage_24h  NUMERIC(20, 8) NOT NULL
);

CREATE TABLE IF NOT EXISTS general_html (
    html text NULL
);

CREATE TABLE IF NOT EXISTS page_visits (
    visit_date VARCHAR(20) NOT NULL,
    visit_count INTEGER
);

ALTER TABLE page_visits ADD COLUMN IF NOT EXISTS visit_ip VARCHAR(40) NULL;
ALTER TABLE page_visits ADD UNIQUE (visit_date, visit_ip);

CREATE TABLE IF NOT EXISTS average_price_per_day (
    pair          varchar(20) NOT NULL,
    date          date,
    average_price NUMERIC(28, 8) NOT NULL
);

ALTER TABLE average_price_per_day ADD UNIQUE (date, pair);

CREATE TABLE IF NOT EXISTS bitcoin_halving_dates (
    halving_date DATE NOT NULL UNIQUE
);

INSERT INTO bitcoin_halving_dates (halving_date) VALUES
('2012-11-28'),
('2016-07-09'),
('2020-05-11'),
('2024-04-20');