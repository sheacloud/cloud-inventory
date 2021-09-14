
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW daily_efs_filesystems AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY report_date) oldest_report_time
	FROM
		efs_filesystems
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)