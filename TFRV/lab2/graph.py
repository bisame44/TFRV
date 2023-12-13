import csv
import sys
import matplotlib.pyplot as plt

algo_list = [
    ('NFDH', 'Next Fit Decreasing Height'),
    ('FFDH', 'First Fit Decreasing Height')
]

header_str = 'Measuring 2D Packing Algo Performance'

desc_map = [
    (3, 'Performance Time'),
    (4, 'Schedule Total Time'),
    (5, 'Schedule Deviation')
]

csv_types = (str, int, int, float, int, float)

def read_file(filename: str) -> list[tuple[any, ...]]:
    with open(filename, 'r', newline='') as f:
        csvreader = csv.reader((row for row in f if not row.startswith('#')), delimiter=' ')
        records = []
        for row in csvreader:
            cells = []
            for i in range(len(row)):
                cells.append(csv_types[i](row[i]))
            records.append(tuple(cells))
    return records


def filter_rows(records: list[tuple[any, ...]], n: int, i: int) -> dict[str, list[tuple[int, any]]]:
    d: dict[str, list[any]] = dict()
    for record in records:
        if record[2] != n:
            continue
        if record[0] not in d:
            d[record[0]] = list()
        d[record[0]].append((record[1], record[i]))
    return d


def build_graph(ax: plt.Axes, d: dict[str, list[tuple[int, any]]], desc: str):
    ax.set_title(desc)
    ax.set_xlabel('Tasks')
    for algo, records in d.items():
        x, y = list(record[0] for record in records), list(record[1] for record in records)
        ax.plot(x, y, label=algo)
    ax.legend()


def main(n: int):
    records = read_file('bench.csv')
    fig, axs = plt.subplots(len(desc_map), 1, figsize=(6, 8))
    fig.suptitle(header_str)
    for i in range(len(desc_map)):
        d = filter_rows(records, n, desc_map[i][0])
        build_graph(axs[i], d, desc_map[i][1])
    fig.tight_layout()
    plt.savefig(f'graph/tmp_{n}.png')


if __name__ == '__main__':
    if len(sys.argv) < 2: 
        print(f'usage: {sys.argv[0]} {{n...}}', file=sys.stderr)
    else:
        for n in sys.argv[1:]:
            main(int(n))
