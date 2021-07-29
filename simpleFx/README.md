# simpleFx
## Provider
```
type SimpleProvider struct {
	AString string
}

func ProvideSimpleProvider() *SimpleProvider {
	return &SimpleProvider{
		"Hi I'm a simple provider",
	}
}
```
## Receiver
```
func SimpleReceiver(sp *SimpleProvider) {
	fmt.Printf("Hi I'm simple receiver, now I'm going to consume simpler provider : %v \n", sp.AString)
}
```
## How to use
1. Register provider and receiver in `fx.New()`
```
fx.New(
    fx.Provide(ProvideSimpleProvider),
    fx.Invoke(SimpleReceiver),
)...
```
2. Run
```
fx.New(
    ...
).Run()
```
![](https://i.imgur.com/3N49NYD.png)