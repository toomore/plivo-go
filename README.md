plivo-go
=========

[![GoDoc](https://godoc.org/github.com/toomore/plivo-go?status.svg)](https://godoc.org/github.com/toomore/plivo-go)

Cmd
----

1. plivoSendMass - [大量發送簡訊](https://godoc.org/github.com/toomore/plivo-go/cmd/plivoSendMass)
2. plivoQuickSend - [快速單發簡訊](https://godoc.org/github.com/toomore/plivo-go/cmd/plivoQuickSend)

Docker
-------

Download image

    docker pull toomore/plivo-go

Send mass SMS with CSV.

    docker run -d -v <local_csv_path>:<container_csv_path> -e PLIVOID=<ID> -e PLIVOTOKEN=<TOKEN> -e PLIVOSRC=<SRC> toomore/plivo-go plivoSendMass -csv=<container_csv_path>

    549fc9920a6ec943a86b2f5afc569ef78d54a67e432e11888f0ddf5081158750

Read send logs

    docker logs 549fc9920a6ec943a86b2f5afc569ef78d54a67e432e11888f0ddf5081158750

License
--------

The MIT License (MIT)

Copyright © 2015 Toomore Chiang, http://toomore.net/ <toomore0929@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

http://toomore.mit-license.org/
