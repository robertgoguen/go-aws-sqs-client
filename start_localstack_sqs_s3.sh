#!/bin/bash

export AWS_DEFAULT_REGION=us-east-1
export AWS_REGION=us-east-1
export AWS_ACCESS_KEY_ID=foobar
export AWS_SECRET_ACCESS_KEY=foobar
export AWS_SDK_LOAD_CONFIG=1
export LOGLEVEL=debug

export SERVICES="sqs,s3"

# start_localstack_sqs_s3.sh 
#
# Example:  ./start_localstack_sqs_s3.sh
# 
# Purpose: This script is intended to be used to do *manual* end-to-end Unit Testing.
#          It will perform the step of starting the localstack AWS emulator, and create the
#          SQS and S3 services.  After running this script, the user
#          can then run the ./create_ahc_s3_bucket_workq_item.sh script to create AWS resources.

/usr/bin/docker run --rm --name localstack --rm -e SERVICES="sqs,s3" -p 4576:4576 -p 4572:4572 -v "/tmp/localstack:/tmp/localstack" -e HOST_TMP_FOLDER="/tmp/localstack" localstack/localstack
