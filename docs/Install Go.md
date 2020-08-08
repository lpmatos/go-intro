# Go on Unix Systems

## Installing

One of the big themes that we'll find as we work with Go is that Go tools can be distributed as static binaries and Go itself works the same way.

Depending on the system that you're using, you need to find the corresponding package from the Go downloads page and unpackage its contents into the /usr/local directory on your machine.

For this course, we'll be using a CentOS 7 cloud server so we'll need the linux-amd64 version of the package.

```bash
cd /tmp
curl -O https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
tar -xzvf go1.11.2.linux-amd64.tar.gz
sudo mv go /usr/local/
cd ~
```

By unpacking this archive into the /usr/local/ directory we now have a /usr/local/go directory that will hold of the Go's internal parts.

## Adding Go Binaries to The $PATH

Before we can actually use the Go tools we'll need to adjust our $PATH to include the binaries distributed as part of Go. We need to at the /usr/local/go/bin directory to our path by modifying our ~/.barshrc file:

```bash
export PATH=$PATH:/usr/local/go/bin
```

Now we're able to run the go tool after reloading our .bashrc file:

```bash
exec $SHELL
$ go version
```

## Setting Up the $GOPATH

Go uses a very specific folder structure and environment variables to know where to find Go code. For the code that we'll be writing, we need to put them into a path that we store as the $GOPATH environment variable. We'll set this directory structure up at $HOME/go and set the $GOPATH within our .bashrc.

Let's create the directory structure first:

```bash
mkdir -p $HOME/go/{bin,src}
```

Each of these directories stores something specific to Go:

* $GOPATH/src - contains Go source files
* $GOPATH/bin - contains executables

Next, let's set the environment variable in our .bashrc:

```bash
# previous lines omitted
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

## Setting Up Other Tools

We don't need very many other tools for our development, but there are some that we'll need:

* Git

```bash
sudo yum install -y git
```
