# URL shortener

My implementation of [Gophercises - Exercise #2: URL Shortener](https://gophercises.com/exercises/urlshort), by Jon Calhoun.

## Exercise details

The goal of this exercise is to create an [http.Handler](https://golang.org/pkg/net/http/#Handler) that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for `/dogs` to `https://www.somesite.com/a-story-about-dogs` we would look for any incoming web requests with the path `/dogs` and redirect them.

Some boilerplate code was provided.


## Notes from research

ServeMux = HTTP request router: it gets the request, looks at the path and check if there's a **handler** to manage that route. (* Remember working with React Router and routes matching precedence).

Handler = responsible for writing headers and body. Virtually any object can be a handler as long as they match the `http.Handler` interface: 

```
ServeHTTP(http.ResponseWriter, *http.Request)
```

HandlerFunc = turns a regular function into satisfying the handler interface (the ServeHTTP signature).


### Flow example

Router receives path ==> Router looks for match ==> Handler for matched route takes over ==> Handler creates header and body ==> Handler is registered for use with Router ==> Server is started and listens for incoming requests

^ This is the code order; when it executes, server is actively listening
