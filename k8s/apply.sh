#!/bin/bash

echo "Creando namespace..."
kubectl apply -f namespace.yaml

echo "Creando Persistent Volume y Claim..."
kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml

echo "Desplegando MongoDB..."
kubectl apply -f mongo-deployment.yaml
kubectl apply -f mongo-service.yaml

echo "Desplegando microservicio p-go-create..."
kubectl apply -f create-deployment.yaml
kubectl apply -f create-service.yaml
kubectl apply -f create-pod.yaml
kubectl apply -f create-replicaset.yaml

echo "Configurando Ingress..."
kubectl apply -f ingress.yaml
