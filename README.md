# API Web Básica em Go

Uma API web simples criada em Go usando o pacote `net/http`. Este projeto demonstra como criar endpoints, lidar com solicitações HTTP e retornar respostas em Go.

## Como Executar

1. Certifique-se de ter o Go instalado no seu sistema.

2. Clone o repositório:

```
git clone https://github.com/seu-usuario/api-web-basica-go
```

3. Navegue até o diretório do projeto:
```
cd api-web-basica-go
```

4. Execute o servidor:
```
go run main.go
```

5. O servidor estará rodando em http://localhost:8080.

- Endpoints
```
/welcome: Retorna uma mensagem de boas-vindas simples.
```

```
/form: Apresenta um formulário HTML onde você pode inserir um nome. Quando o formulário é enviado, ele retorna uma mensagem personalizada com o nome.
```

## Exemplo de Uso
- Acesse http://localhost:8080/welcome para ver a mensagem de boas-vindas.
- Acesse http://localhost:8080/form para acessar o formulário.

## Licença
Este projeto é licenciado sob a MIT License.