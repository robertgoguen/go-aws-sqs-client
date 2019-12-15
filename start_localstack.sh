#!/bin/bash

export AWS_DEFAULT_REGION=us-east-1
export AWS_REGION=us-east-1
export AWS_ACCESS_KEY_ID=foobar
export AWS_SECRET_ACCESS_KEY=foobar
export AWS_SDK_LOAD_CONFIG=1
export LOGLEVEL=debug

export SERVICES=s3

/usr/bin/docker run --rm --name localstack --rm -e SERVICES="s3" -p 4576:4576 -v "/tmp/localstack:/tmp/localstack" -e HOST_TMP_FOLDER="/tmp/localstack" localstack/localstack
