# Pre-requisite 
Make sure you are familiar with the following before taking this course:
- JavaScript 
- React
    - Just the fundamentals 
- JSX
- Thinking like a mobile developer 


# Overview 
## What is react native?
React Native is a cross-platform mobile development framework that lets you build apps for iOS, Android, and web, using JavaScript (mostly with React.js principles) while maintaining a single source code. Its core concepts differ from Go/Python/Bash—they're based on JavaScript's component model, state management, and often leverage JSX syntax.

---
#MindsetShift
- No css or special language for UI, it's done using pure JavaScript/react for styling. 
- The style names and values match css but are in js flavor. 
    - everything is defined in camel-case 
    - uses flexbox for layout 
    - Each component can have its own corresponding stylesheet. 
- React native provides platform-specific styling available for IOS and Android. (_maybe a warapper_)    

---
## React Native components
> components represent small pieces of code that are organised together. 

- Core components - more like standard liberary like golang 
- Community components - more like go package or python lib 
- Custom Native components - bestspoke that you build. 

# Installation 
1. Install Homebrew (Package Manager for macOS)
Homebrew makes it easy to install and manage software on macOS.
- Open Terminal and run:
```bash
$ /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```
After installing, run this to check if it installed correctly:
```bash
$ brew --version
```
2. Install Node.js and npm (Node Package Manager)
React Native requires Node.js, which comes with npm. You can install it via Homebrew or download it directly from Node.js website.
- Using Homebrew:
```bash
$ brew install node
```
- Verify the installation:
```bash
$ node --version
$ npm --version
```
3. Install Watchman (for macOS)
- Watchman is a tool developed by Facebook for watching changes in the filesystem. It’s useful for React Native.
- Install it using Homebrew:
```bash
$ brew install watchman
```
4. Install Xcode and Command Line Tools
React Native requires Xcode to run on iOS simulators and for iOS builds.
- Download and install Xcode from the Mac App Store.
- After installing Xcode, run the following command to install the necessary command line tools:
```bash
$ xcode-select --install
```
- Open Xcode once to accept the license and complete the setup.

5. Install React Native CLI (Optional)
You can choose between using Expo CLI or React Native CLI. If you want more control and flexibility, React Native CLI is the way to go.
- Install the React Native CLI globally using npm:
```bash
npm install react-native-cli
```
6. Create a New React Native Project
- If using Expo CLI (great for easier setup with less configuration):
```bash
npx create-expo-app@latest
```
7. Run the application 
```bash
$ cd my-app
- npm run android
- npm run ios
- npm run web
```

Summary of Tools:
1. Homebrew: Package manager for macOS.
2. Node.js: JavaScript runtime.
3. Watchman: Watches file changes.
4. Xcode: For iOS development.
5. Android Studio: For Android development.
6. React Native CLI: Command line interface for React Native projects.
7. Expo CLI (Optional): For a simpler React Native setup.

---