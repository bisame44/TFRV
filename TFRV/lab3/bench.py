import json
from os import path
import subprocess
import sys

def main(n: int):
    if n <= 0:
        raise Exception(f'n must be positive (got {n} in args)')
    p = subprocess.run(["./main", "-output-time", f"{n}", "3", "4", "5"], capture_output=True)
    if len(p.stderr) != 0:
        print(str(p.stderr))
    algo_ef = json.loads(p.stdout)
    print(f'{n} {algo_ef["TimeSpentSeconds"]}')


if __name__ == '__main__':
    if len(sys.argv) < 2:
        print(f'usage: {sys.argv[0]} {{n...}}', file=sys.stderr)
    else:
        if not path.exists('main'):
            subprocess.run(["go", "build", "main.go"])
        print(f'# Nodes PerformanceTime')
        for n in sys.argv[1:]:
            main(int(n))
