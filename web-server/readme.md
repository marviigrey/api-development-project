Notes: For Building a simple web Server with golang.
- http handleFunc registers a function to a path that we can run with our web browser eg localhost:port/path.
- The server listens on port 10000 and the handler is called when someone visits this page, we use the httpListenAndServe for that purpose.

- http request has a field called body, and body is an io read closer. It implements the interface I/O closer. IT technically means that you can read from any golang standard library stream. One of the stream is the ioutil.ReadAll from the ioutil library.

- created a function named "goodBye" to print some simple data,basically strings. The function is called when we access /goodbye on our web browser.

- ReadAll returns an error so thats why we added the err variable to the code.

- The http.Error function takes three values which are:
    responseWriter, error message, httpStatusCode.

=========================================================
Moving Forward with the web server, we are going to be creating a package for http handlers, we will create a folder called handlers and pass in the files for our http handler codes inside.
- Create a method to define the http handler interface.
- This is done by creating a new type called Handler and implementing the ServeHTTP method from the http.Handler interface.
- An HTTP handler is an interface which has a single method known as the serveHTTP that takes ResponseWriter and points to the *request.
- For us to create a handler,we will need to create a struct which implements the interface HTTP.

- l.logger: we made this implementation of dependency injection for the purpose of handling logging on a broader level, either logging it in stdout or logging to a file.