## liaojuntao简易用户管理系统API文档
前言： 很多地方没有完善的，多多见谅^_^，写word麻烦管理，还是直接用markdown写了。
## API版本说明
api版本 | 时间 | 修订人 | 备注
---|---|---|---
v1.0 | 2022.04.17 | liaojuntao|初步指定框架和功能实现
## 协议与接口标准
- 选用的协议：http1.1
- 请求地址：http://127.0.0.1:8001
- 路径： /xxx/xxx
- 请求：
    - get 传参 /xxx/xxx?key1=value1&key2=value2
    - post 传参 给request.body写入json格式数据
- 标准响应
    - 错误响应： 
    
        |参数 | 数据类型 | 说明 | 备注|
        |---|---|---|---|
        |errCode | 整型 | 响应码 | 请参考下面的标准响应码|
        |errMessage | 字符串 | 错误原因 ||
    - 成功响应: 
    
        |参数 | 数据类型 | 说明 | 备注|
        |---|---|---|---|
        |status | 整型 | 响应码 | 请参考下面的标准响应码|
        |data | json字符串 | 响应数据 ||
- 标准响应码  

    |响应码 |  说明 | 备注|
    |---|---|---|
    |2000 | 成功响应 ||
    |4001  | 参数校验错误 ||
    |4002  | 业务操作错误 ||
    |5000  | 系统内部错误 |反映系统panic了|
## 接口说明
### 创建用户
- 接口名： /user/create
- 请求方法：POST
- 参数列表，需要传json格式：
```
{
  "userName": "test",  
  "birthOfDay": "", 
  "address": "",  
  "description": ""
}
```
参数 | 数据类型 | 说明 | 是否必填 | 备注
---|---|---|---|---
userName | 字符串 | 用户名 | 是| 不允许与现有用户名重复，长度不大于50
birthOfDay | 字符串 | 生日 | 是| 指定年月日，固定格式："2011.05.05"
address | 字符串 | 地址| 是 | 长度不大于100
description | 字符串 | 描述 | 否| 长度不大于150

- 成功响应：
```
{
  "status": xxx,   
  "data": 'string'
}
```
参数 | 数据类型 | 说明 | 备注
---|---|---|---
status | 整型 | 固定值： 2000 | 
data | 字符串 | 固定值： "succeed" | 
- 失败响应：
```
{
  "errCode": xxx,   
  "errMessage": 'err'
}
```
参数 | 数据类型 | 说明 | 备注
---|---|---|---
errCode | 整型 | 参考标准响应码 | 
errMessage | 字符串 | 错误信息 | 

### 更新用户
- 接口名： /user/update
- 请求方法：POST
- 注意：更新接口的参数都必需要传
- 参数列表，需要传json格式：
```
{
  "userId": xxx,   
  "userName": "test",  
  "birthOfDay": "", 
  "address": "",  
  "description": ""
}
```
参数 | 数据类型 | 说明 |是否必填 | 备注
---|---|---|---|---
userName | 字符串 | 用户名|是 | 不允许与现有用户名重复，长度不大于50
birthOfDay | 字符串 | 生日|是 | 指定年月日，固定格式："2011.05.05"
address | 字符串 | 地址|是 | 长度不大于100
description | 字符串 | 描述|是 | 注意：这里不能不填，且长度不大于150

- 成功响应：
```
{
  "status": xxx,   
  "data": 'string'
}
```
参数 | 数据类型 | 说明 | 备注
---|---|---|---
status | 整型 | 固定值： 2000 | 
data | 字符串 | 固定值： "succeed" | 
- 失败响应：
```
{
  "errCode": xxx,   
  "errMessage": 'err'
}
```
参数 | 数据类型 | 说明 | 备注
---|---|---|---
errCode | 整型 | 参考标准响应码 | 
errMessage | 字符串 | 错误信息 | 

### 根据用户id获取用户
- 接口名： /user/getById
- 请求方法：GET
- 参数：用户id:userid
```
/user/getById?userid=1
```
- 成功响应：
```
{
  "status": 2000,   
  "data": 
        {
            "userId": 1,  
            "userName": "test",  
            "birthOfDay": "", 
            "address": "",  
            "description": ""
            "createAt": ""
        }
}
```
参数 | 数据类型 | 说明 | 备注
---|---|---|---
status | 整型 | 固定值： 2000 | 
data | 字符串 | 用户实体的json字符串 | 
- 失败响应：
```
{
  "errCode": xxx,   
  "errMessage": 'err'
}
```
参数 | 数据类型 | 说明 | 备注
---|---|---|---
errCode | 整型 | 参考标准响应码 | 
errMessage | 字符串 | 错误信息 | 

### 根据用户id删除用户
- 接口名： /user/deleteById
- 请求方法：GET
- 参数：用户id: userid
```
/user/deleteById?userid=1
```
- 成功响应：
```
{
  "status": 2000,   
  "data": ""
}
```
参数 | 数据类型 | 说明 | 备注
---|---|---|---
status | 整型 | 固定值： 2000 | 
data | 字符串 | 固定值： "succeed" | 
- 失败响应：
```
{
  "errCode": xxx,   
  "errMessage": 'err'
}
```
参数 | 数据类型 | 说明 | 备注
---|---|---|---
errCode | 整型 | 参考标准响应码 | 
errMessage | 字符串 | 错误信息 | 