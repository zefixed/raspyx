-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS teachers (
    uuid UUID PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    second_name VARCHAR NOT NULL,
    middle_name VARCHAR
);

CREATE TABLE IF NOT EXISTS groups (
    uuid UUID PRIMARY KEY,
    number VARCHAR NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS rooms (
    uuid UUID PRIMARY KEY,
    number VARCHAR NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS subjects (
    uuid UUID PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS subj_types (
    uuid UUID PRIMARY KEY,
    type VARCHAR NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS locations (
    uuid UUID PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS schedule (
    uuid UUID PRIMARY KEY,
    group_uuid UUID NOT NULL REFERENCES groups(uuid),
    subject_uuid UUID NOT NULL REFERENCES subjects(uuid),
    type_uuid UUID NOT NULL REFERENCES subj_types(uuid),
    location_uuid UUID NOT NULL REFERENCES locations(uuid),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    weekday INT NOT NULL,
    link TEXT
);

CREATE TABLE IF NOT EXISTS teachers_to_schedule (
    teacher_uuid UUID NOT NULL REFERENCES teachers(uuid) ON DELETE CASCADE,
    schedule_uuid UUID NOT NULL REFERENCES schedule(uuid) ON DELETE CASCADE,
    PRIMARY KEY (teacher_uuid, schedule_uuid)
);

CREATE TABLE IF NOT EXISTS rooms_to_schedule (
    room_uuid UUID NOT NULL REFERENCES rooms(uuid) ON DELETE CASCADE,
    schedule_uuid UUID NOT NULL REFERENCES schedule(uuid) ON DELETE CASCADE,
    PRIMARY KEY (room_uuid, schedule_uuid)
);

CREATE INDEX IF NOT EXISTS idx_schedule_group_uuid ON schedule(group_uuid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS teachers_to_schedule;
DROP TABLE IF EXISTS rooms_to_schedule;
DROP TABLE IF EXISTS schedule;
DROP TABLE IF EXISTS teachers;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS rooms;
DROP TABLE IF EXISTS subjects;
DROP TABLE IF EXISTS subj_types;
DROP TABLE IF EXISTS locations;
-- +goose StatementEnd