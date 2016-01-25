### How to import into Java project

Android Studio:

* File > New > New Module > Import .JAR or .AAR package
* File > Project Structure > Dependencies
* Access the Go library from anywhere simply by importing *go.android-client.logpackerandroid*

### How to build an *.aar* package from Go package

* golang 1.5+
* go get golang.org/x/mobile/cmd/gomobile
* gomobile init
* Install [Android SDK](https://developer.android.com/sdk/index.html#Other) to ~/android-sdk-linux
* Install java-jdk
* ANDROID_HOME="/home/username/android-sdk-linux" gomobile bind .
