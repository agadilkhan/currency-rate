create table currency
(
    id         serial,
    code       varchar,
    rate       numeric,
    updated_at timestamp default now()
);
