
package main
import (
"fmt"
"mysimplemicroservice/services"
)
var appName = "accountservice"



func main() {
	fmt.Printf("Starting %v\n", appName)
	services.StartWebServer("6767")           // NEW
}