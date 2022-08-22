BEGIN;

create table if not exists users
(
    id         serial
        constraint users_pk
            primary key,
    email      varchar(50) not null,
    fb_id      varchar(50),
    gg_id      varchar(50),
    password   varchar(50) not null,
    salt       varchar(50),
    last_name  varchar(50) not null,
    first_name varchar(50) not null,
    phone      varchar(20),
    role       varchar(50) not null,
    status     integer                  default 1,
    avatar     json,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create unique index if not exists idx_email
    on users (email);

---
create table if not exists restaurants

(
    id                  serial       not null
        constraint restaurants_pk
            primary key,
    user_id             integer,
    name                varchar(50)  not null,
    address             varchar(255) not null,
    city_id             integer,
    latitude            double precision,
    longitude           double precision,
    cover               json,
    logo                json,
    shipping_fee_per_km double precision         default 0,
    status              integer                  default 1,
    created_at          timestamp with time zone default now(),
    updated_at          timestamp with time zone default now()
);

CREATE INDEX IF NOT EXISTS idx_city_id ON restaurants (city_id);
CREATE INDEX IF NOT EXISTS idx_owner_id ON restaurants (user_id);
CREATE INDEX IF NOT EXISTS idx_status ON restaurants (status);
---
create table if not exists restaurant_likes
(
    user_id       integer not null,
    restaurant_id integer not null,
    created_at    timestamp with time zone default now(),
    updated_at    timestamp with time zone default now(),
    constraint restaurant_likes_pk
        unique (restaurant_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_user_id ON restaurant_likes (user_id);


END;