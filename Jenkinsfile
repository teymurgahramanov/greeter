pipeline {
    environment {
        imageName = 'teymurgahramanov/greeter'
        registry = 'https://registry.hub.docker.com'
        registryCredId = "dockerhub-teymurgahramanov"
        slackTokenId = "slack-bot-token"
        slackChannel = "cicd"
        slackMessage = "Project: ${env.JOB_NAME} Build: ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
    }
    options { timestamps() }
    triggers { pollSCM('* * * * *') }
    agent { docker { reuseNode true image 'golang' } }
    stages {
        stage('pre') {
            steps {
                slackSend tokenCredentialId:"${slackTokenId}", channel:"${slackChannel}", color:"warning", message:"🏁 Pipeline started – ${slackMessage}"
            }
        }
        stage('build_code') {
            steps {
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/${JOB_NAME}'
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/${JOB_NAME}'
                sh 'go build -o greeter'
            }
        }
        stage('test_code') {
            steps {
                sh 'go clean -cache'
                sh 'go test ./... -v -short'  
            }
        }
        stage('build_image') {                  
            steps {
                script {    
                    node {
                        checkout scm
                        if ( env.BRANCH_NAME == 'master' || 'main' || 'null' ) {
                            imageTag = 'latest'
                        } else {
                            imageTag = env.BRANCH_NAME
                        }
                        image = docker.build("${imageName}:v2")
                        stage('test_image') {
                            sh "docker network create ${JOB_NAME}"
                            docker.image("${imageName}").withRun("--name ${JOB_NAME} --net ${JOB_NAME}") { test ->
                                docker.image("curlimages/curl").inside("--net ${JOB_NAME}") { 
                                    sh 'curl http://${JOB_NAME}:8080'
                                }
                            }
                        }    
                        stage('push_image') {
                            docker.withRegistry("${registry}","${registryCredId}") {
                                image.push()
                                image.push('latest')
                            }
                        }
                        stage('deploy') {
                            withKubeConfig([credentialsId: 'kubernetes-test']) {
                                sh 'kubectl apply -f k8s.yaml'
                            }
                        }
                    }
                }
            }
        }
    }
    post {
        always {
            sh "docker system prune -af"
        }
        success {
            slackSend tokenCredentialId:"${slackTokenId}", channel:"${slackChannel}", color:"good", message:"👍 Pipeline finished successfully – ${slackMessage}"
        }
        failure {
            slackSend tokenCredentialId:"${slackTokenId}", channel:"${slackChannel}", color:"danger", message:"☠️ Pipeline failed – ${slackMessage}"
        }
    }
}