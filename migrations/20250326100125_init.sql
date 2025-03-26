-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS teachers (
    id UUID PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    second_name VARCHAR NOT NULL,
    middle_name VARCHAR
);

CREATE TABLE IF NOT EXISTS groups (
    id UUID PRIMARY KEY,
    number VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS rooms (
    id UUID PRIMARY KEY,
    number VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS subjects (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS subj_types (
    id UUID PRIMARY KEY,
    type VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS locations (
    id UUID PRIMARY KEY,
    name VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS schedule (
    id UUID PRIMARY KEY,
    teacher_id UUID REFERENCES teachers(id),
    group_id UUID NOT NULL REFERENCES groups(id),
    room_id UUID REFERENCES rooms(id),
    subject_id UUID NOT NULL REFERENCES subjects(id),
    type_id UUID NOT NULL REFERENCES subj_types(id),
    location_id UUID NOT NULL REFERENCES locations(id),
    start_time TIME NOT NULL,
    end_time TIME NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,
    weekday INT NOT NULL,
    link TEXT
);

CREATE INDEX IF NOT EXISTS idx_schedule_teacher_id ON schedule(teacher_id);
CREATE INDEX IF NOT EXISTS idx_schedule_group_id ON schedule(group_id);
CREATE INDEX IF NOT EXISTS idx_schedule_room_id ON schedule(room_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS schedule;
DROP TABLE IF EXISTS teachers;
DROP TABLE IF EXISTS groups;
DROP TABLE IF EXISTS rooms;
DROP TABLE IF EXISTS subjects;
DROP TABLE IF EXISTS subj_types;
DROP TABLE IF EXISTS locations;
-- +goose StatementEnd