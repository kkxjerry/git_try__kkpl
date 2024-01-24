package main

import (
	"fmt"
	"reflect")

type User struct {
	Name  string
	Email string
	}
func main(){
	
	user := User{Name: "John Doe", Email: "johndoe@example.com"}
	fileMtthod(user)
	
}
func fileMtthod(input interface{}){
	getType := reflect.TypeOf(input)
	getValue := reflect.ValueOf(input)
	
	fmt.Println("Type:", getType,"Value:", getValue)
	for i :=0; i<getType.NumField(); i++{
		field := getType.Field(i)
		fmt.Println( getType.Field(i).Name,reflect.ValueOf(getType.Field(i)),getType.Field(i).Type)
	}
	
	
}