# sts - Stats

## Usage

```sh
NAME:
   sts - pronounced 'stats'

   Generate simple stats for a stream of numbers

USAGE:
   seq 1 10 | sts
   sts numbers.txt
   seq 1 10 | sts summary
   seq 1 10 | sts s


COMMANDS:
   summary, s  (default) prints summary statistics for the dataset

GLOBAL OPTIONS:
   --help, -h  show this help message
```

## Example
### Input
```sh
$ seq 1 10 > numbers.txt
$ cat numbers.txt
1
2
3
4
5
6
7
8
9
10

```

### Output
```sh
$ sts numbers.txt
count: 10
min:   1
max:   10
sum:   55

mean:   5.5
stddev: 3.0276503540974917

q1:     3
median: 5
q3:     8
```

or

```sh
cat numbers.txt | sts
count: 10
min:   1
max:   10
sum:   55

mean:   5.5
stddev: 3.0276503540974917

q1:     3
median: 5
q3:     8
```