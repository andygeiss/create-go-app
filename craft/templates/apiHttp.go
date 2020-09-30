package templates

// APIHttp ...
var APIHttp = `
# metrics
POST http://localhost:3000/metrics HTTP/1.1
content-type: application/json

{
}

###

// {{ range $i, $name := .Services }}# {{ lc $name }}
POST http://localhost:3000/{{ lc $name }} HTTP/1.1
content-type: application/json

{
}

###
{{ end }}
`
