module github.com/yeqown/micro-server-demo

go 1.12

require (
	github.com/gin-gonic/gin v1.4.0
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/jinzhu/gorm v1.9.9
	github.com/opentracing/opentracing-go v1.1.0 // indirect
	github.com/sirupsen/logrus v1.2.0
	github.com/urfave/cli v1.20.0
	github.com/yeqown/infrastructure v0.2.0
	google.golang.org/grpc v1.19.0
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)

replace github.com/yeqown/infrastructure => ../infrastructure
