# Remotgo

Send commands over ssh to AWS EC2 instances

### Example
Execute "df -H" command in all instances with core as role and test as environment
```
> go run main.go -role core -environment test -command "df -H"

+---------------------------------------------------+
| ec2-xx-xxx-xxx-xx.eu-west-1.compute.amazonaws.com |
+---------------------------------------------------+
Filesystem      Size  Used Avail Use% Mounted on
/dev/xvda1       43G   27G   14G  68% /
udev            3.9G   13k  3.9G   1% /dev
tmpfs           1.6G  222k  1.6G   1% /run
none            5.3M     0  5.3M   0% /run/lock
none            4.0G     0  4.0G   0% /run/shm
/dev/xvdb        43G   11G   30G  26% /mnt

+---------------------------------------------------+
| ec2-xx-xxx-xx-xx.eu-west-1.compute.amazonaws.com |
+---------------------------------------------------+
Filesystem      Size  Used Avail Use% Mounted on
/dev/xvda1       64G   13G   48G  22% /
udev            2.0G   13k  2.0G   1% /dev
tmpfs           394M  213k  394M   1% /run
none            5.3M     0  5.3M   0% /run/lock
none            2.0G     0  2.0G   0% /run/shm
/dev/xvdb        17G  961M   15G   7% /mnt
```

