aws_service "athena" {
  service_cap_name         = "Athena"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/athena"
  extra_utilized_functions = ["GetDataCatalog", "GetDatabase", "GetWorkGroup", "ListTagsForResource"]
  tag_object_name          = "Tag"

  resource "workgroups" {
    fetch_function        = "ListWorkGroups"
    object_source_name    = "WorkGroup"
    object_singular_name  = "WorkGroup"
    object_plural_name    = "WorkGroups"
    object_unique_id      = "Name"
    object_response_field = "WorkGroups"
    model_only            = true
    pagination            = true
    use_post_processing   = false
    excluded_fields       = []
    convert_tags          = true
    display_fields        = ["Name"]

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }

  resource "data_catalogs" {
    fetch_function        = "ListDataCatalogs"
    object_source_name    = "DataCatalog"
    object_singular_name  = "DataCatalog"
    object_plural_name    = "DataCatalogs"
    object_unique_id      = "Name"
    object_response_field = "DataCatalogsSummary"
    model_only            = true
    pagination            = true
    use_post_processing   = false
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["Name"]
  }

  resource "databases" {
    fetch_function        = "ListDatabases"
    object_source_name    = "Database"
    object_singular_name  = "Database"
    object_plural_name    = "Databases"
    object_unique_id      = "Name"
    object_response_field = "DatabaseList"
    model_only            = true
    pagination            = true
    use_post_processing   = false
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["Name"]

    extra_field {
      name = "DataCatalog"
      type = "string"
    }
  }
}
