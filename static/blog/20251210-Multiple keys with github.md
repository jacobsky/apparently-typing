This is more of a quick tutorial that I am writing for myself so that I do
not forget how to use multiple SSH keys with github easily.

When working with multiple github accounts, it is usually preferable to have
multiple SSH keys over using password authentication (also often disabled
for security reasons).

This can be a problem, especially if you are working with organizations that
require separate accounts for security reasons (such as if a freelancer).

As always, the first step is to create your keys, in this case, I'm going
to call them personal and work for clarity, but the names can be anything.

```sh
ssh-keygen -t ed25519 -C "username@personal.com" -f
~/.ssh/id_ed25519_personal ssh-keygen -t ed25519 -C "username@work.com"
-f ~/.ssh/id_ed25519_work
```

Once the keys have been generated, you can add them as identities with the
`ssh-add [file]` command which does the heavy lifting for us.

If you get an error mentioning that it cannot connect to the agent, you can
force it to start with: ``` eval "$(ssh-agent -s)" ```

Then inside `~/.ssh/config` add in the following entries

```sh Host github.com
    HostName github.com User git IdentityFile ~/.ssh/id_ed25519_personal

Host github-work
    HostName github.com User git IdentityFile ~/.ssh/id_ed25519_work
```

After that, it is just a matter of configuring the remote for your secondary
ID to use the following and then git will automatically use the correct
identity with it.

```sh
git clone git@github-work:workusername/myrepo.git git config set
remote.origin.url git@github-work:workusername/myrepo.git
```

Happy githubbing future me! I know you'll forget how to do it. That's why
I'm writing this down now!!!
