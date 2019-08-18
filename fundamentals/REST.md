# Rest API

### RESTful
RESTful web service -- is based on representational state transfer
	 
Generally preferred to the more robust Simple Object Access Protocol (SOAP) technology because REST leverages less bandwidth, making it more suitable for internet usage.

They use `GET` to retrieve a resource; 
`PUT` to change the state of or update a resource, which can be an object, file or block; 
`POST` to create that resource ;
`DELETE` to remove it..

POST = Create
GET = Read
PUT = Update
DELETE = Delete
... aka CRUD

In a RESTful Web service, requests made to a resource's URI will elicit a response with a payload formatted in HTML, XML, JSON, or some other format

**Stateless** (One request to a service won't talk directly to another request)

**Rest endpoints**
Handle: `POST`, `GET`, `PUT`, & `DELETE` HTTP requests

JSON - Marshalling, Unmarshalling

net/http router or gorilla/mux, third party allows us to easily retrieve path and query parameters.
- Add path variables `router.HandleFun("/article/{id}", getArticleById)`