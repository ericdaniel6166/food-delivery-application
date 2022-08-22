BEGIN;

create table if not exists images
(
    id         bigserial
        constraint images_pk
            primary key,
    url        varchar(255) not null,
    width      integer      not null,
    height     integer      not null,
    status     boolean                  default true,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create table if not exists users
(
    id         bigserial
        constraint users_pk
            primary key,
    email      varchar(255) not null,
    fb_id      varchar(255),
    gg_id      varchar(255),
    password   varchar(255) not null,
    salt       varchar(255),
    last_name  varchar(255) not null,
    first_name varchar(255) not null,
    phone      varchar(255),
    role       varchar(255) not null,
    status     boolean                  default true,
    avatar     json,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create unique index if not exists idx_email
    on users (email);

---
create table if not exists restaurants

(
    id                  bigserial    not null
        constraint restaurants_pk
            primary key,
    user_id             bigint,
    like_count          bigint,
    name                varchar(255) not null,
    address             varchar(255) not null,
    city_id             bigint,
    latitude            double precision,
    longitude           double precision,
    cover               json,
    logo                json,
    shipping_fee_per_km double precision         default 0,
    status              boolean                  default true,
    created_at          timestamp with time zone default now(),
    updated_at          timestamp with time zone default now()
);

CREATE INDEX IF NOT EXISTS idx_city_id ON restaurants (city_id);
CREATE INDEX IF NOT EXISTS idx_owner_id ON restaurants (user_id);
CREATE INDEX IF NOT EXISTS idx_status ON restaurants (status);
---
create table if not exists restaurant_likes
(
    user_id       bigint not null,
    restaurant_id bigint not null,
    created_at    timestamp with time zone default now(),
    updated_at    timestamp with time zone default now(),
    constraint restaurant_likes_pk
        unique (restaurant_id, user_id)
);

CREATE INDEX IF NOT EXISTS idx_user_id ON restaurant_likes (user_id);


END;