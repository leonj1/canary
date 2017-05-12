# Canary
This is a personal version of camelcamelcamel.com
I wrote it because I felt the above mentioned site notified me too slow when a price I was interested in was reached.

# Build
```
go build -v canary.go
```

# Setup
Sample startup script
```
#!/bin/bash

# I'm not a fan of these env vars, but package "go-ses" depends on them
export AWS_ACCESS_KEY_ID=[your_key]
export AWS_SECRET_KEY=[your_secret]
export AWS_SES_ENDPOINT=https://email.us-east-1.amazonaws.com

/opt/canary/current/canary -user=[db_user] -pass=[db_password] -db=[db_name] -port=[web_server_port]

# make sure to chmod +x this script
```

# Usage
## Define a product to track
1. Find the product you are interested in at Amazon site (the only site currently supported, but extendable)
2. Get the URL and POST to your canary web server
```
# example
curl --silent -d '{"name":"SubPac M2 Wearable", "url":"http://amzn.to/2r0njiW", "target_price":"150.00", "website":"amazon"}' http://your_domain.com:8000/products
```

# List Products being tracked
```
# assumes jq is installed to pretty print json in console
curl --silent http://your_domain:8000/products | jq '.'
```

# Healthcheck
## Ensure the canary process is running at a regular interval
```
curl --silent http://your_domain:8000/executions | jq '.'
```

