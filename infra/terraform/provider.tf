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

  backend "azurerm" {
    resource_group_name  = "hub"
    storage_account_name = "terraformvars"
    container_name       = "terraform"
    key                  = "terraform2.tfstate"
  }
}

provider "azurerm" {


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
