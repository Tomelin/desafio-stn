# Github Actions - CI/CD

## Container

Está feito o build do container e jogando no registry, conforme as variáveis definidas no github. 

## Terraform / deploy cloud

Aqui nesse item , também se faz necessário a configuração do Github secrets a variables, para poder fazer a conexão com  a Azure, depois disso faz o deploy.

## Deployment / kubernetes

Esse ponto está pendente fazer a integração (output terraform) pegando o kubeconfig, para fazer o deploy, sem a necessidade de configurar o secrets.

Nesse momento estou tendo problemas para passar o arquivo (sem ser artifacts) para o outro workflow.

## Geral

O workflow de container e terraform estão funcionando corretamente, pendente o de deployment para pegar o kube_config do terraform e aplicar os manifests do ArgoCD