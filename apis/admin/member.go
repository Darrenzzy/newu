package admin

import (
	"encoding/json"
	"errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/models"
	"go-admin/tools"
	"go-admin/tools/app"
	"log"
	"net/http"
)

func GetMemberList(c *gin.Context) {
	var res models.Member
	var err error

	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = tools.StrToInt(err, size)
	}
	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = tools.StrToInt(err, index)
	}
	if name := c.Request.FormValue("username"); name != "" {
		res.Username = name
	}
	if name := c.Request.FormValue("mobile"); name != "" {
		res.Mobile = name
	}
	result, count, err := res.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetMember(c *gin.Context) {
	var res models.Member
	res.ID, _ = tools.StringToInt64(c.Param("id"))
	res, err := res.Get()
	tools.HasError(err, " 无该用户", -1)
	var resp app.Response
	resp.Data = res
	c.JSON(http.StatusOK, resp.ReturnOK())
}

func UpdateMember(c *gin.Context) {
	var data models.Member
	var resp app.Response
	err := c.BindWith(&data, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	if data.ID == 0 {
		err = errors.New("")
		tools.HasError(err, "缺省参数", -1)
		c.JSON(http.StatusOK, resp.ReturnOK())
		return
	}
	data, err = data.Update()
	tools.HasError(err, "数据不存在", -1)
	c.JSON(http.StatusOK, resp.ReturnOK())
}

func DeleteMember(c *gin.Context) {
	var data models.Member
	// data.UpdateBy = tools.GetUserIdStr(c)
	IDS := tools.IdsStrToIdsIntGroup("id", c)
	result, err := data.BatchDelete(IDS)
	tools.HasError(err, "删除失败", 500)
	app.OK(c, result, "删除成功")
}

// @Summary 会员注册
// @Tags 企业网站接口
// @Param username query string false "用户名"
// @Param code query int false "验证码"
// @Param password query string false "密码"
// @Param mobile query string false "手机号"
// @Param email query int false "邮箱"
// @Success 200 {string} string	"{"code": 200, "message": "注册成功"}"
// @Router /api/v1/member/register [post]
func RegisterMember(c *gin.Context) {
	var member *models.Member
	err := c.BindWith(&member, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)

	if member.Mobile == "" || member.Email == "" || member.Password == "" || member.Code == "" {
		err = errors.New("")
		tools.HasError(err, "缺省参数", 500)
	}

	if !tools.VerifyPhone(member.Email) {
		err := errors.New("")
		tools.HasError(err, "邮箱格式错误", 500)
		return
	}

	if !tools.VerifyPhone(member.Mobile) {
		err := errors.New("")
		tools.HasError(err, "手机格式错误", 500)
		return
	}

	if m, ok := Code.mobiles[member.Code]; !ok || m != member.Mobile {
		err = errors.New("")
		tools.HasError(err, "验证码错误", 500)
	}
	id, err := member.Insert()
	tools.HasError(err, "注册失败", 500)
	app.OK(c, id, "注册成功")
}

var Code *AuthCodes

type AuthCodes struct {
	mobiles map[string]string
	Ip      map[string]bool
}

type JsonCode struct {
	Code string `json:"code"`
}

// @Summary 发送验证码
// @Tags 企业网站接口
// @Param mobile query string false "手机号"
// @Success 200 {string} string	"{"code": 200, "message": "发送成功"}"
// @Router /api/v1/sendCode [get]
func SendCode(c *gin.Context) {
	mobile := c.Request.FormValue("mobile")

	if mobile == "" {
		err := errors.New("")
		tools.HasError(err, "缺省手机参数", 500)
		return
	}

	if !tools.VerifyPhone(mobile) {
		err := errors.New("")
		tools.HasError(err, "手机格式错误", 500)
		return
	}

	ip := c.ClientIP()
	Code.Ip[ip] = true
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4GKEQsffZ8REKwGxwm4J", "FbhXk6nJDHebhMf5GGCD29yFXxCJYK")
	tools.HasError(err, "初始化失败", 500)

	code := tools.EncodeToString(6)
	Code.mobiles[code] = mobile
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = mobile
	request.SignName = "诺游"
	request.TemplateCode = "SMS_200180609"
	kv := JsonCode{
		Code: code,
	}
	js, err := json.Marshal(kv)
	tools.HasError(err, "格式化失败", 500)
	request.TemplateParam = string(js)

	response, err := client.SendSms(request)
	if err != nil {
		tools.HasError(err, "验证码发送成功", 200)
	}
	log.Printf("response is %v\n", response.Message)
	app.OK(c, "", "发送成功")

}

// @Summary 会员重置密码
// @Tags 企业网站接口
// @Param mobile query string false "手机号"
// @Param password query string false "密码"
// @Success 200 {string} string	"{"code": 200, "message": "重置成功"}"
// @Router /api/v1/member/reset_pass [post]
func ResetPass(c *gin.Context) {
	var member models.Member
	err := c.BindWith(&member, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	if member.Password == "" || member.Mobile == "" {
		// if member.Password == "" || member.Mobile == "" || member.Code == "" {
		err := errors.New("")
		tools.HasError(err, "缺省参数", 500)
		return
	}
	err = member.ResetPass()
	tools.HasError(err, "", 500)
	member.Password = ""
	app.OK(c, member, "success")
}

// @Summary 会员登录
// @Tags 企业网站接口
// @Param mobile query string false "手机号„"
// @Param password query string false "密码"
// @Success 200 {string} string	"{"code": 200, "message": "登录成功"}"
// @Router /api/v1/member/login [post]
func Login(c *gin.Context) {
	var member models.Member
	err := c.BindWith(&member, binding.JSON)
	tools.HasError(err, "数据解析失败", -1)
	if member.Password == "" || member.Mobile == "" || member.Code == "" {
		err := errors.New("")
		tools.HasError(err, "缺省参数", 500)
		return
	}
	if m, ok := Code.mobiles[member.Code]; !ok {
		err := errors.New("")
		tools.HasError(err, "验证码错误", 500)
		return
	} else if m != member.Mobile {
		err := errors.New("")
		tools.HasError(err, "验证码错误", 500)
	}

	err = member.Login()
	if err != nil {
		tools.HasError(err, err.Error(), 500)
	}

	member.Password = ""
	app.OK(c, member, "success")
}

func init() {
	Code = new(AuthCodes)
	Code.mobiles = make(map[string]string)
	Code.Ip = make(map[string]bool)
}
