# golang

https://ewanvalentine.io/microservices-in-golang-part-1/

https://grpc.io/docs/quickstart/go.html

https://bestonlinecoursescoupon.com/what-is-go-lang/

Golang is an open source programming language which is used to build simple, reliable and efficient more quickly. What is golang? It was formed by combining the performance and security benefits of C++ and includes the speed of Python.

**Packages**

Programs starting running in the package `main`. This program is using the packages with import paths `fmt` and `"math/rand"`. By convention, the package name is the same as the last element of the import path. For instance the `"math/rand"` package comprises files that begin with the statement `package rand`. The environment in which these programs are executed is deterministic, so each time you run the example program `rand.Intn` will return the same number. To see a different number, seed the number generator (`rand.Seed`).

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.


Features Of Go Programming Language
The main character of the Go language is: –

• Cloud Computing: – Datacenter and the clouds are not same. The cloud offers servers on demand. You can let auto-scaling turn the knobs on you. While scaling the instances some of the machine takes a long time to run and parse through every line of code. Go applications can be quickly launched and efficiently respond to servers so that the programming may go faster.

• Multicore Performance: – Go is a scalable language and includes straightforward model, so we don’t need to rewrite our software for additional cores. It doubles the number of cores on a regular basis.

• Microservices: – Use asynchronous input-output for make application to be interactive with some services without blocking web requests. Go helps in building applications as a collection of microservices.

• The Fat Client Renaissance: –  Go language don’t need to be tied with frameworks of yore. It can work with writing APIs and favors an excellent support for multiple clients. It can build the API and then make it flexible for the customers.
• Non-Synchronous: –  Go language runs smoothly on all operating systems and doesn’t block others. Its code is non-synchronous style so that there is no need for callbacks.

• Concurrency: – Go language provides a concurrency primitive comparable to the actors model. These players models help in naming endpoints of the channels.

• Static Binaries: – Go language has no static boundaries. It can be quickly compiled and launched immediately.

What is a microservice? We use to separate codebases. 

Complexity - Splitting features into microservices allows you to split code into smaller chunks. It harks back to the old unix adage of 'doing one thing well'. There's a tendency with monoliths to allow domains to become tightly coupled with one another, and concerns to become blurred. This leads to riskier, more complex updates, potentially more bugs and more difficult integrations.

Scale - In a monolith, certain areas of code may be used more frequently than others. With a monolith, you can only scale the entire codebase. So if your auth service is hit constantly, you need to scale the entire codebase to cope with the load for just your auth service.

First define your service, this should contain the methods that you wish to expose to other services. Then you define your message types, these are effectively your data structure. Protobuf is statically typed, and you can define custom types, as we have done with `Container`. Messages are themselves just custom types.

There are two libraries at work here, messages are handled by protobuf, and the service we defined is handled by a gRPC protobuf plugin, which compiles code to interact with these types, i.e the `service` part of our proto file.


