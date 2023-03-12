-- ユーザーのメディア情報を管理するテーブル
-- +migrate Up
CREATE TABLE IF NOT EXISTS user_medias (
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  uuid VARCHAR(36) NOT NULL UNIQUE,
  user_id INT NOT NULL,
  type INT NOT NULL DEFAULT 0,
  url VARCHAR(255) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY (id),
  INDEX idx_user_medias_user_id (user_id),
);

-- +migrate Down
DROP TABLE IF EXISTS user_medias;
