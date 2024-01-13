
# ssh-geo-ip-blocker
[]
ssh-geo-ip-blockerは、Goをベースにしたアプリケーションで、日本以外からのSSHアクセスをブロックします。地理位置情報にはGeoLite2データベースを使用しています。

## 使い方

ビルド済みのバイナリを取得したら、それを実行可能に設定し、`hosts.deny`ファイルで使用するように設定します。

```shell  
sudo chmod +x /opt/ssh-geo-ip-blocker
```  

次に、`hosts.deny`ファイルに以下の行を追加します：

```text  
sshd: ALL: aclexec /opt/ip-blocking %a  
```  

この行では、`%a`は着信接続のIPアドレスに置き換えられます。(1 (ALLOW) or 0  (DENY)　を返す必要あり)


接続が試みられると、次のようなログエントリ(Auth.log)が表示されます：

```log  
ssh-geo-ip-blocker[1228485]: 2022/08/14 17:47:05 Allow sshd connection from  (Country:JP)  
sshd[1228483]: aclexec returned 1  
```  

接続が日本以外からの場合、接続は拒否され、次のようなログエントリが表示されます：

```log  
ssh-geo-ip-blocker[1228777]: 2022/08/14 17:48:03 Deny sshd connection from 52.229.29.153 (Country:US)  
sshd[1228775]: refused connect from 52.229.29.153 (52.229.29.153)  
ssh-geo-ip-blocker[1228849]: 2022/08/14 17:48:27 Deny sshd connection from 49.88.112.73 (Country:CN)  
sshd[1228847]: refused connect from 49.88.112.73 (49.88.112.73)  
ssh-geo-ip-blocker[1228923]: 2022/08/14 17:48:58 Deny sshd connection from 124.156.0.88 (Country:IN)  
```  

## Link
* [StoneDotのいろいろ: Ubuntu で日本以外からのSSHアクセスを拒否する \(GeoIP2 Python API ver\.\)](http://stonedot.blogspot.com/2014/05/ubuntu-ssh-geoip2-python-api-ver.html)
* [oschwald/geoip2\-golang: Unofficial MaxMind GeoIP2 Reader for Go](https://github.com/oschwald/geoip2-golang)
* [Ubuntu Manpage: hosts\_options \- host access control language extensions](http://manpages.ubuntu.com/manpages/bionic/man5/hosts_options.5.html)

## Disclaimer
This product includes GeoLite2 data created by MaxMind, available from  
[https://www.maxmind.com](https://www.maxmind.com).