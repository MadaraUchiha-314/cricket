# cricket
Cricketing As A Service (CAAS)

## Components

### k8s
Everything is deployed as a k8s cluster.

### Local development

#### Pre-requisites (installation)
- Using [podman](https://podman.io/)
- Using [minikube](https://minikube.sigs.k8s.io/docs/)
- Using [kubectl](https://kubernetes.io/docs/reference/kubectl/)

#### Running locally

```sh
# Initialize podman (needed only once)
podman machine init
# Start podman
podman machine start
# Run minikube using podman as the driver
minikube start --driver podman --container-runtime=cri-o
# Make local podman images discoverable
eval $(minikube podman-env)
kubectl get po -A
minikube dashboard
```

### svc
The micro-service which powers the REST APIs behind Cricketing As A Service.

#### Building

```sh
cd svc
# Build image
podman build -t cricket-svc-image -f ./Containerfile
# Check image has been created
# Should see an entry like: localhost/cricket-svc-image
podman images
# Deploy image as a container on a pod on k8s
kubectl apply -f ./deployment/deployment-local.yaml
# Verify
kubectl get pods
```

### data-store (TBD)
A data store which stores all the raw data which is consumed by other processes to injest into database.