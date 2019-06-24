# Company app

Company application was created as an example how to write layered architecture.
The main idea is to separate application on 3 layers .
1) Http layer. `handler` in this app. 
Gateway to application. As example: all that is related with http has to be stored on this layer.
On this layer we work with requests to our application.
In addition: on this layer we declare interface to the next leyer. And communicate with next layer only via this interface implementation.
Because of this interface we can mock next layer and test only logic in current layer.

2) Business logic layer. `controller` in this app.
All complutation have to be in this layer.
In this layer we implement interface that is declared on layer above. We can change implementation as we need, but we just have to implement
that interface.
As previous layer this layer has interface to the next leyer. And communicate with next layer only via this interface implementation.
And again we can mock next layer and test only logic in current layer. We don't have to thinks which datastore (next layer) will we use, just 
check business logic.

3) Storage layer. `datastore` in this app.
Just write the code for datastore.
Again it doesn't matter which storage we use, it must implement interface that was declared on business logic layer.

## Run app

    go run main.go

## Requests

Create new employee:

    curl -X POST \
        http://localhost:8080/v1/employee \
        -H 'Content-Type: application/json' \
        -H 'cache-control: no-cache' \
        -d '{
            "id": 1,
            "first_name": "name1",
            "last_name": "name2",
            "age": 32,
            "salary": 140
        }'

Get employee by id:

    curl -X GET \
        http://localhost:8080/v1/employee/1 \
        -H 'cache-control: no-cache'

Raise salary:

    curl -X POST \
        http://localhost:8080/v1/employee/raise \
        -H 'Content-Type: application/json' \
        -H 'Postman-Token: c9381f02-3a57-42d0-8416-3dfe32322901' \
        -H 'cache-control: no-cache' \
        -d '{
            "amount": 200,
            "id": 1
        }'

Delete employee:

    curl -X DELETE \
        http://localhost:8080/v1/employee/1 \
        -H 'Postman-Token: 3f81b333-28df-46c9-82fc-8eb059936c07' \
        -H 'cache-control: no-cache'