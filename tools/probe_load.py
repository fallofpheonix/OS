import time
import multiprocessing

def cpu_load():
    while True:
        x = 1000 * 1000

if __name__ == '__main__':
    processes = []
    for _ in range(multiprocessing.cpu_count()):
        p = multiprocessing.Process(target=cpu_load)
        p.start()
        processes.append(p)
    
    time.sleep(10)
    
    for p in processes:
        p.terminate()
