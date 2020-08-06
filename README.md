# Simple users CRUD REST API

Project started from the [template](http://github.com/nejcambrozic/go-dev-template).
The API does not use any rest framework like [Gin](https://github.com/gin-gonic/gin) but uses just a very ligthweight [http router](https://github.com/julienschmidt/httprouter) and `net/http` module from the standard library
Users are stored in a sqlite database.
 
# Get Started

Use make for everything, this assumes you have the code checkout and are in the repo root dir:
Makefile contains help and has all commands documented, to see all the commands simply run `make`

## Start the app with live reloading:

`make dev`: this start the server with live realoading and exposes it on port `3000`

You can check the health endpoint: `http://localhost:3000/health`, it should return `{"Status":"ok"}`

Any change in `.go` file will now trigger a rebuild and you should see this in your terminal:
```bash
api_1  | Running build command!
api_1  | Build ok.
api_1  | Hard stopping the current process..
api_1  | Restarting the given command.
api_1  | stderr: Server started
```

## Debugging in Goland

Is described in [a separate document](doc/debugging-in-ide.md)
 

# Licence
[GNU GPLv3](LICENSE)






 



 