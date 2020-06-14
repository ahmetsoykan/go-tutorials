/*
Context type, which carries deadlines. cancellation signals, and other request-scoped values across API boundaries and between processes.
Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context.
The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, ot WithValue.
When a Context is canceled, all Contexts derived from it are also canceled.

func DoSomething(ctx context.Context, arg Arg) error {
	// ... use ctx ...
}
*/
package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("Server: hello handler started")
	defer fmt.Println("Server:  hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
