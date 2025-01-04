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
# Run minikube using podman as the driver
minikube start --driver podman
kukubectl get po -A
miminikube dashboard
```

### svc
The micro-service which powers the REST APIs behind Cricketing As A Service.

#### Building

```sh
cd svc
# Build image
podman build -t cricket-svc-image .
# Run container
podman run -p 3000:3000 cricket-svc-image
```

### data-store (TBD)
A data store which stores all the raw data which is consumed by other processes to injest into database.