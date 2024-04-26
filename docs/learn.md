# 知识点

## 标准库
* fmt: 实现了类似c语言printf和scanf的格式化i/o,格式化动作源自c语言但更简单
* net/http: 提供了HTTP客户端和服务端的实现

## gin
* gin.Default: 返回Gin的type Engine struct{..},里面包含RouterGroup,
相当于创建一个路由Handlers,可以后期绑定各类的路由规则和函数，中间件等
* router.GET(...){...}创建不同的HTTP方法绑定到Handlers中,也支持 
POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的 Restful 方法
* gin.H{...}就是一个map[string]interface{}
* gin.Context:Context是gin中的上下文，它允许我们在中间件之间传递变量,
管理流，验证json请求，响应json请求等，在gin中包含大量的Context方法，例如
我们常用的DefaultQuery,Query等

## http.server

```golang
type Server struct{
Addr    string
Handler Handler
TLSConfig *tls.Config
ReadTimeout time.Duration
ReadHeaderTimeout time.Duration
WriteTimeout time.Duration
IdleTimeout time.Duration
MaxHeaderBytes int
ConnState func(net.Conn, ConnState)
ErrorLog *log.Logger
}
```
* Addr：监听的 TCP 地址，格式为:8000
* Handler：http 句柄，实质为ServeHTTP，用于处理程序响应 HTTP 请求
* TLSConfig：安全传输层协议（TLS）的配置
* ReadTimeout：允许读取的最大时间
* ReadHeaderTimeout：允许读取请求头的最大时间
* WriteTimeout：允许写入的最大时间
* IdleTimeout：等待的最大时间
* MaxHeaderBytes：请求头的最大字节数
* ConnState：指定一个可选的回调函数，当客户端连接发生变化时调用
* ErrorLog：指定一个可选的日志记录器，用于接收程序的意外行为和底层系统错误；如果未设置或为nil则默认以日志包的标准日志记录器完成（也就是在控制台输出）

## ListenAndServe

```go
func (src *Server) ListenAndServe() error{
	addr := src.Addr
	if add
}
```