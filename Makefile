
# go test 进行单元测试（-gcflags=-l为禁用内联）
test:
	go test -gcflags=-l ./... -cover

# 测试添加用户
test-add-user:
	curl -X POST -k -v \
		-H "Content-Type: application/json"  \
			http://127.0.0.1:8001/user/create \
		-d '{"userName":"用户1","birthOfDate":"2021-02-09","address":"广州","description":"描述1"}'

# 测试修改用户，用户id需要自己修改
test-add-user:
	curl -X POST -k -v \
		-H "Content-Type: application/json"  \
			http://127.0.0.1:8001/user/update \
		-d '{"userId":1,"userName":"用户2","birthOfDate":"2021-02-08","address":"广州","description":"描述1"}'

# 需要自己修改UserId
test-getById:
	curl -X GET -k -v  -H "Content-Type: application/json" http://127.0.0.1:8001/user/getById?userId=1


# 需要自己修改UserId
test-deleteById:
	curl -X DELETE -k -v  -H "Content-Type: application/json" http://127.0.0.1:8001/user/deleteById?userId=1