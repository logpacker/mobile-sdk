[![Build Status](https://travis-ci.org/logpacker/mobile-sdk.svg?branch=master)](https://travis-ci.org/logpacker/mobile-sdk)

Repository contains SDK for Android, iOS and Windows Phone. Android and iOS SDKs are build with help of shared code and gomobile. Windows Phone SDK is located in **wp/** folder and written in Mono C#, it's a portable Class library.

Godoc - https://godoc.org/github.com/logpacker/mobile-sdk

#### How to import into Android Studio (see *screenshots/* folder)

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

#### How to import framework into Xcode (see *screenshots/* folder)

 * Drag *Logpacker.framework* folder into your Xcode's browser
 * Use import *#import "Logpackermobilesdk/Logpackermobilesdk.h"*

#### How to use it in Xcode

```c
#import "ViewController.h"
#import "Logpackermobilesdk/Logpackermobilesdk.h"

@interface ViewController ()

@end

@implementation ViewController

- (void)viewDidLoad {
    [super viewDidLoad];
    GoLogpackermobilesdkClient *client;
    NSError *error;
    GoLogpackermobilesdkNewClient(@"https://logpacker.mywebsite.com", @"dev", [[UIDevice currentDevice] systemVersion], &client, &error);
    GoLogpackermobilesdkMessage *msg;
    msg = client.newMessage;
    msg.message = @"Crash is here!";
    // Use another optional setters for msg object
    GoLogpackermobilesdkResult *result;
    [client send:(msg) ret0_:(&result) error:(&error)];
}
```

#### How to send iOS crashes to LogPacker

You must catch Exceptions and Signals and use LogPacker to send them:

```c
void InstallUncaughtExceptionHandler()
{
    NSSetUncaughtExceptionHandler(&HandleException);
    signal(SIGABRT, SignalHandler);
    signal(SIGILL, SignalHandler);
    signal(SIGSEGV, SignalHandler);
    signal(SIGFPE, SignalHandler);
    signal(SIGBUS, SignalHandler);
    signal(SIGPIPE, SignalHandler);
}

void HandleException(NSException *exception) {
    // Paste here above code to send [exception reason] to the LogPacker Cluster, use msg.logLevel = GoLogpackermobilesdk.fatalLogLevel
}

static void SignalHandler(int signo) {
    // The same
}
```

#### How to import into C# project

 * Add *logpackermobilesdk.dll* into your C# project
 * Add *using logpackermobilesdk;* before you start to use it

#### How to use it in C# code

```cs
using System;
using logpackermobilesdk;

namespace test
{
    class MainClass
	{
		public static void Main (string[] args)
		{
			try {
				Client c = new Client ("https://logpacker.mywebsite.com", "dev", System.Environment.MachineName);
				Event e = new Event ("Crash is here!", "modulename", Event.FatalLogLevel, "1000", "John");
				c.Send (e);
			} catch {
				// Handle connection error here
			}
		}
	}
}
```

#### How to send C# exception to LogPacker

```cs
using System;
using logpackermobilesdk;

namespace test
{
    class MainClass
	{
		public static void Main (string[] args)
		{
			try {
				// Exception thrown here
			} catch (Exception e) {
				Client c = new Client ("https://logpacker.mywebsite.com", "dev", System.Environment.MachineName);
				c.SendException (e);
			}
		}
	}
}
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

#### How to build CS library

 * MonoDevelop
 * Open project (wp folder)
 * Project -> Export

#### How to contribute

* Fork master branch
* Make changes
* Run ./before-commit.sh
* Push and create a Pull Request
