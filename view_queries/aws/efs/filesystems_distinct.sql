
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW distinct_efs_filesystems AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY file_system_id) oldest_report_time
	FROM
		efs_filesystems
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)