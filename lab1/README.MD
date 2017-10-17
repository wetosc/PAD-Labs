# LAB 1:Message Brocker

The first implementation of the system was something like:
- sender publishes a message
- receiver asks for a message
- the "broker" (server) stores the messages in a queue, creating backups on disk sometimes

It worked well, but at higher levels of the lab I had to implement the "publish-subscribe" pattern.

The new architecture passes the responsability of sender and receiver to a new object, "client".
The messages remains almost the same, but the verbs changes completely:
```json
{
    "type": "The verb / action of the message",
    "queue": "Optional parameter, specifies in which queue the message will be published",
    "info": "Optional parameter, the message payload"
}
```
The architecture is inspired a little from MQTT standard  
Possible verbs:  
- **PUBLISH** : Client sends a message to the broker
- **DELIVER** : Broker sends a message to the subscriber
- **SUBSCRIBE** : Subscribes a client to the queue "queue"
- **RESPONSE** : Notifies that a transaction completed succesfully. Usually can be ignored
- **ERROR** : Signifies an error in the request, like trying to subscribe to an unexistent queue.  
  
  
## Eumgent library
---
The library has 2 important entities: Client and Room. The client reads and writes messages. The room accepts connections (clients) and the messages from them. Also, because it stores the clients, it can send messages back. 
  
  
## Project Dependencies:
---
- zerolog ( "github.com/rs/zerolog/log" ) : Used for logging