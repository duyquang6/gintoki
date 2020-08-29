# Gintoki - a best practice for handling high traffic read only service
I create this project to practice a service which only read database. So how I can scale up to 1m req/min(bench on local) ? 
By using local caching and kafka to evict local cache explicitly. I also use big cache to handle gigabytes data cache in case very large of data need to caching
And I believe this approach can scale up to 1m req/sec if it behind an L4 proxy  
# Performance
I only bench on my localmachine (Macbook Pro Core i5 8th quadcore) so the upperbound limit is still a mysterious number

If not using caching(only database), the limit maybe 600-800 req/sec:
```
echo "GET http://192.168.1.104:8080/cache?product_id=35586" | vegeta attack -rate=1000 -duration=60s | vegeta report
Requests      [total, rate, throughput]         60000, 1000.01, 481.02
Duration      [total, attack, wait]             1m28s, 59.999s, 27.757s
Latencies     [min, mean, 50, 90, 95, 99, max]  41.7µs, 1.036s, 46.711ms, 1.275s, 3.384s, 30s, 30.041s
Bytes In      [total, mean]                     24027032, 400.45
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           70.36%
Status Codes  [code:count]                      0:16223  200:42213  400:1564
```

With caching, it can serve upto without 10000 req/s any errors:
```
echo "GET http://192.168.1.104:8080/cache?product_id=35586" | vegeta attack -rate=10000 -duration=60s | vegeta report
Requests      [total, rate, throughput]         600000, 10000.03, 9999.99
Duration      [total, attack, wait]             1m0s, 1m0s, 219.241µs
Latencies     [min, mean, 50, 90, 95, 99, max]  104.294µs, 952.122µs, 209.047µs, 2.928ms, 6.256ms, 9.792ms, 49.215ms
Bytes In      [total, mean]                     340800000, 568.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]
```

And limit when I bench 20000 req/s because localmachine run out of available ports :(

```
echo "GET http://192.168.1.104:8080/cache?product_id=35586" | vegeta attack -rate=20000 -duration=60s | vegeta report
Requests      [total, rate, throughput]         457100, 6391.89, 4168.90
Duration      [total, attack, wait]             1m38s, 1m12s, 26.877s
Latencies     [min, mean, 50, 90, 95, 99, max]  89.174µs, 2.273s, 2.763ms, 608.898ms, 29.903s, 33.46s, 54.552s
Bytes In      [total, mean]                     232979400, 509.69
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           89.73%
Status Codes  [code:count]                      0:46925  200:410175
Get "http://192.168.1.104:8080/cache?product_id=35586": dial tcp 0.0.0.0:0->192.168.1.104:8080: bind: can't assign requested address
```

# Dependencies:
- Shopify/sarama Kafka consumer client.
- allegro/bigcache which can serve upto gigabytes data

# Usage:

In `test-docker` which you can start it, debezium connector will be plug into mysql binlog and produce data changes to kafka.
Downstream service only to consume that table topic and evict cache explicitly

# Contributing
All contributions are welcome.

# License
This library is distributed under MIT license, see LICENSE