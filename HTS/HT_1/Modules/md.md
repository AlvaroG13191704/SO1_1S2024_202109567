obj-m += ram.o

all:
	make -C /usr/lib/modules/6.5.13-7-MANJARO/build M=$(PWD) modules

clean:
	make -C /usr/lib/modules/6.5.13-7-MANJARO/build M=$(PWD) clean
