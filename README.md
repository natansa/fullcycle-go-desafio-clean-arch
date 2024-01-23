Para iniciar a aplicação, estando na raiz do projeto, execute:
1o: docker-compose up -d
2o: go mod tidy
3o: go run main.go wire_gen.go

O Web Server será executado na porta 8000.
O gRPC será executado na porta 50051.
O GraphQL será executado na porta 8080.

Na pasta api, crie algumas ordens com Send Request do arquivo create_order.http.

Na pasta api, liste as ordens que foram criadas com Send Request do arquivo list_order.http.

Realize os mesmos procedimentos com o gRPC e o GraphQL.