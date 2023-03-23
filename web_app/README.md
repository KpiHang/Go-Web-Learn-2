# Go web 一种常用的脚手架模板

## CLD分层
分层设计模式，如MVC为了能够对GUI类型的应用进行方便扩展，将程序划分为：
- 控制器（Controller）：负责转发请求，对请求进行处理。
- 视图（View）：界面设计人员进行图形界面设计。
- 模型（Model）：程序员编写程序应有的功能（实现算法等等）、数据库专家进行数据管理和数据库设计（可以实现具体的功能）。

后来由于前后端分离的流行，前端工程化，可以认为前后分离是把 V 层从 MVC 中抽离单独成为项目，这样一个后端项目一般就只剩下 M 和 C 层了。前后端之间通过 ajax 来交互。

对于复杂的项目，一个 C 和一个 M 层显然是不够用的，现在比较流行的纯后端 API 模块一般采用下述划分方法：
- Controller，控制层，与上述类似，服务入口，负责处理路由，参数校验，请求转发。
- Logic/Service，业务逻辑（服务）层，一般是业务逻辑的入口，可以认为从这里开始，所有的请求参数一定是合法的。业务逻辑和业务流程也都在这一层中。常见的设计中会将该层称为 Business Rules。
- DAO/Repository，DAO（Data Access Object）层，这一层主要负责和数据、存储打交道。将下层存储以更简单的函数、接口形式暴露给 Logic 层来使用。负责数据的持久化工作。

每一层都会做好自己的工作，然后用请求当前的上下文构造下一层工作所需要的结构体或其它类型参数，然后调用下一层的函数。在工作完成之后，再把处理结果一层层地传出到入口。
[![ppdvT56.png](https://s1.ax1x.com/2023/03/23/ppdvT56.png)](https://imgse.com/i/ppdvT56)

## 模块介绍

```
├─controllers
├─dao
│  ├─mysql
│  └─redis
├─logger
├─logic
├─models
├─pkg
├─routes
└─settings
```

- `settings`包中包括一个配置文件`config.yaml`，和`settings.go`使用viper加载配置文件；
- `logger`包中使用zap创建全局Logger，并重构了中间件Logger()和Recovery()；
- `dao`, Data Access Object, 完成了与mysql和与redis的交互；
- `routes`, 将路由注册统一放到routes包中；


## 参考
https://www.jianshu.com/p/403f3316a5fb