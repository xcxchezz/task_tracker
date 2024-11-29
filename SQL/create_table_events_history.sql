CREATE TABLE events_history (
    id BIGSERIAL PRIMARY KEY,
    event_id BIGINT NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    changed_by BIGINT NOT NULL REFERENCES users(id),
    change_time TIMESTAMP NOT NULL DEFAULT now(),
    old_title VARCHAR(128),
    old_description TEXT,
    old_date TIMESTAMP,
    change_type VARCHAR(32) NOT NULL
);
