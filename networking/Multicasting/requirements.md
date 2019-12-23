# MDNS, MDNS-SD Requirements

We're going to implemnent a MDNS Service Discovery using Golang, from scratch following the specifications of [RFC6762 - Multicast DNS](https://tools.ietf.org/html/rfc6762) and [RFC6763 - DNS-Based Service Discovery](https://tools.ietf.org/html/rfc6763)

What problem does Multicast DNS solve:

>the ability to look up DNS resource record data types (including, but not limited to, host names) in the absence of a conventional managed DNS server is useful.

How does a MDNS-SD look like?

[design goals](https://tools.ietf.org/html/rfc6763#page-5)

```
(i) The ability to query for services of a certain type in a
certain logical domain, and receive in response a list of named
instances (network browsing or "Service Instance Enumeration").

(ii) Given a particular named instance, the ability to efficiently
resolve that instance name to the required information a client
needs to actually use the service, i.e., IP address and port
number, at the very least (Service Instance Resolution).

(iii) Instance names should be relatively persistent.  If a user
selects their default printer from a list of available choices
today, then tomorrow they should still be able to print on that
printer -- even if the IP address and/or port number where the
service resides have changed -- without the user (or their
software) having to repeat the step (i) (the initial network
browsing) a second time.
```

> This document describes how clients send DNS-like queries via IP multicast

Get familiar with DNS queries.

## Responsing to a MDNS Query

