package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webGin/controller"
	"webGin/dao/mysql"
	"webGin/logger"
	"webGin/pkg/snowflake"
	"webGin/router"
	"webGin/setting"

	"go.uber.org/zap"
)

// @title 疫情小程序
// @version 1.0
// @description 下面是项目的接口文档
// @termsOfService http://swagger.io/terms/

// @contact.name 18835556819
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1
// @BasePath /api/v1
func main() {
	//初始化Viper
	if err := setting.Init(); err != nil {
		fmt.Printf("setting.Init() failed,err: %v\n", err)
		return
	}
	//初始化日志
	logger_err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode)
	if logger_err != nil {
		fmt.Printf("init() logger failed;err: %v\n", logger_err)
	}
	defer zap.L().Sync()
	//初始化数据库mysql
	if err := mysql.Init(*setting.Conf.MysqlConfig); err != nil {
		fmt.Printf("mysql.Init() failed,err: %v\n", err)
		return
	}
	//初始化Gin框架内置的校验器使用的翻译器
	if trans_err := controller.InitTrans("zh"); trans_err != nil {
		fmt.Printf("init() validator failed;err: %v\n", trans_err)
	}
	//注册路由
	r := router.SetUp(setting.Conf.Mode)
	//注册id生成器
	if createID_err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); createID_err != nil {
		fmt.Printf("init() snowflake failed;err: %v\n", createID_err)
	}

	//启动服务-优雅关机
	// r.Run()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", setting.Conf.AppConfig.Port),
		Handler: r,
	}
	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
