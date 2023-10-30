package main
import "test.com/test/intro/initializers"
import "test.com/test/intro/models"
import "fmt"


func init(){
	initializers.ConnectToDb()
}


func main(){
	 // Migrate the schema
	 fmt.Println( initializers.DB)
	 initializers.DB.AutoMigrate(&models.Post{})
}
