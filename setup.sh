#!/bin/bash
echo "[Start.....]"
echo "[Build Mongo DB.....]"

docker build --tag mongo-db:1.0 mongodb/

echo "[Build GO .....]"

docker build --tag go-web:1.0 .

echo "[Run Images .....]"

docker-compose up -d

echo "[Done.....]"
