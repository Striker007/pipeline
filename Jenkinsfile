pipeline {
  agent any
  stages {
    stage('tests') {
      steps {
        sh 'make test'
      }
    }
    stage('run') {
      steps {
        sh 'make'
      }
    }
  }
}