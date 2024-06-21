terraform {
  required_version = ">= 1.4"
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "3.108.0"
    }
    helm = {
      source  = "hashicorp/helm"
      version = "2.14.0"
    }
  }
}

provider "azurerm" {
  subscription_id = "ea9f2737-3006-4d2f-b375-177c70866743"
  client_secret   = "Slx8Q~lt6wdZ52XsouZTYFR.snnTUDyL~kHlWdwC"
  client_id       = "1d6c772d-1255-44f8-b100-e8f2a963ec4c"
  tenant_id       = "95b4c5e3-043d-4721-af9f-de5039ff7051"

  features {}
}

provider "helm" {
  debug = true
  kubernetes {
    host                   = local.kube.kube_config[0].host
    token                  = local.kube.kube_config[0].password
    cluster_ca_certificate = base64decode(local.kube.kube_config[0].cluster_ca_certificate)
  }
}