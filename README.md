# 关于Hade

hade是一个注重协议，注重开发效率的前后端一体化框架。hade框架的目标在于将go+vue模型的应用开发最简化，并且提供统一，一体化的脚手架工具促进业务开发。

我们相信在GO的框架开发中，指定协议比实现更为重要。hade框架对于具体应用开发者来说最便利的功能，在于其制定了一系列的基本协议，在具体的业务逻辑中，可以通过每个协议独特的Key来从全局容器中获取已经注入的服务实例。所有的具体应用开发，在业务逻辑中，都是按照hade约定的协议进行逻辑处理，从而脱离了具体的每个服务所定义的个性化差异。

## Contract

hade框架目前定义的协议都在目录（framework/contract）下：

* app: 定义了整体应用框架，包括应用版本，应用基础路径，应用配置，日志存放路径等。
* env: 定义了对应的环境配置，默认环境配置存放在app中定义的basePath下，使用.env进行存储，单行以key=value为格式
* config: 定义了获取配置信息的服务，可以从app中定义的应用配置路径获取具体的配置，按照"[文件名].key1.key2"的方式获取配置信息，获取的配置值支持string,bool,int,float,time,[]int,[]string,map[string]string,map[string]interface{}, object 等多种格式
* log: 定义了日志记录格式，定义7种级别的错误日志打印，支持控制台、文件、文件切割等日志打印方式

## ServiceProvider

hade框架使用ServiceProvider机制来满足协议的，通过serviceProvder提供某个协议服务的具体实现。这样如果开发者对具体的实现协议的服务类的具体实现不满意，则可以很方便的通过切换具体协议的ServiceProvider来进行具体服务的切换。

hade的路由，controller的定义是选择基于gin框架进行扩展的。所有的gin框架的路由，参数获取，验证，context都和gin框架是相同的。唯一不同的是gin的全局路由gin.Engine实现了hade的容器结构，可以对gin.Engine进行服务提供的实例化，且可以从context中获取具体的服务。

hade提供两种服务实例化的方法：
* Bind: 将一个ServiceProvider绑定到容器中，可以控制其是否是单例
* Singleton: 将一个单例ServiceProvider绑定到容器中

hade提供了三种服务获取的方法：
* Make: 根据一个Key获取服务，获取不到获取报错
* MustMake: 根据一个Key获取服务，获取不到返回空
* MakeNew: 根据一个Key获取服务，每次获取都实例化，对应的ServiceProvider必须是以非单例形式注入

一个ServiceProvider是一个单独的文件夹，它包含服务提供和服务实现。具体可以参考framework/provider/demo

service.go提供具体的实现，它提供一个实例化的方法
```
func NewDemoService(params ...interface{}) (interface{}, error) {
}

```

provider.go 提供服务适配的实现，实现一个Provider必须实现对应的五个方法
```
// provider的Key，根据这个Key从容器中对应服务
func (sp *DemoServiceProvider) Name() string {
	return "demo"
}

// 注册服务的实例化方法
func (sp *DemoServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewDemoService
}

// 实例化是否延后，服务是调用时实例化还是注入的时候实例化
func (sp *DemoServiceProvider) IsDefer() bool {
	return true
}

// 服务实例化的参数
func (sp *DemoServiceProvider) Params() []interface{} {
	return []interface{}{sp.C}
}

// 服务实例化之前的启动函数，可以在这个启动函数中做一些初始化的操作
func (sp *DemoServiceProvider) Boot(c framework.Container) {
	fmt.Println("demo service boot")
}
```

一个SerivceProvider就是一个独立的包，这个包可以作为插件独立地发布和分享。

你也可以定义一个无contract的ServiceProvider，其中的Name()需要保证唯一。
目前hade中默认的Key有：

* app
* env
* config
* log
* gorm

# 代码框架
```
app  # 应用代码存放目录
  - console # 控制台应用
    - command # 具体的控制台应用
    kernel.go # 控制台应用的核心代码，如果写了新的控制台应用，需要在这个文件中注入
  - http # web应用
    - controllers # web应用控制器
    kernel.go # web应用控制器核心代码，如果写了新的serviceProvider， 需要在这个文件中注入
  - models # 具体模型
  - providers # 自定义方法
config # 配置文件目录
framework # 框架核心代码
public # 对外访问路径，也是编译后的前后端存放路径
routes # web应用路由
src # 前端vue的代码
  - assets  # 前端的资源文件
  - components  # 前端的vue组件
  App.vue # 组件入口
  main.js # 总入口
storage # 存储目录
  - cache # 缓存文件存储
  - coverage # 代码覆盖率报告存储地址
  - logs # 日志存储地址
testdata # 测试需要的数据文件
tests # 独立的测试用例目录
.env # 环境变量
babel.config.js # 前端babel文件
converage.sh # 代码覆盖率脚本
hade # 初始化可执行的hade命令，下载框架后，运行build就生成
main.go # 后端主要入口
package.json # 前端的包管理文件
```

# 具体协议

## app

## env

## config

## log

# 中间件

由于gin框架地址的迁移和修改，https://github.com/gin-contrib 下面的 gin 中间件插件并不能直接使用。需要
* 拷贝到 github.com/jianfengye/hade/framework/middleware 下面
* 去除go.mod和go.sum和.git
* 将文件夹中所有文件import处 github.com/gin-gonic/gin 替换为 github.com/jianfengye/hade/framework/gin

也可以直接使用命令

`./hade middleware [add/remove/all] [repo]`
将 https://github.com/gin-contrib 下面的中间件项目迁移到当前框架下


# 待办事项

- 尝试一个命令同时提供web和api服务
- 尝试将vue-element-admin嵌入进框架
- 前后端框架名字进行管理
- 前端webpack增加，日常常用的vue webpack插件增加