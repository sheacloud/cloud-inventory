
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW current_iam_policies AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY report_date) oldest_report_time
	FROM
		iam_policies
	WHERE (report_date = current_date)
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)