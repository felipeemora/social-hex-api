

### Run in localhost
#### pre-requisites
1. go: 1.24.1 or latest 
2. docker & docker-compose `brew install docker-compose`
3. golang-migrate `brew install golang-migrate`

#### Steps:
1. make up
2. make migrate-up
3. make seed
4. make run-dev
5. docker ps y obtener el id de redis
5. docker exec -it d97733d3e4de redis-cli GET "user:61"
7. Tambié está Redis  Comander en el http://localhost:8081

### Links
- Swagger: https://localhost:8080/v1/swagger/index.html

- Para medir tiempos de respuesta 
npx autocannon http://localhost:8080/v1/users/61 --connections 10 --duration 5 -H "Authorization: Bearer {{token}}"

