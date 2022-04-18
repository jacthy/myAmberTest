
# go test 进行单元测试（-gcflags=-l为禁用内联）
test:
	go test -gcflags=-l ./... -cover
	
test-add-user:
	curl -X POST -k -v \
		-H "Content-Type: application/json"  \
			http://127.0.0.1:8001/user/create \
		-d '{"userName":"用户1","birthOfDate":"2021-02-09","address":"广州","description":"描述1"}'

