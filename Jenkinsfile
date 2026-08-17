pipeline {
    agent {
        label 'ubuntu'
    }

    environment {
        APP_NAME    = 'go-web-app'
        APP_VERSION = "1.0.${BUILD_NUMBER}"
    }

    stages {
        stage('Checkout & Info') {
            steps {
                echo "=================================================="
                echo " Building Branch   : ${env.BRANCH_NAME}"
                echo " Job Name          : ${env.JOB_NAME}"
                echo " Build Number      : ${env.BUILD_NUMBER}"
                echo " Git Commit SHA    : ${env.GIT_COMMIT}"
                echo "=================================================="
                checkout scm
            }
        }

        stage('Compile & Build') {
            steps {
                echo "==> Compiling Go Application..."
                sh '''
                    go version
                    go build -v -o bin/${APP_NAME} main.go
                '''
            }
        }

        stage('Unit Tests') {
            steps {
                echo "==> Running Unit Tests & Generating Coverage..."
                sh '''
                    go test -v -cover -coverprofile=coverage.out ./...
                '''
            }
        }

        // =========================================================================
        // Stage 1: Runs ONLY on 'dev' branch (or feature branches if configured)
        // =========================================================================
        stage('Dev Deployment & Verification') {
            when {
                anyOf {
                    branch 'dev'
                    branch 'development'
                    branch 'feature/*'
                }
            }
            steps {
                echo "=================================================="
                echo " [DEV STAGE] Executing for Branch: ${env.BRANCH_NAME}"
                echo "=================================================="
                sh '''
                    echo "[DEV] Deploying to Dev/Staging Environment..."
                    echo "[DEV] App Version: ${APP_VERSION}"
                    echo "[DEV] Running Integration & Smoke tests on Dev..."
                    # Mock dev deployment step
                    export APP_ENV="dev"
                    echo "[DEV] Deployment successful to Dev cluster!"
                '''
            }
        }

        // =========================================================================
        // Stage 2: Runs ONLY on 'main' branch
        // =========================================================================
        stage('Production Deployment & Release') {
            when {
                anyOf {
                    branch 'main'
                    branch 'master'
                }
            }
            steps {
                echo "=================================================="
                echo " [PROD STAGE] Executing for Branch: ${env.BRANCH_NAME}"
                echo "=================================================="
                sh '''
                    echo "[PROD] Preparing Release Artifacts for ${APP_NAME}:${APP_VERSION}..."
                    echo "[PROD] Running Security scan and Production Checks..."
                    echo "[PROD] Deploying to Production Live Cluster..."
                    export APP_ENV="production"
                    echo "[PROD] Production deployment finished successfully!"
                '''
            }
        }
    }

    post {
        always {
            echo "Pipeline run completed for branch: ${env.BRANCH_NAME}"
            cleanWs deleteDirs: true, notFailBuild: true, patterns: [[pattern: 'bin/**', type: 'INCLUDE']]
        }
        success {
            echo "BUILD SUCCESSFUL: ${env.JOB_NAME} #${env.BUILD_NUMBER} (${env.BRANCH_NAME})"
        }
        failure {
            echo "BUILD FAILED: ${env.JOB_NAME} #${env.BUILD_NUMBER} (${env.BRANCH_NAME})"
        }
    }
}
