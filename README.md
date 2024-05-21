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
Para execução dos testes manuais, basta usar o arquivo .http na raíz do projeto, e executar as chamadas

Para execução dos testes automatizados basta executar o comando abaixo:

```
make test
```
Explicação dos testes:
- Teste de Rate Limit por IP: Testa se o rate limit por IP está funcionando corretamente

Para essa funcionalidade, dentro do arquivo de test existe uma automatização de 20 requests sem a utilização de token.
Com isso as 10 primeiras requests, o que é permitida pela variável de ambiente, são realizadas com sucesso.
As próximas 10 requests, que ultrapassam o limite, retornam um erro 429.

- Teste de Rate Limit por Token: Testa se o rate limit por Token está funcionando corretamente

Para essa funcionalidade, dentro do arquivo de test existe uma automatização de 200 requests com a utilização de um token que é permitido.
Com isso as 100 primeiras requests, o que é permitida pela variável de ambiente, são realizadas com sucesso.
As próximas 100 requests, que ultrapassam o limite do token, retornam um erro 429.

