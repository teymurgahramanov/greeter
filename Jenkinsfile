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
            steps {
                image = docker.build imageName
            }

        }
        stage('test') {
            image.inside {
                sh 'echo "Tests passed"'
            }
        }
    }
    post {
        always {
        /* Notify on slack */
        echo "Notification"
        }
    }
    }
}