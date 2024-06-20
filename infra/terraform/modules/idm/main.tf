/**
 * # Main title
 *
 * Azure Identity
 *
 * Module to deploy Azure identity
 */

resource "azurerm_user_assigned_identity" "idm" {
  name                = "${var.name}-idm"
  resource_group_name = var.resource_group
  location            = var.location
  tags                = var.tags

  lifecycle {
    ignore_changes = [
      tags["created_at"]
    ]
  }
}