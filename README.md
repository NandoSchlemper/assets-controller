# assets-controller
Projeto para controle de bens ativos de uma empresa

### Requisitos Funcionais
1. Cadastrar Produtos
[ID, user_id, name, price, creation_date]

2. Cadastrar Baixas & Cadastrar Multiplas Baixas [<quantidade>]
[ID, produto_id, patrimony_id, condition, creation_date]

3. Cadastrar Usuários
[ID, name, creation_date]

4. Relatórios
[SELECT * FROM Produtos WHERE creation_date >= <data_inicial> AND creation_date <= <data_final>]
etc . . .



### Requisitos Não Funcionais
1. Armazenar os dados em algum banco
