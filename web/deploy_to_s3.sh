#!/bin/bash
set -x

aws s3 sync ./build/ s3://$S3_DEPLOY_BUCKET/ --delete
INVALIDATION_ID=$(aws cloudfront create-invalidation --distribution-id $CLOUDFRONT_DISTRIBUTION_ID --paths "/*" | jq -r '.Invalidation.Id')
aws cloudfront wait invalidation-completed --distribution-id $CLOUDFRONT_DISTRIBUTION_ID --id $INVALIDATION_ID