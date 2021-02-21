# htpasswd-crd-go
The following CRD/Operator is created with the purpose of create secrets objects with basic auth.

You can use them, for instance to authenticate users in your Kubernetes Ingress(NginX). 

Created with [kubebuilder](https://github.com/kubernetes-sigs/kubebuilder) framework.

## Example manifest
The has the following structure

```yaml
apiVersion: security.htpasswd-crd-go/v1
kind: Htpasswd
metadata:
  name: htpasswd-basic-auth
spec:
  user: admin
  password: admin123.
  namespace: default
```

## Get htpasswd
As CRD you can get them using kubectl and executing

```kubectl
kubectl get htpasswd
```

## Development
To install this CRD in your Kubernetes DEV cluster, just need to execute the following commands:
```make
make install
```
```make
make run
```

## Installation
To install this CRD in your Kubernetes cluster, just need to execute the following commands:
```make
make run
```


## Make useful commands
Building and deploying the Operator

```make
make run:- Run the on the default Kubernetes cluster
make install:- Install the CRD into the cluster
make uninstall:- Uninstall the CRD
make deploy:- Deploy the Operator into the cluster
make manifests:- Generate the YAML files
make generate:- Generate source codes
make docker-build:- Build the docker image
make docker-push:- Push the docker image to the specified registry

```

---
Jose Ramón Mañes