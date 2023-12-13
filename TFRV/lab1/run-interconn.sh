#!/bin/bash

mb_seq=(1 2 4 8 16 24 32 40 48 56 64)

printf "%s\0" "${mb_seq[@]}" | xargs -0 -P "${#mb_seq[@]}" -n 1 -I {} sbatch -J conntest-{}mb-${pr} -n 2 --ntasks-per-node=2 --ntasks-per-socket=2 -o ./csv/.conntest-{}mb-onsocket -W --wrap="mpiexec ./conntest {}"
printf "%s\0" "${mb_seq[@]}" | xargs -0 -P "${#mb_seq[@]}" -n 1 -I {} sbatch -J conntest-{}mb-${pr} -n 2 --ntasks-per-node=2 --ntasks-per-socket=1 -o ./csv/.conntest-{}mb-onnode -W --wrap="mpiexec ./conntest {}"
printf "%s\0" "${mb_seq[@]}" | xargs -0 -P "${#mb_seq[@]}" -n 1 -I {} sbatch -J conntest-{}mb-${pr} -n 2 -N 2 -o ./csv/.conntest-{}mb-internode -W --wrap="mpiexec ./conntest {}"
echo "Размер пакета (в МБ),Среднее время (в сек.)" > ./csv/conntest-onsocket.csv
echo "Размер пакета (в МБ),Среднее время (в сек.)" > ./csv/conntest-onnode.csv
echo "Размер пакета (в МБ),Среднее время (в сек.)" > ./csv/conntest-internode.csv
for data in `ls -a -v ./csv/.conntest-*mb-*`; do
    mb=$(echo "${data}" | sed 's/^.*-\([[:digit:]]*\)mb-.*$/\1/')
    conntype=$(echo "${data}" | sed 's/^.*-\(.*\)$/\1/')
    echo "${mb},$(cat ${data})" >> ./csv/conntest-${conntype}.csv
    rm "${data}"
done
