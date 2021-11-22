# Greeter 👋
Experimental project to train on building CI/CD pipeline. Stack includes Golang, Jenkins, Slack, Docker, Kubernetes and Helm. Feel free to make your contribution!

---

## Pipeline
- Multibranch
- Conditional events depending on branch name and commit message.
- Slack notifications

### Simple Go application
Just makes greeting when accessing page ¯\\__(ツ)_/¯. Displays user's X-Forwarded-For IP, hostname, and generates funny string consist of random adjectives and sciencist names (Used [namesgenerator](https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go) package which used by Docker to generate random container names)

### Jenkins
Installed outside of K8S cluster using Docker Compose. Also installed [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/) and [helm](https://helm.sh/docs/intro/install/) binaries on the same server and exposed them to Jenkins container.

Pipeline fully described in Jenkinsfile no configurations need on server side. \
\
Used plugins:
- Docker Pipeline
- Kubernetes CLI
- Slack Notification

Used credentials:
- Docker Hub (Username/Password)
- Kubernetes config file (Secret file)
- Slack bot token (Secret text)

### Helm and Kubernetes
- Custom values and templates
