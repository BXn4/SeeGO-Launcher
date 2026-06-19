
<div align="center">
    <img src="https://i.imgur.com/xqcd4sK.png" alt="logo" width="120" height="120"/>
<h1>SeeGO Launcher</h1>
</div>
<p>
An alternative open-source launcher for the SeeRPG server that uses WebView instead of Electron.
The launcher was developed in Go using the Wails framework.
</p>

# Contents
- [Information](#information)
- [Development](#development)
  * [Running at locally](#running-at-locally)

<div align="center">
  <img src="https://i.imgur.com/Iim5Dk9.png" alt="screenshot" />
</div>

## Information
<p>
This project was created as an independent development and does not use the source code of the official SeeRPG launcher.
This software is not a modified version of the official launcher, but rather a standalone implementation designed exclusively for compatibility with SeeRPG's services.</p>
<p>This project is unofficial. It is not affiliated with SEE-ONLINE Kft. or the operators of SeeRPG, and does not have their support, approval, or partnership.
This project does not include, distribute, or use the source code of SeeRPG's official client, launcher, or other software.</p>
<p>The purpose of this project is to ensure compatibility and improve the user experience for the SeeRPG community. Since the launcher was originally created for anti-cheat purposes, I decided to add several useful features to it, to make it more usefull.</p>

Key differences from Electron:
| Aspect   |      Wails      |  Electron |
|----------|:-------------:|------:|
| Browser | OS-provided WebView| Bundled Chromium (~100MB) |
| Backend | Go (compiled) | Node.js (interpreted) |
| Communication | In-memory bridge | IPC (inter-process) |
| Bundle Size | ~15MB | ~150MB |
| Memory | ~10MB | ~100MB+ |
| Startup | <0.5s | 2-3s |


### Development
**Depends on:**:
- [Go](https://go.dev/dl/) `>= 1.26`
- [Node.js](https://nodejs.org/) `>= 26.2`
- [Wails](https://v3.wails.io/quick-start/installation/) `v3`
```bash
go install github.com/wailsapp/wails/v3/cmd/wails@latest
```

**Operating System:**

| OS | WebView |
|--|--|--|
| Windows 10/11 | WebView2 (Edge) |
| macOS 12+ | WKWebView (Safari) |
| Linux (GTK) | WebKitGTK |


### Running at locally
**1. Clone this repo**

```bash
git clone https://gitea.com/bxn4/seego-launcher.git
cd seego-launcher
```

**2. Install dependencies**

```bash
go install
```

**3. Running it**

```bash
wails3 dev
```

**4. Building**
Before you build it, please place the Tebex account in the .env OA varible. You can get this from the website store request.
```bash
./build.sh // only for WINDOWS!
```

The compiled binary is in the `./bin/` folder.
