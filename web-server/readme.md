Notes: For Building a simple web Server with golang.
- http handleFunc registers a function to a path that we can run with our web browser eg localhost:port/path.
- The server listens on port 10000 and the handler is called when someone visits this page, we use the httpListenAndServe for that purpose.

- http request has a field called body, and body is an io read closer. It implements the interface I/O closer. IT technically means that you can read from any golang standard library stream. One of the stream is the ioutil.ReadAll from the ioutil library.

- created a function named "goodBye" to print some simple data,basically strings. The function is called when we access /goodbye on our web browser.

- ReadAll returns an error so thats why we added the err variable to the code.

- The http.Error function takes three values which are:
    responseWriter, error message, httpStatusCode.
