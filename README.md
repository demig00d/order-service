## Usage:
```sh
# Postgres, JetStream
$ make compose-up
# Run app with migrations
$ make run-subscriber
# Run publishing script
$ make run-publisher
```

---
## HTTP Endpoints:
* ```GET localhost/order/:id``` - *display order by id*
* ```GET localhost/healthz``` - *liveness probe*
* ```GET localhost/swagger/*any``` - *docs*

---
## Demo
![](https://github.com/demig00d/order-service/blob/master/demo.gif)

--- 
Created from [Go-clean-template](https://evrone.com/go-clean-template?utm_source=github&utm_campaign=go-clean-template) which is created & supported by [Evrone](https://evrone.com/?utm_source=github&utm_campaign=go-clean-template).
