# GO GRAPHQL

## Requirements

If you are wanting to build and develop this, you will need the following items installed.

- Go version 1.12+
- MySql server

# Install the dependencies

```
go get -u "github.com/rsvijay2009/go-graphql"

go get -u "github.com/99designs/gqlgen/handler"

go get -u "github.com/go-sql-driver/mysql"

go get -u "github.com/vektah/gqlparser"

```

# Up and running

Follow the steps to run the application

- Update your database credentials in `run.sh` file

- Run the below command to import the database

  `mysql -u username -p database_name < db_scripts.sql`

- Run `./run.sh`

- Your GrapQL server is up and running in http://localhost:8080/

# Example queries

# Get users

```
query users {
  users {
    first_name
  }
}
```

# Create user

```
mutation createUser {
  createUser(input:{
    first_name:"your_first_name",
    last_name:"your_last_name",
    email: "your_email",
    password:"your_password"
  }){
    first_name
    last_name
    email
  }
}
```

# Schema modification

- If you want to change any schema change, you should run `./run.sh` and need to compare the generated files with existing files. For ex: (resolver.go etc..)
