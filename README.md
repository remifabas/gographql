# gographql
Some graphql implementation in go

# go get :
~~~
go get	"github.com/friendsofgo/graphiql"
go get	"github.com/gographql/queries"
go get	"github.com/graphql-go/graphql"
go get	"github.com/graphql-go/handler"
~~~

# dev

## launch database before
docker run -d -p 27017:27017 --name mongodb mongo:latest

## build :
go build -o /home/remi/go/bin/graphqlgo main.go

## launch
/home/remi/go/bin/graphqlgo

