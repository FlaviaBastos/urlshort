# URL shortener

My implementation of [Gophercises - Exercise #2: URL Shortener](https://gophercises.com/exercises/urlshort), by Jon Calhoun.

## Exercise details

The goal of this exercise is to create an [http.Handler](https://golang.org/pkg/net/http/#Handler) that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for `/dogs` to `https://www.somesite.com/a-story-about-dogs` we would look for any incoming web requests with the path `/dogs` and redirect them.

Some boilerplate code was provided.
