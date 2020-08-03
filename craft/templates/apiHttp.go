package templates

// APIHttp ...
var APIHttp = `
# metrics
POST http://localhost:3000/metrics HTTP/1.1
content-type: application/json

{
}

###

# status
POST http://localhost:3000/status HTTP/1.1
content-type: application/json

{
}

###

`
