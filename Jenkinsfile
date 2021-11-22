void NotifyOnSlack(nToken,nChannel,nColor,nMessage) {
    slackSend tokenCredentialId:nToken, channel:nChannel, color:nColor, message:nMessage
}
pipeline {
    environment {
        mainBranch = 'main'
        registry = 'https://registry.hub.docker.com'
        registryCredId = "dockerhub-teymurgahramanov"
        slackTokenId = "slack-bot-token"
        slackChannel = "cicd"
        slackMessage = "Project: ${env.JOB_NAME} Build: ${env.BUILD_NUMBER} (<${env.BUILD_URL}|Open>)"
    }
    options {
        timestamps()
        disableConcurrentBuilds()
        skipDefaultCheckout(true)
    }
    triggers { pollSCM('* * * * *') }
    agent { label 'master' }
    stages {
        stage('build_image') {
            steps {
                script {
                    cleanWs()
                    checkout scm
                    result = sh (script: "git log -1 | grep '.*\\[ci_skip\\].*'", returnStatus: false)
                    if (result == 0) {
                        echo ("'ci skip' spotted in git commit. Aborting.")
                        success ("'ci skip' spotted in git commit. Aborting.")
                    }
                    NotifyOnSlack("${slackTokenId}","${slackChannel}","warning","🏁 Pipeline started – ${slackMessage}")
                    helmChart = readYaml file: "${WORKSPACE}/k8s/greeter/Chart.yaml"
                    helmValues = readYaml file: "${WORKSPACE}/k8s/greeter/values.yaml"
                    imageName =  "${helmValues.image.repository}"
                    if (env.BRANCH_NAME == "${mainBranch}") {
                        imageTag = "${helmChart.appVersion}"
                    } else {
                        imageTag = env.BRANCH_NAME
                    }
                    image = docker.build("${imageName}:${imageTag}")
                }
            }
        }
        stage('push_image') {
            steps {
                script {
                    docker.withRegistry("${registry}","${registryCredId}") {
                        image.push("${imageTag}")
                        if (env.BRANCH_NAME == "${mainBranch}") {
                            image.push("latest")
                        } 
                    }
                }
            }
        }
        stage('lint_chart') {
            steps {
                script {
                    withKubeConfig([credentialsId: 'kubernetes-test']) {
                        sh "helm lint ${WORKSPACE}/k8s/greeter"
                    }
                }
            }
        }
        stage('deploy_to_dev') {
            steps {
                script {
                    withKubeConfig([credentialsId: 'kubernetes-test']) {
                        sh "helm upgrade --install --set image.tag=${imageTag} greeter ${WORKSPACE}/k8s/greeter"
                    }
                }
            }
        }
        stage('deploy_to_prod') {
            when {
                branch 'main'  
            }
            steps {
                script {
                    withKubeConfig([credentialsId: 'kubernetes-prod']) {
                        sh "helm upgrade --install greeter ${WORKSPACE}/k8s/greeter"
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
            NotifyOnSlack("${slackTokenId}","${slackChannel}","green","👍 Pipeline finished successfully – ${slackMessage}")
        }
        failure {
            NotifyOnSlack("${slackTokenId}","${slackChannel}","danger","☠️ Pipeline failed – ${slackMessage}")
        }
        aborted {
            NotifyOnSlack("${slackTokenId}","${slackChannel}","danger","❕ Pipeline aborted – ${slackMessage}")
        }
    }
}