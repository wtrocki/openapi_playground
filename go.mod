module github.com/wtrocki/openapi_playground

go 1.18

replace github.com/wtrocki/openapi_playground/contenttype => ./contenttype

require golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be

require (
	github.com/golang/protobuf v1.4.2 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)
