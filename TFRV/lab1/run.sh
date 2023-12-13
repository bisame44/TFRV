#!/bin/bash

proc=(2)
mb_seq=(1 2 4 8 16 24 32 40 48 56 64)

for pr in ${proc[@]}; do
    printf "%s\0" "${mb_seq[@]}" | xargs -0 -P "${#mb_seq[@]}" -n 1 -I {} sbatch -J conntest-{}mb-${pr} -N 2 -n ${pr} -o ./csv/.conntest-{}mb-${pr} -W --wrap="mpiexec ./conntest {}"
    echo "Размер пакета (в МБ),Среднее время (в сек.)" > ./csv/conntest-${pr}.csv
    for data in `ls -a -v ./csv/.conntest-*mb-${pr}`; do
        mb=$(echo "${data}" | sed 's/^.*-\([[:digit:]]*\)mb-.*$/\1/')
        echo "${mb},$(cat ${data})" >> ./csv/conntest-${pr}.csv
        rm "${data}"
    done
done
