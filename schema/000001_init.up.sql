CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    photo_url      varchar(255) not null ,
    user_vk_id     int       not null unique
);

CREATE TABLE trip_items
(
    id          serial       not null unique,
    author_id    int not null,
    date       varchar(255) not null,
    route       varchar(255) not null,
    count       varchar(255) not null,
    type       varchar(255) not null
);
