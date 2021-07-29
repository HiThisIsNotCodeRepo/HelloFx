package main

import (
	"fmt"
	"go.uber.org/fx"
)

type SimpleProvider struct {
	AString string
}

func ProvideSimpleProvider() *SimpleProvider {
	return &SimpleProvider{
		"Hi I'm a simple provider",
	}
}

func SimpleReceiver(sp *SimpleProvider) {
	fmt.Printf("Hi I'm simple receiver, now I'm going to consume simpler provider : %v \n", sp.AString)
}
func main() {
	fx.New(
		fx.Provide(ProvideSimpleProvider),
		fx.Invoke(SimpleReceiver),
	).Run()
}
