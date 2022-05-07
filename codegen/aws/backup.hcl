aws_service "backup" {
  service_cap_name         = "Backup"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/backup"
  extra_utilized_functions = ["ListTags", "ListBackupSelections"]

  resource "vaults" {
    fetch_function        = "ListBackupVaults"
    object_source_name    = "BackupVaultListMember"
    object_singular_name  = "BackupVault"
    object_plural_name    = "BackupVaults"
    object_unique_id      = "BackupVaultArn"
    object_response_field = "BackupVaultList"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["BackupVaultName"]

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }

  resource "plans" {
    fetch_function        = "ListBackupPlans"
    object_source_name    = "BackupPlansListMember"
    object_singular_name  = "BackupPlan"
    object_plural_name    = "BackupPlans"
    object_unique_id      = "BackupPlanArn"
    object_response_field = "BackupPlansList"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["BackupPlanName"]

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }

    child {
      object_source_name = "BackupSelectionsListMember"
      new_field_name     = "Selections"
      field_type         = "list"
    }
  }
}
