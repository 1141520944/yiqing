package router

import (
	"net/http"
	"webGin/controller"
	_ "webGin/docs" // 千万不要忘了导入把你上一步生成的docs
	"webGin/logger"
	"webGin/mildware"

	swaggerFiles "github.com/swaggo/files"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
)

// SetUp 初始化路由
func SetUp(mode string) *gin.Engine {
	//判断是否是发布模式
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	//使用中间件
	r.Use(CORSMiddleware(), logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//注册路由
	v1 := r.Group("/api/v1")
	{
		//测试
		v1.GET("hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "成功",
			})
		})
		//用户注册
		v1.POST("/signup", controller.SingUpHandlerInstructor)
		//用户登录
		v1.POST("/login", controller.LoginHandlerInstructor)
		//辅导员
		in := v1.Group("/instructor")
		//使用JWT中间件
		in.Use(mildware.JWTAuthMiddleware())
		{
			//检查是否登录成功
			in.GET("/ping", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "pong")
			})
			//操作学生
			{
				in.POST("/student/add", controller.AddInformationStudent)
				in.POST("/student/update", controller.UpdateInformationStudent)
				in.GET("/student/delete/:uid", controller.DeleteInformationStudent)
				in.GET("/student/deleteC/:uid", controller.DeleteCompletelyInformationStudent)
				in.GET("/student/all", controller.SelectPageInformationStudent)
			}
			//操作super 学生
			{
				in.POST("/super/add", controller.AddInformationSuper)
				in.POST("/super/update", controller.UpdateInformationSuper)
				in.GET("/super/all", controller.SelectPageInformationSuper)
				in.GET("super/deleteC/:uid", controller.DeleteInformationSuper)
			}
			//操作班级
			{
				in.POST("/class/add", controller.AddInformationClass)
				in.POST("/class/update", controller.UpdateInformationClass)
				in.GET("/class/all", controller.SelectPageInformationClass)
				in.GET("class/deleteC/:uid", controller.DeleteInformationClass)
			}
			//查看核酸数目
			{
				in.POST("/nucleicAcid/number", controller.GetClassStudentNucleicAcid)
			}
		}
		su := v1.Group("/super")
		//操作员
		{
			su.POST("/login", controller.LoginHandlerSuper)
			su.Use(mildware.JWTAuthMiddleware())
			//检查是否登录成功
			su.GET("/ping", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "pong")
			})
			//super学生操作核酸
			nu := su.Group("/nucleicAcid")
			{
				nu.POST("/add", controller.AddInformationNucleicAcid)
			}
		}
		//普通学生
		st := v1.Group("/student")
		{
			// st.GET("/ping")
			st.POST("/add", controller.RegistStudent)
			st.POST("/one", controller.IsOrNoRegistStudent)
			st.GET("/getOne", controller.GetOneByUidStudent)
			st.GET("/class/getall", controller.GetAllClasses)
			st.GET("/instructor/getall", controller.GetAllInstructors)
			st.GET("/class/one", controller.GetOneByClassIDClass)
			st.GET("/instructor/one", controller.GetOneByInstructorIDInstructor)
			st.Use(mildware.JWTAuthMiddleware())
			//检查是否登录成功
			st.GET("/ping", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "pong")
			})
			nuc := st.Group("/nucleicAcid")
			{
				nuc.GET("/all", controller.SelectPageInformationNucleicAcid)
			}
		}
		ad := v1.Group("/admin")
		{
			ad.POST("/login", controller.LoginHandlerAdmin)
			ad.POST("update", controller.UpdateInformationInstructor)
			ad.GET("/all", controller.SelectPageInformationInstructor)
			ad.GET("/deleteC/:uid", controller.DeleteInformationInstructor)
			{
				ad.GET("/super/all", controller.SelectPageInformationSuperAll)
			}
		}
		wx := v1.Group("/wx")
		{
			wx.POST("/openid", controller.WxGetOpenID)
		}

	}
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
