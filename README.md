[![Build Status](https://travis-ci.org/logpacker/mobile-sdk.svg?branch=master)](https://travis-ci.org/logpacker/mobile-sdk)

#### How to import into Java project

Android Studio (see scnreenshots/ folder):

* File > New > New Module > Import .JAR or .AAR package
* File > Project Structure > app -> Dependencies -> Add Module Dependency
* Add import: *import go.logpackermobilesdk.Logpackermobilesdk;*

#### How to use it in Java:

```java
import go.logpackermobilesdk.Logpackermobilesdk;
// ...
try {
    client = Logpackermobilesdk.NewClient("https://logpacker.mywebsite.com", "dev", android.os.Build.MODEL);

    msg = client.NewMessage();
    msg.setMessage("Crash is here!");
    // Use another optional setters for msg object

    client.Send(msg); // Send will return Cluster response
} catch (Exception e) {
    // Cannot connect to Cluster or validation error
}
```

#### How to send Android crashes to LogPacker

You must catch uncaughtException in your application and use LogPacker to send the exception:

```java
// ...
import go.logpackermobilesdk.Logpackermobilesdk;

public class MyApplication extends Application {
    public void onCreate () {
        // Setup handler for uncaught exceptions.
        Thread.setDefaultUncaughtExceptionHandler (new Thread.UncaughtExceptionHandler() {
            @Override
            public void uncaughtException (Thread thread, Throwable e) {
                // Paste here above code to send e.getMessage() to the LogPacker Cluster, use msg.setLogLevel(Logpackermobilesdk.FatalLogLevel)
            }
        });
    }
}
```

#### How to import into iOS project

Untar Logpackermobilesdk.framework.tar to the root of your project. Or drag Logpackermobilesdk.framework folder to the Xcode's file browser.

#### How to use in Xcode

```c
#import "Logpackermobilesdk/Logpackermobilesdk.h"
```

#### How to build an *.aar* or *.framework* packages from Go package

* golang 1.5+
* go get golang.org/x/mobile/cmd/gomobile
* gomobile init
* Install [Android SDK](https://developer.android.com/sdk/index.html#Other) to ~/android-sdk
* ~/android-sdk/tools/android sdk
* java-jdk
* export ANDROID_HOME=$HOME"/android-sdk" && gomobile bind --target=android .
* Find *.aar* file in working folder
* gomobile bind --target=ios .
* Find Logpackermobilesdk.framework folder

#### How to contribute

* Fork master branch
* Make changes
* Run ./before-commit.sh
* Push and create a Pull Request
