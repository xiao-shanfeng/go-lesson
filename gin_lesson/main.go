package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	// 禁止控制台显示颜色
	//gin.DisableConsoleColor()
	// 强制控制台显示颜色
	gin.ForceConsoleColor()
	// gin默认中间件
	r := gin.Default()
	//r := gin.New()

	// demo
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "success",
		})
	})

	// AsciiJson
	r.GET("/asciiJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "go语言",
			"tag":  "<br>",
		}
		//c.JSON(200, data)
		// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON
		c.AsciiJSON(200, data)
	})

	// html模版
	//r.LoadHTMLGlob("template/*")
	r.LoadHTMLFiles("template/index.tmpl")
	r.GET("/html", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title": "Go",
		})
	})

	// post请求参数	Multipart/Urlencoded 绑定
	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		// 显式绑定参数
		/*
			postman中使用post请求form-data、x-www-form-urlencoded传参数
		*/
		//err := c.ShouldBindWith(&form, binding.Form)
		// 简单的使用ShouldBind
		err := c.ShouldBind(&form)
		if err == nil {
			if form.User == "Tom" && form.Password == "123456" {
				c.JSON(200, gin.H{"message": "login success"})
			} else {
				c.JSON(401, gin.H{"message": "unauthorized"})
			}
		} else {
			fmt.Println(err)
		}
	})

	// post请求参数	Multipart/Urlencoded 表单
	r.POST("/formPost", func(c *gin.Context) {
		/*
			postman中使用post请求form-data、x-www-form-urlencoded传参数
		*/
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "success",
			"message": message,
			"nick":    nick,
		})
	})

	// PureJSON
	r.GET("/json", func(c *gin.Context) {
		// unicode实体
		//c.JSON(200, gin.H{
		//	"html": "<h1>Hello,World!</h1>",
		//})

		// 字面字符
		c.PureJSON(200, gin.H{
			"html": "<h1>Hello,World!</h1>",
		})
	})

	// 请求参数	Query 和 post form
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.Query("page")
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.JSON(200, gin.H{
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
	})

	// SecureJSON
	r.GET("/someJson", func(c *gin.Context) {
		names := []string{"tom", "jack", "mary"}
		//c.JSON(200, names)
		c.SecureJSON(200, names)
	})

	//XML/JSON/YAML/ProtoBuf 渲染
	r.GET("/moreJson", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Tom"
		msg.Message = "Success"
		msg.Number = 200
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "Hi", "status": http.StatusOK})
	})
	r.GET("/someYaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hi", "status": http.StatusOK})
	})
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}

		c.ProtoBuf(http.StatusOK, data)
	})

	// 上传文件
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// 单个文件
		//file, _ := c.FormFile("file")
		//fmt.Println(file.Filename)
		//dst := "./" + file.Filename
		// 多个文件
		form, _ := c.MultipartForm()
		files := form.File["files"]
		for _, file := range files {
			dst := "./" + file.Filename
			_ = c.SaveUploadedFile(file, dst)
		}
		//_ = c.SaveUploadedFile(file, dst)
		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
