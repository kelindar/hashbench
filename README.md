## Experiment: Benchmarks of Hash Functions in Go

This repository contains a small experiment of mine, trying out different hash functions and comparing their performance against each other on different data sizes. Note that this is not a reason to choose one hash function over another, as there's the collisions and randomness that need to be taken into consideration. 

```
  32-bit |                 crc |   10 B |          17.40 ns/op |   574.59 MB/s  
  32-bit |                 crc |   25 B |          15.53 ns/op |  1609.75 MB/s  
  32-bit |                 crc |  100 B |          24.53 ns/op |  4077.08 MB/s  
  32-bit |                 crc | 4.0 kB |          163.5 ns/op | 24461.54 MB/s  
  32-bit |                 crc |  10 MB |         463880 ns/op | 21557.31 MB/s  
  32-bit |                 crc |  50 MB |        2756338 ns/op | 18140.01 MB/s  
  32-bit | dgryski/go-marvin32 |   10 B |          6.694 ns/op |  1493.84 MB/s  
  32-bit | dgryski/go-marvin32 |   25 B |          10.61 ns/op |  2357.37 MB/s  
  32-bit | dgryski/go-marvin32 |  100 B |          31.98 ns/op |  3126.93 MB/s  
  32-bit | dgryski/go-marvin32 | 4.0 kB |           1176 ns/op |  3402.06 MB/s  
  32-bit | dgryski/go-marvin32 |  10 MB |        3006061 ns/op |  3326.61 MB/s  
  32-bit | dgryski/go-marvin32 |  50 MB |       15063943 ns/op |  3319.18 MB/s  
  32-bit |    mtchavez/jenkins |   10 B |          11.98 ns/op |   834.73 MB/s  
  32-bit |    mtchavez/jenkins |   25 B |          27.03 ns/op |   924.73 MB/s  
  32-bit |    mtchavez/jenkins |  100 B |          111.1 ns/op |   900.02 MB/s  
  32-bit |    mtchavez/jenkins | 4.0 kB |           4353 ns/op |   918.90 MB/s  
  32-bit |    mtchavez/jenkins |  10 MB |       10980498 ns/op |   910.71 MB/s  
  32-bit |    mtchavez/jenkins |  50 MB |       54597529 ns/op |   915.79 MB/s  
  32-bit |        twmb/murmur3 |   10 B |          8.488 ns/op |  1178.15 MB/s  
  32-bit |        twmb/murmur3 |   25 B |          12.63 ns/op |  1978.82 MB/s  
  32-bit |        twmb/murmur3 |  100 B |          32.41 ns/op |  3085.30 MB/s  
  32-bit |        twmb/murmur3 | 4.0 kB |           1090 ns/op |  3669.56 MB/s  
  32-bit |        twmb/murmur3 |  10 MB |        2663511 ns/op |  3754.44 MB/s  
  32-bit |        twmb/murmur3 |  50 MB |       13308251 ns/op |  3757.07 MB/s  
  32-bit |     dgryski/go-farm |   10 B |          7.941 ns/op |  1259.27 MB/s  
  32-bit |     dgryski/go-farm |   25 B |          12.22 ns/op |  2045.20 MB/s  
  32-bit |     dgryski/go-farm |  100 B |          20.59 ns/op |  4856.51 MB/s  
  32-bit |     dgryski/go-farm | 4.0 kB |          587.7 ns/op |  6805.68 MB/s  
  32-bit |     dgryski/go-farm |  10 MB |        1458989 ns/op |  6854.06 MB/s  
  32-bit |     dgryski/go-farm |  50 MB |        7373852 ns/op |  6780.72 MB/s  
  64-bit |                 crc |   10 B |          18.65 ns/op |   536.29 MB/s  
  64-bit |                 crc |   25 B |          41.12 ns/op |   607.93 MB/s  
  64-bit |                 crc |  100 B |          53.58 ns/op |  1866.47 MB/s  
  64-bit |                 crc | 4.0 kB |           1707 ns/op |  2343.22 MB/s  
  64-bit |                 crc |  10 MB |        4095172 ns/op |  2441.90 MB/s  
  64-bit |                 crc |  50 MB |       20528605 ns/op |  2435.63 MB/s  
  64-bit |                fnv1 |   10 B |          10.38 ns/op |   963.59 MB/s  
  64-bit |                fnv1 |   25 B |          19.78 ns/op |  1263.97 MB/s  
  64-bit |                fnv1 |  100 B |          82.33 ns/op |  1214.55 MB/s  
  64-bit |                fnv1 | 4.0 kB |           3517 ns/op |  1137.40 MB/s  
  64-bit |                fnv1 |  10 MB |        8792132 ns/op |  1137.38 MB/s  
  64-bit |                fnv1 |  50 MB |       43779073 ns/op |  1142.10 MB/s  
  64-bit |               fnv1a |   10 B |          9.441 ns/op |  1059.16 MB/s  
  64-bit |               fnv1a |   25 B |          20.00 ns/op |  1250.14 MB/s  
  64-bit |               fnv1a |  100 B |          81.70 ns/op |  1223.94 MB/s  
  64-bit |               fnv1a | 4.0 kB |           3523 ns/op |  1135.51 MB/s  
  64-bit |               fnv1a |  10 MB |        8799812 ns/op |  1136.39 MB/s  
  64-bit |               fnv1a |  50 MB |       43906431 ns/op |  1138.79 MB/s  
  64-bit |             maphash |   10 B |          15.63 ns/op |   639.94 MB/s  
  64-bit |             maphash |   25 B |          15.86 ns/op |  1575.95 MB/s  
  64-bit |             maphash |  100 B |          27.89 ns/op |  3586.14 MB/s  
  64-bit |             maphash | 4.0 kB |          677.8 ns/op |  5901.29 MB/s  
  64-bit |             maphash |  10 MB |        1670025 ns/op |  5987.93 MB/s  
  64-bit |             maphash |  50 MB |        8379192 ns/op |  5967.16 MB/s  
  64-bit |      cespare/xxhash |   10 B |          5.015 ns/op |  1993.99 MB/s  
  64-bit |      cespare/xxhash |   25 B |          6.351 ns/op |  3936.31 MB/s  
  64-bit |      cespare/xxhash |  100 B |          12.83 ns/op |  7791.79 MB/s  
  64-bit |      cespare/xxhash | 4.0 kB |          226.2 ns/op | 17681.05 MB/s  
  64-bit |      cespare/xxhash |  10 MB |         617066 ns/op | 16205.73 MB/s  
  64-bit |      cespare/xxhash |  50 MB |        3410148 ns/op | 14662.12 MB/s  
  64-bit |    dgryski/go-metro |   10 B |          5.474 ns/op |  1826.70 MB/s  
  64-bit |    dgryski/go-metro |   25 B |          7.320 ns/op |  3415.30 MB/s  
  64-bit |    dgryski/go-metro |  100 B |          14.08 ns/op |  7099.92 MB/s  
  64-bit |    dgryski/go-metro | 4.0 kB |          207.3 ns/op | 19291.71 MB/s  
  64-bit |    dgryski/go-metro |  10 MB |         595203 ns/op | 16800.98 MB/s  
  64-bit |    dgryski/go-metro |  50 MB |        3720653 ns/op | 13438.50 MB/s  
  64-bit |        twmb/murmur3 |   10 B |          6.658 ns/op |  1501.89 MB/s  
  64-bit |        twmb/murmur3 |   25 B |          8.257 ns/op |  3027.60 MB/s  
  64-bit |        twmb/murmur3 |  100 B |          16.06 ns/op |  6225.24 MB/s  
  64-bit |        twmb/murmur3 | 4.0 kB |          441.8 ns/op |  9052.85 MB/s  
  64-bit |        twmb/murmur3 |  10 MB |        1118167 ns/op |  8943.21 MB/s  
  64-bit |        twmb/murmur3 |  50 MB |        5754655 ns/op |  8688.62 MB/s  
  64-bit |   minio/highwayhash |   10 B |          45.22 ns/op |   221.14 MB/s  
  64-bit |   minio/highwayhash |   25 B |          46.43 ns/op |   538.40 MB/s  
  64-bit |   minio/highwayhash |  100 B |          52.96 ns/op |  1888.39 MB/s  
  64-bit |   minio/highwayhash | 4.0 kB |          249.1 ns/op | 16059.83 MB/s  
  64-bit |   minio/highwayhash |  10 MB |         615502 ns/op | 16246.89 MB/s  
  64-bit |   minio/highwayhash |  50 MB |        3323853 ns/op | 15042.78 MB/s  
  64-bit |    dgryski/go-sip13 |   10 B |          10.81 ns/op |   924.95 MB/s  
  64-bit |    dgryski/go-sip13 |   25 B |          13.05 ns/op |  1915.75 MB/s  
  64-bit |    dgryski/go-sip13 |  100 B |          26.12 ns/op |  3827.98 MB/s  
  64-bit |    dgryski/go-sip13 | 4.0 kB |          703.3 ns/op |  5687.29 MB/s  
  64-bit |    dgryski/go-sip13 |  10 MB |        1740708 ns/op |  5744.79 MB/s  
  64-bit |    dgryski/go-sip13 |  50 MB |        8709560 ns/op |  5740.82 MB/s  
  64-bit |        dgryski/tsip |   10 B |          6.988 ns/op |  1431.08 MB/s  
  64-bit |        dgryski/tsip |   25 B |          9.719 ns/op |  2572.22 MB/s  
  64-bit |        dgryski/tsip |  100 B |          20.73 ns/op |  4825.03 MB/s  
  64-bit |        dgryski/tsip | 4.0 kB |          645.0 ns/op |  6201.13 MB/s  
  64-bit |        dgryski/tsip |  10 MB |        1609975 ns/op |  6211.28 MB/s  
  64-bit |        dgryski/tsip |  50 MB |        8120107 ns/op |  6157.55 MB/s  
  64-bit |     dgryski/go-farm |   10 B |          5.747 ns/op |  1740.12 MB/s  
  64-bit |     dgryski/go-farm |   25 B |          6.572 ns/op |  3803.75 MB/s  
  64-bit |     dgryski/go-farm |  100 B |          27.64 ns/op |  3617.79 MB/s  
  64-bit |     dgryski/go-farm | 4.0 kB |          310.1 ns/op | 12901.04 MB/s  
  64-bit |     dgryski/go-farm |  10 MB |         861088 ns/op | 11613.21 MB/s  
  64-bit |     dgryski/go-farm |  50 MB |        4535989 ns/op | 11022.95 MB/s  
```