package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/go-playground/validator/v10"
	"github.com/op/go-logging"
	"io"
	"net/http"
	"os"
	"time"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type PostParams struct {
	Name string `json:"name"`
	Age  int    `json:"age" binding:"required,mustBig"`
	Sex  bool   `json:"sex"`
}

func mustBig(fl validator.FieldLevel) bool {
	fmt.Println(fl)
	fmt.Println(fl.Field())
	fmt.Println(fl.Field().Interface())
	fmt.Println(fl.Field().Interface().(int))
	if fl.Field().Interface().(int) > 18 {
		return true
	}
	return false
}

func middle() func(*gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("进入中间件1")
		t := time.Now()
		c.Next()
		fmt.Println("中间件1时间", time.Since(t))
		fmt.Println("离开中间件1")
	}
}

func middle1() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		fmt.Println("进入中间件2")
		t := time.Now()
		c.Next()
		fmt.Println("中间件2时间", time.Since(t))
		fmt.Println("离开中间件2")
	}
}

var log = logging.MustGetLogger("goLogging")
var format = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)

func main() {
	// go-logging使用
	backend := logging.NewLogBackend(os.Stderr, "", 0)

	backendLevel := logging.AddModuleLevel(backend)
	backendLevel.SetLevel(logging.ERROR, "")

	log.Info("info")
	log.Error("error")
	log.Warning("warning")
	log.Debug("debug")
	log.Notice("notice")
	log.Critical("critical")
	// 禁止控制台显示颜色
	//gin.DisableConsoleColor()
	// 强制控制台显示颜色
	//gin.ForceConsoleColor()

	// gin日志
	//gin.DisableConsoleColor()

	//file, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	// gin默认中间件
	r := gin.Default()
	//r := gin.New()

	// 自定义日志格式
	//r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
	//	// 自定义格式
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		params.ClientIP,
	//		params.TimeStamp.Format(time.RFC1123),
	//		params.Method,
	//		params.Path,
	//		params.Request.Proto,
	//		params.StatusCode,
	//		params.Latency,
	//		params.Request.UserAgent(),
	//		params.ErrorMessage)
	//}))
	//r.Use(gin.Recovery())

	//r := gin.New()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mustBig", mustBig)
	}

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

	r.POST("/testParams", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindJSON(&p)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"msg":  "接口失败",
				"data": gin.H{},
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "接口成功",
				"data": p,
			})
		}
	})

	// 回显
	r.POST("/testUpload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		out, _ := file.Open()
		defer out.Close()
		in, _ := os.Create("./" + file.Filename)
		defer in.Close()
		io.Copy(in, out)
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment;filename=%s", file.Filename))
		c.File("./" + file.Filename)
	})

	v1 := r.Group("v1").Use(middle()).Use(middle1())
	v1.GET("test", func(c *gin.Context) {
		fmt.Println("我是分组方法内部")
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}
