create table users
(
    id         serial not null,
    name       varchar(255)                          not null,
    email      varchar(255)                          not null,
    password   varchar(255)                          not null,
    role       varchar(20)                           not null,
    created_at timestamp,
    updated_at timestamp,
    primary key (id),
    CONSTRAINT user_email_unique UNIQUE (email)
);

create table posts
(
    id         serial not null,
    url        varchar(1024)                         not null,
    title      varchar(1024)                         not null,
    content    text,
    created_by bigint                                not null REFERENCES users (id),
    created_at timestamp,
    updated_at timestamp,
    primary key (id)
);

ALTER SEQUENCE users_id_seq RESTART WITH 101;
ALTER SEQUENCE posts_id_seq RESTART WITH 101;
