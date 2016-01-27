[![Build Status](https://travis-ci.org/logpacker/android-client.svg?branch=master)](https://travis-ci.org/logpacker/android-client)

#### How to import into Java project

Android Studio:

* File > New > New Module > Import .JAR or .AAR package
* File > Project Structure > app -> Dependencies -> Add Module Dependency
* Add import: *import go.logpackerandroid.Logpackerandroid;*

#### How to use it in Java:

```java
String logPackerClusterURL = "https://logpacker.mywebsite.com";

Logpackerandroid.Client logPackerClient;
Logpackerandroid.Message logMessage;

try {
    logPackerClient = Logpackerandroid.NewClient(logPackerClusterURL);

    logMessage = new Logpackerandroid.Message();
    logMessage.setMessage("Crash is here!");
    logMessage.setTagName("myapp");
    logMessage.setSource("paymentmodule");
    logMessage.setUserID("1001");
    logMessage.setUserName("John");

    logPackerClient.Send(logMessage);
} catch (Exception e) {
    // Cannot connect to Cluster
    // Or validation error
}
```

#### How to build an *.aar* package from Go package

* golang 1.5+
* go get golang.org/x/mobile/cmd/gomobile
* gomobile init
* Install [Android SDK](https://developer.android.com/sdk/index.html#Other) to ~/android-sdk-linux
* Install java-jdk
* ANDROID_HOME="/home/username/android-sdk-linux" gomobile bind .
