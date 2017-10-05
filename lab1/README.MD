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
- **UNSUBSCRIBE** : Unsubscribes a client from the queue "queue"
- **RESPONSE** : Notifies that a transaction completed succesfully. Usually can be ignored
- **ERROR** : Signifies an error in the request, like trying to subscribe to an unexistent queue.


## Project Dependencies:
---
- zerolog ( "github.com/rs/zerolog/log" ) : Used for logging