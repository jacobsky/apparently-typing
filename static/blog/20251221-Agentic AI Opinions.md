Recently I have spent a fair bit of time with `opencode` over the past few days experimenting and trying to fit it into my neovim development workflow.

[`opencode`](opencode.ai) is an opensource coding agent aimed at running AI agents via the commandline and is compatible with all kinds of coding agents both hosted online and hosted locally.

While I have historically not had great experiences with AI agents, but I continue to hear about how important it is to have Agentic AI in your workflows so I figured I'd give it another shake. Maybe the previous ones that I used were just bad.

`opencode` fit the bill for a few reasons:

- installable via the Arch (btw) User Repository `yay -Sy opencode-bin`
- [`opencode.nvim`](https://github.com/NickvanDyke/opencode.nvim) fits into an nvim workflow.
- Terminal oriented
- Relatively fast
- Model and vendor agnostic (I should probably write my thoughts about why almost everyone should avoid vendor locking)
- Potential to support locally hosted models (massive+ because I hate being bound to cloud providers who may vanish at a moment's notice.)

Upon installing, I picked an older project that I wanted to work on and spent some time using it's features.

## It's not about the task, it's about the plan

One feature of `opencode` that differentiates it from my last experience with agentic AI is the difference between planning mode and building mode. Planning mode puts your LLM in read-only mode (can only index and read files in your repository) and orients everything around the LLM asking you questions about the goal to help refine a plan that it can tackle at a later state.

My _first_ attempt at using things I started directly in build mode and just gave it a simple task like "create unit tests for each module of the project, do not update any existing code, only create new files" and this went... exceptionally poorly. The LLM itself basically went around creating tests and altering changing files willy nilly to make things pass even going so far as to try to duplicate the work of code generation tools like `sqlc` that I have been using instead of an ORM.

(NOTE: I am aware that general long form tasks are the kind of task that _LLMs are the worst at_, but this is a test and I feel it's necessary to give them the harder problems instead of the easier problems.)

I _almost_ gave up on the project here, but read a little more into it and realized that the problem was probably that I didn't use planning mode prior. Working with planning mode largely resolved this. While the code and test quality is a _bit dubious at times_ (this could be model specific as I'm not paying for this), the big thing that changed is that it actually _did_ what I asked. Which was a pretty great improvement.

## Implementation

`opencode` is considerably better at implementing code than when I last worked with Agentic AI which primarily delivered broken code that I had to completely audit. Now at least my project would start up and run after giving the agent a task.

I will say that I was moderately impressed at how well it did build up scaffolding once I build up an example of what I wanted. I was also impressed when I was able to take my old project and just task the AI with converting the `log.Printf` calls to a proper `slog` call. While not all of the logic was readily apparent, I was able to understand it easily enough while code reviewing.

This did have the downside of, every time I give the agent a task, I need to do a code review before committing and moving forward. While `opencode` comes with a nice `/review` command, there is a bit of a problem with having an LLM review it's own work. It _might_ pick up on things, but it's still got the weakness of potentially explaining or judging things wrong so whenever I use this, I do feel that I'm stuck spending a lot of time just reading through code. It works, but it's not what I really love doing.

## An Enhanced Rubber Duck

The more I experimented, the more that I found that planning mode was more of a game changer. Working solo on things, the biggest thing that I struggle with is reviewing my code. It's not always possible to get a second pair of eyes on my code or ideas, so I end up spending a fair bit of time searching around on the internet to look for stuff and trying to digest best practices for go. One of my favorite things that I have been able to use `opencode` Being able to say something like "I'd like to write unit tests" I was actually it was actually a good way to get a first eye on the problem. It functions a lot like an enhanced rubberduck debugger. Most of the time, I'm basically muttering to myself or a cat to think through a problem anyways, so being able to yeet out some questions or thoughts into the void _and actually get an answer_ can be helpful.

It's also less burdensome than code since, there's no testing or maintenance required afterwards. In this way, I'm actually really convinced that AI reviewing should be a major part of my workflow. When an Agent reviews my current commit and gives me some wrong information like "`time.Ticker` does not always call `Ticker.Stop()` and is a major memory leak" forcing me to prove my assumptions by digging into the actual library code to find the comment "Since go 1.23 the garbage collector can now collect unreferenced timers even when they are not stopped" was a _major learning point_ that forced me to validate my understanding.

This is amazing and exactly what I need as someone that primarily works solo.

## It's not all perfect

While it was exceedingly better than my prior experiences (i.e. I am now moved from a clear "I hate AI" position to more of a "I don't love it, but it might be useful" position), I've still got some major concerns.

My first major issue is that I feel zero sense of ownership or satisfaction from the code produced by the agent. This means that I dread committing the work more than when I'm actually working on the project myself. No context, no understanding, no feeling of desire to improve it or change it. It's just code that exists and I feel the possible maintenance burden (i.e. I still very much dislike the agents going off and coding stuff.) It gives me a feeling of apathy towards improving the codebase or trying to find good abstractions to work with as the AI can just kind of write a bunch of code _anyways_. I feel like a passenger just attempting to point an unlearning student driver in the vaguely correct direction instead of working as a technical leader who is actively participating in the process.

I also learned nothing about golang while using the agent for actually programming. Even in the code reviews, while I remember reading through the diffs and checking the code compiled, I don't really remember anything specific that it generated. The closest being that it logged certain status codes 400 or higher as `slog.Error` instead of `slog.Info` which I only remember because I thought "Well, really we should only log them as an error when they are 500 or above" for metrics purposes. The closest thing to knowledge was that I was able to reinforce what I had learned, but not really learning new things.

Much of this concern for learning comes from first hand experience in a different domain, second language learning (specifically reading/writing). As I previously worked for an AI company that, not only allowed, but encouraged the heavy use of AI tools in our workflows, my skills in Japanese largely stagnated and atrophied during my tenure at that company. This is because learning to read takes time and effort and using AI is _fast_. The AI translation and writing tools do not have to be perfect, they have to be faster and less effort than doing it manually. This is not a high bar for someone that is not particularly skilled in the language. Given how fast companies want to move, this temptation grows with every single task. And the temptation never goes away.

I also find that AI reduces the resilience required to get over seeing a long paragraph of text and having to dive in to read it. We also see this with the tendency to trend towards AI summarization and automatic drafting, etc. We as humans need a level of resilience to look at a task to dive in.

This experience with language learning is very relevant to the conversation as programming is much more closely related to linguistics than it is related to math (though it does borrow heavily from it). Learning programming languages is very much akin to learning a new context-free language with all the syntax. Oddly enough, you can relate a lot of programming _to_ regular language constructs with things like the formal grammar, informal grammar (native extensions and standard library), dialects (user packages), and slang (user built shortcuts/functions/macros). The temptations and methods of learning are very _very_ similar with a lot of overlap.

## In Summary: Not great, not terrible

Consider me officially convinced to add it into a workflow, but unconvinced on the actual code generation. I can see myself using AI for really boring stuff like scaffolding (once I figure out a good agentic prompt to scaffold my projects in a way that _I_ like consistently); converting between logging solutions across a project (like a better grep/find replace); as a reviewer for my commits and work; and finally hyper specific function level implementations for specific flows that I find exceptionally tedious an uninteresting (i.e. JWT authorization, writing specific test cases, etc).

I plan to generally avoid it for doing more than that. I don't plan to use it for asking questions about the code base itself and really only plan to ask it when I would normally turn to google because the included documentation isn't sufficient.

Thank you `opencode` team for making an actually usable agent service.
