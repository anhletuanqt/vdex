CREATE TABLE worker_configs (
  "id" SERIAL PRIMARY KEY,
  "block_number" integer
);

CREATE TABLE users (
  "id" SERIAL PRIMARY KEY,
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp DEFAULT NOW(),
  "address" varchar(255) NOT NULL
);

CREATE TABLE products (
  "id" varchar(255) PRIMARY KEY,
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp DEFAULT NOW(),
  "base_currency" varchar(255) NOT NULL,
  "quote_currency" varchar(255) NOT NULL,
  "base_min_size" decimal(32,16) NOT NULL,
  "base_max_size" decimal(32,16) NOT NULL,
  "base_scale" integer NOT NULL,
  "quote_scale" integer NOT NULL,
  "quote_increment" double precision NOT NULL,
  "quote_min_size" decimal(32,16) NOT NULL,
  "quote_max_size" decimal(32,16) NOT NULL
);

CREATE TABLE orders (
  "id" BIGSERIAL PRIMARY KEY,
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp DEFAULT NOW(),
  "product_id" varchar(255) NOT NULL,
  "user_id" BIGSERIAL NOT NULL,
  "size" decimal(32,16) NOT NULL DEFAULT '0.0000000000000000',
  "funds" decimal(32,16) NOT NULL DEFAULT '0.0000000000000000',
  "filled_size" decimal(32,16) NOT NULL DEFAULT '0.0000000000000000',
  "executed_value" decimal(32,16) NOT NULL DEFAULT '0.0000000000000000',
  "price" decimal(32,16) NOT NULL DEFAULT '0.0000000000000000',
  "fill_fees" decimal(32,16) NOT NULL DEFAULT '0.0000000000000000',
  "type" varchar(255) NOT NULL,
  "side" varchar(255) NOT NULL,
  "time_in_force" varchar(255) DEFAULT NULL,
  "status" varchar(255) NOT NULL,
  "settled" boolean NOT NULL DEFAULT 'false',
  "created_tx_hash" varchar(255) NOT NULL 
);

CREATE TABLE trades (
  "id" BIGSERIAL PRIMARY KEY,
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp DEFAULT NOW(),
  "product_id" varchar(255) NOT NULL,
  "time" timestamp,
  "taker_order_id" bigint NOT NULL,
  "maker_order_id" bigint NOT NULL,
  "price" decimal(32,16) NOT NULL,
  "size" decimal(32,16) NOT NULL,
  "side" varchar(255) NOT NULL,
  "log_offset" bigint NOT NULL DEFAULT '0',
  "log_seq" bigint NOT NULL DEFAULT '0'
);

CREATE TABLE ticks (
  "id" BIGSERIAL PRIMARY KEY,
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp DEFAULT NOW(),
  "product_id" varchar(255) NOT NULL,
  "granularity" bigint NOT NULL,
  "time" bigint NOT NULL,
  "open" decimal(32,16) NOT NULL,
  "high" decimal(32,16) NOT NULL,
  "low" decimal(32,16) NOT NULL,
  "close" decimal(32,16) NOT NULL,
  "volume" decimal(32,16) NOT NULL,
  "log_offset" bigint NOT NULL DEFAULT '0',
  "log_seq" bigint NOT NULL DEFAULT '0'
);

CREATE TABLE fills (
  "id" BIGSERIAL PRIMARY KEY,
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp DEFAULT NOW(),
  "user_id" bigint NOT NULL,
  "order_id" bigint NOT NULL DEFAULT '0',
  "product_id" varchar(255) NOT NULL,
  "size" decimal(32,16) NOT NULL,
  "price" decimal(32,16) NOT NULL,
  "funds" decimal(32,16) NOT NULL DEFAULT '0.0000000000000000',
  "fee" decimal(32,16) NOT NULL DEFAULT '0.0000000000000000',
  "settled" boolean default FALSE,
  "side" varchar(255) NOT NULL,
  "done_reason" varchar(255) NOT NULL,
  "log_offset" bigint NOT NULL DEFAULT '0',
  "log_seq" bigint NOT NULL DEFAULT '0'
);

CREATE TABLE trade_transactions (
  "id" BIGSERIAL PRIMARY KEY,
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp DEFAULT NOW(),
  "trade_id" bigint NOT NULL,
  "tx_hash" varchar(255) NOT NULL DEFAULT '',
  "settled" boolean default FALSE
);

CREATE UNIQUE INDEX idx_a ON users (address);
CREATE INDEX idx_uspsi ON orders (user_id, product_id, status, side, id);
CREATE INDEX idx_u ON orders (user_id);
CREATE UNIQUE INDEX idx_cth ON orders (created_tx_hash);
CREATE UNIQUE INDEX idx_p_g_t ON ticks (product_id, granularity, time);
CREATE UNIQUE INDEX idx_o_l ON fills (order_id, log_offset);
CREATE INDEX idx_oudr ON fills (order_id, user_id, done_reason);
CREATE UNIQUE INDEX idx_t ON trade_transactions (trade_id);

INSERT INTO "products" ("id","created_at","updated_at","base_currency","quote_currency","base_min_size","base_max_size","base_scale","quote_scale","quote_increment","quote_min_size","quote_max_size") values
('BNB-VIC',null,null,'BNB','VIC',0.00001,10000,4,2,0.01,0E-16,0E-16);