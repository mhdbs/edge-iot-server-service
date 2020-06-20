#!/bin/bash
set -eux
#For the first time build and deploy
docker build -t edge-iot-server-service .
docker image tag edge-iot-server-service localhost:5000/edge-iot-server-service
docker push localhost:5000/edge-iot-server-service
docker rmi localhost:5000/edge-iot-server-service
kubectl apply -f pods.yml
kubectl apply -f service.yml