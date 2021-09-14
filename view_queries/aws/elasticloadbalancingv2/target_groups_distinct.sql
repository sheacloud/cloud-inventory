
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW distinct_elasticloadbalancingv2_target_groups AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY target_group_arn) oldest_report_time
	FROM
		elasticloadbalancingv2_target_groups
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)