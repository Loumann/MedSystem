package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"somename/configs"
	"somename/handler"
	"somename/models"
	"somename/repository"
)

var upgrader = websocket.Upgrader{}

func main() {
	config, err := configs.GetConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	db, err := repository.GetDatabaseConnection(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := repository.GetRepository(db)
	h := handler.GetHandler(r)

	router := GetRouter(h)
	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func GetRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("template/*.html")
	router.Static("/assets", "assets")

	router.GET("/", h.LoginPage)
	router.GET("/index", h.IndexPage)
	router.POST("/create-user", h.CreateUser)
	router.POST("/delete-user", h.DeleteUser)

	router.POST("/sign-in", h.SignIn)
	router.GET("/users", h.GetUsers)
	router.GET("/users/:search", h.GetUsers)
	router.GET("/analyses/:id", h.GetAnalyse)

	//хранилище для куки
	router.GET("/wait-analyse", waitAnalyse)
	router.GET("/ws", websocketHandler)

	return router
}

func waitAnalyse(c *gin.Context) {
	// Обработка нажатия кнопки "Wait analyse"
}

// передача по порту
func websocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	for {
		// Читаем сообщение от клиента
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		// декомпиляция json
		var analysis models.Analysis
		err = json.Unmarshal(message, &analysis)
		if err != nil {
			log.Println(err)
			continue
		}

		// Записываем данные в базу данных
		//err = writeDataToDB(analysis)
		//if err != nil {
		//	log.Println(err)
		//	continue
		//}

		deviceMessage := fmt.Sprintf("Анализ пользователя %s был успешно сохранен на сервер!")

		// oтправляем сообщение клиенту
		err = conn.WriteMessage(websocket.TextMessage, []byte(deviceMessage))
		if err != nil {
			log.Println(err)
			break
		}
	}
}
