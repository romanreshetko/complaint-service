CREATE TABLE review_complaints (
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    review_id  BIGINT    NOT NULL,
    author_id  BIGINT    NOT NULL,
    reason     TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE comment_complaints (
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    comment_id BIGINT    NOT NULL,
    author_id  BIGINT    NOT NULL,
    reason     TEXT      NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE TABLE moderation_logs (
    id           BIGSERIAL NOT NULL PRIMARY KEY,
    actor_id     BIGINT    NOT NULL,
    action_time  TIMESTAMP NOT NULL,
    content_type TEXT      NOT NULL CHECK (content_type IN ('review', 'comment')),
    content_id   BIGINT    NOT NULL,
    result       TEXT      NOT NULL CHECK (result IN ('published', 'blocked', 'error'))
);