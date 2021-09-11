
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW current_ec2_managed_prefix_lists AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY report_date) oldest_report_time
	FROM
		ec2_managed_prefix_lists
	WHERE (report_date = current_date)
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)