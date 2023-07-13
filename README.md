
# Payment Gateway

This is an implementation of a payment gateway technical test, ref: https://github.com/processout-hiring/payment-gateway-hassan-Sabeh/tree/master

### File structure
the Go framework used in this project is Echo and so the project structure is compliant with the good practices of an Echo project.

#### Files:
- cmd/service/main.go: the main binray to launch the payment gateway server.
- hanlders: package for handler functions containing business logic for requests.
- database: a database simulation object and methods, mocking a database for the payment gateway.
- helpers: All helpers used throughout the project.
- bank_gateway_simulation: the interface and implementation of the bank gateway, in this case the bank is a simulation.
- models: Model objects use in requests, json binding and db.


## Run
After cloning the repo to run the server, you should have Go installed version 1.13 or higher.

Install dependencies at root:

```
$ go mod tidy
```
Run the server:

```
$ go run cmd/service/main.go
```
It will take a few seconds to launch, by default the server will run on port:1234 on your local host, changing the port can be done in the main.go.

### Docker
To run with docker, build the image (make sure you are at the root of the project):

```
docker build -t payment_gateway .
```

After the image is successfully built, you can run a container based on the image, the -p flag is to map the port inside the container to outside on the localhost, run the following command:

```
docker run --publish 1234:1234 payment

```
**Note**: if you choose to change the port, make sure you do it on the application level as mentioned earlier before publishing it with the run command.



## API reference

Now that your application is ready, you can start testing, open a new terminal for curl or any other http client. Below is the API reference.

#### Make payment

`POST /process-payment`

#### JSON body

| Fields | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `card_number` | `string` | **Required**  masked card number |
| `expiry_month` | `int` | **Required**  card expiration month |
| `expiry_year` | `int` |  **Required**  card expiration year |
| `amount` | `float` |  **Required**  payment amount |
| `currency` | `string` |  **Required**  transaction currency |
| `cvv` | `string` |  **Required**  card sercret key |

Example:

    curl -X POST -H "Content-Type: application/json" -d '         {                                     
        "card_number": "1234567891234567",
        "expiry_month": 12,
        "expiry_year": 2013,
        "amount": 99.99,
        "currency": "euro",
        "cvv": "123"
    }' http://localhost:1234/process-payment

### Response

    {
       "data":{
            "payment_id":"c59e293b-caae-49b6-a7d8-4d4c3ad66c64",
          "card_info":{
             "card_number":"XXXXXXXXXXXX4567",
             "expiry_month":12,
             "expiry_year":2013,
             "amount":99.99,
             "currency":"euro",
             "cvv":"123"
          },
          "processed_at":"2023-07-13T14:03:22.  621455276+02:00",
          "payment_status":"Fail"
       }
    }

#### Get Payment

`GET /payment/${id}`

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required**. Payement id |

Example:

    curl -X GET http://localhost:1234/payment/c59e293b-caae-49b6-a7d8-4d4c3ad66c64

#### JSON Response

| Fields | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `card_number` | `string` |  masked card number |
| `expiry_month` | `int` | card expiration month |
| `expiry_year` | `int` |  card expiration year |
| `amount` | `float` |  payment amount |
| `currency` | `string` |  transaction currency |
| `cvv` | `string` |  card sercret key |

Example:

    {
       "data":{
            "payment_id":"c59e293b-caae-49b6-a7d8-4d4c3ad66c64",
          "card_info":{
             "card_number":"XXXXXXXXXXXX4567",
             "expiry_month":12,
             "expiry_year":2013,
             "amount":99.99,
             "currency":"euro",
             "cvv":"123"
          },
          "processed_at":"2023-07-13T14:03:22.  621455276+02:00",
          "payment_status":"Fail"
       }
    }

#### Assumptions
The following assumptions were taken during the development process:
- The credit card number is 16 digits
- The bank interface only has one concrete method which is for processing the payment.
- Some validation is done on the client side for credit card form, otherwise gateway returns bad request.
- Failed payments should also be persisted for tracebility and troubleshooting.
- Payment success and Failure rate with the bank interface is random for testing and simulation.

#### Improvements
The following are a few suggestions for improvement on the service
- **Interface for the Database** should be added for modularity and flexibility if db changes.
- **Docker** for Containerization and deployment.
- Unit testing with **Ginkgo**.
- Instead of implementing Go interfaces for bank simulation and DB, use **gRPC**.
- Use a **JSON schema validator/generator** such as **go-swager** or **gojsonschema**.
- Use a **Logging** package or enhance middleweare in Echo for logging services.

#### Cloud technologies
- **AWS**: Cloud services solution, quick and easy to scale, especially for payment transaction services
- **CockroachDB**: A DB that supports PostgreSQL, fast and easy intigration with go, highly resiliant with multiple nodes and clusters, very optimized written in go.
- **Prometheus and Grafana**: Easy integration for monitoring and visualization into cloud services, helps alot in troubleshooting issues with services and nodes in the cloud. 
