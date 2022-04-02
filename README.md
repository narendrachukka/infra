# appconfig

`appconfig` is an opinionated configuration loading library for long running applications
(sometimes called services, daemons, etc).

Configuration is loaded in order from the following sources:

1. default values defined in a struct literal
2. config file (YAML)
3. environment variables
4. overrides (often from a command line flag)


`appconfig` supports:

* generating markdown documentation for a configuration struct
* nested structs
* use reasonable conventions to remove the need for many struct field tags
