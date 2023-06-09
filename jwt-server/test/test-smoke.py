import multiprocessing
from time import time
import requests
from sys import argv

def test(cycles, proc_num, return_list:list):
    times = []
    for _ in range(cycles):
        start = time()
        res = requests.get(
            "http://localhost:8080/ping",
            # headers=headers,
            # params=params        
        )
        end = time()
        times.append(end-start)

    sum_times = sum(times)
    avg = sum_times/cycles
    return_list[proc_num] = avg

if __name__ == '__main__':
    cycles = int(argv[1])
    p_num = int(argv[2])
    processes = []
    return_list = multiprocessing.Manager().dict()

    for i in range(p_num):
        p = multiprocessing.Process(target=test, args=(cycles, i, return_list))
        processes.append(p)
        p.start()

    for p in processes:
        p.join()
    total = 0
    for i in return_list.values():
        total += i

    print (f"It took on average {(total/p_num)*1000}ms")
    print (return_list)

