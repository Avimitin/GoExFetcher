pipeline {
	agent {
		docker {
			image 'golang:1.15-buster'
		}
	}

	environment{
		XDG_CACHE_HOME='/tmp/.cache'
	}

	stages{
		stage('Setup') {
			steps{
				sh 'go get ./...'
			}
		}

		stage('Test') {
			steps {
				sh 'go test -v ./...'
			}
		}

		stage('Build') {
			steps{
				sh 'echo "building"'
				sh '''
				go build -ldflags '-w -s' -o bin/osuAPI-linux-amd64
				tar -C bin -zcvf osuAPI-linux-amd64.tar.gz osuAPI
				'''
			}
		}
	}
}
