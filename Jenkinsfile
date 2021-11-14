pipeline {
    agent {
        docker { image 'node:14-alpine' } 
    }
    triggers { pollSCM('* * * * *') }
    stages {
        stage('Test') {
            steps {
                sh 'node --version'
            }
        }
    }
}