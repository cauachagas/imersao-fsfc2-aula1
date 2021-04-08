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

Acesse container `kafka_kafka_1 bash` em e rode no primeiro terminal, que chamaremos de `producer`

```bash
kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
```

No segundo terminal, que chamaremos de `consumer`

```bash
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal
```

Acesse em um terceiro terminal o container `simulator` e rode

```bash
go run main.go
```

Agora a simulação dos pedidos está sendo feita

Para imprimirmos as informação dos pedidos (posição e se o pedido chegou ao destino), deveremos, após essas etapas, digitar no terminal intitulado `producer`

```json
{"clientId":"1","routeId":"1"}
{"clientId":"2","routeId":"2"}
{"clientId":"3","routeId":"3"}
```

No terminal intitulado `consumer` teremos as seguintes respostas

![](https://media.giphy.com/media/ozpMfaPuDHkYGaZySl/giphy.gif)
