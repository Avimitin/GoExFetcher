pipeline {
	agent {
		docker {
			image 'golang:1.15-buster'
		}
	}
	stages{
		stage('Setup') {
			step{
				sh 'go get ./...'
			}
		}

		stage('Test') {
			step {
				sh 'go test -v ./...'
			}
		}

		stage('Build') {
			step{
				sh 'echo "building"'
				sh '''
				go build -ldflags '-w -s' -o bin/osuAPI-linux-amd64
				tar -C bin -zcvf osuAPI-linux-amd64.tar.gz osuAPI
				'''
			}
		}
	}
}
