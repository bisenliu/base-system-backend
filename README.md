# base-system-backend

## 1. 基本介绍

### 1.1 项目介绍

> base-system-backend 是一个基于 Gin 开发的后端基础脚手架，集成了 JWT 鉴权和基础 RBAC（Role-Based Access Control）功能。

## 2. 使用说明

- golang版本 >= v1.20
- IDE推荐：Goland
- 使用版本号接口,需要打开 .git 钩子 .git/hooks/pre-commit.sample

  ```bash
  # 复制 pre-commit
  cp .git/hooks/pre-commit.sample .git/hooks/pre-commit
  # 编辑 pre-commit
  vim .git/hooks/pre-commit
  # 写入以下内容
  ```

  ```bash
  # 获取当前时间戳
  timestamp=$(date +%s)
  # 写入到 ./server/version.txt 文件里
  printf "%s" "$timestamp" > ./server/version.txt
  # 提交
  git add ./server/version.txt
  ```

### 2.1 server项目

使用 `Goland` 或其他编辑工具，打开项目

#### 2.1.1 本地启动项目

```bash
# 克隆项目
git clone -b develop https://github.com/bisenliu/base-system-backend.git
# 进入 server 文件夹
cd base-system-backend/server
# 运行初始化脚本
chmod +x ./project_init.sh && bash ./project_init.sh

# 使用 go mod 并安装go依赖包
go mod download

# 第一次运行请初始化(自动创建表以及账号角色,账号 root 密码 123456)
# 可在 initialize/internal/default_data.go 去修改
go run main.go --env local --system_init true

#后续使用下面启动项目，否则每次启动项目都会重置相关表--env local/dev/test/product（缺省默认是local）
go run main.go --env local

# 编译 
go build -o server main.go (windows编译命令为go build -o server.exe main.go )

# 运行二进制
./server (windows运行命令为 server.exe)
```

#### 2.1.2 Docker-compose 启动项目

- 使用前请安装Docker和Docker-compose

- 请不要移动docker-compose.yml,移动需修改docker-compose.yml配置

```bash
# 克隆项目
git clone -b develop https://github.com/bisenliu/base-system-backend.git

# 进入 deploy 文件夹
cd base-system-backend/deploy

# 部署项目
docker-compose up -d

# 安装完成后修改 docker-compose.yml base-system-backend 容器的 command(删除---system_init true) 否则每次重启容器都会重置相关表
sed -i "s#./server --env dev --system_init true#./server --env dev#g" ./docker-compose.yml

# 修改完成后重启容器
docker-compose up -d
```

### 2.3 swagger自动化API文档

#### 2.3.1 安装 swagger

##### （1）可以访问外国网站

````
go get -u github.com/swaggo/swag/cmd/swag
````

##### （2）无法访问外国网站

由于国内没法安装 go.org/x 包下面的东西，推荐使用 [goproxy.cn](https://goproxy.cn)
或者 [goproxy.io](https://goproxy.io/zh/)

```bash
# 如果您使用的 Go 版本是 1.13 - 1.15 需要手动设置GO111MODULE=on, 开启方式如下命令, 如果你的 Go 版本 是 1.16 ~ 最新版 可以忽略以下步骤一
# 步骤一、启用 Go Modules 功能
go env -w GO111MODULE=on 

# 步骤二、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# 使用如下命令下载swag
go get -u github.com/swaggo/swag/cmd/swag
```

#### 2.3.2 生成API文档

```` shell
cd server
swag init --parseVendor --parseInternal --parseDependency
````

> 执行上面的命令后，server目录下会出现docs文件夹里的 `docs.go`, `swagger.json`, `swagger.yaml` 三个文件更新，启动go服务之后,
> 在浏览器输入 [http://localhost:8888/swagger/index.html](http://localhost:8888/swagger/index.html) 即可查看swagger文档

## 3. 技术选型

- 后端：用 [Gin](https://gin-gonic.com/) 快速搭建基础restful风格API，[Gin](https://gin-gonic.com/) 是一个go语言编写的Web框架。
- 数据库：采用`PostgreSql` > (15.3) 版本 数据库引擎 InnoDB，使用 [gorm](http://gorm.cn) 实现对数据库的基本操作。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- API文档：使用`Swagger`构建自动化文档。
- 配置文件：使用 [viper](https://github.com/spf13/viper) 实现`yaml`格式的配置文件。
- 日志：使用 [zap](https://github.com/uber-go/zap) 实现日志记录。
- 限流：使用 [Ratelimit](https://github.com/juju/ratelimit) 令牌桶，限制到达系统的并发请求数。

## 4. 项目架构

### 4.1 目录结构

```
    ├── base-system-backend
        ├── deploy              (docker 部署)
        ├── server
            ├── api             (api层)
            │   └── v1          (v1版本接口)
            ├── config          (配置包)
            ├── core            (核心文件)
            ├── docs            (swagger文档目录)
            ├── constants       (常量文件)
            ├── global          (全局对象)                    
            ├── initialize      (初始化)                        
            │   └── internal    (初始化内部函数)                            
            ├── middleware      (中间件层)                        
            ├── model           (模型层)                    
            │   ├── request     (入参结构体)                        
            │   └── response    (出参结构体)                                                      
            ├── router          (路由层)              
            ├── service         (service层)                               
            └── utils           (工具包)                                     
```

## 5. 主要功能

- 权限管理：基于`jwt`和`装饰器`实现的权限管理 (后期使用 casbin 进行权限管理)。

- 用户管理：系统管理员分配用户角色和角色权限。

- 角色管理：创建权限控制的主要对象，可以给角色分配不同API权限。

- API 管理：不同用户可调用的 API 接口的权限不同。

## 6.说明

在 model/common/field/aes 文件中，自定义了字段加密相关的功能，可以对一些敏感数据进行加密，并在返回数据时进行解密。

```go
// model/common/field/aes
// 自定义字段,对于一些敏感数据可对其进行加密,返回数据时进行解密
// field.PlainEncrypt 对内容进行整体加密
// field.SplitEncrypt 对单个字符进行加密,并组合(需要进行模糊查询用此字段)

type User struct {
Phone       field.PlainEncrypt           `gorm:"column:phone;size:11;comment:手机号"`
Phone2      field.SplitEncrypt           `gorm:"column:phone2;size:11;comment:手机号2"`
}

```
