pipeline {
    options { timestamps() }
    triggers { pollSCM('* * * * *') }
    environment {
        image = ''
        imageName = 'teymurgahramanov/onewaymail'
        imageTag = 'latest'
        registryCred = 'dockerhub-teymurgahramanov'
    }
    agent none
    stages {
        stage('build') {
            agent { docker { image 'golang' reuseNode true } }
            steps {
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/${JOB_NAME}'
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/${JOB_NAME}'
                sh 'go build'
            }
        }
        stage('test') {
            image.inside {
                sh 'go clean -cache'
                sh 'go test ./... -v -short'  
            }
        }
    }
    post {
        always {
            echo "Slack Notification"
        }
    }
}