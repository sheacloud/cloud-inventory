
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW distinct_iam_instance_profiles AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY instance_profile_id) oldest_report_time
	FROM
		iam_instance_profiles
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)