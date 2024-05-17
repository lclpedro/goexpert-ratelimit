## Aplicação de Rate Limit

### Para execução: 
Você primeiramente deve definir as variáveis de ambientes no docker-compose

```
TOKEN_RATE_LIMIT: <Quantidade de Request para Chamadas com Token>
IP_RATE_LIMIT: <Quantidade de Request para Chamadas sem Token>
TOKENS_PERMITED: AAAAA|BBBBB <List dde tokens permitidos>
EXPIRATION_TIME: <Tempo de expiração em segundos do cache>
```

Após isso, basta executar o comando abaixo:

```
docker-compose up --build
```
### Tests
Para execução dos testes, basta usar o arquivo .http na raíz do projeto, e executar as chamadas
