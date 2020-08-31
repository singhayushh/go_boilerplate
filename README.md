# boilerplate_golangFront
Boilerplate for serving static pages using golang

### Directory Structure:
> As this is a boilerplate for just rendering static files, I have used just a single handler sub-directory
/
    - /handlers: package for handling the api calls. This is where the actual server side code lies.
    - /views: this directory is home to frontend files - index.html, css, js, img etc.
    - main.go: this is the entry point to the server.

### How to start the web server:
- Clone this repository to your local system and open the terminal in the cloned directory.
- Run the command: `go run main.go`
- Head to `127.0.0.1:8080` on your browser as directed after running the above command to see some busy gophers!