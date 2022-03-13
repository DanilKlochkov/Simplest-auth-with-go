# Simplest-auth-with-go
<img alt="GitHub watchers" src="https://img.shields.io/github/watchers/DanilKl4/Simplest-auth-with-go?style=social">
It's a learning project. To run it, navigate to a folder in your console. 
Then type to start the server:

`go run main.go`

Then go to http://localhost:5000/sign-up-form

You should see the following:
***
![image](https://user-images.githubusercontent.com/72443284/158067597-26f597e2-206e-464e-9d08-cb4d3552c4fe.png)
***
You can also change the listening server in [main.go](./main.go):


```
func main() {
    http.HandleFunc("/", userHandler)
	  http.ListenAndServe("YourServer:Port", nil)
}
```
