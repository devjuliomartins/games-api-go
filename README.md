# API de Jogos
## Requisitos

Antes de rodar o projeto, você precisa ter os seguintes softwares instalados na sua máquina:

- [Go](https://go.dev/dl/)
- [Git](https://git-scm.com/)
- [PostgreSQL](https://www.postgresql.org/download/) rodando localmente (ver .env)
<br>
<br>

## Clonando o projeto

Abra o terminal e execute:

```
git clone https://github.com/devjuliomartins/games-api-go.git

```
```
cd games-api-go

```
<br>
<br>

### Configuração do ambiente
O projeto já contém um arquivo .env com as variáveis necessárias para conectar ao banco de dados.


Rodando o projeto
No terminal, dentro da pasta do projeto, execute:

```
go run main.go
```
A API estará disponível em: http://localhost:8080
<br>
<br>

### Testando a API
Você pode usar o Postman ou o Insomnia para testar os endpoints da API.

Também é possível utilizar o terminal com curl ou bibliotecas como httpie.
<br>
<br>

## Planejamento de Melhorias

* [ ] Implementar autenticação com JWT.
* [ ] Mudar id de int para uuid.
* [ ] Estudar uma forma diferente de retorno das funções.
* [ ] Adição de Documentação com Swagger.
<br>
<br>

## Para Contribuir com o Projeto

Estou utilizando o Padrão de Commits Semânticos para gerênciar as alterações do projeto e facilitar a revisão do código.

Vou estipular algumas regras

- As pull requests devem ser feitas para a hmg.
- Buscar seguir a tabela de commits semânticos que vou deixar discriminado a baixo.
- Caso tenham alguma ideia para melhorar o projeto ou caso tenha dúvida [abra uma issue](https://github.com/devjuliomartins/games-api-go/issues/new).
<br>
<br>

### Passos

1. Crie um fork deste repositório.
2. Envie seus commits em português.
3. Insira um checkbox no Planejamento de Melhorias Marcado e com data e nome.
4. Crie um pull request.
<br>
<br>

### Padrão Commits Semânticos

Atualmente seguiremos esta documentação do [iuricode/padroes-de-commits](https://github.com/iuricode/padroes-de-commits).

E para a criação de branches seguiremos estes padrões:

1. **bugfix/**[nome-da-brach]: Para branch de resolução de BUG.
2. **feature/**[nome-da-branch]: Para branch de uma feature que será adicionada ao projeto.
3. **hotfix/**[nome-da-branch]: Para branch de correção de cor, textos, alterações não tão urgentes.
4. **improvement/**[nome-da-branch]: Para branch de melhoria de algo que já exista, seja performance, escrita, layout, etc.
