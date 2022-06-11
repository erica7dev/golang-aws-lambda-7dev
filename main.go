package main

import(
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type myEvent struct {
	Name string `json:"what is your name?"`
	Age int `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer"`
}

func handlerLambdaEvent(event myEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}
func main(){
	lambda.Start(handlerLambdaEvent) 
}