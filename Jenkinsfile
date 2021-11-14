pipeline {
    agent {
        docker { image 'node:14-alpine' } 
    }
    triggers { pollSCM('* * * * *') }
    stages {
        stage('Test') {
            steps {
                sh 'node --version'
                sh 'echo test'
            }
        }
        stage('Build image') {
            // If you have multiple Dockerfiles in your Project, use this:
            // app = docker.build("my-ubuntu-base", "-f Dockerfile.base .")

            app = docker.build("test")
        }

        stage('Test image') {
            app.inside {
                sh 'echo "Tests passed"'
            }
        }
    }
}