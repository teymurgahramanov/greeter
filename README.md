# Greeter 👋
Experimental project to train on building CI/CD pipeline. Stack includes Golang, Jenkins, Slack, Docker, Kubernetes and Helm. Feel free to make your contribution!

---

## Pipeline
Multistage pipeline with conditional events depending on branch name and ~~tag~~. Will send notifications to Slack on start,success,fail and abort. Deploy to cluster performs with Helm.

### Simple Go application
Just makes greeting when accessing page ¯\\__(ツ)_/¯. Displays user's X-Forwarded-For IP, hostname, and generates funny string consist of random adjectives and sciencist names (Used [namesgenerator](https://github.com/moby/moby/blob/master/pkg/namesgenerator/names-generator.go) package which used by Docker to generate random container names)

### Jenkins
Installed outside of K8S cluster using Docker Compose. Also installed [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/) and [helm](https://helm.sh/docs/intro/install/) binaries on the same server and exposed them to Jenkins container.
```
version: '3.7'
volumes:
  data:

services:
  jenkins:
    image: jenkins/jenkins:lts
    privileged: true
    user: jenkins # useradd jenkins -M -s /sbin/nologin usermod -aG docker jenkins
    environment:
      - TZ=Asia/Baku
    ports:
      - 8081:8080
      - 50000:50000
    container_name: jenkins
    volumes:
      - data:/var/jenkins_home
      - /var/run/docker.sock:/var/run/docker.sock
      - /usr/bin/docker:/usr/bin/docker
      - /usr/local/bin/kubectl:/usr/local/bin/kubectl
      - /usr/local/bin/helm:/usr/local/bin/helm
      - /etc/localtime:/etc/localtime:ro
```
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

