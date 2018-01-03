pipeline {
  agent any
  stages {
    stage('tests') {
      steps {
        sh '''ls -la .
make test'''
      }
    }
    stage('run') {
      steps {
        sh 'make'
      }
    }
  }
}