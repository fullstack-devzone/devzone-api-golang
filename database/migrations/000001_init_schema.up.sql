create table links
(
    id serial primary key,
    url varchar(1024) not null,
    title varchar(1024),
    created_at timestamp,
    updated_at timestamp
);
