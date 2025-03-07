#流程简析
#1.GitHub Webhook设置：在GitHub仓库设置webhook，指向Argo Sensor服务。
#2.监听Push事件：Argo Sensor监听GitHub push事件。
#3.触发Workflow：当检测到push事件时，Argo Sensor触发预定义的Argo Workflow。
#4.自动化构建与部署：Argo Workflow执行自动化构建和部署任务。
#5.访问应用：用户通过Ingress或LoadBalancer访问部署的应用。


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
kubectl apply -n argo -f https://github.com/argoproj/argo-workflows/releases/download/v3.6.4/install.yaml

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

root@master:/home/eilinge/argo-cd# kubectl -n argocd get svc
NAME                                      TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                      AGE
argocd-applicationset-controller          ClusterIP   10.43.24.111    <none>        7000/TCP,8080/TCP            11d
argocd-dex-server                         ClusterIP   10.43.40.214    <none>        5556/TCP,5557/TCP,5558/TCP   11d
argocd-metrics                            ClusterIP   10.43.73.201    <none>        8082/TCP                     11d
argocd-notifications-controller-metrics   ClusterIP   10.43.65.142    <none>        9001/TCP                     11d
argocd-redis                              ClusterIP   10.43.107.228   <none>        6379/TCP                     11d
argocd-repo-server                        ClusterIP   10.43.176.112   <none>        8081/TCP,8084/TCP            11d
argocd-server                             NodePort    10.43.238.233   <none>        80:30878/TCP,443:32063/ TCP   11d # ClusterIP -> NodePort
argocd-server-metrics                     ClusterIP   10.43.82.129    <none>        8083/TCP                     11d

# 获取argocd admin 密码
root@master:/home/eilinge/argo-cd# kubectl -n argocd get secret argocd-initial-admin-secret --output=jsonpath={.data.password} |base64 -d
```


#### Argo CD UI


### Uninstall

```shell script
kubectl delete ns argocd
kubectl delete ns argo
kubectl delete ns argo-events
```
