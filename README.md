## Server API

### NOTE

These APIs dont use gRPC


API will expose end-points to retrieve port information from a static data JSON file.

```
host:9000/port/{key}
```
Retrieves an individual port, using a specific id (key)

```
host:9000/listports
```
Retrieves a list of all port ids (key)

```
host:9000/listports/all
```
Retrieves a list of all ports details

```
host:9000/help
```
Retrieves a help response, showing end-points

