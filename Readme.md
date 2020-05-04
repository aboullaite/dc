`dc` is a command line tool to continuously build and deploy docker-compose services in development mode. It take care of the repeated tasks (build containers, stop container, start containers) letting you to focus on what matters the most: coding.

You need to have `docker-compose` already installed and configured in your machine. `dc` simply makes system calls calls to your `docker-compose` executable so you can focus on fixing bugs or adding features. 

It doesn't require any config file. It simply use `docker-compose.yaml` file to get insight about your services, by extracting the context path. for the moment, it is necessary to to specify `context` value for each service that you wanna track.

![](https://media.giphy.com/media/M9HnSBnQtqYzwzvhlF/giphy.gif)
