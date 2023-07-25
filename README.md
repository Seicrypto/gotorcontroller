# gotorcontroller

now support: [torcontroller](https://github.com/Seicrypto/torcontroller) v1.0-1

[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2FSeicrypto%2Fgotorcontroller&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=hits&edge_flat=false)](https://github.com/Seicrypto/gotorcontroller)

With this package, you can easily control your backend application to use different IP addresses through Tor's service. This enables users to leverage multiple IPs, making it ideal for tasks such as web scraping and other applications that require a large number of IP addresses.

## QuickStart

Step1. Make sure your program works environment already follow [torcontroller](https://github.com/Seicrypto/torcontroller) and install it.

step2. get gotorcontroller lib

```bash
# Terminal
go get -u github.com/Seicrypto/gotorcontroller
# go ... gotorcontroller v1.0.0
```

step3. call torcontroller function.

Just call the function while you need a different ip address. There is an example building a get required and using an ip switch.

```go
func main(){
 // Start tor service.
 gotorcontroller.StartTor()
 //now 8118 port exit tor router ip: {"ip":"23.137.249.150"}
 gotorcontroller.SwitchIP()
 //now 8118 port exit tor router ip: {"ip":"2.58.56.37"}
}
```

See more in [examples](./examples/)

### Get API Example

There is an example building a get required and using an ip switch:

[GetapiWithTorcontrollerInGo](./examples/exampleGetapi.go)

You can run it with Github codespaces:

![GetAPIWithTor](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExd3h6ZjkyYjd1dDV2ODl2MnozY2k2ejE2YzFoYTYwbnF2dHplY2loeiZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/EdJIvL6y5eAWdnJq7k/giphy.gif)

You can try my example on github codespaces.

![Github Codespace](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExYjJkY2U3MXp0ZjJuYzh4ejFyb25neHRwNTdweWR3ZzlsOHlmYWhmbCZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9Zw/tqMK5aGxsyWvOCP9Z5/giphy.gif)

## Reference

This go lib basic on [torcontroller](https://github.com/Seicrypto/torcontroller), if you're looking for devolep with different languages or more feature. Please read it.
