
# ssh-geo-ip-blocker
[]
ssh-geo-ip-blockerは、Goをベースにしたアプリケーションで、日本以外からのSSHアクセスをブロックします。地理位置情報にはGeoLite2データベースを使用しています。

## 使い方

ビルド済みのバイナリを取得したら、それを実行可能に設定し、`hosts.allow`ファイルで使用するように設定します。

```shell  
sudo chmod +x /opt/ssh-geo-ip-blocker
```  

次に、`hosts.allow`ファイルに以下の行を追加します：

```text  
sshd: ALL: aclexec /opt/ip-blocking %a  
```  
この行では、`%a`は着信接続のIPアドレスに置き換えられます。(0(ALLOW) or 1 (DENY)　を返す必要あり)

`hosts.deny`ファイルに以下の行を追加します：
```text
sshd : all
```
許可されていない接続をブロックします



## Link
* [StoneDotのいろいろ: Ubuntu で日本以外からのSSHアクセスを拒否する \(GeoIP2 Python API ver\.\)](http://stonedot.blogspot.com/2014/05/ubuntu-ssh-geoip2-python-api-ver.html)
* [oschwald/geoip2\-golang: Unofficial MaxMind GeoIP2 Reader for Go](https://github.com/oschwald/geoip2-golang)
* [Ubuntu Manpage: hosts\_options \- host access control language extensions](http://manpages.ubuntu.com/manpages/bionic/man5/hosts_options.5.html)

## Disclaimer
This product includes GeoLite2 data created by MaxMind, available from  
[https://www.maxmind.com](https://www.maxmind.com).