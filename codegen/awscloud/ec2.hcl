aws_service "ec2" {
    sdk_url = "github.com/aws/aws-sdk-go-v2/service/ec2"

    resource "vpcs" {
        fetch_function = "DescribeVpcs"
        object_name = "Vpc"
        object_unique_id = "VpcId"
        result_object_path = ["Vpcs"]
        model_only = false
        pagination = true
    }
}