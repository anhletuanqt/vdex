ALTER TABLE orders ADD public_id uuid DEFAULT gen_random_uuid();
ALTER TABLE fills ADD public_id uuid DEFAULT gen_random_uuid();
ALTER TABLE trades ADD public_id uuid DEFAULT gen_random_uuid();
ALTER TABLE users ADD public_id uuid DEFAULT gen_random_uuid();

ALTER TABLE fills DROP COLUMN log_seq;
ALTER TABLE ticks DROP COLUMN log_seq;
ALTER TABLE trades DROP COLUMN log_seq;

CREATE UNIQUE INDEX user_pid ON users (public_id);
CREATE UNIQUE INDEX trade_pid ON trades (public_id);
CREATE UNIQUE INDEX fill_pid ON fills (public_id);
CREATE UNIQUE INDEX order_pid ON orders (public_id);
