# Servidor Go

Este é um simples servidor web escrito em Go, que apresenta duas rotas:

- /: Esta rota apresenta um formulário HTML onde o usuário pode inserir seu nome.

- /bemvindo: Esta rota recebe o nome inserido pelo usuário no formulário e exibe uma mensagem de boas-vindas personalizada.

## Requisitos

```Certifique-se de ter Go instalado em sua máquina antes de executar este código.```

### Como executar

1. Clone este repositório para o seu ambiente local ou baixe o arquivo main.go.
2. Abra o terminal e navegue até o diretório onde o arquivo main.go está localizado.
3. Execute o seguinte comando para iniciar o servidor:

    ````go
    go run main.go
    ````

4. O servidor começará a ser executado em <http://localhost:8080>.

### Uso

1. Acesse <http://localhost:8080> em seu navegador.

2. Insira seu nome no formulário e envie.

3. Você será redirecionado para a página de boas-vindas, onde verá uma mensagem personalizada.

### Estrutura do código

- O código consiste em dois manipuladores de função principais:

  - ```formHandler```: Lida com requisições GET e POST para a rota raiz (/). Retorna um formulário HTML para GET e processa os dados do formulário para POST.

  - ```welcomeHandler```: Lida com a rota /bemvindo e exibe uma mensagem de boas-vindas personalizada, utilizando o nome fornecido pelo usuário.

### Observações

Se o usuário acessar diretamente <http://localhost:8080/bemvindo>, sem fornecer um nome na query string, ele será recebido como "Desconhecido".

Os arquivos HTML e CSS relacionados ao formulário não estão incluídos neste código, mas espera-se que estejam presentes e funcionais no mesmo diretório em que o servidor está sendo executado. Certifique-se de ter esses arquivos presentes para que o formulário funcione corretamente.
