package cmd
import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)
func main(){
	err:=godotenv.Load()

	if err!=nil{
		fmt.Println(err)
	}
	app:=fiber.New();
	app.Use(logger.New())
}