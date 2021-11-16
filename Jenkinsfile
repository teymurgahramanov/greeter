pipeline {
    environment {
        imageName = 'teymurgahramanov/greeter'
        registry = 'https://hub.docker.com'
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
                        if ( env.BRANCH_NAME == 'master' || env.BRANCH_NAME == 'main' ) {
                            imageTag = 'latest'
                        }
                        else {
                            imageTag = env.BRANCH_NAME
                        }                         
                        def image
                        checkout scm
                        image = docker.build("${imageName}")
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
                                image.push("${imageTag}")
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
            sh "docker network rm ${JOB_NAME}"
        }
    }
}