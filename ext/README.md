# External Projects (Git Submodules)


## Adding a submodule
```sh
# Select a directory and add the submodule
cd ext
git submodule add https://github.com/muzammilar/geomrpc.git
# Add with a tag
git submodule add --branch 0.0.1 https://github.com/muzammilar/geomrpc.git

# Check submodule status (see `git submodule -h` for all commands)
git submodule status

```

## Updating a submodule
```sh
# Update using the path
git submodule update --remote geomrpc

# Change branch or tag (and check status)
git submodule set-branch --branch master geomrpc
git submodule update --remote geomrpc
git submodule status

# Change branch to tag (and check status)
git submodule set-branch --branch 0.0.1 geomrpc
git submodule update --remote geomrpc
git submodule status
```

## Deteling a submodule

```sh

```
