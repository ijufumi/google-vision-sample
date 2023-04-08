create table if not exists jobs
(
    id                varchar(26)  not null primary key,
    name              varchar(255) not null,
    original_file_key varchar(255) not null,
    status            varchar(10)  not null,
    created_at        timestamp    not null,
    updated_at        timestamp    not null,
    deleted_at        timestamp
);

create index if not exists idx_jobs_deleted_at on jobs (deleted_at);

create table if not exists input_files
(
    id           varchar(26)  not null primary key,
    job_id       varchar(26)  not null references jobs (id),
    page_no      integer      not null default 1,
    file_key     varchar(255) not null,
    file_name    varchar(255) not null,
    size         bigint       not null default 0,
    width        bigint       not null default 0,
    height       bigint       not null default 0,
    content_type varchar(100) not null default 'application/octet-stream',
    status       varchar(10)  not null,
    created_at   timestamp    not null,
    updated_at   timestamp    not null,
    deleted_at   timestamp
);

create table if not exists output_files
(
    id            varchar(26)  not null primary key,
    job_id        varchar(26)  not null references jobs (id),
    input_file_id varchar(26)  not null references input_files (id),
    file_key      varchar(255) not null,
    file_name     varchar(255) not null,
    size          bigint       not null default 0,
    width         bigint       not null default 0,
    height        bigint       not null default 0,
    content_type  varchar(100) not null default 'application/octet-stream',
    created_at    timestamp    not null,
    updated_at    timestamp    not null,
    deleted_at    timestamp
);

create table if not exists extracted_texts
(
    id             varchar(26)      not null primary key,
    job_id         varchar(26)      not null references jobs (id),
    input_file_id  varchar(26)      not null references input_files (id),
    output_file_id varchar(26)      not null references output_files (id),
    text           varchar(255)     not null,
    top            double precision not null,
    bottom         double precision not null,
    "left"         double precision not null,
    "right"        double precision not null,
    created_at     timestamp        not null,
    updated_at     timestamp        not null,
    deleted_at     timestamp
);

create index if not exists idx_extracted_texts_deleted_at on extracted_texts (deleted_at);
create index if not exists idx_extracted_texts_job_id on extracted_texts (job_id);
create index if not exists idx_extracted_texts_input_file_id on extracted_texts (input_file_id);
create index if not exists idx_extracted_texts_output_file_id on extracted_texts (output_file_id);
