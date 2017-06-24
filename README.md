# Hitsmon

A service that pick hits from Redis and save them into a database. Principle: the hits are stored into Redis by some external 
process. The Hitsmon collector will save these into a database and clean the Redis keys at a fixed interval.

## Supported databases

- Postgresql
- Influxdb
- Rethinkdb

## Configuration

   ```javascript
   {
	"type": "influxdb",
	"addr":"http://localhost:8086",
	"user":"admin",
	"password":"admin",
	"db": "hits",
	"table": "hits",
	"domain": "localhost",
	"frequency": 1
   }
   ```

Frequency is the interval between the runs to save the data: 1 means the process will run every second.

## Input sources

Django middleware to save hits into Redis: [django-hitsmon](https://github.com/synw/django-hitsmon)
