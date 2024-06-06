## Orders API

A simple [GO tutorial](https://www.youtube.com/playlist?list=PL4cUxeGkcC9iImF8w9FbFOc2UntutL9Wv) from the Net Ninja Youtube channel

### Launch the app:

1. `docker run -p 6379:6379 redis:latest`
2. To test that redis is working open another terminal windor
   - `redis-cli`
   - `KEYS * `
3. In another terminal `go run main.go`

### Request

#### POST

`curl -X POST -d '{"customer_id:"'$uuidgen'","line_items":[{"item_id":"'$uuidgen'","quantity":5,"price":1999}]}' localhost:3000/orders`

To check if it worked

1. Copy the id from the response
2. `redis-cli`
3. `GET "order_id"`
4. `SMEMBERS` to check the set of orders

#### GET ALL

`curl -sS localhost:3000/orders | jq`
`curl -sS localhost:3000/orders?cursor=<next> | jq`

#### GET

`curl -sS "localhost:3000/orders/<order_id>" | jq`

### POST

`curl -X POST -d '{"customer_id:"'$uuidgen'","line_items":[{"item_id":"'$uuidgen'","quantity":5,"price":1999}]}' localhost:3000/orders` create the item
`curl -X POST -d '{"status": "shipped"}' -sS localhost:3000/orders/<order_id> | jq` set the shipped status
`curl -X POST -d '{"status": "shipped"}' -sS localhost:3000/orders/<order_id> | jq` try again the shipped status -> 400
`curl -X POST -d '{"status": "completed"}' -sS localhost:3000/orders/<order_id> | jq`

#### DELETE

`curl -X DELETE localhost:3000/orders/<order_id>`

#### CONFIG

`SERVER_PORT=8080 go run main.go`

## TODO

- [ ] Handle nice error message on BadRequest
- [ ] Add GoDotEnv package to load automatically config from a .env file
- [ ] Try swapping to a new DB like Postegres
- [ ] Automate E2E
