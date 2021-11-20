# Simple Debian Packaging

A basic example of debian packaging with dh-sysuser (to create a user).

```sh
# Use docker to build container (it contains all build dependencies)
# (for ease of use, it's a make command shortcut)
make buildcontainer

# make the debian
make deb

# check the contents of the debian
dpkg --contents dist/*.deb
dpkg -I dist/*.deb

# install the deb on the target machine
```
