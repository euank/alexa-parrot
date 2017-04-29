#!groovy

properties([
    buildDiscarder(logRotator(daysToKeepStr: '20', numToKeepStr: '30')),

    [$class: 'GithubProjectProperty',
     projectUrlStr: 'https://github.com/euank/alexa-parrot'],

    pipelineTriggers([
      // Pull requests, with whitelisting/auth
      [$class: 'GhprbTrigger',
       adminlist: 'euank',
       cron: 'H/1 * * * *',
       permitAll: false,
       displayBuildErrorsOnDownstreamBuilds: true,
       gitHubAuthId: '894ea30e-d9bd-456c-a090-291b3c339ada'],

      // master branch
      pollSCM('H/15 * * * *')
    ])
])


node('docker') {
  stage('SCM') {
    checkout scm
  }

  stage('Build') {
    sh '''#!/bin/bash -ex
      docker run -v $(pwd):/go/src/github.com/euank/alexa-parrot \
        -w /go/src/github.com/euank/alexa-parrot \
        golang:1.8 \
        make
    '''

    if(env.BRANCH_NAME == 'master') {
      withCredentials([usernamePassword(credentialsId: 'quay-alexaparrot', 
                      passwordVariable: 'DOCKER_PASS', 
                      usernameVariable: 'DOCKER_USER')]) {
        sh '''#!/bin/bash -ex
          export HOME=${WORKSPACE}
          docker login --username "${DOCKER_USER}" --password "${DOCKER_PASS}"
          make docker-push
        '''
      }
    }
  }
}
