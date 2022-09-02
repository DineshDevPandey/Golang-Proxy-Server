pipeline {
    agent any 
    tools {
        go '1.19'
    }
    parameters{choice(choices: ['a', 'b', 'c'], description: 'choice', name: 'ENV')}
    stages {
        stage('Running Build...') {
            steps {
                echo 'Hello, Build'
                sh 'go build main/proxy.go'
            }
        }
        stage('Running Test...') {
            steps {
                echo 'Hello, Test'
                sh 'cd main; go test'
            }
        }
        stage('Running Deploy...') {
            steps {
                echo 'Hello, Deploy'
            }
        }
    }
}
