# Lock Order Problem (Kitara Store)

- Your mom has an online store name is "Kitara Store".
- She has one big problem, in many time customer order requested, the products will produce minus
quantity.
- When the product stock is limited and many customers order in the same time, the issue always occur.

## Solution
- use many layer checking & transactional on order

## Requirements

- Go 1.13+
- Go mod (for dependencies)

update dependencies:
- go mod download
- go mod verify
- go mod tidy

## Technology
- Go 1.13+
- Mysql 
- GORM
- Go Gin Framework
- Go Mock 

## Run Application
### Run Migration
You need to ensure that mysql already running on your machine and edit environment file `.env` with your setting.

Run migration with command `go run main.go migrate`

### Run Application
ensuring all dependencies with go mod

build application with script `go build main.go`

run application `./main`

run app with auto build `go run main.go`

## Run Tests
### Using Framework
execute this command to run test with go test framework

`go test github.com/mariojuzar/kitara-store/tests`


## API Documentation
1. Get All Product `GET /api/v1/products`
   
   This endpoint is to show the seed data products that available from migration. 
   You can use this product id to test for the lock order
   
2. Lock Order `POST /api/v1/order/lock`

   This endpoint used to lock order from customer. This endpoint accept user_id and product that want to order.
   This endpoint support lock order for multiple product.
   
   Sample request:
   
   ```
   {
     "user_id": 2,
     "products": [
       {
         "product_id": 1,
         "quantity": 2
       },
       {
         "product_id": 2,
         "quantity": 1
       }
     ]
   }
   ```
   
   When succes lock the order, the expected response:
   
   ```
   {
     "server_time": "0001-01-01T00:00:00Z",
     "code": 200,
     "message": "OK",
     "data": {
       "order_id": 11,
       "order_details": [
         {
           "order_detail_id": 21,
           "product_id": 1,
           "product_name": "masker",
           "quantity": 2
         },
         {
           "order_detail_id": 22,
           "product_id": 2,
           "product_name": "sikat gigi",
           "quantity": 1
         }
       ]
     }
   }
   ```
   
   If some product not available then the expected result:
   
   ```
   {
     "server_time": "0001-01-01T00:00:00Z",
     "code": 400,
     "message": "some product not available",
     "data": null
   }
   ```
   
   But, this endpoint only support for lock order for single store. If many store in request then system will reject it
   
   ```
   {
     "server_time": "0001-01-01T00:00:00Z",
     "code": 400,
     "message": "product must in same store",
     "data": null
   }
   ```