Sheathe (WIP)
=============

Sheathe is a URL shortening service built with Go, focused on infrastructure freedom.


# What makes Sheathe different?

Sheathe leverages the principles of Hexagonal Architecture to provide a technology-independent application layer that communicates with two ports:

1. A `LinkRepository` to persist link translations (programmers can write particular implementation to store links on MySQL, Redis, the Filesystem, etc.
2. An `EventNotifier` to which domain events (`LinkCreated`, `LinkRetrieved`) are dispatched. It may be connected to an enterprise log, a custom analytics service, a visual dashboard, etc.)

With this simple [DIP](http://en.wikipedia.org/wiki/Dependency_inversion_principle) implementation, we allow any user to build


# How to use it

TODO: Write how the repository can be configured or added new adapters


# How to deploy it

TODO: Write how the repository can be deployed via Docker


