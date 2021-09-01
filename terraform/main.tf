provider "aws" {
  region = "us-east-1"

  assume_role {
    role_arn     = "arn:aws:iam::306526781466:role/gitlab-automation-role"
    session_name = "GITLAB_PIPELINE"
  }
}
