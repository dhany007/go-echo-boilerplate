# Go Echo Framework Boilerplate
## Prerequisites
- Go 1.20.4
- Echo https://echo.labstack.com/
- not use database, just variable for

## API Specs
### Add Products
- Endpoint: ```POST {{host}}/products```
- Param: -
- Body:
  - `id: string`
  - `name: string`
  - `quantity: int64`
  - `price: int64`
  ```
  {
    "id": "P001",
    "name":"jambu",
    "quantity": 100,
    "price": 1000
  }
  ```
- Response:
  ```
  {
    "status": 201,
    "messages": "success",
    "data": {
      "id": "P001",
      "name": "jambu",
      "quantity": 100,
      "price": 1000
    }
  }
  ```
### Get List Products
- Endpoint: ```GET {{host}}/products```
- Param: -
- Body: -
- Response:
  ```
  {
    "status": 200,
    "messages": "success",
    "data": [
      {
        "id": "P001",
        "name": "jambu",
        "quantity": 100,
        "price": 1000
      },
      {
        "id": "P002",
        "name": "jambu air",
        "quantity": 100,
        "price": 1000
      }
    ]
  }
  ```
### Get Detail Products
- Endpoint: ```GET {{host}}/products/:id```
- Param:
  - `id: string`
- Body:-
- Response:
  ```
  {
    "status": 200,
    "messages": "success",
    "data": {
      "id": "P001",
      "name": "jambu",
      "quantity": 100,
      "price": 1000
    }
  }
  ```
### Update Products
- Endpoint: ```PUT {{host}}/products/:id```
- Param:
  - `id: string`
- Body:
  - `name: string`
  - `quantity: int64`
  - `price: int64`
  ```
  {
    "name":"jambu air baru",
    "quantity": 100,
    "price": 1000
  }
  ```
- Response:
  ```
  {
    "status": 200,
    "messages": "success",
    "data": {
      "id": "P002",
      "name": "jambu air baru",
      "quantity": 10,
      "price": 20000
    }
  }
  ```
### Remove Products
- Endpoint: ```DELETE {{host}}/products/:id```
- Param: -
- Body: -
- Response:
  ```
  {
    "status": 200,
    "messages": "success",
    "data": "P002"
  }
  ```

