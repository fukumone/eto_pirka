
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE messages
(
  id int NOT NULL auto_increment COMMENT 'ID',
  name varchar(128) NOT NULL COMMENT '名前',
  body varchar(128) NOT NULL COMMENT '本文',
  user_id int,
  community_id int,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  CONSTRAINT pk_messages PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE messages;
