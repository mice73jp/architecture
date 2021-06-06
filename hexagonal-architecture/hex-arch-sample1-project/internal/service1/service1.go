package service1

import (
	"context"
	"fmt"
)

// AppCoreLogciIn .
type AppCoreLogicIn struct {
	From    string
	Message string
}

// AppCoreLogic
func AppCoreLogic(ctx context.Context, in AppCoreLogicIn) {
	fmt.Println("--------------------------------------------------")
	fmt.Println("service1")
	fmt.Println("this is an application core logic-1.")
	fmt.Printf("from: %s, message: %s\n", in.From, in.Message)
	fmt.Println("--------------------------------------------------")
}
