create table if not exists jobs
(
    id         varchar(26) not null primary key,
    status     varchar(10) not null,
    created_at timestamp   not null,
    updated_at timestamp   not null,
    deleted_at timestamp
);

create index if not exists idx_jobs_deleted_at on jobs (deleted_at);

create table if not exists extracted_texts
(
    id         varchar(26)      not null primary key,
    job_id     varchar(26)      not null references jobs (id),
    text       varchar(255)     not null,
    top        double precision not null,
    bottom     double precision not null,
    "left"     double precision not null,
    "right"    double precision not null,
    created_at timestamp        not null,
    updated_at timestamp        not null,
    deleted_at timestamp
);

create index if not exists idx_extracted_texts_deleted_at on extracted_texts (deleted_at);
create index if not exists idx_extracted_texts_job_id on extracted_texts (job_id);

create table if not exists job_files
(
    id           varchar(26)  not null primary key,
    job_id       varchar(26)  not null references jobs (id),
    is_output    boolean      not null default false,
    file_key     varchar(255) not null,
    file_name    varchar(255) not null,
    size         bigint       not null default 0,
    content_type varchar(100) not null default 'application/octet-stream',
    created_at   timestamp    not null,
    updated_at   timestamp    not null,
    deleted_at   timestamp
);

create index if not exists idx_job_files_deleted_at on job_files (deleted_at);
create index if not exists idx_job_files_job_id on job_files (job_id);
