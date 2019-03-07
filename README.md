This repository is my practice space for GraphQL.

## dotenv example

```
MYSQL_ROOT_PASSWORD=password
MYSQL_DATABASE=goraphql
DB_USER=root
DB_PASSWORD=password
DB_NAME=goraphql
```

## Request example

### Get single user
```
$ curl -X POST -d '{ user(id: 1) { id, name, created_at }}' localhost:8080/graphql | jq
```

### Create user
```
$ curl -X POST -d 'mutation{ createUser(name: "💩") { id, name }}' localhost:8080/graphql | jq
```

### Get user list
```
$ curl -X POST -d '{ userList { id, name, created_at } }' localhost:8080/graphql | jq
```