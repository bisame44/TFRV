import csv
import sys
import matplotlib.pyplot as plt

header_str = 'Measuring Iterative Braun Method Performance'

csv_types = (int, float)

def read_file(filename: str) -> list[tuple[int, float]]:
    with open(filename, 'r', newline='') as f:
        csvreader = csv.reader((row for row in f if not row.startswith('#')), delimiter=' ')
        records = []
        for row in csvreader:
            cells = []
            for i in range(len(row)):
                cells.append(csv_types[i](row[i]))
            records.append(tuple(cells))
    return records


def build_graph(ax: plt.Axes, d: list[tuple[int, float]], desc: str):
    ax.set_title(desc)
    ax.set_xlabel('Nodes')
    ax.set_ylabel('Time (s.)')
    x, y = list(record[0] for record in d), list(record[1] for record in d)
    ax.set_xbound(0, max(x))
    ax.set_ybound(0, max(y))
    ax.plot(x, y)
    ax.legend()


def main():
    records = read_file('bench.csv')
    fig, axs = plt.subplots(1, 1, figsize=(8, 6))
    fig.suptitle(header_str)
    build_graph(axs, records, 'C_1 = 3; C_2 = 5; C_3 = 6.')
    fig.tight_layout()
    plt.savefig(f'graph/tmp.png')


if __name__ == '__main__':
    main()
