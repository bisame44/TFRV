#include <iomanip>
#include <iostream>
#include <vector>
#include <mpi.h>

int main(int argc, char *argv[]) {
    int rank, commsize;

    MPI_Init(&argc, &argv);
    MPI_Comm_size(MPI_COMM_WORLD, &commsize);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);

    const int n_mb = (argc >= 2) ? atoi(argv[1]) : 4;
    const int n = n_mb * 1024 * 1024;

    std::vector<uint8_t> sbuf(n);
    std::vector<uint8_t> rbuf(n);
    std::vector<MPI_Request> reqs(2 * commsize);
    std::vector<MPI_Status> stats(2 * commsize);
    double T, Tmax;

    T = MPI_Wtime();

    int i = rank;
    do {
        MPI_Isend(sbuf.data(), n, MPI_UINT8_T, i, 1, MPI_COMM_WORLD, &reqs[i]);
        MPI_Irecv(rbuf.data(), n, MPI_UINT8_T, i, 1, MPI_COMM_WORLD, &reqs[commsize + i]);
        i = (i + 1) % commsize;
    } while (i != rank);
    MPI_Waitall(commsize * 2, reqs.data(), stats.data());

    T = MPI_Wtime() - T;
    MPI_Reduce(&T, &Tmax, 1, MPI_DOUBLE, MPI_MAX, 0, MPI_COMM_WORLD);
    if (rank == 0) {
        std::cout << std::setprecision(6) << Tmax << std::endl;
    }

    MPI_Finalize();

    return 0;
}
