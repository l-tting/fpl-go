package main 

import(
	"fmt"
	"github.com/l-tting/fpl-app/internal/db"

)

func main(){
	fmt.Println("FPL kicking!!!")
	db.Connect()
}
