drop index if exists idx_extracted_texts_extraction_result_id;
drop index if exists idx_extracted_texts_deleted_at;
drop index if exists idx_extraction_results_deleted_at;

drop table if exists extracted_texts;
drop table if exists extraction_results;