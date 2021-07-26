# NaffEncryption
Very basic example of using AES symmetric encryption

Very basic example of AES symmetric encryption, just messing about with some encryption code. God's sake don't even think about copying this, it's so insecure but the principles are just about useful to get an idea!

```
C:\go\ > .\enc01.exe -genkey
Encryption key: 169c561cedd20c94c42bb93d9a58f448f45083b7f65991976cf43e430b341b8c
C:\go\ > .\enc01.exe -encrypt -key 169c561cedd20c94c42bb93d9a58f448f45083b7f65991976cf43e430b341b8c -msg "Stuff"
0000000000000000000000001023a0e983e7659e3fcd659763332814631892f1ae
C:\go\ > .\enc01.exe -decrypt -key 169c561cedd20c94c42bb93d9a58f448f45083b7f65991976cf43e430b341b8c -msg 0000000000000000000000001023a0e983e7659e3fcd659763332814631892f1ae
Stuff
```
