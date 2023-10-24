

### Imagem dockerfile
1. librdkafka-dev= biblioteca que implementa o protocolo Apache Kafka fornecendo clientes produtores, consumidores e administradores;
2. 


### Gerar os arquivos do servico gRPC
```bash
protoc --go_out=application/grpc/pb --go_opt=paths=source_relative --go-grpc_out=application/grpc/pb --go-grpc_opt=paths=source_relative --proto_path=application/grpc/protofiles application/grpc/protofiles/*.proto
```

### Acessar o Evans para testar o gRPC
```bash
evans application/grpc/protofiles/pixkey.proto
```

### Command line interface (Cobra)

### inicializar 
```bash
cobra-cli init
```

### Adicionar cmd
```bash
cobra-cli add grpc
```

### Rodar
```bash
 go run main.go grpc
```

### Configurar comando no arquivo grpc.go 
## Rodar com porta customizada
```bash
go run main.go grpc -p 50052
```

