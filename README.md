# go-cloud-infra: A Go Web App with Complete DevOps Pipeline

This project demonstrates the implementation of DevOps practices for a Go web application- personal portfolio written in Golang and utilizing HTML/CSS. The focus lies on establishing a robust development and deployment workflow for the web app.

## Project Overview

The project leverages the following DevOps tools and technologies:

 - Docker for Containerization
 - Kubernetes and EKS
 - NGINX Ingress Controller   
 - Helm Chart
 - Git Actions for Continues Integration
 - ArgoCD for Continues Deployment

## Diagram




## Complete Flow to use the Project

### First, Run the Application Locally

Prerequisites: Golang installed

Command to Run the Go application:

```bash
go run main.go
```

Access the application:
Open http://localhost:8080/home in your web browser to view the portfolio website.

You will see the home page like this: 


### Create Dockerfile

The Dockerfile is used to build a Docker image. The Docker image contains the Go web application and its dependencies. The Docker image is then used to create a Docker container. The Dockerfile utilizes a multi-stage build:

  - Base stage: Installs necessary dependencies for building the Go application.

  - Distroless stage: Copies the compiled application binary and minimizes the image size by using a slim base image. This approach reduces the final image size and improves security by excluding unnecessary components.

### Containzerization

Containerization is the process of packaging an application and its dependencies into a container. The container is then run on a container platform such as Docker. Containerization allows you to run the application in a consistent environment, regardless of the underlying infrastructure.

Uses Docker to containerize the Go web application. Docker is a container platform that allows to build, ship, and run containers.

#### Commands to build the Docker container:
  
  ```bash
  docker build -t <your-docker-username>/go-web-app .
  ```
  for Example: 
  ```bash
  docker build -t urvishkumar0409/go-cloud-infra:v1 .
  ```
  


#### Command to run the Docker container:

```bash
docker run -p 8080:8080 <your-docker-username>/go-web-app
```
for Example: 
  ```bash
  docker run -p 8080:8080 -it urvishkumar0409/go-cloud-infra:v1
  ```
#### Command to push the Docker container to Docker Hub:

```bash
docker push <your-docker-username>/go-web-app
```
for Example: 
  ```bash
  docker push urvishkumar0409/go-cloud-infra:v1
  ```

### Kubernetes Deployment on EKS

#### Prerequisites:
  
  - kubectl (Command line tool for working with kubernetes clusters)
  - eksctl (To work with EKS clusters that automate many individuale task)
  - aws cli (Command line tool to use AWS services- authencticate with aws account using access key and seceret key to work with aws servicecs, including Amazon EKS)

#### Create an EKS cluster using eksctl.

```bash
eksctl cerate cluster --name demo-cluster --region us-east-1
```

#### Build Kubernetes manifests (deployment, service, and ingress YAML files).

```bash
kubectl apply -f k8s/manifests/deployment.yaml
```
```bash
kubectl apply -f k8s/manifests/service.yaml
```
```bash
kubectl apply -f k8s/manifests/ingress.yaml
```
```bash
kubectl get pods
```

Access the application externally

Obtain the External IP, Target Port from running pods to check the web app working locally with EKS.

Update your local driver's hosts file to map the IP to a desired domain name mentioned in service file for DNS Mapping.


### Helm Chart Management
Helm facilitates managing application configurations and deployments.

#### Helm Chart Structure:

 - chart.yaml: Defines the chart metadata.
 - templates: Contains deployment templates with placeholders for configuration values.
 - values.yaml: Provides default configuration values for the application.

Deployment and Uninstallation:

To install the application using Helm:  

```bash
helm install go-cloud-infra ./go-cloud-infra-chart
```
To uninstall the application: 

```bash
helm delete go-cloud-infra
```


### Continuous Integration & Delivery (CI/CD) Pipeline

#### CI with GitHub Actions:

The workflow triggers on pull requests and pushes to the main branch.
Builds the Docker image.
Runs tests on the application.
Pushes the image to a container registry on successful builds.

#### CD with Argo CD:

Argo CD leverages GitOps principles to synchronize the desired application state with the running Kubernetes cluster.
Upon successful CI pipeline completion, the updated Helm chart triggers a deployment in Argo CD.

## Summary
This project showcases a comprehensive DevOps approach for a Go web application. It demonstrates leveraging containerization, Kubernetes, Helm for package management, and CI/CD for automated builds and deployments.
