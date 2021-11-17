def imageName = 'teymurgahramanov/greeter'
def registry = 'https://registry.hub.docker.com'
def registryCredId = 'dockerhub-teymurgahramanov'
def slackTokenId = 'slack-bot-token'
def slackChannel = '#cicd'
def slackMessageStart = "🏁 Pipeline started – Project: ${env.JOB_NAME} Build: ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
def slackMessageSuccess = "👍 Pipeline finished successfully – Project: ${env.JOB_NAME} Build: ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
def slackMessageFailure = "☠️ Pipeline failed – Project: ${env.JOB_NAME} Build: ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
def slackNotify(caseName,caseMessage) {
    slackSend token:slackTokenId, channel:slackChannel, color:caseName, message:caseMessage
}
pipeline {
    /*
    environment {
        imageName = 'teymurgahramanov/greeter'
        registry = 'https://registry.hub.docker.com'
        registryCredId = 'dockerhub-teymurgahramanov'
        slackTokenId = ''
        slackeChannel = ''
        slackMessage = "Project: ${env.JOB_NAME} Build: ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
    }
    */
    options { timestamps() }
    triggers { pollSCM('* * * * *') }
    agent { docker { reuseNode true image 'golang' } }
    stages {
        stage('pre') {
            steps {
                slackNotify("warning", slackMessageStart)
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
                        image = docker.build("${imageName}:${imageTag}")
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
                            }
                        }
                        stage('deploy') {
                            withKubeConfig([credentialsId: 'kubernetes-test', serverUrl: 'https://192.168.120.207:6443/']) {
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
            slackSend color:"good", message:"👍 Pipeline finished successfully – ${slackMessage}"
        }
        failure {
            slackSend color:"danger", failOnError: true, message:"☠️ Pipeline failed – ${slackMessage}"
        }
    }
}