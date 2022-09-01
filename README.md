# Testing performance of http servers with different languages

The test includes sending jel library ( https://github.com/LumaRay/jel ) to a client and running a simple command to show a header.

The test tool is wrk: https://github.com/wg/wrk

The command: _~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:<port\>_

The testing has been performed on Windows 7x64 host with Ubuntu 22.04 x64 virtual machine with 8Gb RAM and 4 cores, Intel-VT enabled.

Host:
- CPU: Intel Core i7-4790 3.6 GHz 4 cores / 8 threads
- RAM: 32 Gb

Test results are given in ascending order, worst to best.

## test-python-flask

```
test@test-virtual-machine:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:5000
Running 30s test @ http://127.0.0.1:5000
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   119.16ms   58.65ms   1.81s    98.87%
    Req/Sec   283.62    132.85   490.00     60.49%
  33746 requests in 30.09s, 215.88MB read
  Socket errors: connect 0, read 84, write 0, timeout 23
Requests/sec:   1121.35
Transfer/sec:      7.17MB
```


## test-python-waitress

```
test@test-virtual-machine:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:8080
Running 30s test @ http://127.0.0.1:8080
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    59.64ms   32.10ms   1.09s    90.79%
    Req/Sec   549.04    249.02     1.18k    53.17%
  49234 requests in 30.07s, 313.13MB read
Requests/sec:   1637.05
Transfer/sec:     10.41MB
```


## test-python-fastapi

```
test@test-virtual-machine:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:44777
Running 30s test @ http://127.0.0.1:44777
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    76.49ms   21.16ms 149.69ms   65.05%
    Req/Sec     1.31k   227.90     2.01k    66.81%
  156571 requests in 30.05s, 0.97GB read
Requests/sec:   5211.12
Transfer/sec:     33.14MB
```


## test-python-blacksheep

```
test@test-virtual-machine:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:44777
Running 30s test @ http://127.0.0.1:44777
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    56.88ms   21.00ms 130.29ms   65.21%
    Req/Sec     1.77k   258.39     2.74k    69.12%
  210692 requests in 30.03s, 1.31GB read
Requests/sec:   7015.40
Transfer/sec:     44.51MB
```


## test-java-spring-boot

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:8080
Running 30s test @ http://127.0.0.1:8080
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    14.01ms    8.64ms 129.23ms   75.54%
    Req/Sec     7.11k     1.58k   12.80k    71.95%
  847129 requests in 30.07s, 5.25GB read
Requests/sec:  28172.08
Transfer/sec:    178.70MB
```


## test-rust-astra

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:8080
Running 30s test @ http://127.0.0.1:8080
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.91ms    0.89ms  54.12ms   94.73%
    Req/Sec    14.52k    13.47k   31.25k    68.17%
  867941 requests in 30.09s, 5.38GB read
Requests/sec:  28848.86
Transfer/sec:    183.10MB
```


## test-cpp-lithium 4 thread (but seems to use 1 core)

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:7774
Running 30s test @ http://127.0.0.1:7774
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    77.70us   59.66us  16.84ms   96.28%
    Req/Sec    42.90k     2.80k   47.54k    78.67%
  1280489 requests in 30.08s, 7.95GB read
Requests/sec:  42567.00
Transfer/sec:    270.71MB
```


## test-go-http

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:3333
Running 30s test @ http://127.0.0.1:3333
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.48ms   12.23ms 149.25ms   85.00%
    Req/Sec    11.26k     2.31k   20.31k    69.48%
  1345604 requests in 30.10s, 8.43GB read
Requests/sec:  44708.65
Transfer/sec:    286.95MB
```


## test-cpp-drogon 1 thread (uses 1 core)

```
test@test-virtual-machine:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:8080
Running 30s test @ http://127.0.0.1:8080
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    13.21ms   64.24ms   1.45s    98.78%
    Req/Sec    13.88k     1.82k   24.75k    90.13%
  1653428 requests in 30.09s, 10.34GB read
Requests/sec:  54950.97
Transfer/sec:    351.82MB
```


## test-cs-dotnet-core-webapp

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:5000
Running 30s test @ http://127.0.0.1:5000
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     6.03ms    5.14ms 114.32ms   90.08%
    Req/Sec    16.72k     2.74k   27.06k    71.33%
  1994978 requests in 30.07s, 12.41GB read
Requests/sec:  66351.23
Transfer/sec:    422.57MB
```


## test-go-fasthttp

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:8080
Running 30s test @ http://127.0.0.1:8080
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.04ms    6.02ms 101.22ms   90.54%
    Req/Sec    22.34k     4.66k   39.40k    69.60%
  2662898 requests in 30.07s, 16.80GB read
Requests/sec:  88570.12
Transfer/sec:    572.35MB
```


## test-rust-axum

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:3000
Running 30s test @ http://127.0.0.1:3000
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.65ms    3.32ms  67.89ms   85.83%
    Req/Sec    29.49k     4.47k   50.08k    70.90%
  3522348 requests in 30.10s, 21.82GB read
Requests/sec: 117013.61
Transfer/sec:    742.20MB
```


## test-rust-warp

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:3030
Running 30s test @ http://127.0.0.1:3030
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.29ms    3.13ms  64.53ms   87.94%
    Req/Sec    32.49k     4.54k   52.20k    70.73%
  3884225 requests in 30.10s, 24.06GB read
Requests/sec: 129039.95
Transfer/sec:    818.49MB
```


## test-rust-actix

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:8080
Running 30s test @ http://127.0.0.1:8080
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.07ms    4.78ms 205.72ms   85.74%
    Req/Sec    32.59k     6.28k   67.34k    73.28%
  3893787 requests in 30.10s, 24.17GB read
Requests/sec: 129365.80
Transfer/sec:    822.40MB
```


## test-rust-minihttp-json

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:8080
Running 30s test @ http://127.0.0.1:8080
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.88ms    3.09ms 207.19ms   87.50%
    Req/Sec    40.07k     7.65k   82.20k    73.76%
  4783994 requests in 30.10s, 29.74GB read
Requests/sec: 158945.49
Transfer/sec:      0.99GB
```


## test-rust-hyper

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:5555
Running 30s test @ http://127.0.0.1:5555
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.81ms    2.97ms  43.76ms   88.71%
    Req/Sec    41.01k     6.65k   68.07k    70.56%
  4895365 requests in 30.08s, 30.34GB read
Requests/sec: 162735.60
Transfer/sec:      1.01GB
```


## test-cpp-drogon 4 threads (uses 4 cores)

```
test@test-virtual-machine:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:8080
Running 30s test @ http://127.0.0.1:8080
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.42ms    2.74ms  96.71ms   91.13%
    Req/Sec    43.00k     8.90k   71.02k    68.57%
  5138322 requests in 30.09s, 32.13GB read
Requests/sec: 170738.60
Transfer/sec:      1.07GB
```


## test-rust-minihttp

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:8080
Running 30s test @ http://127.0.0.1:8080
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.99ms    4.31ms 107.89ms   87.82%
    Req/Sec    51.80k    13.37k  120.95k    83.47%
  6182331 requests in 30.09s, 38.14GB read
Requests/sec: 205475.87
Transfer/sec:      1.27GB
```


---

## Some other unrelated test

## rust-actix-service

```
test@ubuntu:~$ ~/wrk/wrk -t4 -c400 -d30s http://127.0.0.1:8000
Running 30s test @ http://127.0.0.1:8000
  4 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.52ms    4.22ms  54.02ms   84.65%
    Req/Sec    42.43k     7.70k   81.40k    74.18%
  5071450 requests in 30.10s, 473.98MB read
Requests/sec: 168486.24
Transfer/sec:     15.75MB
```
