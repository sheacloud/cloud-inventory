
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW distinct_rds_db_clusters AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY db_cluster_arn) oldest_report_time
	FROM
		rds_db_clusters
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)