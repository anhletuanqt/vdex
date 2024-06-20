ALTER TABLE trade_transactions RENAME TO settlements;

ALTER TABLE settlements RENAME COLUMN trade_id TO ref_id;
ALTER TABLE settlements ADD type varchar(16) DEFAULT 'trade';

DROP INDEX IF EXISTS idx_t;
CREATE UNIQUE INDEX idx_settlements_rt ON settlements (ref_id,type);