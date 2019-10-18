 ## Requirements
 - golang
 - docker 
 - postgres
 
 
 ## Install and Run
- go installation https://golang.org/doc/install
- docker install https://docs.docker.com/install/
- postgresql https://hub.docker.com/_/postgres


## Development
- Run -> go run main.go (https://www.golang-book.com/books/intro/2)
- Build  
    - go build (This will generate go-docker file)
    - ./api-example  (This will run application) 
- Create and Run Docker Image
    - docker build -t api-example .
    - docker images (This will list docker images, api-example as well)
    - docker run -d -p 9001:9001 (Then go to localhost:9001, you know what will be next :) )
