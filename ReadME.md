# Our design rules
## Error handling
### SSL Keys.
1. if either the key or the pem weren't found, we will terminate the server as we cannot establish an https server.

### Main page
1. if the index.html was not found, we will terminate the session.
2. if the index.html was found but it had any execution errors in terms of templating we will terminate the server session 

**This is as the index.html file is the most important file. without it, there is no meaning in running the server, so we will gracefully shut it down.**
### Error pages
1. if there is an issue with a client's request, it will be a network status 400 **Bad Request**.
2. if the user is using a method that's not allowed to the endpoint, we give it network status 405 **Method Not Allowed**.
3. if there is an issue with the server parsing or executing the error pages it will throw an 500 **internal server error** network status.
4. if the client access a page that's not found, we will throw an 404 **Page Not Found** error page.


