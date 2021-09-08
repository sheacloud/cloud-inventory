#!/bin/bash

docker build -t cloud-inventory-scraper -f Dockerfile .

REGION="us-east-1"
REPO_URL=$(aws sts get-caller-identity --output text --query Account).dkr.ecr.$REGION.amazonaws.com/cloud-inventory-scraper
aws ecr get-login-password --region=$REGION |  docker login --username AWS --password-stdin $(aws sts get-caller-identity --output text --query Account).dkr.ecr.$REGION.amazonaws.com
docker tag cloud-inventory-scraper:latest $REPO_URL
docker push $REPO_URL