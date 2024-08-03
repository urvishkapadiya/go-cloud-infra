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




## DevOps Pipeline

### 1) First, Run the Application Locally

Prerequisites: Golang installed

Command to Run the Go application:

```bash
go run main.go
```

Access the application:
Open http://localhost:8080/home in your web browser to view the portfolio website.

You will see the home page like this: 

![homepage](https://github.com/user-attachments/assets/8567d52f-612d-4a7f-96c1-5902ab6a4afc)



### 2) Create Dockerfile

The Dockerfile is used to build a Docker image. The Docker image contains the Go web application and its dependencies. The Docker image is then used to create a Docker container. The Dockerfile utilizes a multi-stage build:

  - Base stage: Installs necessary dependencies for building the Go application.

  - Distroless stage: Copies the compiled application binary and minimizes the image size by using a slim base image. This approach reduces the final image size and improves security by excluding unnecessary components.



#### 3. Containerization

Containerization is the process of packaging an application and its dependencies into a container, enabling consistent and reliable deployment across various environments. Docker is a widely-used platform that facilitates building, shipping, and running containers.

**Steps to Containerize the Go Web Application:**

1. **Build the Docker Container:**

    The Dockerfile utilizes a multi-stage build process to create a streamlined and secure Docker image. The `base` stage installs necessary dependencies for building the Go application, while the `distroless` stage copies the compiled application binary and minimizes the image size by using a slim base image.

    ```bash
    docker build -t <your-docker-username>/go-web-app .
    ```

    Example:

    ```bash
    docker build -t urvishkumar0409/go-cloud-infra:v1 .
    ```

2. **Run the Docker Container:**

    Once the Docker image is built, you can run it locally to ensure it works as expected. The following command maps port 8080 on your local machine to port 8080 in the container:

    ```bash
    docker run -p 8080:8080 <your-docker-username>/go-web-app
    ```

    Example:

    ```bash
    docker run -p 8080:8080 -it urvishkumar0409/go-cloud-infra:v1
    ```

3. **Push the Docker Image to Docker Hub:**

    To make the Docker image accessible from anywhere, push it to Docker Hub. This step is crucial for deploying the container on cloud platforms or sharing it with others.

    ```bash
    docker push <your-docker-username>/go-web-app
    ```

    Example:

    ```bash
    docker push urvishkumar0409/go-cloud-infra:v1
    ```


### 4. Creating Kubernetes Manifests and EKS Configuration for Clustering

Containerized applications need orchestration tools for deployment and scaling. Kubernetes is a powerful orchestration tool that manages containerized applications, and Amazon Elastic Kubernetes Service (EKS) simplifies running Kubernetes on AWS without needing to install and operate your own Kubernetes control plane or nodes.

**Prerequisites:**

- **kubectl**: Command-line tool for interacting with Kubernetes clusters.
- **eksctl**: Command-line tool to create and manage EKS clusters.
- **AWS CLI**: Command-line tool to interact with AWS services, including Amazon EKS. Authenticate using AWS Access Key and Secret Key.

**Steps to Create Kubernetes Manifests and EKS Configuration:**

1. **Create an EKS Cluster:**

    After writing your Kubernetes manifests files, use `eksctl` to create an EKS cluster on AWS:

    ```bash
    eksctl create cluster --name demo-cluster --region us-east-1 
    ```

2. **Deploy Kubernetes Manifests:**

    Build and deploy the Kubernetes manifests, including deployment, service, and ingress YAML files:

    ```bash
    kubectl apply -f k8s/manifests/deployment.yaml
    ```

    ```bash
    kubectl apply -f k8s/manifests/service.yaml
    ```

    ```bash
    kubectl apply -f k8s/manifests/ingress.yaml
    ```

    Verify the deployed pods:

    ```bash
    kubectl get pods
    ```

3. **Optional: Test Externally Using NodePort:**

    Before configuring the NGINX controller to get the address of the Network Load Balancer, you can run the web app externally to confirm the successful creation of Kubernetes resources with EKS.

    **Note:** Edit the `service.yaml` file, changing the port type from `ClusterIP` to `NodePort`. Rebuild and apply the Kubernetes manifests:

    ```bash
    kubectl apply -f k8s/manifests/deployment.yaml
    ```

    ```bash
    kubectl apply -f k8s/manifests/service.yaml
    ```

    ```bash
    kubectl apply -f k8s/manifests/ingress.yaml
    ```

    Get the external IP address and target port from the running nodes:

    ```bash
    kubectl get svc
    ```

    ```bash
    kubectl get nodes -o wide
    ```

    Navigate to the application using the external IP address and port number:

    ```http
    http://YOUR-PODS-EXTERNAL-IP:TARGET-PORT-NUMBER/home
    ```

    Example:

    ```http
    http://34.161.26.159:32001/home
    ```

    If the home page is displayed, the web app is successfully running on the Kubernetes cluster. Revert the `service.yaml` file's port type back to `ClusterIP` and reapply the Kubernetes manifests.


### 5. NGINX Ingress Controller

The NGINX Ingress Controller is used to manage ingress resources in Kubernetes, creating a Network Load Balancer (NLB) to expose IP addresses for external access.

**Steps to Set Up NGINX Ingress Controller:**

1. **Install the NGINX Ingress Controller on AWS:**

    Deploy the NGINX Ingress Controller using the provided manifest file:

    ```bash
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.11.1/deploy/static/provider/aws/deploy.yaml
    ```

2. **DNS Mapping:**

    The NGINX Ingress Controller creates a Network Load Balancer for the running cluster nodes. Retrieve the IP address of the created load balancer to set up DNS mapping.

    ![dns_command](https://github.com/user-attachments/assets/af230f42-6b82-4ee5-a82f-90ddc378364d)

    Add the IP address to the hosts file for DNS mapping:

    - Path to hosts file on Windows: `C:\Windows\System32\drivers\etc\hosts`
    - Path to hosts file on macOS/Linux: `/etc/hosts`

    Add the following entry to the hosts file:

    ```
    <IP-ADDRESS> <MAPPING-NAME>
    ```

    Example entry:

    ```
    34.161.26.159 go-cloud-infra
    ```

    ![host_path](https://github.com/user-attachments/assets/529235fd-3d59-4d55-8406-dce4e7313cc6)

    ![dns_mapping](https://github.com/user-attachments/assets/cfaf4e06-ca60-4b73-8c92-a24c40d32885)

3. **Access the Application:**

    Navigate to the application using the mapped DNS name:

    ```http
    http://go-cloud-infra.local/home
    ```

    Confirm that the web app is working successfully with the created load balancer and EKS.


### 6. Helm Chart Management

Helm is a powerful tool that simplifies the management of Kubernetes applications by streamlining configuration and deployment processes.

#### Helm Chart Structure

A typical Helm chart includes the following components:

- **chart.yaml**: Contains metadata about the chart, such as its name, version, and description.
- **templates/**: Houses the deployment templates with placeholders for configuration values.
- **values.yaml**: Specifies default configuration values for the application.

To use Helm for managing your Kubernetes resources, follow these steps:

#### Delete Existing Kubernetes Components

Before deploying with Helm, remove any previously created Kubernetes components:

![delete_k8s](https://github.com/user-attachments/assets/0693295c-4b39-492d-9a8a-2c4cea9d4120)

#### Deployment and Uninstallation

1. **Deploy the Application Using Helm:**

    Run the following command to install the application using Helm:

    ```bash
    helm install go-cloud-infra ./go-cloud-infra-chart
    ```

2. **Uninstall the Application:**

    To uninstall the application, use the command:

    ```bash
    helm delete go-cloud-infra
    ```

#### Updating Kubernetes Manifests

After deploying with Helm, you may need to update the Kubernetes manifests (deployment, service, and ingress YAML files). Ensure that the tag values in `values.yaml` are properly updated to reflect these changes.



### 7. Continuous Integration & Delivery (CI/CD) Pipeline

#### Continuous Integration (CI) with GitHub Actions

Continuous Integration (CI) automates the integration of code changes into a shared repository. This practice helps catch bugs early in the development process and ensures that the code remains in a deployable state.

The CI workflow includes:

- Triggering on pull requests and pushes to the main branch.
- Building the Docker image.
- Running tests on the application.
- Pushing the image to a container registry on successful builds.
- updating helm chart

#### Continuous Deployment (CD) with Argo CD

Continuous Deployment (CD) automates the deployment of code changes to a production environment, reducing the time between code changes and their delivery to users. Argo CD is a declarative, GitOps continuous delivery tool for Kubernetes. It uses Git as the source of truth to deploy applications to Kubernetes clusters.

**Argo CD** leverages **GitOps principles** to synchronize the desired application state with the running Kubernetes cluster. Upon successful completion of the CI pipeline, the updated Helm chart triggers a deployment in Argo CD.

To configure Argo CD, follow these steps:

1. **Patch the Argo CD Server and Retrieve the Secret Key:**

    Use the following command to patch the Argo CD server and obtain the secret key for the 'admin' username:

    ![argoCD_command](https://github.com/user-attachments/assets/5db1cea2-79d4-413e-a9b6-eedfd0b455b3)

2. **Decrypt the Base64 Encrypted Password (Secret Key):**

    Use this command to decrypt the base64 encrypted password:

    ```bash
    [System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String("<ENCRYPTED-PASSWORD>"))
    ```

3. **Authenticate with Argo CD:**

    - Use the obtained Argo CD credentials to authenticate.
    - Create the application in Argo CD and link it with the CI and GitHub repository to automate the deployment.
    - On every successful CI, Argo CD will run the described steps to deploy the container and display the following chart:

    ![argocd_chart](https://github.com/user-attachments/assets/7606623f-4a7e-4f09-bca9-13730152507f)

An example of a successful CI/CD pipeline execution on every commit PR to the main branch can be seen [here](https://github.com/urvishkapadiya/go-cloud-infra/actions/runs/10189106603).

   
## Summary
This project showcases a comprehensive DevOps approach for a Go web application. It demonstrates leveraging containerization, Kubernetes, Helm for package management, and CI/CD for automated builds and deployments.
