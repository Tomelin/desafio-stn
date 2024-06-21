# Main title

Group of module to deploy any resource in Azure to support Kubernetes Servuce

OBS.: Em ambiente produtivo, alteraria o source para a url do git.  Não fiz para otimizar o tempo

## Requirements

The following requirements are needed by this module:

- <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) (>= 1.4)

- <a name="requirement_azurerm"></a> [azurerm](#requirement\_azurerm) (3.108.0)

- <a name="requirement_helm"></a> [helm](#requirement\_helm) (2.14.0)

## Providers

No providers.

## Modules

The following Modules are called:

### <a name="module_aks"></a> [aks](#module\_aks)

Source: ./modules/aks

Version:

### <a name="module_argocd"></a> [argocd](#module\_argocd)

Source: ./modules/argoproj/argocd

Version:

### <a name="module_identity"></a> [identity](#module\_identity)

Source: ./modules/idm

Version:

### <a name="module_ingress"></a> [ingress](#module\_ingress)

Source: ./modules/ingress

Version:

### <a name="module_log_analytics"></a> [log\_analytics](#module\_log\_analytics)

Source: ./modules/law

Version:

### <a name="module_private_dns"></a> [private\_dns](#module\_private\_dns)

Source: ./modules/private_dns

Version:

### <a name="module_resource_group"></a> [resource\_group](#module\_resource\_group)

Source: ./modules/rsg

Version:

## Resources

No resources.

## Required Inputs

No required inputs.

## Optional Inputs

The following input variables are optional (have default values):

### <a name="input_default_node_pool"></a> [default\_node\_pool](#input\_default\_node\_pool)

Description: AKS default node pool

Type:

```hcl
object({
    name                   = string
    node_count             = number
    enable_host_encryption = optional(string)
    os_disk_size_gb        = optional(string)
    os_disk_type           = optional(string)
    os_sku                 = optional(string)
    ultra_ssd_enabled      = optional(bool)
    vm_size                = optional(string)
    pod_subnet_id          = optional(string)
    zones                  = list(string)
  })
```

Default:

```json
{
  "enable_host_encryption": false,
  "name": "dpool",
  "node_count": 2,
  "zones": [
    "1",
    "2",
    "3"
  ]
}
```

### <a name="input_dns_prefix"></a> [dns\_prefix](#input\_dns\_prefix)

Description: DNS prefix to deploy with kubernetes

Type: `string`

Default: `"internal"`

### <a name="input_enable_auto_scaling"></a> [enable\_auto\_scaling](#input\_enable\_auto\_scaling)

Description: Should the Kubernetes Auto Scaler be enabled for this Node Pool

Type: `bool`

Default: `false`

### <a name="input_enabled_argocd"></a> [enabled\_argocd](#input\_enabled\_argocd)

Description: ArgoCD

Type:

```hcl
object({
    name             = optional(string)
    repository       = optional(string)
    chart            = optional(string)
    version          = optional(string)
    namespace        = optional(string)
    installCRDs      = optional(bool)
    enabled          = bool
    create_namespace = optional(bool)

  })
```

Default:

```json
{
  "chart": "argo-cd",
  "create_namespace": true,
  "enabled": true,
  "name": "argocd"
}
```

### <a name="input_enabled_ingress_controller"></a> [enabled\_ingress\_controller](#input\_enabled\_ingress\_controller)

Description: Enable to install nginx ingress controller

Type:

```hcl
object({
    name       = string
    repository = string
    chart      = string
    enabled    = bool

  })
```

Default:

```json
{
  "chart": "nginx-ingress-controller",
  "enabled": false,
  "name": "nginx-ingress-controller",
  "repository": "https://charts.bitnami.com/bitnami"
}
```

### <a name="input_law_sku"></a> [law\_sku](#input\_law\_sku)

Description: Specifies the SKU of the Log Analytics Workspace

Type: `string`

Default: `"PerGB2018"`

### <a name="input_location"></a> [location](#input\_location)

Description: The Azure region to deploy resources

Type: `string`

Default: `"East US"`

### <a name="input_name"></a> [name](#input\_name)

Description: Name of Azure resource

Type: `string`

Default: `"tomelin"`

### <a name="input_name_id"></a> [name\_id](#input\_name\_id)

Description: ID to generate a different resources

Type: `number`

Default: `3`

### <a name="input_network_profile"></a> [network\_profile](#input\_network\_profile)

Description: Network profile kubernetes

Type:

```hcl
object({
    load_balancer_sku = optional(string)
    network_plugin    = optional(string)
    network_policy    = optional(string)
    outbound_type     = optional(string)
    pod_cidr          = optional(string)
    service_cidr      = optional(string)
    dns_service_ip    = optional(string)
  })
```

Default: `{}`

### <a name="input_node_pool"></a> [node\_pool](#input\_node\_pool)

Description: AKS node pool

Type:

```hcl
list(object({
    name                   = string
    node_count             = number
    enable_host_encryption = optional(string)
    os_disk_size_gb        = optional(string)
    os_disk_type           = optional(string)
    os_sku                 = optional(string)
    ultra_ssd_enabled      = optional(bool)
    vm_size                = optional(string)
    pod_subnet_id          = optional(string)
    zones                  = list(string)
    tags                   = optional(map(string))
  }))
```

Default:

```json
[
  {
    "name": "internal",
    "node_count": 1,
    "tags": {
      "Environment": "Production"
    },
    "vm_size": "Standard_DS2_v2",
    "zones": [
      "1",
      "2",
      "3"
    ]
  },
  {
    "name": "client",
    "node_count": 1,
    "tags": {
      "Environment": "Development"
    },
    "vm_size": "Standard_DS2_v2",
    "zones": [
      "1",
      "2",
      "3"
    ]
  }
]
```

### <a name="input_private_domain"></a> [private\_domain](#input\_private\_domain)

Description: Name to create a private dns zone

Type: `string`

Default: `""`

### <a name="input_retention_in_days"></a> [retention\_in\_days](#input\_retention\_in\_days)

Description: Number of days to retention the logs

Type: `number`

Default: `30`

### <a name="input_scale_down_mode"></a> [scale\_down\_mode](#input\_scale\_down\_mode)

Description: Specifies the autoscaling behaviour of the Kubernetes Cluster

Type: `string`

Default: `"Deallocate"`

### <a name="input_tags"></a> [tags](#input\_tags)

Description: List of tags to resource group.

Type: `map(string)`

Default: `{}`

### <a name="input_ultra_ssd_enabled"></a> [ultra\_ssd\_enabled](#input\_ultra\_ssd\_enabled)

Description: Used to specify whether the UltraSSD is enabled in the Default Node Pool.

Type: `bool`

Default: `false`

## Outputs

The following outputs are exported:

### <a name="output_aks"></a> [aks](#output\_aks)

Description: Created a kubernetes service

### <a name="output_created"></a> [created](#output\_created)

Description: Created time

### <a name="output_identity"></a> [identity](#output\_identity)

Description: Created a identity

### <a name="output_ingress"></a> [ingress](#output\_ingress)

Description: Values about the ingress controller

### <a name="output_kube_config"></a> [kube\_config](#output\_kube\_config)

Description: n/a

### <a name="output_log_analytics"></a> [log\_analytics](#output\_log\_analytics)

Description: Created a log analytics

### <a name="output_private_dns"></a> [private\_dns](#output\_private\_dns)

Description: Created a private dns

### <a name="output_resource_group"></a> [resource\_group](#output\_resource\_group)

Description: Created a resource group

### <a name="output_tags"></a> [tags](#output\_tags)

Description: Tags to created
