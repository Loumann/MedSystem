package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"somename/configs"
	"somename/handler"
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
	router.DELETE("/delete-user/:id", h.DeleteUser)

	router.POST("/sign-in", h.SignIn)
	router.GET("/users", h.GetUsers)
	router.GET("/users/:search", h.GetUsers)
	router.GET("/analyses/:id", h.GetAnalyse)

	//хранилище для куки
	//router.POST("/wait-analyse/:id")
	//router.POST("/analysis/:id", writeAnalysis)

	return router
}

// передача по порту
//func writeAnalysis(c *gin.Context) {
//	idstr := c.Param("id")
//	fmt.Println(idstr)
//
//	// принять ID пользователя
//
//	response, e := http.Get("https://localhost:7118/api/urs")
//	if e != nil {
//		fmt.Println(e)
//		return
//	}
//
//	bytes, e := io.ReadAll(response.Body)
//	if e != nil {
//		fmt.Println(e)
//	}
//	analysis := &models.Analysis{}
//	if err := json.Unmarshal(bytes, analysis); err != nil {
//		log.Fatal(err)
//	}
//
//	db, err := sql.Open("postgres", "host=10.14.206.28 user=postgres password=*sJ#44dm dbname=medbase sslmode=disable")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	stmt, err := db.Prepare(`INSERT INTO "Analise" ("Date", "Bld", "Ubg", "Bil", "Pro", "Nit", "Ket", "Glu", "PH", "SG", "Leu")
//                             VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning "ID"`)
//	if err != nil {
//		log.Fatal(err)
//	}
//	row := stmt.QueryRow(analysis.Date, analysis.Bld, analysis.Ubg, analysis.Bil, analysis.Pro, analysis.Nit, analysis.Ket, analysis.Glu, analysis.PH, analysis.SG, analysis.Leu)
//	err = row.Scan(&analysis.ID)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	//row := db.Exec(`INSERT INTO "UserAnalise"`)
//
//	c.Status(http.StatusOK)
//}
