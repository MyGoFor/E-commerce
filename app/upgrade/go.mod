module github.com/MyGoFor/E-commerce/app/upgrade

go 1.23.4

replace (
    github.com/MyGoFor/E-commerce/app/casbin => ../../app/casbin
	github.com/MyGoFor/E-commerce/rpc_gen => ../../rpc_gen
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)


require github.com/golang/protobuf v1.5.4 // indirect