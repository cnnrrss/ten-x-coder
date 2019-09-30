# Concurrency in Python

### I/O vs CPU Bound Processes

Speeding up an I/O-Bound Process involves overlapping the times spent wiating for devices.

Speeind up CPU-Bound Process invloves finding ways to do more computations in the same amount of time.

**Overlapping Requests**: For I/O bound processes you can spin up a worker pool (in python 5-10 is a good place to start)
so that you can begin making requests before the previous request returns.

```python
import concurrent.futures
import threading

with concurrent.futures.ThreadPoolExecutor(max_workers=5) as executor:
    executor.map(download_data, data_obj)
```



