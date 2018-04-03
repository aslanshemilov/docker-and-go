# golang

https://www.golang-book.com/public/pdf/gobook.0.pdf

https://ewanvalentine.io/microservices-in-golang-part-1/

https://grpc.io/docs/quickstart/go.html

https://bestonlinecoursescoupon.com/what-is-go-lang/

Set up Go Workspace: 

`$ brew update`

`$ brew install golang`

`$ vi .bashrc`

# This is actually your .bashrc file

`export GOPATH=$HOME/go-workspace` # don't forget to change your path correctly!

`export GOROOT=/usr/local/opt/go/libexec`

`export PATH=$PATH:$GOPATH/bin`

`export PATH=$PATH:$GOROOT/bin`

`$ mkdir -p $GOPATH $GOPATH/src $GOPATH/pkg $GOPATH/bin`

`$GOPATH/src` : Where your Go projects / programs are located

`$GOPATH/pkg` : contains every package objects

`$GOPATH/bin` : The compiled binaries home
