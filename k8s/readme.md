To generate kubernetes manifest artifact for dev

kubectl kustomize overlays/dev > manifest-dev.yaml

To generate kubernetes manifest artifact for prod
kubectl kustomize overlays/prod > manifest-prod.yaml

if you dont want to save kubernetes manifest artifact
kubectl apply -k overlays/dev


