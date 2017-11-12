# Go-Revel-Tutorial

Revel is A web framework for the [Go language](http://www.golang.org/).

This application is a mini project in which I realized a method POST which create a new Model and a method GET which can show a Model with its Id and to simplify the configuration I stored the models in the cache.

After you clone these project
### You can  Start the web server:

   revel run Go-Revel-Tutorial

### Go to http://localhost:9000/ and you'll explore the Application


## Code Layout

The directory structure of this application:

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        views/        Templates directory
        models/       Contains the model of The application

    messages/         Message files

    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files

    tests/            Test suites




## Ressources

* The [Getting Started with Revel](http://revel.github.io/tutorial/gettingstarted.html).
* The [Revel guides](http://revel.github.io/manual/index.html).
* The [Revel sample apps](http://revel.github.io/examples/index.html).
* The [API documentation](https://godoc.org/github.com/revel/revel).

