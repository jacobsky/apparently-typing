It's been a little less than a month since my last brief experiment with
tailwindcss (the component based solution to CSS). At the end of that,
while I was impressed by the features of tailwind, I ended up realizing that
the work that I wanted to concentrate on at the time both _didn't need it_
and tailwindcss was _getting in the way_ of progress on the actual meat of
my applications. As a result, I ended up going with the (nearly) classless
CSS framework [picocss](https://picocss.com).

Well, after reaching the limit with pico in my current project, I am back
to tailwind and gave all my apps a refresh with a unified style for all the
projects under the `apparently-typ.ing` domain. This time around, as I had a
proper app to work with and style (rather than attempting to style something
while you build it). I feel that most of my assumptions regarding tailwind
were largely validated by this second attempt.

#### Great for beginners

While it may seem odd, the inline markup is actually great for beginners
(provided that you are willing to read the documentation rather
thoroughly). There are two major reason that I found.

1. Style on the element directly makes it trivial to understand _why_
something is happening (e.g. if you see something sized incorrectly,
open the inspector, check the size of each element until you find wrong
one, then check the tailwind elements) 2. Component libraries, such as
[daisyui](https://daisyui.com) which I am using, have heaps of prebuilt
useful components for you to use to help accelerate the development to
(almost) semantic HTML levels (once you understand what they require and
how they work with tailwind).

The one big caveat is that you really need to understand tailwind to make
use of it, but -- personally -- I don't find that to be a compelling argument
against it's use (you need to have a good understanding of CSS to use it in
a traditional way anyways). But the biggest win is how self explanatory each
piece of it is.

For example, what does this component do?

```
<section class="mx-auto my-4">
    <article class="prose lg:prose-xl">
        <h3> header </h3> <div class="divider"></div> <h4> sub header</h4>
        This is some content
    </article> <div class="flex">
        <ul class="menu menu-horizontal rounded-box justify-center">
            <li>
                <a href="#">Back</a>
            </li> <li>
                <a href="#">Top</a>
            </li> <li>
                <a href="#">Next</a>
            </li>
        </ul>
    </div>
</section>
```

Well, it is what renders this blog post. But going top to bottom, you can
see it does the following:

1. Margin settings of auto for x direction and 4 for y direction 2. It is
using the prose plugin for rendering generic HTML with special settings
for LG and up screens 3. It has a menu at the bottom in a flex container,
styling links using daisyui's `menu` component with generic links that get
rendered via the parent menu styling.

#### Great for advanced usage

Given that you can make so much, tailwind also ends up being absolutely
fantastic for anyone that needs custom components or working on complex
features (it was my latest project that features a chat window that pushed
me to look into tailwind).

#### Great for hypermedia approaches

Given the hypermedia focus on simplicity and Locality of Behavior design
principle, tailwind tends to match well. While some argument can be made
that more complex custom pieces are going to be "heavy" over the wire,
I think that's going to have to be evaluated on a case by case basis and
will have to be properly measured to tell. As a whole, it allowed me to
better leverage the templ (my template engine of choice for go) to iterate
on and fix any issues that arose. (Template engines also reduce a lot of
the concern about keeping things rather DRY)

#### The merits of classes CSS frameworks

While not tailwind directly, this entire experience also gave me a really
good perspective on when a classless CSS framework should be used with a
project. In short, working with pico, which gave me a lot of power for very
little investment into learning it, also left my frontend code in a place
where removing my pico dependency and then adding in the tailwind markup
required minimal refactoring.

At this point, it mostly reinforces my process with regards to web
development that the best way to work is to start with a classless
(or low class) CSS framework like [picocss](https://picocss.com),
[water.css](https://watercss.kognise.dev/), [beercss](https://www.beercss.com),
or any of the other classeless CSS options out there. The primary reason is
that style (usually) is not the selling point of your app, so it is always
a later concern. Using classless CSS focused on semantic HTML allows you to
get the business value first, then do an iteration in another framework to
spruce it up later.

All in all, very happy with tailwind, will be sticking with it in the future
(even if it isn't how I start an app).
