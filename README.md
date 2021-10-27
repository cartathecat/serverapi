## Server API

### NOTE

These APIs dont use gRPC


API will expose end-points to retrieve port information from a static data JSON file.

portno defaults to 9000 if not set using in the environment variable API_PORT

```
host:{apiport}/port/{key}
```
Retrieves an individual port, using a specific id (key)

```
host:{apiport}/listports
```
Retrieves a list of all port ids (key)

```
host:{apiport}/listports/all
```
Retrieves a list of all ports details

```
host:{apiport}/help
```
Retrieves a help response, showing end-points

