# Sheath [![Build Status](https://semaphoreci.com/api/v1/projects/1d052568-9ffd-4166-ab9b-1ad76917a1c3/425484/badge.svg)](https://semaphoreci.com/medbrain/sheath)      

Sheath is a URL shortening service built with Go, focused on high adaptability.


## What makes Sheath different?

Sheath leverages the principles of Hexagonal Architecture to provide a technology-independent application layer that relies on the following interfaces:

1. A `LinkRepository` to persist link translations. Developers can write particular implementation to store links on MySQL, Redis, the Filesystem, etc.
2. A `Notifier` to which domain events (`LinkCreated`, `LinkRetrieved`) are dispatched. It may be connected to an enterprise log, a custom analytics service, a visual dashboard, etc.
3. A `Validator` that forbids the creation of certain kinds of links. It may be used to prevent a redirect to certain protocols, the creation of spammy links (via a blacklist), etc.

With this simple [DIP](http://en.wikipedia.org/wiki/Dependency_inversion_principle)-based implementation, Sheath allows for a great level of adaptability. It provides the core functionality of a URL shortener, and the framework to extend it or fine-tune it at will.



## How to use it (out of the box)

The default configuration of Sheath (no modification required) provides a functional service that will listen on localhost:1827, and act on the following endpoints:

* A `POST /` request with param `url=http://www.any-url.com` in the post form will shorten the provided url and return the stub identifying it. Stubs are, by default, a 10-character-long alphanumeric string
* A `GET /:stub` request will return a redirect to the URL corresponding to the provided stub

### Default Link repository
The default configuration relies on a Redis database running on `host:port`, where the host is defined by the environment variable `SHEATH_REDIS_HOST` (`localhost` if such environment variable is not defined), and the port is defined by `SHEATH_REDIS_PORT` (`6379` if it is not defined).

### Default Notifier
The default notifier discards all the domain events it receives. The decision not to implement a smarter notifier out of the box was taken because the analytical part is the most subjective use that can be made of a URL shortening service. Thus, we leave advanced users the possibility of implementing their own notifier logic. 

### Default Validators
By default, the system only accepts `http://` and `https://` urls. This is a reasonable default, since Sheathe's main purpose is to handle browser requests.



## How to use it (adapted to a particular use)

To run Sheath with a different configuration or extended functionality, the following path may be followed:

* Fork the repository, write additional adapters in place, and change the final configuration. We encourage you not to modify the core logic or structure of the repository, but to work within the provided framework.
* Use Sheath as an imported library. Create a Sheath application instance and expose ports using both built-in adapters and additional adapters you implement within your repository.



## How to collaborate

All kinds of collaboration are welcomed:
 
* If you detect a bug, you may write and pull-request a test reproducing the bug.
* If you write an adapter (a repository, notifier or validator) you think could be of use to the community, please share it!
* If you have any doubt or suggestion, do open an issue!
