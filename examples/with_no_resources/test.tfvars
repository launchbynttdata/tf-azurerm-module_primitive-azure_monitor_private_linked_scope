resource_names_map = {
  resource_group = {
    name       = "rg"
    max_length = 80
  }
  monitor_private_link_scope = {
    name       = "ampls"
    max_length = 80
  }
  application_insights = {
    name       = "appins"
    max_length = 80
  }
}
instance_env            = 0
instance_resource       = 0
logical_product_family  = "launch"
logical_product_service = "ampls"
class_env               = "gotest"
location                = "eastus"
ingestion_access_mode   = "PrivateOnly"
query_access_mode       = "PrivateOnly"
