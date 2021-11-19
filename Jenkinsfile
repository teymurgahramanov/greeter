pipeline {
    environment {
        registry = 'https://registry.hub.docker.com'
        registryCredId = "dockerhub-teymurgahramanov"
        slackTokenId = "slack-bot-token"
        slackChannel = "cicd"
        slackMessage = "Project: ${env.JOB_NAME} Build: ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
    }
    options {
        timestamps()
        disableConcurrentBuilds() 
    }
    triggers { pollSCM('* * * * *') }
    agent { label 'master' }
    stages {
        stage('build') {
            steps {
                script {
                    helmChart = readYaml file: "${WORKSPACE}/k8s/greeter/Chart.yaml"
                    helmValues = readYaml file: "${WORKSPACE}/k8s/greeter/values.yaml"
                    imageName =  "${helmValues.image.repository}"
                    imageTag = "${helmChart.appVersion}"
                    image = docker.build("${imageName}:${imageTag}")
                }
            }
        }
        stage('push') {
            steps {
                script {
                    docker.withRegistry("${registry}","${registryCredId}") {
                        image.push()
                        image.push('latest')
                    }
                }
            }
        }
        stage('deploy') {
            steps {
                script {
                    withKubeConfig([credentialsId: 'kubernetes-test']) {
                        sh "helm upgrade --install greeter ${WORKSPACE}/k8s/greeter"
                    }
                }
            }
        }
    }
}