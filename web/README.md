# HTTP verbs demo

The [client](verb.go) sends requests to a [local demo server](verbserver/main.go).
Run both commands from this directory, in separate terminals:

1. Start the server: `go run ./verbserver`.
2. Run the client: `go run verb.go`.

The server listens on `127.0.0.1:3000`. `GET /` returns a greeting,
`POST /post` echoes a JSON body, and `POST /postform` echoes form fields as JSON.
If the client reports `connection refused`, check that the server is still
running. Stop the server with Ctrl+C when finished.
