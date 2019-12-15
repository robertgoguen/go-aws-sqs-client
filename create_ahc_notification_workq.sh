#!/bin/bash

/usr/bin/aws --endpoint-url=http://localhost:4576 sqs create-queue --queue-name ahc-notification-work-queue
