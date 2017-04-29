#!groovy

properties([
    buildDiscarder(logRotator(daysToKeepStr: '20', numToKeepStr: '30')),

    [$class: 'GithubProjectProperty',
     projectUrlStr: 'https://github.com/euank/alexa-parrot'],

    pipelineTriggers([
      [$class: 'GhprbTrigger',
       adminlist: 'euank',
       cron: 'H/1 * * * *',
       permitAll: false,
       displayBuildErrorsOnDownstreamBuilds: true,
       gitHubAuthId: '894ea30e-d9bd-456c-a090-291b3c339ada']
    ])
])


node('docker') {
  stage('SCM') {
    checkout scm
  }

  stage('Build') {
    sh '''bash -ex
      docker run -v $(pwd):/go/src/github.com/euank/alexa-parrot \
        -w /go/src/github.com/euank/alexa-parrot \
        golang:1.8 \
        make
    '''
  }
}