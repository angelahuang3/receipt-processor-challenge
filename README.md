# **Receipt Processor Challenge**

## **Overview**
This project implements a receipt processing API that:
1. Accepts receipt data in JSON format via REST API  
2. Calculates reward points based on predefined business rules  
3. Built with Go (Golang) and Gin framework 
4. Includes unit tests for API validation  

## **Tech Stack**
1. Language: Go
2. Framework: Gin
3. Storage: In-memory Database
4. Test Suite: Go built-in testing package
## **Installation**

### **1. Clone the repository**

```sh
git clone https://github.com/angelahuang3/receipt-processor-challenge.git
cd receipt-processor-challenge
```

### **2. Install all require dependencies**

```sh
go mod tidy
```

### **3. Run the server**

```sh
go run main.go
```
The server will start in http://localhost:8081/     

## **API Endpoints & Testing**

### **1. Run unit tests**

```
go test ./tests/...
```

### **2. Test APIs with cURL**

### ***1. Send a Receipt for Processing***
This process send the receipt and return generated unique id. <br/>

a. Endpoint:
```
POST /receipts/process
```
b. Example:

```
curl -X POST "http://localhost:8081/receipts/process" \
     -H "Content-Type: application/json" \
     -d '{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}'
```

### _2. copy the generated id_

```
{"id": "generated-uuid"}
```

### _3. Get reward points for the receipt_
This endpoint returns the calculated reward points for a populated receipt. <br/>

a. Endpoint:
```
GET /receipts/{id}/points
```
b. Example:
```
curl -X GET "http://localhost:8081/receipts/<RECEIPT_ID>/points"
```

### Expected result
```
{"points":109,"receipt_id":"generated-uuid"}
```
