pipeline {
    agent any
   // agent {
   //    label 'my-agent'
    //}
    // We split the work into 3 stages:
   
    
    stages {
        // 1. Checkout the files from Git
        stage ('Git Checkout') {
            steps {
                checkout scm
            }
        }
        // 2. Check if 'my-code.c' exists, 
        stage ('Build') {
            steps {
                script {
                    sh 'pwd'
                    sh 'docker build -t go-hello-world:v1 .'
                }
            }
        }
        stage ('Push') {
            steps {
              script {
                        sh 'docker push -t charlesparasa/go-hello-world:v1 .'
                      }
                   }
         }
        // 3. Dummy deploy
        // Print a message (only done if the build is stable)
        stage ('Deploy') {
            steps {
                echo 'Deploying it gently...'
            }
        }
       
    }
   
}
