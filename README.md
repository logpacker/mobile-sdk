[![Build Status](https://travis-ci.org/logpacker/mobile-sdk.svg?branch=master)](https://travis-ci.org/logpacker/mobile-sdk)

#### How to import into Java project

Android Studio:

* File > New > New Module > Import .JAR or .AAR package
* File > Project Structure > app -> Dependencies -> Add Module Dependency
* Add import: *import go.logpackermobilesdk.Logpackermobilesdk;*

#### How to use it in Java:

```java
String logPackerClusterURL = "https://logpacker.mywebsite.com";
String logPackerEnv = "development";
String logPackerAgent = "Nexus";

Logpackermobilesdk.Client client;
Logpackermobilesdk.Message msg;

try {
    client = Logpackermobilesdk.NewClient(logPackerClusterURL, logPackerEnv, logPackerAgent);

    msg = Logpackermobilesdk.NewMessage();
    msg.setMessage("Crash is here!");
    msg.setTagName("myapp");
    msg.setSource("paymentmodule");
    msg.setUserID("1001");
    msg.setUserName("John");

    client.Send(msg);
} catch (Exception e) {
    // Cannot connect to Cluster
    // Or validation error
}
```

#### How to send Android crashes to LogPacker

You must catch uncaughtException in your application and use LogPacker to send the exception:

```java
public class MyApplication extends Application {
    public void onCreate () {
        // Setup handler for uncaught exceptions.
        Thread.setDefaultUncaughtExceptionHandler (new Thread.UncaughtExceptionHandler() {
            @Override
            public void uncaughtException (Thread thread, Throwable e) {
                handleUncaughtException (thread, e);
            }
        });
    }

    public void handleUncaughtException (Thread thread, Throwable e) {
        String logPackerClusterURL = "https://logpacker.mywebsite.com";
        String logPackerEnv = "development";
        String logPackerAgent = "Nexus";

        Logpackermobilesdk.Client client;
        Logpackermobilesdk.Message msg;

        try {
            client = Logpackermobilesdk.NewClient(logPackerClusterURL, logPackerEnv, logPackerAgent);

            msg = Logpackermobilesdk.NewMessage();
            msg.setMessage(e.getMessage());
            msg.setTagName("crash");
            msg.setSource("crash");
            msg.setUserID("1001");
            msg.setUserName("John");

            client.Send(msg);
        } catch (Exception e) {
            // Cannot connect to Cluster
            // Or validation error
        }

        System.exit(1); // kill off the crashed app
    }
}
```

#### How to import into iOS project

Untar Logpackermobilesdk.framework.tar to the root of your project. Or drag Logpackermobilesdk.framework folder to the Xcode's file browser.

#### How to use in Xcode

```c
#import "Logpackermobilesdk/Logpackermobilesdk.h"
```

#### How to build an *.aar* package from Go package

* golang 1.5+
* go get golang.org/x/mobile/cmd/gomobile
* gomobile init
* Install [Android SDK](https://developer.android.com/sdk/index.html#Other) to ~/android-sdk
* ~/android-sdk/tools/android sdk
* java-jdk or
* export ANDROID_HOME=$HOME"/android-sdk" && gomobile bind --target=android .
* Find *.aar* file in working folder
* gomobile bind --target=ios .
* Find Logpackermobilesdk.framework folder

#### How to contribute

* Fork master branch
* Make changes
* Run ./before-commit.sh
* Push and create a Pull Request
