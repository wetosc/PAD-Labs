## Message Broker Protocol
Code name: eumgemt (eugenius message agent)

The messages are using json formatting, in the following format:
```
{
    type:   String    (required)
    queue:  String    (optional)
    info:   String    (optional)
}
```

Possible verbs for the "type" parameter:  
- **PUBLISH** : Client sends a message to the broker
- **DELIVER** : Broker sends a message to the subscriber
- **SUBSCRIBE** : Subscribes a client to the queue "queue"
- **UNSUBSCRIBE** : Unsubscribes a client from the queue "queue"
- **RESPONSE** : Notifies that a transaction completed succesfully. Usually can be ignored
- **ERROR** : Signifies an error in the request, like trying to subscribe to an unexistent queue.
