So, most of my experience is in the weeds of programming. Not nearly as much is in styling and CSS, but have been learning.

Interestingly enough, CSS is still a similar beast to what it has been since the dawn of styling on the internet, where general layouts are relatively simple to implement using semantic HTML, but specialized widgets can be a nightmare of complexity.

When I first studied this in the early 2010s, it was convention to use classes to layout specific components, possibly extending to use multiple CSS and SASS files to lay everything out, but utility frameworks have been on the rise.

Now there is also a competing idea of Utility-class frameworks which -- in my opinion -- is probably the best way to do complex CSS. While the "external CSS file" is often very clean and pretty i.e.

```
<div class="mystyle" id="mydiv"></div>
```


The challenge is that often trying to mix styles is difficult with conflicting meanings. `mystyle` could have any number of CSS and could apply to any number of different types of styles with any number of additional effects.

This makes it harder to know what exactly is gong on with any given style.

Utility-class frameworks like tailwindcss solve this by heavily leaning into composition so your classes would look like something similar to:

```
<div class="container flex  bg-gray-50 mx-8 px-8"></div>
```

As can be seen in the above, it's inline with the div, it defines that it is a flex container with the color gray-50 with margins and padding. It does get uglier and much more verbose as you go which clutters the markup. This is a fact, and it is a major point of contention among critics that has a lot of merit (as can be seen by the excessive number of component libraries out there like tailwindplus, daisyui, shadcn, basecoat, templui, etc). The verbosity also can lead to quite the learning curve (e.g. the Rust of CSS if you will) as you cannot simply look at the class and get a general understanding of what it does, you need to _know_ it or it is utterly incomprehensible.


While valid, I honestly can't think of a better idea for the complex web applications. There is an incredible amount that can go into styling web applications with very specific requirements and a class based approach like the [csszengarden](https://csszengarden.com/) is very difficult to scale cleanly _anyways_. There is the added challenge that updating or altering individual components can affect other components in separate areas of the site. As a result, utility classes manage to overcome this by moving all the complexity out into the markup instead of hiding it away.

Actually, speaking of the csszengarden, there has also been a minimalist movement gaining traction with the apearance of classless CSS libraries aimed at adding styling based (almost) exclusively on semantic HTML. This site uses one such framework [picocss](https://picocss.com/) exclusively because it is _dead simple_. I imagine that most webapps should aggressively attempt to use a classless CSS framework (either found online or made on their own) as long as possible before adding tailwind. Tailwind seems to still make excessive use of `<div>` when modern HTML has so many better ways to define the type of content in the tags like `<header>`, `<main>`, `<section>`, etc that minimalism should be embraced before complexity arrives.


That being said, I think there's some application ideas that I want to try where tailwind would be quite useful. Will see if I can try that later

