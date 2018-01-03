pipeline {
  agent any
  stages {
    stage('tests') {
      steps {
        sh 'cmake test'
      }
    }
    stage('run') {
      steps {
        sh 'cmake'
      }
    }
  }
}