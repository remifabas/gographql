# gographql
Some graphql implementation in go

# go get :
~~~
go get	"github.com/friendsofgo/graphiql"
go get	"github.com/gographql/queries"
go get	"github.com/graphql-go/graphql"
go get	"github.com/graphql-go/handler"
~~~

# GO GRAPHQL
This repository is a test ! Always WIP !  
Feel free to tell me if you thinks something is wrong  

I tried to build a graphql api : https://graphql.org/

## Development commands
~~~
# Initialize docker container
docker run -d -p 27017:27017 --name mongodb mongo:latest

# Relauch/start/stop docker container
docker stop mongodb
docker start mongodb

# build application:
go build -o /home/remi/go/bin/graphqlgo main.go

# launch application
/home/remi/go/bin/graphqlgo
~~~
