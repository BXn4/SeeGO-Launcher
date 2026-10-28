<div align="center">
  <img src="https://i.imgur.com/FPArktl.png" alt="logo" width="120" height="120" />
  <img src="https://i.imgur.com/xqcd4sK.png" alt="logo" width="120" height="120" />
  <h1>SeeGO Launcher</h1>
  <p>An alternative open-source launcher for the SeeRPG server that uses WebView instead of Electron.</p>
    <a href="https://go.dev">
      <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=white" alt="Golang" /></a>
    <a href="https://v3.wails.io/">
      <img src="https://img.shields.io/badge/Wails-DF0000?style=for-the-badge&logo=Wails&logoColor=white" alt="Wails" /></a>
  </p>
</div>

# Contents
- [Information](#information)
- [Development](#development)
  * [Running locally](#running-locally)

<div align="center">
  <img src="https://i.imgur.com/rRcbSy9.png" alt="screenshot" />
</div>

<div align="center">
  <img src="https://i.imgur.com/1UGOknM.png" alt="screenshot" />
</div>


## Information
This project was created as an independent development and does not use the source code, modify, distribute, decompile, reverse enginner, or include any source from the official SeeRPG launcher.
This software is not a modified version of the official launcher, and the purpose of this project is to ensure compatibility and improve the user experience for the community.

Difference from the original:
| Aspect | SeeGO | Official Launcher |
|---|---|---|
| Frontend | Svelte | HTML, CSS, JavaScript |
| Backend | Go (compiled binary) | Node.js (JavaScript runtime) |
| Framework | Wails | Electron |
|Browser | System Webview | Chromium |

**Webview by the OS:**

| OS | WebView |
|--|--|
| Windows 10/11 | WebView2 (Edge) |
| macOS 12+ | WKWebView (Safari) |
| Linux (GTK) | WebKitGTK |


### Development
**Depends on:**
- [Go](https://go.dev/dl/) `>= 1.26`
- [Node.js](https://nodejs.org/) `>= 26.2`
- [Wails](https://v3.wails.io/quick-start/installation/) `v3`
```bash
go install github.com/wailsapp/wails/v3/cmd/wails@latest
```

### Running locally
**1. Clone this repo**

```bash
git clone https://github.com/bxn4/seego-launcher.git
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
Before you build it, please place the Tebex account in the .env TEBEXACCOUNT varible. You can get this from the website store http request.
Starts with: "11osk"
```bash
./build.sh
```

The compiled binary is in the `./bin/` folder.
