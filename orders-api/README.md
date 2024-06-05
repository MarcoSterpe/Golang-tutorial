## Orders API

A simple [GO tutorial](https://www.youtube.com/playlist?list=PL4cUxeGkcC9iImF8w9FbFOc2UntutL9Wv) from the Net Ninja Youtube channel

To launch the app:
1. `docker run -p 6379:6379 redis:latest`
2. To test that redis is working open another terminal windor
   - `redis-cli`
   - `KEYS * ` 
3. In another terminal `go run main.go`