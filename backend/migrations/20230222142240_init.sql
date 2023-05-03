-- +goose Up
-- +goose StatementBegin
create table users
(
    id           bigserial
        primary key,
    name         text            not null
        unique,
    passwordHash     text            not null,
    createdAt timestamp default now() not null,
    isAdmin      boolean   default false not null
);

create table inviteTokens
(
    id bigint generated by default as identity primary key,
    token          text not null,
    createdById      bigint       not null
        references users,
    usedById         bigint
        references users,
    expiresAt timestamp default (now() + '14 days'::interval)
);

create table bahnStations (
    id bigint generated by default as identity primary key,
    externalId varchar(255) unique,
    name text not null
);

create table connections (
    id bigint generated by default as identity primary key,
    trackedById bigint references users on delete cascade not null,
    fromId bigint references bahnStations on delete cascade not null,
    toId bigint references bahnStations on delete cascade not null,
    departureMarginMinutes int default 10 not null,
    departureInfoHistoryDays int default 30 not null,
    unique (trackedById, fromId, toId)
);

create table departures (
    id bigint generated by default as identity primary key,
    connectionId bigint references connections on delete cascade not null,
    departure time not null,
    nextCheck timestamp not null
);

create table departureInfos (
    departureId bigint references departures on delete cascade,
    day date not null,
    scheduledTime timestamp not null,
    actualTime timestamp,
    primary key (departureId, day)
);

create function delete_old_departure_infos() returns trigger language plpgsql as $$
declare
    historyDays int;
    targetConnectionId int;
begin
    select
        c.id, c.departureInfoHistoryDays
    into targetConnectionId, historyDays
    from connections c
        join departures d on c.id = d.connectionId
    where d.id = old.departureId;

    delete
    from departureInfos i
    using
        departures d
    where
        i.departureId = d.id and
        d.connectionId = targetConnectionId and
        i.day < (now() :: date) - historyDays;

    return old;
end;
$$;

create trigger delete_old_departure_infos_trigger after insert on departureInfos for each row execute function delete_old_departure_infos();

create table pushNotificationSubs (
    id bigint generated by default as identity primary key,
    ownerId bigint references users on delete cascade not null,
    createdAt timestamp default now() not null,
    rawSubscription jsonb not null,
    name text not null,
    isEnabled boolean default true not null
);

create table alarms (
    id bigint generated by default as identity primary key,
    receiverId bigint references users on delete cascade not null,
    createdAt timestamp default now() not null,
    alarmData jsonb not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table departureInfos, departures, connections, bahnStations, users, inviteTokens, pushNotificationSubs, alarms;
-- +goose StatementEnd
