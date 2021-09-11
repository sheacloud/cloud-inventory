
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW distinct_ec2_vpc_peering_connections AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY vpc_peering_connection_id) oldest_report_time
	FROM
		ec2_vpc_peering_connections
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)