#! /bin/bash
##
##
for((i=1;i<=6000;i++))
do
	dd if=/dev/urandom of=./a/a-$i.bin bs=4k count=1
done
for((i=1;i<=6000;i++))
do
	dd if=/dev/urandom of=./b/b-$i.bin bs=4k count=1
done
for((i=1;i<=6000;i++))
do
	dd if=/dev/urandom of=./c/b-$i.bin bs=4k count=1
done
for((i=1;i<=6000;i++))
do
	dd if=/dev/urandom of=./d/b-$i.bin bs=4k count=1
done
