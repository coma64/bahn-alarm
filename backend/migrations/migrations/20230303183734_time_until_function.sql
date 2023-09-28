-- +goose Up
-- +goose StatementBegin
create function seconds_until_departure(t time) returns integer language plpgsql as $$
    declare
        seconds integer;
    begin
        seconds = (select extract(epoch from t));

        if t < current_time then
            -- add one day as it already passed today
            seconds = seconds + 60 * 60 * 24;
        end if;

        return seconds - extract(epoch from current_time);
end
$$;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop function seconds_until_departure;
-- +goose StatementEnd
