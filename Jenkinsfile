pipeline {
  agent any
  stages {
    stage('tests') {
      steps {
        sh 'make test'
      }
    }
    stage('build') {
      steps {
        sh 'make build'
        archiveArtifacts(artifacts: 'main', caseSensitive: true, fingerprint: true, onlyIfSuccessful: true)
      }
    }
  }
}