## 系统介绍
go backend test demo from liaojuntao，一个简易的用户管理系统demo,
server端算是一个对自己的challenge,自己实现的简易web框架
### 系统简析
![image](https://note.youdao.com/s/DS5d5tY7)
- 系统主要分三层：接口层，业务层，基础服务层
    - 接口层主要是注册路由，规范restful服务接口，以及拦截器（后期可添加限流，熔断，鉴权等高级特性）
    - 业务层主要是处理业务逻辑，然后调用基础服务层（其实就是dto层）实现持久化存储
    - 基础服务层，这里抽象了一层repo，为的是将数据库解耦，便于插拔式剪裁组件。这里用作demo，简易为主，用了sqlite做存储，后期可以改用mysql，redis等实现存储
### tips for developer
1. 接口层是一个基于go的基础包http实现的一个简易web框架，只添加了两个组件router和interceptor,分别实现路由配置，以及拦截器功能
    - 注册路由及handler
    ```
    // UpdateUserRouter 更新用户的router
    func (u *UserRouter) DeleteByIdRouter() Router {
        u.path = "/user/deleteById" // set path
        u.handler = deleteByIdHandler // set handler
        return u
    }
    ```
    - 加载到server路由处理器中
    ```
    // loadRouter 加载路由
    func loadRouter() []router.Router {
        r := make([]router.Router, 0, routerNum)
        return append(r,
            router.GetUserRouter().CreateUserRouter(),
            router.GetUserRouter().UpdateUserRouter(),
            router.GetUserRouter().GetByIdRouter(),
            router.GetUserRouter().DeleteByIdRouter(),
        )
    }
    ```
    - 注册拦截器
    ```
    // serverInterceptors 设置拦截器,利用链式调用实现解耦，插件式添加拦截器
    func serverInterceptors() []Interceptor {
        i := make([]Interceptor, 0, interceptorNum)
        return append(i,
            defaultInterceptor,
            safeValidInterceptor,
        )
    }
    ```
2. controller 业务逻辑层，依赖基础设施层提供持久化服务，并且通过依赖抽象接口的方式解耦
    - 处理业务逻辑
    ```
    // CreateUser 创建用户业务逻辑
    func (u *UserCtl) CreateUser(user *infrastruct.User) error {
    	isNotExist, err := u.repo.NotExistByName(user.UserName)
    	if err != nil {
    		return err
    	}
    	if !isNotExist {
    		return errors.New("该用户名已存在")
    	}
    	return u.repo.Create(user)
    }
    ```
3. infrastruct基础设施层，目前只有sqlite实现的存储功能
    - 将业务对象进行持久化处理
        ```
        // Create 将新建对象进行持久化
        func (u *UserRepo) Create(user *infrastruct.User) error {
        	model := toSqliteModel(user) // 转换为dto层的对象进行存储
        	currentTime := time.Now()
        	model.CreateAt = &currentTime
        	return u.db.Create(model).Error
        }
        ```
### 单元测试简介
1. 单元测试设计原则：server层与controller层解耦，controller层与repo层解耦
2. 测试方法参考：可使用Makefile中的test命令，直接执行命令make test
2. 选用gomonkey作为打桩工具使用，实现对部分方法的mock实现,注意点：一些内联优化的函数是mock不了的，所以go test的时候要禁用内联。
    ```
        patch := NewPatches()
        defer patch.Reset()
    
        // 打桩controller层，router.handler的业务应与controller层解耦，所以测试用例也应该解藕
        patch.ApplyMethod(reflect.TypeOf(controller.NewUserController(nil)), "DeleteUserById",
            func(_ *controller.UserCtl, mockId int) error {
                println("heated it")
                if mockId == 222 {
                    return errors.New("controller err")
                }
                return nil
            })
    ```
3. 选用convey进行一个测试用例的阐述说明
4. 通过接口+依赖注入方式实现的repo，会直接mock一个对象进行打桩
### 接口设计文档参考
[api接口文档](./api_design_document.md)