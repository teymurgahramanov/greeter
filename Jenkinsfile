pipeline {
    environment {
        imageName = 'teymurgahramanov/greeter'
        registry = 'https://registry.hub.docker.com'
        registryCred = 'dockerhub-teymurgahramanov'
    }
    options { timestamps() }
    triggers { pollSCM('* * * * *') }
    agent { docker { reuseNode true image 'golang' } }
    stages {
        stage('build_code') {
            steps {
                slackSend color:"warning", message:"started ${env.JOB_NAME} ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
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
                            docker.withRegistry("${registry}","${registryCred}") {
                                image.push()
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
            slackSend color:"good", message:"Build deployed successfully - ${env.JOB_NAME} ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
        }
        failure {
            slackSend color:"danger", failOnError: true, message:"Build failed  - ${env.JOB_NAME} ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
        }
    }
}