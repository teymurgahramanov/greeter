pipeline {
    options { timestamps() }
    triggers { pollSCM('* * * * *') }
    environment {
        imageName = 'teymurgahramanov/greeter'
        imageTag = 'latest'
        registryCred = 'dockerhub-teymurgahramanov'
    }
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
        stage('dockerize') {
            steps {
                script {
                    node {
                        def image
                        checkout scm
                        stage('build_image') {
                            image = docker.build("${imageName}")
                        }
                        stage('test_image') {
                            image.inside {
                                sh 'curl http://localhost:8080'
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