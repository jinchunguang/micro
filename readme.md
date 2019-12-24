 
## Switch to "release" mode in production
 
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

## build

    docker build -t go/micro-kernel:[tag] .
    docker run -d  -p 8080:8080 --name api_server go/micro-kernel:[tag]
    
