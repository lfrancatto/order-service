# Order Service (Go + Kafka + Outbox Pattern)

Projeto de estudo com foco em arquitetura limpa, event-driven e boas práticas de produção.

## 🚀 Stack

- Golang
- Apache Kafka
- PostgreSQL
- Goose (migrations)

## 🧱 Arquitetura

- Clean Architecture (DDD + Ports/Adapters)
- Outbox Pattern (consistência entre DB e Kafka)
- Producer Kafka idempotente

## 🔄 Fluxo

1. API recebe pedido
2. Salva no PostgreSQL (orders)
3. Salva evento na tabela outbox
4. Worker publica no Kafka
5. Evento é marcado como processado

## 📦 Estrutura
