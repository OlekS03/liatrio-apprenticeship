# Liatrio Apprenticeship Interview Exercise

## Overview
This project is part of the **Liatrio Apprenticeship Interview Exercise**, showcasing the ability to **build**, **containerize**, **test**, and **deploy** a simple web application using modern DevOps tools and workflows.

The application is a **Golang web service** built with the [Fiber](https://gofiber.io/) framework. It exposes a single HTTP endpoint (`/`) that returns a **minified JSON** object containing:

- the developer’s name, and  
- a dynamically generated timestamp (in milliseconds).

Example response:
```json
{"message":"My name is Oleksandr Shestopalov","timestamp":1729460000000}
```

## Live Deployment
The latest deployed version of this app can be accessed at:
https://liatrio-app-76565816101.us-central1.run.app

Every time new changes are pushed to the main branch in this GitHub repository, the CI/CD pipeline automatically:

Builds a new Docker image.

Runs automated endpoint tests.

Pushes the image to Docker Hub.

Deploys the updated image to Google Cloud Run.

## Local Development
1. Clone the repository
```
git clone https://github.com/OlekS03/liatrio-apprenticeship.git
cd liatrio-apprenticeship
```
2. Run locally
```
go run main.go
```
The app will start on port 80 (or another port if you set PORT):
```
PORT=8080 go run main.go
```
Then visit http://localhost:80 or your chosen port.

## Running with Docker
Build the image
```
docker build -t liatrio-app .
```
Run the container
```
docker run -p 80:80 liatrio-app
```
Test the endpoint
```
curl http://localhost:80/
```
Expected output:

```json
{"message":"My name is Oleksandr Shestopalov","timestamp":<current_time_in_ms>}
```

## CI/CD Workflow (GitHub Actions)
The GitHub Actions workflow automatically runs on each push to main.

Workflow Stages
Checkout code -> gets the latest source from GitHub.

Build Docker image -> compiles and containerizes the Go app.

Run Liatrio endpoint tests -> validates the / endpoint behavior.

Push to Docker Hub -> uploads tagged images.

Deploy to Google Cloud Run -> publishes the latest image automatically.

Workflow File
You can find it under:
```
.github/workflows/main.yml
```
## Forking & Using GitHub Secrets
If you fork this repository to your own GitHub account, you’ll need to set up secrets in your fork so the workflow can:

Log in to Docker Hub

Authenticate to Google Cloud

Deploy to Cloud Run

Secrets are not copied when you fork a repo.
You must add them manually.

### Required GitHub Secrets

| Secret Name         | Purpose                                      | Example |
|----------------------|----------------------------------------------|----------|
| `DOCKERHUB_USERNAME` | Your Docker Hub username                     | `olekshestopalov` |
| `DOCKERHUB_TOKEN`    | Docker Hub access token (for login/push)     | *(from Docker Hub -> Settings -> Security -> New Access Token)* |
| `GCP_SA_KEY`         | JSON key for your Google Cloud service account | *(Paste the full JSON from your GCP service account key)* |


## How to Add Secrets
Go to your forked repository on GitHub.

Click Settings -> Secrets and variables -> Actions.

Click New repository secret.

Add each secret above with the correct name and value.

Once added, your forked project will build, test, and deploy correctly when you push to main.

## Setting Up Google Cloud (for your own fork)
To deploy on your own GCP project:

Create a new project in Google Cloud Console.

Enable Cloud Run API and Artifact Registry API.

Create a Service Account with these roles:

roles/run.admin

roles/iam.serviceAccountUser

roles/storage.admin

Create a JSON key for that service account.

Copy the key’s contents and paste them as your `GCP_SA_KEY` secret in GitHub.
