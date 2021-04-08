# imersao-fsfc2-desafio1
Código com o desafio 1 da imersão Full Stack &amp; Full Cycle


## Subindo os containers

Na raiz e .docker/kafka

```bash
docker-compose up -d
```

## Acessando bash no container

```bash
docker exec -it simulator bash 
docker exec -it kafka_kafka_1 bash
```

## Visualizando mensagens de um tópico

Acesse o container `kafka_kafka_1 bash` e rode

```bash
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=readtest
```

Acesse em outro terminal o container `simulator` e rode

```bash
go run main.go
```