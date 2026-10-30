package main 
import "fmt"

type users []user;
	
type user struct{
	Name,password string
	id uint
}
func (a *users) create(b user){
	*a = append(*a,b)
}
func  (a *user) reduce(new_password string) {
	a.password = new_password
}
func (a *user) update(new_name,new_password string, new_id uint){
	a.Name = new_name
	a.password = new_password
	a.id = new_id
	}	
func (a *users) delete(n int) error{
	if n < 0 || n >= len(*a) {
		return fmt.Errorf("index %d out of range [0, %d)", n, len(*a))
	}
	*a = append((*a)[:n], (*a)[n+1:]...)
	return nil
}
