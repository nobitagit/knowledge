.# Architecting with Google Kubernetes Engine

From this [course](https://app.pluralsight.com/player?course=architecting-google-kubernetes-engine-foundations&author=google-cloud&name=ee573538-7487-4b59-a338-c27fbd55fa14&clip=0&mode=live).

Kubernetes is **a software layer that sits between an application and the hardware infra**.
This abstraction layer makes managing the application easer.

## Cloud computing

- On demand: no human interaction needed to access resources
- Access from anywhere
- Providers set up a pool of resources and allocate chunks of it to the various customers
- Can scale up or down

## GCP solutions

### Compute Engine

Google's IaaS solution. You manage the server instance yourself.

### Kubernetes Engine (GKE)

Run cointainerised applications managed by Google and administered by yourself. You package your code in containers and Kubernetes operates as the orchestrator.

### App Engine

GCP PaaS. Run code in the cloud without worrying about the infra.
Google provisions and manages resources.

### Cloud Functions

Google's Lambda.

## Regions and zones

The World is divided in 3 multi-regional areas.

- America
- Europe
- Asia/Pacific

Every region is the divided into regions, which are independent geographic areas within the same continent.
Within a region there's fast network latency.
Every region is divided into zones, which are deployment areas. A zone is, in practice, a **datacenter** (although this is a simplification since 1:1 mapping zone:datacenter is not always guaranteed).

Deploying across multiple zones enables **fault tolerance** and **high availability**.

### Using and organising regions and zones

A project organises how resources and services are deployed.
Projects can be grouped into folders. A folder mirrors the structure of a company.
For instance:

```
              FOLDERS

 Department A        Department B
      |                    |
  |-------|            |-------|
  |       |            |       |
Team 1  Team 2       Team 3  Team 4
  |                            |
  |           PROJECTS         |
  |                            |
 App            Test Env  ---------- Production
```
