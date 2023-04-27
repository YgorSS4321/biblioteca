create table if not exists book
(
    id                      serial primary key,
    title                   varchar(100) not null,
    pages_quantity          int          not null,
    author_name             varchar(50)  not null,
    publishing_company_name varchar(100) not null,
    publication_year        int          not null
);