#!/bin/bash
CGO_ENABLED=0 GOOS=linux go build -o bin/k8sClientListPodsImages.linux-amd64  .

