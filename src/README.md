# Micro serviço para integração com o Kubernetes

## Estudos

Algumas funcionalidades foram baseadas nos estudos que fiz a algum tempo atrás
https://github.com/Tomelin/kubernetes-controllers-controller-expose
https://github.com/Tomelin/golang-kubernetes/

## Do funcionamento

O micro serviço conecta no Kubernetes atráves do kube_config ou rest.InClusterConfig, quando se está dentro do cluster.   Se for usado o rest.InClusterConfig, precisamos aplicar o service account, role e role binding.  Essas configuraÇoes do kubernetes serão aplicadas através dos manifestos via ArgoCD.

## REST

Tem o swagger para poder validar os paths e está semi pronto o middleware para aceitar autenticação, pois nesse momento ficou aberta as requests (sem autenticação)

Foram criados paths, para listar, count, filter por nome e criação de forma básica, validando nesse primeiro momento, apenas o nome do object.

Como serviço de http, foi implementado o GIN.

## Config

As configurações são lidas através do viper, passando o path do arquivo de config.  por default é em /app/config.

A variável que identifica se tem a config em diretório diferente é: PATH_CONFIG

Essa config, pode ser gerada de forma automatica, através do hashicorp vault, colocando um init container e mapeando o volume e colocando o arquivo dentro, depois o container mapeia o volume com o arquivo, podendo mapear apenas como leitura.

Outra opção, mas naào implementada, é o micro serviço conectar diretamente no hashicorp vault, usando o package "github.com/hashicorp/vault/api", nesse caso, o ideal é criar uma biblioteca para o mesmo.

## healthz

Não foi implementado o /healthz, por isso o mesmo tb não será implementado no deploymentt. Porém the best prectices, recomenda fortemente criar o servição de valid'ão do serviço está up ou não

## Makefile

Foi criado o Makefile, para apoiar no dia-a-adia, durante o desenvolvimento.

## Soluções

Outra solução que poeria ser adotada de forma mais simples e rápida, era criar o proketo através do kubebuilder, onde cria toda a strutura de conexão com o kubernetes, ficando apenas a lógica de publicação de api a ser implementada.