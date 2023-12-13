# Теория функционирования распределённых вычислительных систем

## Лабораторная работа №1.

**[Задание](../docs/dcsft-lab1.pdf)**

## Процедура запуска

Тестировалось на кластере с Linux, OpenMPI 4, Slurm; на хосте установлены CPython 3.9 (с *matplotlib*), компилятор C++17.

```sh
make
./run-interconn.sh
python graph-interconn.py csv/conntest-internode.csv csv/conntest-onnode.csv csv/conntest-onsocket.csv
```

На выходе графики по пути:
* `./graph/conntest-internode.png` *(между узлами)*;
* `./graph/conntest-onnode.png` *(на одном узле, между сокетами)*;
* `./graph/conntest-onsocket.png` *(на одном сокете)*.
