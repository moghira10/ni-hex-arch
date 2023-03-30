## CheckIn REST API

Simple CheckIn API like Foursquare which takes in location data and writes into the datastore.

A sample request looks like

`curl -X POST \
  http://localhost:8080/addCheckIn \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
"userId": "a0mv239242",
"place": {
"placeId":"8203",
"name": "Bistro 65",
"lng": 31.112232,
"lat": 20.123221,
"category": "restaurant"
},
"checkinTimestamp": "2019-06-17T12:42:31Z"
}'`
