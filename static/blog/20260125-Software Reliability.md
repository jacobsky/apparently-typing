As I dive deeper and deeper into the AI bubble and hype, one thing that I keep coming back to is "reliability".
It may be something that the CEO of Cursor mentioned ["We build a browser with GPT-5.2 in Cursor"](https://x.com/mntruell/status/2011562190286045552) which, based on the claims, is pretty amazing.

Kind of bonkers that AI is _just so good_ isn't it? Let's take a [looksee at the source they provided](https://github.com/wilsonzlin/fastrender).

Well, looking into it, it's not really impressing me. Aside from the dependencies appearing to be a _mess_ and stealing servo's homework (what with using their HTML parser) there's a lot of things that just don't make a lot of sense when you compare it to the existing options.

That and -- when I went to build it -- it wasn't working half the time ([I also couldn't get my mouse clicks to register on google homepage like this issue](https://github.com/wilsonzlin/fastrender/issues/105)). There's also an interesting report in the actions showing [1484 runs successful runs](https://github.com/wilsonzlin/fastrender/actions?query=is%3Asuccess) and [58,153 failures](https://github.com/wilsonzlin/fastrender/actions?query=is%3Afailure). 

Wait, what? Let me double check.

O_O;;

That is indeed correct. Holy moly that's just not something that I would judge to be acceptable based on their announcement.

In their defense, [this was an experiment in scaling agents](https://cursor.com/blog/scaling-agents), but I'm not very impressed and it looks like 3 million lines of AI slop. They have successfully proven that AI can write a lot of half baked code, and -- well -- I knew that already.

One reason that I rarely use coding agents for more than auto complete tasks is that they are generally pretty bad at anything _meaningful_ and things are _rarely_ a one shot success. Most of the time it requires babysitting and getting the code to the point that I could write it (design decided, interfaces and scaffolding prepped, psuedo code written, etc) before the agent has a shot at saving me time. With articles continuing to tell me that I am being obsolete, I do feel the need to address the one elephant in the room that no one seems to be considering in the slightly.

Reliability.

"What does reliability mean?" to all these AI enthusiasts. I don't know, but I know what it means to me.

My earliest notion of reliability was in my highschool band. Notably, our conductor/teacher said "Music isn't like other classes. 90% may get you an A, but it sounds awful" and there's a good reason for that. For simplicity, let's use Beethoven's 5th symphony (sometimes refered to as the much cooler name "Fate Symphony") which has a runtime of about 30-40 minutes. If you assume a 99% reliability rate that each instrument will play _all the proper notes_ each player would be allowed to play incorrect notes for 18 seconds (assuming a conductor wants the speedier 30 minute runtime). When you extrapolate that out to _every instrument_ (21 different scores) that's a grant total of 378 seconds worth of wrong notes spread throughout the 30 minutes. You either end up with a single 18 seconds of the most shrill and vile sounds imaginable or you have as _pleathora of small things_ that sound wrong and off and make the whole performance rather...  disappointing.

Now, back to tech, this is not too dissimilar to most reliability metrics. The first is usually uptime. For example: most cloud services will offer a minimum service level agreement of 99.9% which translates to approximately 8.76 hours of downtime per year. This scales up with every additional metric, such as 99.95 being half that at 4.38 hours, 99.99 being just shy of 53 minutes per year of downtime, and 99.999% downtime being 5.26 _minutes per year_ of downtime.

This is actually pretty important because the difficulty in orchestrating this kind of reliability becomes orders of magnitude more complex and expensive.

Even students and hobbyists can offer a 90% uptime guarantee. In a year at 90% you have 36.5 days of downtime allowed. You have an _entire month_ where you can be down for any reason be it maintenance, cloudflare/AWS/Azure(but why?)/GCP(lol) outage, bad push to production, etc. But as you push upwards, it gets increasingly harder. Getting to 99% reliability (still that mess I illustrated earlier) means dropping the entire downtime by a factor of 10 with each extra amount of reliability added.

Reliability comes in other forms as well, such as applications loading, applications working the same way every time, applications not throwing errors constantly, etc. Some of the _worst experiences_ that I have had with software that I've ever used has been bad only because it was inconsistent. Dynamics applications taking excessive time to load _anything_ where saving is a suggestion; losing work progress because of an application grinding to a halt for unexplicable reasons; web apps that may or may not require me to refresh my page N times to get it to load; etc. What connects a lot of these is that they fall in the 99% of the app is reliable, but this 1% is _intensely aggravating_. And this is software that _mostly works_.

The reason for this long diatribe is that I'm rather disappointed at the hype cycle around AI. What was created by Cursor _should be embarrassing_. It does not _even come close_ to working _at all_. This project they are proud enough to announce and say "It *kind of* works!" is a stretch. _Kind of_ working in a software sense means that it is fully functional and is usable with missing use cases. It does not mean "Cannot even click the search bar on the most used search engine on the planet".

There's also probably another bit that needs to be said about the difficulty in achieving reliability and how AI often doesn't really help with that, but I'll leave that for a future post.

