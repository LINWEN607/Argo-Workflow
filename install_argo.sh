
#!/bin/bash

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

#sa
kubectl create serviceaccount argoweb -n argo

#ClusterRole
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: argoweb-clusterrole-full-access
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["*"]
EOF

#ClusterRoleBinding
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: argoweb-clusterrolebinding-full-access
subjects:
- kind: ServiceAccount
  name: argoweb
  namespace: argo  # 确保命名空间正确
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: argoweb-clusterrole-full-access

  #TOKEN
  ARGO_TOKEN="Bearer $(kubectl get secret argoweb.service-account-token -n argo -o=jsonpath='{.data.token}' | base64 --decode)"
  echo $ARGO_TOKEN
