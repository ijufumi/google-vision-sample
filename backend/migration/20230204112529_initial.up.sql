create table extraction_results
(
    id         varchar(26)  not null primary key,
    status     varchar(10)  not null,
    image_uri  varchar(255) not null,
    created_at timestamp    not null,
    updated_at timestamp    not null,
    deleted_at timestamp,
);

create table extracted_texts
(
    id                   varchar(26)      not null primary key,
    extraction_result_id varchar(26)      not null references extraction_results (id),
    text                 varchar(255)     not null,
    top                  double precision not null,
    bottom               double precision not null,
    left                 double precision not null,
    right                double precision not null,
    created_at           timestamp        not null,
    updated_at           timestamp        not null,
    deleted_at           timestamp
);