package main 
import "fmt"


func main(){
	newUser3 := user{"Dima","Soss",4}
	allUsers := users{{"Gleb","Hleb1337",1},{"Fillip","fillipinI23",2},{"Jhon","Chikago2007",3}}
	allUsers.create(newUser3)
	allUsers[1].reduce("JhonLoshok")
	allUsers[0].update("Artem","DotaBest39",5)
	allUsers.delete(2)
	fmt.Println(allUsers)
}