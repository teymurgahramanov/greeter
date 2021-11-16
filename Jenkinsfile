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
            agent { docker { reuseNode true image 'golang' } }
            steps {
                sh 'cd ${GOPATH}/src'
                sh 'mkdir -p ${GOPATH}/src/${JOB_NAME}'
                sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/${JOB_NAME}'
                sh 'go build -o greeter'
            }
        }
        stage('test') {
            steps {
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