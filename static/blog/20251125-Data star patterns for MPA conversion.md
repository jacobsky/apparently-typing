After getting a bit more experience, I think I have to revise my understanding of datastar a little bit.

One of the big hypermedia oriented approaches to web dev is to build an MPA and then sprinkle in your frontend interactivity where you need it. When initially working with datastar (specifically signals and server sent events) I expressed the feeling that unlike most hypermedia applications it was better to build datastar applications _with datastar in mind_ from the start and that it's not great for adding reactity to MPAs.

While there are obvious benefits to starting with datastar, it's also perfectly acceptable to add reactivity to your MPA application.

There are a few tricks to doing so.

### Trick #1 - Differentiate by route wildcards
The first simplest trick is to simply use a bunch of different sub routes, such as by using specific wildcards.

For this blog, I have the 

```
func (h *handler) servehttp(w http.responsewriter, r *http.request) {
	switch r.PathValue("id") {
	case "all":
	    // continously read all posts
        // .. snip
	case "latest":
	    // Shows only the latest post
        // .. snip

	case "":
        // If empty show the base page

	default:
        // For anything else, assume number and post the post
        // .. snip
	}
}
```

This typically works, but does require a bit more reworking when adding the reactivity to your application. While useful and a good starting point there a few better tricks we can use to handle this.

### Trick #2 - Content Negotiation

[Content negotiation](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Content_negotiation) is a ~~trick~~ mechanism of of the web that allows user agents to specify which representation of a resource on the server is desired. This is typically via the use of headers such as `Accept`   which specifies which type of content that it wants such as `text/html`, `text/json`, `text/pdf` and then the server responds with the header `Content-Type` to inform which type of content is actually being returned. 

Datastar automatically send the `Datastar-Request: true` header with every request and it is also possible to send a `headers` object with a request to decorate your requests to include any and all negotiation.

This allows for you to have a handler such as the following:

```
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.Header.Get("Datastar-Request") != "" {
        // Do the datastar events path
    } else {
        // Emit the base HTML (probably the full HTML with no morphing)
    }
}
```


### Trick #3 - Query Parameters

I know, this is yet another web standard, but it's worth mentioning anyways, it is also pretty easy to differentiate the kinds of requests via query parameters such as the following:

```
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Has("Flag") {
        // Do some datastar events
    } else if param := r.URL.Query().Get("Parameter") {
        // Do some datastar events with the query parameters
    }
    // Emit the base HTML (probably the full HTML with no morphing)
}
```

By mixing and matching these three tricks, it makes it relatively easy to build up MPA apps and then convert them into a more rich and reactive datastar application.
