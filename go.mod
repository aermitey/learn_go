module learngo

go 1.16

require (
	github.com/armstrongli/go-bmi v0.0.1
	github.com/favadi/protoc-go-inject-tag v1.3.0 // indirect
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-gonic/gin v1.7.7
	github.com/go-sql-driver/mysql v1.6.0
	github.com/spf13/cobra v1.3.0
	golang.org/x/net v0.0.0-20220325170049-de3da57026de
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	gorm.io/driver/mysql v1.3.2
	gorm.io/gorm v1.23.2
	learngo.tools v0.0.0-00010101000000-000000000000

)

replace learngo.tools => ../learngo.tools
