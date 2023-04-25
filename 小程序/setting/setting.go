package setting

//读取配置文件 viper
import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Conf 全局配置文件信息结构体 -内存对齐
var Conf = new(App)

type App struct {
	*AppConfig   `mapstructure:"app"`   //app
	*AuthConfig  `mapstructure:"auth"`  //auth
	*LogConfig   `mapstructure:"log"`   //log
	*MysqlConfig `mapstructure:"mysql"` //mysql
	*RedisConfig `mapstructure:"reds"`  //redis
	*Admin       `mapstructure:"admin"` //admin
}
type Admin struct {
	Username string `mapstructure:"username"` //用户名
	Password string `mapstructure:"password"` //密码
}

// AppConfig 配置项目的基本信息
type AppConfig struct {
	Name      string `mapstructure:"name"`       //项目名
	Mode      string `mapstructure:"mode"`       //开发模式
	Version   string `mapstructure:"version"`    //版本
	StartTime string `mapstructure:"start_time"` //项目开始时间
	Port      int    `mapstructure:"port"`       //项目运行端口
	MachineID int64  `mapstructure:"machine_id"` //机器ID -雪花算法
}

// AuthConfig 配置JWT登录的配置
type AuthConfig struct {
	JwtExpire int `mapstructure:"jwt_expire"` //jwt使用的校验数字
}

// LogConfig 配置日志文件的配置信息
type LogConfig struct {
	Level      string `mapstructure:"level"`    //日志显示的等级
	Filename   string `mapstructure:"filename"` //日志文件的名字
	MaxSize    int    `mapstructure:"max_size"` //日志文件的最大数目
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// MyslConfig 配置Mysql
type MysqlConfig struct {
	Host         string `mapstructure:"host"`           //ip地址
	Port         int    `mapstructure:"port"`           //端口
	User         string `mapstructure:"user"`           //用户名
	Password     string `mapstructure:"password"`       //密码
	DBname       string `mapstructure:"dbname"`         //数据库名
	MaxOpenConns int    `mapstructure:"max_open_conns"` //连接池最大数
	MaxIdleConns int    `mapstructure:"max_idle_conns"` //mysql实例对象数
}

// RedisConfig 配置Redis
type RedisConfig struct {
	Host     string `mapstructure:"host"`      //ip地址
	Port     string `mapstructure:"port"`      //端口
	DB       int    `mapstructure:"db"`        //选择的数据库
	Password string `mapstructure:"password"`  //密码
	PoolSzie int    `mapstructure:"pool_size"` //连接池的数目
}

// Init 初始化viper 读取配置文件
func Init() (err error) {
	//指定文件名称
	viper.SetConfigName("config")
	//指定文件类型
	viper.SetConfigType("yaml")
	//指定文件路径
	viper.AddConfigPath(".")
	//处理配置文件
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		//panic(fmt.Errorf("Fatal error config file: %s \n", err))
		fmt.Printf("viper.ReadInConfig() failed,err:%v\n", err)
		return
	}
	//把读取到的配置信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	//热加载-配置文件实时的监控配置
	viper.WatchConfig()
	//钩子函数-当config改变时：
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})
	return
}
