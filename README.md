Opm-Server-Golang
=================

Opera Mini Server Mirror with Golang

##Requirements

Go compiler toolchains <http://golang.org/doc/install>

Android apktool <https://code.google.com/p/android-apktool/>

##Usage

###Server

Build & Run

```
go build server.go
./server &
```

Set iptables

```
iptables -t nat -A PREROUTING -p tcp -m tcp --dport 9003 -j DNAT --to-destination 141.0.11.253:1080
iptables -t nat -A POSTROUTING -p tcp -m tcp --dport 9003 -j SNAT --to-source YOUR-SERVER-IP
```

###Client(for Android)

Download operamini.apk

`wget http://m.opera.com/android/Ow7Good/operamini.apk`
  
Decode apk file

`apktool d operamini.apk`
  
Replace proxy server with your server `opera.example.com:8080`

```
sed -i 's/"http:\/\/"/"http:\/\/opera.example.com:8080\/"/g' `find operamini -name '*.smali'`
sed -i 's/"socket:\/\/"/"socket:\/\/opera.example.com:9003\/"/g' `find operamini -name '*.smali'`  
```
  
Build apk file

`apktool b operamini`

Sign apk file

`java -jar signapk.jar testkey.x509.pem testkey.pk8 operamini/dist/operamini.apk operamini_mod.apk`

Enjoy~
