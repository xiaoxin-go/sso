APP=sso

build:
        @go build -o ${APP} main.go
windows:
        @GOOS=windows go build -o ${APP}-windows
linux:
        @GOOS=linux go build -o ${APP}-linux
darwin:
        @GOOS=darwin go build -o ${APP}-darwin