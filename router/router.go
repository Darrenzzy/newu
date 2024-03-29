package router

import (
	"go-admin/apis/admin"
	"go-admin/pkg/jwtauth"
	jwt "go-admin/pkg/jwtauth"

	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
)

// 路由示例
func InitExamplesRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	// 无需认证的路由
	examplesNoCheckRoleRouter(r)
	// 需要认证的路由
	examplesCheckRoleRouter(r, authMiddleware)

	MemberRouter(r)
	WorthRouter(r)
	WebRouter(r)
	MemberTransaction(r)
	AppointmentRouter(r)
	ContactsRouter(r)
	ResumeBaseRouter(r)
	return r
}

// 用户管理
func MemberRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1/member")
	v1.GET("/data/:id", admin.GetMember)
	v1.GET("/list", admin.GetMemberList)
	v1.PUT("/", admin.UpdateMember)
	v1.DELETE("/:id", admin.DeleteMember)
	v1.POST("/register", admin.RegisterMember)
	v1.POST("/login", admin.Login)
	v1.POST("/reset_pass", admin.ResetPass)

}

// 企业网站 路由
func WebRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.GET("/sendCode", admin.SendCode)

}

// 交易记录
func MemberTransaction(r *gin.Engine) {
	v1 := r.Group("/api/v1/member_transaction")
	v1.GET("/list", admin.GetMemberTransactionList)
	v1.GET("/data/:id", admin.GetMemberTransaction)
	v1.PUT("", admin.UpdateMemberTransaction)
	v1.DELETE("/:id", admin.DeleteMemberTransaction)
	v1.POST("/add", admin.AddMemberTransaction)
}

// 净值
func WorthRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1/netWorth")
	v1.PUT("/", admin.UpdateNetWorth)
	v1.GET("/data/:id", admin.GetNetWorth)
	v1.DELETE("/:id", admin.DeleteWorth)
	v1.GET("/list", admin.GetNetWorthList)
	v1.POST("/", admin.InsertNetWorth)
}

// 预约
func AppointmentRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1/appointment")
	v1.PUT("/", admin.UpdateAppointment)
	v1.GET("/data/:id", admin.GetAppointment)
	v1.DELETE("/:id", admin.DeleteAppointment)
	v1.GET("/list", admin.GetAppointmentList)
	v1.POST("", admin.InsertAppointment)
}

// 联系我们
func ContactsRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1/contacts")
	v1.PUT("/", admin.UpdateContacts)
	v1.GET("/data/:id", admin.GetContacts)
	v1.DELETE("/:id", admin.DeleteContacts)
	v1.GET("/list", admin.GetContactsList)
	v1.POST("", admin.InsertContacts)
}

// github 访问接口
func ResumeBaseRouter(r *gin.Engine) {
	r.GET("/get_data", admin.GetResumeData)
	r.POST("/push_data", admin.PushResumeData)
}

// 无需认证的路由示例
func examplesNoCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")
	// 空接口防止v1定义无使用报错
	v1.GET("/nilcheckrole", nil)

	// {{无需认证路由自动补充在此处请勿删除}}
}

// 需要认证的路由示例
func examplesCheckRoleRouter(r *gin.Engine, authMiddleware *jwtauth.GinJWTMiddleware) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")
	// 空接口防止v1定义无使用报错
	v1.GET("/checkrole", nil)

	// {{认证路由自动补充在此处请勿删除}}
}
