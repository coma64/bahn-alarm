-- +goose Up
-- +goose StatementBegin
create view fatDepartures as
select
    d.*,
    f.externalId fromStationId,
    f.name fromStationName,
    t.externalId toStationId,
    t.name toStationName,
    c.departureMarginMinutes,
    c.trackedById,
    coalesce(floor((extract(epoch from i.actualTime) - extract(epoch from i.scheduledTime)) / 60), 0) as delayMinutes,
    case
        when i.departureId is null then 'not-checked'
        when i.actualTime = i.scheduledTime then 'on-time'
        when i.actualTime is null then 'canceled'
        else 'delayed'
        end status
from departures d
         inner join connections c on d.connectionId = c.id
         inner join bahnStations f on c.fromId = f.id
         inner join bahnStations t on c.toId = t.id
         left outer join departureInfos i on i.departureId = d.id and i.day = now() :: date;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view fatDepartures;
-- +goose StatementEnd
