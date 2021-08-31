# k8s-helloworld

[![Go](https://github.com/machine-brest/k8s-helloworld/actions/workflows/go.yml/badge.svg)](https://github.com/machine-brest/k8s-helloworld/actions/workflows/go.yml)

Sample App for CI/CD Demo

### Development

```bash
minikube start

# start
skaffold debug -f skaffold.yaml \
  --port-forward=true \
  --wait-for-deletions-max=2m0s \
  --status-check=true \
  --verbosity warn
```
