Sheathe (WIP)
=============

Sheathe is a URL shortening service built with Go.


# What makes Sheathe different?

Sheathe leverages the principles of Hexagonal Architecture to provide a technology-independent application layer that communicates with three ports:

1. A `LinkRepository` to persist link translations (programmers can write particular implementation to store links on MySQL, Redis, the Filesystem, etc.
2. A `CommandMarshaler` to call the application layer from any protocol (a RESTful HTTP service, a CLI, AMQP, etc.)
3. An `EventNotifier` to which domain events (`LinkCreated`, `LinkRetrieved`) are dispatched. It may be connected to an enterprise log, a custom analytics service, a visual dashboard, etc.)

Thus, Sheathe's primary focus is on the developer's flexibility to implement URL shortening with the technologies they want.


# How to use it

...

