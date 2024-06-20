ALTER TABLE settlements ADD taker_fee decimal(32,16) NOT NULL DEFAULT '0.0000000000000000';
ALTER TABLE settlements ADD maker_fee decimal(32,16) NOT NULL DEFAULT '0.0000000000000000';
ALTER TABLE settlements ADD time timestamp;