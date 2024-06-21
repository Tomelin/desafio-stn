# Deployment

Nesse ponto, foi criado dois diretórios, um com o gitops, para ter o project e application do ArgoCD e/ou FluxCD e outro com o manifest.  Referente ao manifest, eu recomendo criar um helm chart, para que tenhamos apenas um arquivos de values.yaml.

Deixarei o exemplo de manifests, apenas para nível de conhecimento., mas poderia ficar dentro do chart.


## Argo Project
Eu normalmente não crio o ArgoCD project junto com o application, pois dependente a implementação de pipeline, pode remover o ArgoCD project e tudo que está dentro.