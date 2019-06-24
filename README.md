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
