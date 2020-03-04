# JavaGolangCallDynamicLib

 javac -d . -classpath jna-5.5.0.jar HelloWorld.java
 
 java -classpath .:jna-5.5.0.jar HelloWorld
 
 go build -o client.so -buildmode=c-shared client.go
 
 gcc Hello.c -fPIC -shared -o libHello.so
