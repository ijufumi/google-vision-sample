drop index if exists idx_files_extraction_result_id;
drop index if exists idx_files_deleted_at;
drop index if exists idx_extracted_texts_extraction_result_id;
drop index if exists idx_extracted_texts_deleted_at;
drop index if exists idx_jobs_deleted_at;

drop table if exists files;
drop table if exists extracted_texts;
drop table if exists jobs;