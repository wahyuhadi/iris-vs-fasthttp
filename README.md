# iris-vs-fasthttp-vs-net/http

##  iris

 wrk -c500 -t 4 --latency -d 60s http://207.148.119.124:8080/hello/wahyu

	[rwx dir ~ ]$ wrk -c500 -t 4 --latency -d 60s http://207.148.119.124:8080/hello/wahyu
	Running 1m test @ http://207.148.119.124:8080/hello/wahyu
	  4 threads and 500 connections
	  Thread Stats   Avg      Stdev     Max   +/- Stdev
	    Latency    35.50ms   53.69ms 729.21ms   95.39%
	    Req/Sec   246.05    121.77   550.00     62.01%
	  Latency Distribution
	     50%   22.95ms
	     75%   25.70ms
	     90%   30.46ms
	     99%  270.95ms
	  57403 requests in 1.00m, 4.11MB read
	Requests/sec:    955.27
	Transfer/sec:     69.97KB
	
## fasthttp

wrk -c500 -t 4 --latency -d 60s http://207.148.119.124:8080/hello/wahyu

	[rwx dir ~ ]$ wrk -c500 -t 4 --latency -d 60s http://207.148.119.124:8080/hello/wahyu
	Running 1m test @ http://207.148.119.124:8080/hello/wahyu
	  4 threads and 500 connections
	  Thread Stats   Avg      Stdev     Max   +/- Stdev
	    Latency    32.27ms   47.29ms 948.98ms   96.62%
	    Req/Sec   290.76    164.26   636.00     57.37%
	  Latency Distribution
	     50%   22.91ms
	     75%   25.35ms
	     90%   29.31ms
	     99%  265.08ms
	  68444 requests in 1.00m, 9.73MB read
	Requests/sec:   1339.03
	Transfer/sec:    165.74KB
	
## net/http
 wrk -c500 -t 4 --latency -d 60s http://207.148.119.124:8080/hello

	[rwx dir ~ ]$ wrk -c500 -t 4 --latency -d 60s http://207.148.119.124:8080/hello

	Running 1m test @ http://207.148.119.124:8080/hello
	  4 threads and 500 connections
	  Thread Stats   Avg      Stdev     Max   +/- Stdev
	    Latency    64.90ms  105.14ms   1.66s    85.96%
	    Req/Sec   261.69    190.00     0.87k    60.61%
	  Latency Distribution
	     50%   25.29ms
	     75%   30.16ms
	     90%  254.25ms
	     99%  487.28ms
	  60264 requests in 1.00m, 4.31MB read
	Requests/sec:   1003.41
	Transfer/sec:     73.49KB
	
	
# server 

pengujian ini dilakukan pada server dengan 1 core cpu dan ram 1 GB


