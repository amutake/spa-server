spa-server
==========

Simple static file server for single-page application

```sh
$ tree
.
|-- index.html
`-- assets
    |-- js
    |   `-- main.js
    `-- css
        `-- main.css
$ sap-server 5050
...
$ curl http://localhost:5050/
=> ./index.html
$ curl http://localhost:5050/assets/js/main.js
=> ./assets/js/main.js
$ curl http://localhost:5050/index.html
=> ./index.html
$ curl http://localhost:5050/page1
=> ./index.html
$ curl http://localhost:5050/page2/123
=> ./index.html
```
