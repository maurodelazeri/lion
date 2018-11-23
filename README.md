# lion

Zinnion Lion is a set of shared libraries to work across projects

```
cd /opt
wget https://repo.zabbix.com/zabbix/4.0/ubuntu/pool/main/z/zabbix-release/zabbix-release_4.0-2+xenial_all.deb
dpkg -i zabbix-release_4.0-2+xenial_all.deb
apt update
apt install -y zabbix-agent


hostname=$(hostname -s)
sed -i "s/Server=127.0.0.1/Server=192.168.1.202/g" /etc/zabbix/zabbix_agentd.conf
sed -i "s/ServerActive=127.0.0.1/ServerActive=192.168.1.202/g" /etc/zabbix/zabbix_agentd.conf
sed  -i "s/Hostname=Zabbix server/Hostname=${hostname}/g" /etc/zabbix/zabbix_agentd.conf
systemctl restart zabbix-agent
systemctl enable zabbix-agent
```

```
hostname=$(hostname -s | tr '[:lower:]' '[:upper:]')
```

```
/root/work/src/github.com/maurodelazeri/winter/winter --venues $(hostname -s | tr '[:lower:]' '[:upper:]')
```

```
#!/bin/bash

# -x flag only match processes whose name (or command line if -f is
# specified) exactly match the pattern.

if pgrep -x "winter" > /dev/null
then
    echo "Running"
else
    echo "Stopped"
    /root/work/src/github.com/maurodelazeri/winter/winter --venues $(hostname -s | tr '[:lower:]' '[:upper:]')
fi
```
