# delivery-much
Tech Challenge for Delivery Much

Para compilar o projeto basta executar o comando **make build**.

Para rodar a aplicação basta executar o comando **make run**.

A aplicação pode ser acessada na URL http://localhost:3000 e possui a seguinte API:

## /recipes/
### Parâmetros da Query
**i:**: Ingrediêntes da Receita - min: 1 - max: 3
### Descrição
Busca por receitas que contenham os ingredientes informados
### Exemplo
http://localhost:3000/recipes/?i=egg,bacon

Para configurar a chave de acesso do Giphy, basta editar o arquivo **config.development.json**, e inserir sua chave na propriedade **apiKey** do JSON.
