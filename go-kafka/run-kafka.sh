#!/bin/bash

# Iniciar o Zookeeper
echo "Iniciando o Zookeeper..."
./bin/zookeeper-server-start.sh config/zookeeper.properties &

# Aguardar um breve momento para garantir que o Zookeeper esteja iniciado antes de iniciar o Kafka
sleep 5

# Iniciar o Kafka
echo "Iniciando o Kafka..."
./bin/kafka-server-start.sh config/server.properties