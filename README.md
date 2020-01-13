# sts - Stats

## Usage

```sh
NAME:
   sts - stats

Generate simple stats for a stream of numbers

USAGE:
   sts numbers.txt or cat numbers.txt | sts

GLOBAL OPTIONS:
   --help, -h  show this help message
```

## Example
### Input
```sh
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
count:	10
min:	1
q1:	3
median:	5.5
q3:	8
max:	10
mean:	5.5
sum:	55
stddev:	3.0276503540974917
```

or

```sh
cat numbers.txt | sts
count:	10
min:	1
q1:	3
median:	5.5
q3:	8
max:	10
mean:	5.5
sum:	55
stddev:	3.0276503540974917
```