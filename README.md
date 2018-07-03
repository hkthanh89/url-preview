# Setup development environment

- Install dependencies packages:
```
go get github.com/gorilla/mux
go get github.com/stretchr/testify
go get github.com/codegangsta/gin
go get github.com/PuerkitoBio/goquery
```

- Start server
```
go run main.go
```

- Start server with live reload
```
gin -p 8888 -a 8889 run main.go
```

- Run test

Cd to folders have `*_test.go` file and run:
```
go test --cover -v
```

# API
#### Url Preview

###### Request
```sh
GET /preview?url={url}
```

###### Success response
```json
{
   "code":200,
   "result":{
      "object":{
         "og:url":"http://apple.com",
         "og:title":"Apple",
         "og:description":"Discover the innovative world of Apple and shop everything iPhone, iPad, Apple Watch, Mac, and Apple TV, plus explore accessories, entertainment, and expert device support.",
         "og:image":"https://www.apple.com/ac/structured-data/images/open_graph_logo.png?201709101434"
      }
   }
}
```