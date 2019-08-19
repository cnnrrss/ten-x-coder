# CAP Theorem

**Consistency** - all nodes see the same data even at the same time with concurrent updates.
**Availability** - a guarantee that every request receives a response about whether it was successful or failed.
**Partition tolerance** - the system continues to operate despite arbitrary message loss or failure of part of the system.

The CAP acronym corresponds to these 3 guarantees. This theorem has created the base for modern distributed computing approaches. Worlds most high volume traffic companies (e.g. Amazon, Google, Facebook) use this as basis for deciding their application architecture. Its important to understand that only two of these three conditions can be guaranteed to be met by a system.