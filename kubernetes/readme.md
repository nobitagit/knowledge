# Kubernetes

## Features

- Storage orchestration
- Automate rollouts/rollbacks (for instance, ensure a 0 downtime deployment by deploying an app without destroying the currently deployed one)
- Self-healing
- Management of secrets and configuration
- Horizontal scaling

## Benefits

On top the benefits of containers, such as:

- Get new devs up to speed up and running
- Environment consistency
- Ship to prod faster

K8s can:

- Orchestrate these containers
- Zero downtime deployments
- Self healing if a container fails
- Easy to scale
- Create an e2e test env
- Ensure application scales properly
- Ensure secrets and config are working properly for all containers
- Perf test

## Installation

Easiest way to have it locally is to use Docker Desktop and enable Kubernetes from there. See [here](https://www.techrepublic.com/article/how-to-add-kubernetes-support-to-docker-desktop/

## Tools

- kubectl: the K8s command line interface.
- web ui dashboard: https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/

To install the ui:

- `kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0-beta8/aio/deploy/recommended.yaml`
- `k describe secret -n kube-system | less` and copy the token of type `Type: kubernetes.io/service-account-token`
- `kubectl proxy`
- open http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login
- paste the token copied above
