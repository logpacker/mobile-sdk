[![Build Status](https://travis-ci.org/logpacker/mobile-sdk.svg?branch=master)](https://travis-ci.org/logpacker/mobile-sdk)
[![Gitter](https://badges.gitter.im/logpacker/mobile-sdk.svg)](https://gitter.im/logpacker/mobile-sdk?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=body_badge)

Repository contains SDK for Android, iOS and Windows Phone. Android and iOS SDKs are build with help of shared code and gomobile. Windows Phone SDK is located in **wp/** folder and written in C#, it's a portable Class library.

Godoc - https://godoc.org/github.com/logpacker/mobile-sdk

#### How to import into Android Studio (see *screenshots/* folder)

* File > New > New Module > Import .JAR or .AAR package
* File > Project Structure > app -> Dependencies -> Add Module Dependency
* Add import: *import go.logpackermobilesdk.Logpackermobilesdk;*

#### How to use it in Java:

```java
import go.logpackermobilesdk.Logpackermobilesdk;

// It's possible to catch all app's crashes via Thread.setDefaultUncaughtExceptionHandler and send it to LogPacker
try {
    Client client = Logpackermobilesdk.newClient("https://logpacker.mywebsite.com", "dev", android.os.Build.MODEL);
    client.setCloudKey("");

    Message msg = client.newMessage();
    msg.setMessage("Crash is here!");
    // Use another optional setters for msg object

    client.send(msg); // Send will return Cluster response
} catch (Exception e) {
    // Cannot connect to Cluster or validation error
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
    client.cloudKey = @"";
    GoLogpackermobilesdkMessage *msg;
    msg = client.newMessage;
    msg.message = @"Crash is here!";
    // Use another optional setters for msg object
    GoLogpackermobilesdkResult *result;
    [client send:(msg) ret0_:(&result) error:(&error)];
}

// It's possible to catch all app's crashes via signal(SIGSEGV, SignalHandler) and send it to LogPacker from SignalHandler func
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
				Client c = new Client ("https://logpacker.mywebsite.com", "dev", System.Environment.MachineName, "");
				Event e = new Event ("Crash is here!", "modulename", Event.FatalLogLevel, "1000", "John");
				c.Send (e);
			} catch {
				// Handle connection error here
			}
		}
	}
}

// It's possible to catch all app's crashes via global try-catch block and send it to LogPacker
```

#### How to build an *.aar* or *.framework* packages from Go package

* golang 1.7+
* `go get golang.org/x/mobile/cmd/gomobile`
* `gomobile init`
* Install [Android SDK](https://developer.android.com/sdk/index.html#Other) to ~/android-sdk
* ~/android-sdk/tools/android sdk
* Install `java-jdk`
* `export ANDROID_HOME=$HOME"/android-sdk"`
* gomobile bind --target=android .
* Find *.aar* file in working folder
* Install XCode
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
