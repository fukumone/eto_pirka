
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE user_communities
(
  id int NOT NULL auto_increment COMMENT 'ID',
  user_id int,
  community_id int,
  created_at timestamp not null default current_timestamp,
  updated_at timestamp not null default current_timestamp on update current_timestamp,
  CONSTRAINT pk_user_communities PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE user_communities;
