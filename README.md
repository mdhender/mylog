MyLog is a simple implementation of a log.

The log supports adding and retrieving data.

Data can be retrieved by in the order added or
by unique `id`. The `id` is a sequence number
based on the order added.

This is not a production quality log. It does not
copy the payload, so it is the client's responsibility
to treat it as immutable data.