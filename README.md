# envoyctl
cli app to control the Envoy Project app

## Resources
* [envoyproject.com](http://envoyproject.com)
* [API reference](https://envoyproject.com/apioverview)

## Requirements


To build this application, you need Go version 1.7 or higher


## Build and run


```
go install .
envoyctl
```

## Configuration

The app tries to read the configuration file ~/.envoyctl.yml . To see an example configuration, you can run ```envoyctl exampleconfig``` . 

To save it, you can run ```envoyctl exampleconfig > ~/.envoyctl.yml```