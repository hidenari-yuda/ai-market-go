-- ユーザー情報を管理するテーブル
-- +migrate Up
CREATE TABLE IF NOT EXISTS content_details (
  id INT NOT NULL AUTO_INCREMENT UNIQUE,
  uuid VARCHAR(36) NOT NULL UNIQUE,
  content_id INT NOT NULL,
  content INT NOT NULL DEFAULT 0,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  PRIMARY KEY (id)
  INDEX idx_content_details_content_id (content_id)
);

ALTER TABLE content_details 
  ADD CONSTRAINT fk_content_details_content_id
  FOREIGN KEY (content_id) 
  REFERENCES details(id) 
  ON DELETE CASCADE
  ON UPDATE CASCADE;

-- +migrate Down
ALTER TABLE content_details DROP FOREIGN KEY fk_content_details_content_id;

DROP TABLE IF EXISTS content_details;
