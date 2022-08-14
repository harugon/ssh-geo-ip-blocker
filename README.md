# ssh-geo-ip-blocker
日本以外からのSSHアクセスを拒否する

## geoip2

[GeoLite2 Free Geolocation Data \| MaxMind Developer Portal](https://dev.maxmind.com/geoip/geolite2-free-geolocation-data?lang=en)
にてGeoLite2-Country.mmdbをダウンロードし配置

```go
//go:embed src/GeoLite2-Country.mmdb
var mmdb []byte
```
ファイルはバイナリに埋め込まれます。

## 使う
ビルド
```shell
go build -o ssh-geo-ip-blocker main.go
```
``ip-blocking``が作成される

```shell
sudo chmod +x /opt/ssh-geo-ip-blocker
```
実行できるようにする

hosts.deny に追記
```text:hosts.deny
sshd: ALL: aclexec /opt/ip-blocking %a

```
(%a　にアクセス元IPアドレスが入る)

### Auth.logに出力される
```log
ssh-geo-ip-blocker[1228485]: 2022/08/14 17:47:05 Allow sshd connection from  (Country:JP)
sshd[1228483]: aclexec returned 1
~~~~
ssh-geo-ip-blocker[1228777]: 2022/08/14 17:48:03 Deny sshd connection from 52.229.29.153 (Country:US)
sshd[1228775]: refused connect from 52.229.29.153 (52.229.29.153)
ssh-geo-ip-blocker[1228849]: 2022/08/14 17:48:27 Deny sshd connection from 49.88.112.73 (Country:CN)
sshd[1228847]: refused connect from 49.88.112.73 (49.88.112.73)
ssh-geo-ip-blocker[1228923]: 2022/08/14 17:48:58 Deny sshd connection from 124.156.0.88 (Country:IN)
```
## memo

```text
1 (ALLOW) or 0  (DENY)　を返す必要あり
```

```shell
$env:GOOS="linux"
$env:GOARCH="amd64"
```
windowsでbuildする場合

## Link
* [StoneDotのいろいろ: Ubuntu で日本以外からのSSHアクセスを拒否する \(GeoIP2 Python API ver\.\)](http://stonedot.blogspot.com/2014/05/ubuntu-ssh-geoip2-python-api-ver.html)
* [oschwald/geoip2\-golang: Unofficial MaxMind GeoIP2 Reader for Go](https://github.com/oschwald/geoip2-golang)
* [Ubuntu Manpage: hosts\_options \- host access control language extensions](http://manpages.ubuntu.com/manpages/bionic/man5/hosts_options.5.html)
