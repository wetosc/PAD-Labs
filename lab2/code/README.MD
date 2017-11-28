# LAB 2: Distributed Data Collections 

The requierements for this project can be found at: [github.com/Alexx-G/PAD-labs](https://github.com/Alexx-G/PAD-labs/blob/master/PAD-lab2.md).

---

**Client:** requires data from mediator.  
**Mediator:** requires data from central nodes and sends it back to the client.  
**Node:** Accepts GET and REPLY requests. GETs data from connected nodes and sends it as REPLY upstream.  
  
More features:  
- The node implements a filter mechanism. 
- The mediator can serialize data in JSON or XML, on demand.
- The client validates the document (XSD or JSON-Schema).
- The data from nodes is retrieved recursively.

## tcpClient library
---
The library works as a helper on top of net package, oriented on TCP connexions. It provides methods for Read, Write and Async Read, Async Write.
  
  
## Project Dependencies:
---
- zerolog ( "github.com/rs/zerolog/log" ) : Used for logging