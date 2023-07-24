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

### Get API Example

There is an example building a get required and using an ip switch:

```go
func TorGet(targetURL string) string {
 // Set proxy, this sample was working in docker golang:bullseye
 // Using 8118 port 'cause torcontroller setting provixy config to 8118 port.
 proxyURL, err := url.Parse("http://127.0.0.1:8118")
 if err != nil {
  fmt.Println("Failed to parse proxy URL:", err)
  os.Exit(1)
 }
 // build HTTP client
 httpTorClient := &http.Client{
  Transport: &http.Transport{
   Proxy: http.ProxyURL(proxyURL),
  },
 }
 // Send GET require.
 resp, err := httpTorClient.Get(targetURL)
 if err != nil {
  fmt.Println("Failed to send GET request:", err)
  os.Exit(1)
 }
 defer resp.Body.Close()

 // Read the response body.
 body, err := ioutil.ReadAll(resp.Body)
 if err != nil {
  fmt.Println("Failed to read response body:", err)
  os.Exit(1)
 }
 // Print the response body.
 return string(body)
}

func main() {
 // Start tor service.
 gotorcontroller.StartTor()
 fmt.Println(TorGet("https://api64.ipify.org?format=json"))
 //{"ip":"23.137.249.150"}
 //A example current tor exit router ip which responded by ipify.
 gotorcontroller.SwitchIP()
 fmt.Println(TorGet("https://api64.ipify.org?format=json"))
 //{"ip":"2.58.56.37"}
 //A example current tor exit router ip which responded by ipify.
}
```

## Reference

This go lib basic on [torcontroller](https://github.com/Seicrypto/torcontroller), if you're looking for devolep with different languages or more feature. Please read it.
