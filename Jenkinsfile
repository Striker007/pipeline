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
        archiveArtifacts(artifacts: 'main', caseSensitive: true, fingerprint: true, onlyIfSuccessful: true)
      }
    }
  }
}