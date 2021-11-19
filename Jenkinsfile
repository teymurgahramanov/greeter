def NotifyOnSlack(token,channel,color,message) {
    slackSend tokenCredentialId:token, channel:channel, color:color, message:message
}

pipeline {
    environment {
        registry = 'https://registry.hub.docker.com'
        registryCredId = "dockerhub-teymurgahramanov"
        imageName =  "${helmValues.image.repository}"
        imageTag = "${helmChart.appVersion}"
        slackTokenId = "slack-bot-token"
        slackChannel = "cicd"
        slackMessage = "Project: ${env.JOB_NAME} Build: ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
    }
    options {
        timestamps()
        skipDefaultCheckout true
        disableConcurrentBuilds() 
    }
    triggers { pollSCM('* * * * *') }
    agent { label 'master' }
    stages {
        stage('build_image') {                  
            steps {
                script {                
                    NotifyOnSlack("${slackTokenId}","${slackChannel}","warning","🏁 Pipeline started – ${slackMessage}")
                    sh 'rm -rf *'
                    checkout scm
                    helmChart = readYaml file: "${WORKSPACE}/k8s/greeter/Chart.yaml"
                    helmValues = readYaml file: "${WORKSPACE}/k8s/greeter/values.yaml"
                    image = docker.build("${imageName}:${imageTag}")
                    stage('push_image') {
                        docker.withRegistry("${registry}","${registryCredId}") {
                            image.push()
                            image.push('latest')
                        }
                    }
                    stage('deploy') {
                        withKubeConfig([credentialsId: 'kubernetes-test']) {
                            sh "helm upgrade --install greeter ${WORKSPACE}/k8s/greeter"
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
            slackSend tokenCredentialId:"${slackTokenId}", channel:"${slackChannel}", color:"good", message:"👍 Pipeline finished successfully – ${slackMessage}"
        }
        failure {
            slackSend tokenCredentialId:"${slackTokenId}", channel:"${slackChannel}", color:"danger", message:"☠️ Pipeline failed – ${slackMessage}"
        }
        aborted {
            slackSend tokenCredentialId:"${slackTokenId}", channel:"${slackChannel}", color:"danger", message:"❕ Pipeline aborted – ${slackMessage}"
        }
    }
}