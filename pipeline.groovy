pipeline {
    agent any
    
    parameters {
        string(name: 'godaddyApiSecret', defaultValue: 'Enter Go Daddy Api Secret', description: 'GoDaddy API Secret')
        string(name: 'godaddyApiKey', defaultValue: 'Enter Go Daddy Api Key', description: 'GoDaddy API Key')
        string(name: 'smtpUser', defaultValue: 'Enter SMTP User', description: 'SMTP User')
        string(name: 'smtpPassword', defaultValue: 'Enter SMTP Password', description: 'SMTP Password')
        string(name: 'ceebit_license_key', defaultValue: 'Enter Ceebit License Key', description: 'Ceebit License Key')
        string(name: 'ceebit_license_secret', defaultValue: 'Enter Ceebit License Secret', description: 'Ceebit License Secret')
        string(name: 'ceebit_code_user', defaultValue: 'Enter Ceebit Code User', description: 'Ceebit Code User')
        string(name: 'ceebit_code_password', defaultValue: 'Enter Ceebit Code Password', description: 'Ceebit Code Password')
    }
    stages {
        stages {
        stage('Clone Repository') {
            steps {
                echo "Cloning the code from GitHub..."
                git credentialsId: 'GIT_HUB_CREDENTIALS', url: 'https://github.com/SaurabhSK123/jenkins-k8s-ci-cd.git'
            }
        }
        stage('Build and Run') {
            steps {
                echo "Running Terraform..."
                sh "go run main.go -godaddyApiSecret ${params.godaddyApiSecret} -godaddyApiKey ${params.godaddyApiKey} -smtpUser ${params.smtpUser} -smtpPassword ${params.smtpPassword} -ceebit_license_key ${params.ceebit_license_key} -ceebit_license_secret ${params.ceebit_license_secret} -ceebit_code_user ${params.ceebit_code_user} -ceebit_code_password ${params.ceebit_code_password}"
                }
            }
        }
    }
}