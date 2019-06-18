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


- ### How would you route the object to another service\Module? What approach\tools you will be using?

 The routing works using the popular gorilla/mux library for multiplexing different request to corresponding handlers. The Checkin object is dispatched to the CheckinHandler to be consumed by the appropriate service.
  
 - ### What are your concerns about performance and memory usage?
 
 The major performance bottlenecks are usually high object allocations, high GC latencies, blocking calls. Identifying and fixing these considerably increase performance. 
 
 - ### What framework(s) you recommend to use?
 Golang has built in tools for the modern web applications and the capabilities to plug and play external packages makes it extremely powerful language. Much of the work can be done using the standard library, and the rest can utilise the external libraries. However for quicker developement we can alway utilise the pre existing frameworks. Gin and Beego are the popular web frameworks, however for microservices the GoKit is highly recommended framework.
 
 - ### As a team leader, how much estimation you will ask to deliver a fully-fledged component that is ready for deployment to production?
 
