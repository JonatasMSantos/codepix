# Projeto de Integração com Kafka, gRPC e Kubernetes

Este é um projeto que envolve várias tecnologias e componentes diferentes para criar uma aplicação de integração que utiliza Kafka, gRPC e Kubernetes. A estrutura do projeto inclui os seguintes componentes:

- `nestjs`: Projeto em NestJS que se comunica com o aplicativo em Go (`central-bank`) via gRPC e troca mensagens com o Kafka.
- `nextjs`: Frontend da aplicação.
- `k8s`: Scripts do Kubernetes para implantar os diferentes componentes do projeto em ambientes separados.
- `kafka`: Docker Compose com todas as dependências do Kafka para suportar a integração.
- `central-bank`: Aplicativo em Go, inicializado com Cobra, que fornece serviços acessíveis via gRPC.

## Requisitos

Antes de começar, você deve garantir que os seguintes requisitos estejam atendidos:

- Docker instalado e em execução.
- [Docker Compose](https://docs.docker.com/compose/install/) instalado.
- Kubernetes configurado, se você deseja implantar nos ambientes Kubernetes.
- Node.js e NPM instalados, se você deseja trabalhar no projeto `nextjs`.
- Go instalado, se você deseja trabalhar no projeto `central-bank`.

## Configuração do Kafka

Antes de executar o projeto, siga estas etapas:

1. Navegue até a pasta `kafka`:
```bash
cd kafka
```
2. Inicialize o ambiente Kafka com todas as dependências (ZooKeeper, Kafka, etc.):
```bash
docker-compose up --build
```

### Configuração do Projeto NestJS

O projeto NestJS se comunica com o central-bank via gRPC e troca mensagens com o Kafka. Certifique-se de que a configuração apropriada esteja em vigor, incluindo os arquivos de configuração do Kafka e gRPC.

### Inicialização do Projeto Central Bank (Go)
O projeto central-bank é um aplicativo Go que pode ser inicializado com o seguinte comando:

```bash
go run main.go all
```

> Certifique-se de que todas as dependências Go estejam instaladas antes de executar o projeto.

Dentro do arquivo `go.mod` apague todo o conteúdo de `require` e no terminal rode o comando abaixo para instalar todas as depêndencias:

`go mod init`

Com isso a aplicação deve funcionar corretamente.

### Execução do Projeto Next.js (Frontend)

> Navegue até a pasta nextjs:

```bash
cd nextjs
```

> Instale as dependências:

```bash
npm install
```

> Inicie o aplicativo Next.js:

```bash
npm run dev
```

### Implantação em Kubernetes

Se desejar implantar os componentes em Kubernetes, navegue até a pasta k8s e use os scripts fornecidos para implantar os serviços e recursos necessários.

Observação: Certifique-se de ter o Kubernetes configurado e os objetos necessários criados no cluster.