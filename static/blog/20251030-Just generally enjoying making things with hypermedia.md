Over the last week, I have had a very pleasant time working with
datastar. Originally, I had intended to stick primarily to HTMX and Alpine
for the frontend reactivity, but Datastar has largely won me over (for
personal projects at least) due to it's heavy emphasis on easily including
both aspects of reactivity into my applications. In fact, I got another idea
that is probably only possible with datastar (or a SPA, but let's be real,
SPAs are just not what I'm interested in).

What I've really kind of nailed down over the last week is that I feel that
the discussion with the web is _usually_ all about trade-offs in performance,
simplicity, richeness, etc. But the strangest thing about datastar is that
it feels like... I get everything and I get it all for free? Which makes me
kind of question the modern web, like really?

In reality this is kind of the perfect approach where the tool itself exposes
very simple tooling that can be used to make very expressive web applications.

**NOTE**: Datastar is still 100% hypermedia focused. Interactions are
not driven by the client, but rather driven by and streamed from the
server. Datastar is not going to replace some of the frontend frameworks
that can turn the web into a thick client delivery system (though, I
feel like a native app is a better fit than a web browser in this case,
but there's probably some economic or install base incentive at work I am
unaware of). _Anyways_ onto some of the fun that is the GoTD (Go + Templ +
Datastar) stack.

In addition to the normal `sse.PatchElement()` the go SDK directly ships with
a convenience method for templ `sse.PatchTempl()` which makes writing SSEs
via HTML fragments entirely unnecesary. You can just pass in the component
you are wanting to patch directly into the event like so:

```go

func myhandler (w http.ResponseWriter, r *http.Request) {
    sse := datastar.NewSSE(w, r) _ = sse.PatchTempl(MyComponent())
}
```

The convenience functionality for working with signals also makes a lot
of the server-driven reactivity absolutely trivial such as this quick and
dirty example of reading the signals and patching them. Note: There's a bit
of an unspoken order of operations that you must read the signals _before_
opening an Server Sent Event, so some of the error handling can be a smidge
clunky around this specific aspect.

**NOTE** Datastar is **not** expecting that all the signals be defined in
`MySignals` while it will send all signals (unless filtered) the SDK is smart
enough to ensure that only the specific keys listed will be taken. In this
example, if there were actually `third` and `fourth` defined as signals in
the Templates, there will be no errors in running the following.

```go type
MySignals struct {
    First string `json:"first"` Second bool `json:"second"`
}

func myhandler (w http.ResponseWriter, r *http.Request) {
    signals := &MySignals{}
 _ := datastar.ReadSignals(r, signals)
    sse := datastar.NewSSE(w, r) // Just some operations signals.First =
    "New Value" signals.Second = !signals.Second
 output, _ := json.Marshal(signals)
    _ := sse.PatchSignals(output)
}
```

With these two building blocks it makes it trivial to drive any and all
behavior with the server. For example, in my first actual project (an LLM
wrapper app, plz don't judge) the entire transaction is served via signals to
indicate the current step of the process, from loading to actually assembling
the final output and updating the DOM with the final result of the transaction
(all in one single post request).

What impresses me is that I started this knowing basically _nothing_ and
stumbled into a reasonable streaming based approach to incrementally send
updates to the frontend from the server. With minimal additional code. It
also -- in many ways -- feels simpler than working with HTMX and the HTML
fragments (templ is lovely and does this fantastically with HTMX, but somehow
the "simpler" solution of HTMX feels a bit more clunky since each endpoint
_must_ return HTML and you need something else to support event streams
like websockets)

That being said, while I have had nothing but good experiences with datastar,
it's not without it's hangups. In particular, I ran into three minor issues
as I worked on my initial prototype:

1. SDKs documentation can be a little variable (the go SDK documentation
seemed a bit sparse of examples). This was mostly rectified by me just using
my LSP to dig, but I could see it being a bit of a hangup with adoption
into organizations.  2. There seems to be a minor disagreement with how two
of the 1.0 RCs work that have some cryptic (or non-existent) syntax errors
(particularly `data-attr-myattr` vs `data-attr:myattr`). It was nothing
major, but lead to some initial hangups and confusion. 3. The debugger
being blocked behind pro might hinder adoption slightly. I don't think it's
entirely necessary, but -- at the very least -- the build in developer tools
did not seem to be as consistent with catching the signals.

This doesn't really detract majorly from my enjoyment of using it. It's an
incredible tool that really makes building webapps a joy while making much
of the complex reactivity relatively trivial to implement. It also makes me
excited to see what else I can build with it.

For now, it's probably going to be the tool that I lean on (alongside go and
templ) when building internal tooling. I think it's still a bit too early
to go _all in_ on it for critical workloads whereas something a little more
mature like HTMX might be a better fit, _for now_.
