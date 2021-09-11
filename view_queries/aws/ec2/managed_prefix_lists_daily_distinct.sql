
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW daily_distinct_ec2_managed_prefix_lists AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY prefix_list_arn, report_date) oldest_report_time
	FROM
		ec2_managed_prefix_lists
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)