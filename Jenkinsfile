pipeline {
    environment {
        imageName = 'teymurgahramanov/greeter'
        imageTag = 'latest'
        registryCred = 'dockerhub-teymurgahramanov'
    }
    options { timestamps() }
    triggers { pollSCM('* * * * *') }
    agent { docker { reuseNode true image 'golang' } }
    stages {
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
                        def image
                        checkout scm
                        image = docker.build("${imageName}")
                        stage('test_image') {
                            docker.image("${imageName}").withRun("-n ${JOB_NAME}")
                            docker.image("curlimages/curl").inside {
                                sh 'curl http://${JOB_NAME}:8080'
                            }
                        }
                        stage('push_image') {
                            docker.withRegistry("${registryCred}") {
                                app.push("latest")
                            }
                        }
                    }
                }
            }
        }
    }
    post {
        always {
            echo "Slack Notification"
        }
    }
}