[![Build Status](https://travis-ci.org/donachys/kubevalonline.svg?branch=master)](https://travis-ci.org/donachys/kubevalonline) [![Coverage Status](https://coveralls.io/repos/github/donachys/kubevalonline/badge.svg?branch=master)](https://coveralls.io/github/donachys/kubevalonline?branch=master)
# kubevalonline

`kubevalonline` is an experimental project to use garethr's [kubeval](https://github.com/garethr/kubeval) to verify kubernetes configuration files as a web app similar to [JSONLint](jsonlint.com).

## Usage

```
$ make docker
$ docker run -p 5000 donachys/kubevalonline:latest
```

use docker ps to see which port 5000 was mapped to and then open your browser to `localhost:mappedport`



