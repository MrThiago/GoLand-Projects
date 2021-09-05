## Go and Gin Learning Path

Tutorial by https://golang.org/doc/tutorial/web-service-gin
and some samples from
https://gin-gonic.com/docs/testing/ 
https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e

Create new working directory
```
mkdir web-service-gin
cd web-service-gin
```

Create module to manage dependencies 
```
go mod init <package>/web-service-gin
```

Create new main file 
```
touch main.go
```

Add code to main file 
first line must be: `package main`

Add the import
```
import (
    "net/http"
    "github.com/gin-gonic/gin"
)
```

Run code to get the libs
```
go get .
```

To Run code
```
go run .
```

Test APIS
Get
```
curl http://localhost:8080/albums
```
Post
```
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```

Get by ID
```
curl http://localhost:8080/albums/2
```

Get Ping
```
curl http://localhost:8080/ping
```

## Using Docker
Build
```
docker build . -t go-dock
```

Run
```
docker run -p 8080:8080 go-dock
```

Test 
```
curl http://localhost:8080/ping
```

### Docker tips

See all containers
```
docker ls -a
```

Clean up all images
```
docker system prune --all
```

**Docker Cheat Sheet**
![Docker Cheat Sheet](https://raw.githubusercontent.com/sangam14/dockercheatsheets/master/dockercheatsheet8.png)
> from https://dockerlabs.collabnix.com/docker/cheatsheet/
