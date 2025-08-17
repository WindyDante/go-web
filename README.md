# 后端

# start
go mod tidy
在根目录创建.env并编写对应的环境变量
go run cmd/gen/generate.go
go run cmd/oa/main.go