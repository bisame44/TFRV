#!usr/bin/env python

import csv
import os.path
import re
from typing import List
import matplotlib.pyplot as plt
import sys

mb_axis_title = 'Размер пакета (в МБ)'
time_axis_title = 'Среднее время (в сек.)'

def main(*args: str):
    for csvfilename in args:
        x_axis: List[int] = []
        y_axis: List[float] = []
        with open(csvfilename) as csvfile:
            reader = csv.DictReader(csvfile, delimiter=',')
            for row in reader:
                x_axis.append(int(row[mb_axis_title]))
                y_axis.append(float(row[time_axis_title]))

        csvbasename = os.path.basename(csvfilename)
        procnum = re.search(r'(\d+)\.csv', csvbasename).group(1)
        plt.plot(x_axis, y_axis, color='blue', linestyle='--', linewidth=1, marker='o', markersize=4, mec='black', mfc='black', label=f'Процессов: {procnum}')
        plt.xlim(0, max(x_axis))
        plt.xticks(range(0, max(x_axis) + 1, max(x_axis) // 8))
        plt.ylim(0, max(y_axis))
        plt.xlabel(mb_axis_title)
        plt.ylabel(time_axis_title)

        plt.title(f'Среднее время передачи пакета\nот одного процесса другому (для {procnum} процессов)')

        plt.savefig('graph/' + csvbasename[:-4] + '.png')

if __name__ == '__main__':
    main(*sys.argv[1:])
