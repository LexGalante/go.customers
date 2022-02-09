# GO.CUSTOMERS API
API de exemplo em golang

# REQUISITOS
 - go 1.17.*
 - docker

# PREPARAÇÃO AMBIENTE
O projeto tem como banco de dados o postgres e usa a mensageria rabbitmq
 - execute o comando abaixo para preparar o ambiente:
    ```
    docker-compose up -d
    ```
 - acesse [localhost:15672](http://localhost:15672/) com usuário: __guest__ e senha: __guest__
 - na aba exchange do rabbitmq crie duas novas exchanges do tipo fanout com seguintes nomes:
   - new_customers
   - change_customers
 


