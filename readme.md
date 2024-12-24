# argo ci/cd

## Argo CD

Argo CD is a declarative, GitOps continuous delivery tool for Kubernetes.

### Installation

```shell script
kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

# argo workflow
ARGO_WORKFLOWS_VERSION="v3.6.2"
kubectl create namespace argo
kubectl apply -n argo -f "https://github.com/argoproj/argo-workflows/releases/download/${ARGO_WORKFLOWS_VERSION}/quick-start-minimal.yaml"

# argo events
kubectl create namespace argo-events
kubectl apply -f https://raw.githubusercontent.com/argoproj/argo-events/stable/manifests/install.yaml
# Install with a validating admission controller
kubectl apply -f https://raw.githubusercontent.com/argoproj/argo-events/stable/manifests/install-validating-webhook.yaml

# argo events eventbus
kubectl apply -n argo-events -f https://raw.githubusercontent.com/argoproj/argo-events/stable/examples/eventbus/native.yaml

 # sensor rbac
kubectl apply -n argo-events -f https://raw.githubusercontent.com/argoproj/argo-events/master/examples/rbac/sensor-rbac.yaml
 # workflow rbac
kubectl apply -n argo-events -f https://raw.githubusercontent.com/argoproj/argo-events/master/examples/rbac/workflow-rbac.yaml
```

#### result

```shell script
root@master:/home/eilinge/argo-cd/events# kubectl -n argo-events get pod
NAME                                         READY   STATUS      RESTARTS         AGE
controller-manager-666764f7b8-phvh7          1/1     Running     0                5h19m
eventbus-default-stan-0                      2/2     Running     30 (5h23m ago)   4d22h
eventbus-default-stan-1                      2/2     Running     30 (5h23m ago)   4d22h
eventbus-default-stan-2                      2/2     Running     32 (5h23m ago)   4d22h
events-webhook-54d6d574d7-8k7g7              1/1     Running     0                5h17m

root@master:/home/eilinge/argo-cd/events# kubectl -n argo get pod
NAME                                     READY   STATUS      RESTARTS      AGE
argo-server-67bfcbc559-bxqwd             1/1     Running     3 (28h ago)   8d
workflow-controller-b84cc4f5b-fg5ss      1/1     Running     0             5h20m

root@master:/home/eilinge/argo-cd/events# kubectl -n argocd get pod
NAME                                                READY   STATUS    RESTARTS        AGE
argocd-application-controller-0                     1/1     Running   2 (28h ago)     10d
argocd-applicationset-controller-684cd5f5cc-h78fl   1/1     Running   2 (28h ago)     10d
argocd-dex-server-77c55fb54f-tgc2z                  1/1     Running   2 (28h ago)     10d
argocd-notifications-controller-69cd888b56-frrwd    1/1     Running   8 (5h23m ago)   4d4h
argocd-redis-855694d977-gmzmb                       1/1     Running   3 (28h ago)     10d
argocd-repo-server-584d45d88f-88hkp                 1/1     Running   5 (5h23m ago)   4d4h
argocd-server-8667f8577-whgwn                       1/1     Running   4 (5h24m ago)   4d4h
```


#### Argo CD UI


### Uninstall

```shell script
kubectl delete ns argocd
kubectl delete ns argo
kubectl delete ns argo-events
```
