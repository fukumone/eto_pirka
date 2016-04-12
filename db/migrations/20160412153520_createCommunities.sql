
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE communities
(
  id int NOT NULL auto_increment COMMENT 'ID',
  name varchar(128) NOT NULL COMMENT '名前',
  description varchar(128) NOT NULL COMMENT '説明',
  administrator_id int,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  CONSTRAINT pk_communities PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE communities;
