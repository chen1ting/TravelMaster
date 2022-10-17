Deploying

```
❯ npm run build
❯ scp -r ./build/* VMadmin@172.21.148.164:/home/VMadmin/build

VMadmin@CZ3002-0001:~$ sudo rm -rf /usr/share/nginx/html/
VMadmin@CZ3002-0001:~$ sudo mkdir /usr/share/nginx/html
VMadmin@CZ3002-0001:~$ sudo mv build/* /usr/share/nginx/html
```
