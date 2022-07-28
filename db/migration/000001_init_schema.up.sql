BEGIN;

create table if not exists restaurants

(
    id                  serial       not null
        constraint restaurants_pkey
            primary key,
    owner_id            integer,
    name                varchar(50)  not null,
    addr                varchar(255) not null,
    city_id             integer,
    lat                 double precision,
    lng                 double precision,
    cover               json,
    logo                json,
    shipping_fee_per_km double precision default 0,
    status              integer          default 1,
    created_at          timestamp        default now(),
    updated_at          timestamp        default now()
);

CREATE INDEX IF NOT EXISTS idx_city_id ON restaurants(city_id);
CREATE INDEX IF NOT EXISTS idx_owner_id ON restaurants(owner_id);
CREATE INDEX IF NOT EXISTS idx_status ON restaurants(status);

END;