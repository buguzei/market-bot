CREATE TABLE users (
    userid   bigint primary key,
    username text,
    invoice  int
);

CREATE TABLE goods (
    sellerid    bigint references users (userid),
    goodid      bigint,
    price       bigint,
    description text,
    photo       text
);

CREATE TABLE reviews (
    mark        int2,
    description text,
    photo       text
);