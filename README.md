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

---
Jose Ramón Mañes