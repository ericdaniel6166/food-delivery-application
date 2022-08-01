BEGIN;

create table if not exists restaurants

(
    id                  serial       not null
        constraint restaurants_pk
            primary key,
    owner_id            integer,
    name                varchar(50)  not null,
    address             varchar(255) not null,
    city_id             integer,
    latitude            double precision,
    longitude           double precision,
    cover               json,
    logo                json,
    shipping_fee_per_km double precision default 0,
    status              integer          default 1,
    created_at          timestamp        default now(),
    updated_at          timestamp        default now()
);

CREATE INDEX IF NOT EXISTS idx_city_id ON restaurants (city_id);
CREATE INDEX IF NOT EXISTS idx_owner_id ON restaurants (owner_id);
CREATE INDEX IF NOT EXISTS idx_status ON restaurants (status);
---
create table if not exists restaurant_likes
(
    user_id     integer not null,
    restaurant_id     integer not null,
    created_ate timestamp default now(),
    updated_at  timestamp default now(),
    constraint restaurant_likes_pk
        unique (restaurant_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_user_id ON restaurant_likes (user_id);


END;