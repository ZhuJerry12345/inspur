#! /bin/bash
##
##
a=1
b=1
c=1
d=1
for((i=1;i<=6000;i++))
do
	seed=$[$RANDOM%4]
	if ((seed==0));then
		dd if=/dev/urandom of=./a/a-$a.bin bs=4k count=1
		a=$[$a+1]
	elif ((seed==1));then
		dd if=/dev/urandom of=./b/b-$b.bin bs=4k count=1
		b=$[$b+1]
	elif ((seed==2));then
		dd if=/dev/urandom of=./c/c-$c.bin bs=4k count=1
		c=$[$c+1]
	else
		dd if=/dev/urandom of=./d/d-$d.bin bs=4k count=1
		d=$[$d+1]	
	fi
done

