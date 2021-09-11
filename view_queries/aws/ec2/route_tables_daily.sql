
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW daily_ec2_route_tables AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY report_date) oldest_report_time
	FROM
		ec2_route_tables
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)