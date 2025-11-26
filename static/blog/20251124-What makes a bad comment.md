Capturing a conversation between myself and one of my friends about over commenting in code bases.

One of my key arguments for good comments is intention. It's most important that each comment in code be intentionally placed so that they don't become additional noise.

Unfortunately, there's no right answer in how to write comments such that they are not overcommenting (as the rules may vary by team and need), but rather, I think it can be nice to have some examples of what I consider "useless comments"

### Useless Comment #3 - Committing commented out code

Outside of specific build by build debugging (though please, just use a debugger or improve your logging), comments that leave a bunch of additional loose code should just generally not be something that is strewn about the codebase. Few things are more distracting than having to skip through vaste swathes of a function just to find what you are looking for.

Note, examples are much different than commenting out chunks of code (such as tests), examples are working code that can be used to demonstrate concepts and are typically stored separately, such as in a module or package comment block. Some languages will even allow this to be compiled to confirm that it is working (such as the rust compiler) but otherwise it can just be stored in an examples directory with the project.

### Useless Comment #2 - Rephrasing a function or package

This should be relatively obvious, but there's no reason to comment `Reverses the datatype` for a reverse function _unless_ there is something novel about it that warrants it.

For example the difference between the following

```
// Sorts the slice
func sort(a []Type) {
    ...
}

// Implements an in place quicksort algorithm with no additional memory allocations
func sort(a []Type) {
    ...
}
```

### Useless Comment #3 - Explaining the syntax

The less obvious, but still wasteful comments are 
Consider the following golang snippet

```golang
// Iterates through the slice
for i := range len(connections) {
    // if the element is currently closed, prune
    if connections[i].IsDisconnected() {
        // close
        close(connections[i])
    }
}
```

For inexperienced or non-technical readers, these comments may be welcome. The main issue is that, for _most_ code bases, people should be able to understand what happens from the syntax alone.

Much more useful would be explaining _why_ we are iterating through our connections.

Unless you are aiming at educating users on the syntax of a language, this is typically unnecessary bloat to the code base (and let's be honest, the target audience for most comments are yourself)

Something like the following would be more useful in order to outline the intent of why the connections are being iterated.

```
// This is to ensure periodic cleanup of connections that have previously been closed
for i := range len(connections) {
    // Due to the way that connections 
    if connections[i].IsDisconnected() {
        // close and prune the associated listener 
        close(connections[i])
    }
}
```

Remember, comments code syntax self documents _what is happening_ but doesn't document the intent of the code. Documenting the intent is the difference between immediately understanding why a conclusion has been reached and having to reinvent it only to end up with the same solution.
