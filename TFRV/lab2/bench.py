import json
from os import path
import random
import subprocess
import sys

algo_list = [
    ('FFDH', 'First Fit Decreasing Height', '-ffdh'),
    ('NFDH', 'Next Fit Decreasing Height', '-nfdh')
]
mr = range(500, 5001, 500)
tj = range(1, 101)


def buildRandomTasks(n: int, path: str):
    rj = range(1, n+1)
    for m in mr:
        with open(path, 'w') as f:
            for _ in range(m):
                print(f'{random.choice(rj)} {random.choice(tj)}', file=f)
        yield m


def main(n: int):
    if n <= 0:
        raise Exception(f'n must be positive (got {n} in args)')
    random.seed(n)
    for m in buildRandomTasks(n, 'tmp.csv'):
        for algo in algo_list:
            p = subprocess.run(["./main", algo[2], "-n", f"{n}", "-f", "tmp.csv"], capture_output=True)
            if len(p.stderr) != 0:
                print(str(p.stderr))
                break
            algo_ef = json.loads(p.stdout)
            print(f'{algo[0]} {m} {n} {algo_ef["PerformanceTime"]:.6f} {algo_ef["ScheduleTotalTime"]} {algo_ef["ScheduleDeviation"]:.6f}')


if __name__ == '__main__':
    if len(sys.argv) < 2:
        print(f'usage: {sys.argv[0]} {{n...}}', file=sys.stderr)
    else:
        if not path.exists('main'):
            subprocess.run(["go", "build", "main.go"])
        print(f'# Algorithm Tasks Rank PerformanceTime ScheduleTotalTime ScheduleDeviation')
        for n in sys.argv[1:]:
            main(int(n))
