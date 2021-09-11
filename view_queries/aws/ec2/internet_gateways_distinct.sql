
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW distinct_ec2_internet_gateways AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY internet_gateway_id) oldest_report_time
	FROM
		ec2_internet_gateways
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)