
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW daily_distinct_iam_groups AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY group_id, report_date) oldest_report_time
	FROM
		iam_groups
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)