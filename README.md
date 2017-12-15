# kubevalonline

`kubevalonline` is an experimental project to use garethr's [kubeval](https://github.com/garethr/kubeval) to verify kubernetes configuration files as a web app similar to [JSONLint](jsonlint.com).

## Usage

```
$ make docker
$ docker run -p 5000 donachys/kubevalonline:latest
```

use docker ps to see which port 5000 was mapped to and then open your browser to localhost:<portnum>



