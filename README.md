# micro-server-demo
A project about of micro-server's code architecture


### arch

```sh
.
├── Dockerfile      // Dockerfile
├── LICENSE
├── Makefile        // start and other command
├── README.md       // project intro
├── api             // proto related
├── cmd             // command line interface entry. to start the project and control it
├── configs         // config files
├── global          // nesscarry global instances, like: dbConn, redisClient, sessionManager, singleton
├── go.mod
├── go.sum
├── internal        // application logic modules
├── model           // models include: forms, request, response 
├── pkg             // exportable methods of current project
├── repository      // DAO layer
├── router          // router (HTTP and gRPC) + middlewares
└── scripts         // scripts: SQL, shell, py etc
```