CREATE TABLE user_segments (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "type" smallint default 0,
  "volume" decimal(32,16) NOT NULL,
  "time" bigint NOT NULL,
  "created_at" timestamp DEFAULT NOW(),
  "updated_at" timestamp DEFAULT NOW(),
  "log_offset" bigint NOT NULL DEFAULT '0'
);

CREATE UNIQUE INDEX user_segment_idx_u_t ON user_segments (user_id, time);