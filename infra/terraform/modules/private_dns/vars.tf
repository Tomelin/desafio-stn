variable "name" {
  type        = string
  description = "Private DNS domain"

  validation {
    condition     = can(regex("^([a-z0-9]+\\.)+([a-z0-9]+\\.)+[a-z]{1,3}$", var.name))
    error_message = "The private dns domain should be a correct subdomain+domain"
  }
}

variable "tags" {
  type        = map(string)
  description = "List of tags to identity."
}


variable "resource_group" {
  type        = string
  description = "Name of resource group to deploy a resource"

  validation {
    condition     = length(var.resource_group) == 16
    error_message = "The resource group name should be a 16 character"
  }


}