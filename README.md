# API da web básica em Go
Uma API da Web simples incorporada ao Go usando o pacote `net/http`.Este projeto demonstra como criar pontos de extremidade, lidar com solicitações HTTP e retornar respostas no GO.

## Como executar

### 1. Assegure-se de ter Go instalado no seu sistema.

### 2. Clone o repositório:

```
git clone https://github.com/ntsation/api_web.git
```

### 3. Navegue até o diretório do projeto:
```
cd api_web
```

### 4. Execute o servidor:
```
go run main.go
```

### 5. O servidor estará em execução em http://localhost:8080.

## Endpoints
```
/welcome: Retorna uma mensagem de boas-vindas simples.
```

```
/form: Apresenta um formulário HTML onde você pode inserir um nome.Quando o formulário é enviado, ele retorna uma mensagem personalizada com o nome.
```

## Exemplo de uso
- Visite http://localhost:8080/welcome para ver a mensagem de boas -vindas.
- Visite http://localhost:8080/form para acessar o formulário.

## Licença
Este projeto está licenciado sob a licença do MIT.