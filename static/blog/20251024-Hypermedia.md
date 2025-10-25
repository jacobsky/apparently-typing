
### Web 1.0, HTMX and Datastar

I feel like I'm in a rather unique position going back to review web technologies this late in the game. I spent the better portion of my career doing embedded device engineering and working on cloud based reliability and support, so trying my hand at full stack this late in the game (mostly for fun and profit) it has been startling to see that I seemingly went from Web 1.0 (front end is your HTML template language) skipped over all the SPA fad (frontend is a thick javascript client) to "Web 1.1" (frontend is HTML in your template language again).

My recent projects have shifted away from Rust (as much as I love that language) towards go which is a bit quicker and simpler to write (despite having some very specific warts) and alongside it `templ` which is probably the best templating library that I've ever used. For simplicity, I opted to try two of the leading "hypermedia" branded technologies (Htmx and Datastar) which each have their benefits.

#### Why not [Insert SPA] Framework?
Aside from being more interested in the backend server/database/infrastructure, I am naturally skeptical of SPA frameworks. While I will entirely admit that they are engineering marvels, the primary problem mostly boils down to the following three points:

1. Complexity (most have huge build chains and tools that are required, along with a need to synchonize client and server state)
2. Eschews native controls for javascript controls
3. Inflexible, due to RPC-like architecture.

The above effectively make SPA frameworks eschew the major benefits of the web with very little gain. This is especially apparent when reading about "hydration" which seems to be a problem of synchronization, but while SPAs are thick clients, they don't seem to be implementing a lot of the important thick client features (see online games and netcode techniques like rollback netcode for fighting games or long term planning like in monster hunter). Considering they put such a heavy emphasis on duplicating the state between the client and server state, it's odd that there's not a bigger emphasis on other techniques to increase the negotiation instead of the nebulous "hydration" techniques.

Regardless, most of the web does not need the added complexity and I find that a hypermedia approach is much more relevant to the problems we are trying to solve on the web (a see of create-read-update-delete operations, with small pockets of interactivity)

#### HTMX

The successor to intercooler.js (and the one that reignited the hypermedia revolution) the emphasis is on returning the web to 2004 with a small selection of additional HTML extensions (hence the name HTMX).

Essentially, it's the kind of thing that would allow you to make a button that expands the code from the server.

```
<button id="expand_content" hx-get="/expand" hx-swap="outerHTML">My button which expands to something else</button>
```

The declarative nature of the syntax allows for simple directives to be embedded that enable rich interactivity for the web at minimal cost. To be honest, it's just about the perfect extension to HTML for _most_ applications, though you may need another library like alpine.js or hyperscript to get more clientside interactivity (such as animations and such)

I think that the biggest point that I adore about HTMX is that the design philophy is "Build your webapp without it first, then sprinkle it in where you need additional client side reactivity.

#### Datastar

While HTMX is rooted in a desire for the simplicity of the past, Datastar is firmly rooted in the modern web as a fullstack framework that you can leverage for a modern feeling web app with the simplicity of HTMX.

Similar syntax to the above, the main difference is that the endpoint itself is more directly responsible for _how_ it addresses this. An `ElementPatch` call would allow for the same functionality as in HTMX example above

```html
<button id="expand_content" data-on:click="@get('/expand')"></button>
```
But it's also possible to do more with simple json signals themselves that allows for trivial databinding (useful for things like verification)

```html
<input type="text" data-signals:valid="true" data-on:change="@get('/verify')"/>
<label data-bind="$valid">Invalid!</label>
```

This can be a super efficient way to set up endpoints and create a reactive frontend with minimal effort or hassle.

In particular the Server Sent Events implementation of datastar makes it an incredibly flexible tool to update small portions of the dom dynamically without having to duplicate the state on the client. 

While Datastar does allow -- to some extent -- while you can follow the HTMX approach of "build it without reactivity and then sprinkle it in" that isn't really leveraging the true power of Datastar. Actually adding in more of the functionality requires a bit more forethought to make the most of the server sent events directly in your application. Building with Datastar in mind if much more "required" when starting out if you want to unlock it's full potential. As contrasted with HTMX which benefits considerably more from the "sprinkling in reactivity" approach.

Both technologies are fantastic web technologies that manage to be simple, elegant, _and_ powerful.

As with everything they aren't a silver bullet, but they are super fun to use and I'm looking forward to making more things with them.

Maybe later I'll be able to talk about more of the examples.

