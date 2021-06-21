# Dependencies

* Minikube
    * https://minikube.sigs.k8s.io/docs/start/
- Helm
  - `curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash`
- Dapr
  - `wget -q https://raw.githubusercontent.com/dapr/cli/master/install/install.sh -O - | /bin/bash`

# Getting started

```bash
minikube start
eval $(minikube docker-env)
dapr init --kubernetes --wait

# https://github.com/bitnami/charts/tree/master/bitnami/rabbitmq/#installing-the-chart
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install rabbitmq bitnami/rabbitmq \
  --set auth.username=username \
  --set accessKey.forcePassword=true\
  --set auth.password=password \
  --set secretKey.forcePassword=true
```
# Setup

```bash
kubectl apply -f ./components
docker build -t examples/code-image .
kubectl apply -f ./items
# Open rabbitmq dashboard to view status
# kubectl port-forward --namespace default svc/rabbitmq 15672:15672
```