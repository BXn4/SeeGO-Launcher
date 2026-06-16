# SeeGO Launcher

<img src="https://i.imgur.com/77sR4yP.png" align="right"
     alt="SeeGO Launcher logo by Fodor Bence" width="120" height="178">

Egy alternatív nyílt forráskódú (open source) launcher a SeeRPG szerverhez, ami electron helyett webviewet használ.
A launcher Go nyelven, a Wails framework felhasználásával készült.

Key differences from Electron:
| Aspect   |      Wails      |  Electron |
|----------|:-------------:|------:|
| Browser | OS-provided WebView| Bundled Chromium (~100MB) |
| Backend | Go (compiled) | Node.js (interpreted) |
| Communication | In-memory bridge | IPC (inter-process) |
| Bundle Size | ~15MB | ~150MB |
| Memory | ~10MB | ~100MB+ |
| Startup | <0.5s | 2-3s |

A SeeGO jelenleg a háttérben 8mb memóriát kér. Ezen még lehet karcsúsítani ;), illetve az egész app csupán ~13mb, de ezen is lehet még karcsúsítani ;)

Hogyan működik?
Minden egyes indításnál ellenőrzi, hogy a helyi (local) filek frissek-e (md5 hashel validálás) a távoli (remote helyről). Ha eltér, vagy nem létezik, akkor lekéri (reuqest-eli) a távoli helyről.

Mivel egy két funkciót nem tehettem be az appba, ezért a fórum, illetve a(z) UCP csak a böngészőböl érhető el.
De minden más a launcherben megtalálható. MÉG A BOLT IS!!

A projekt független fejlesztésként jött létre, és nem használja fel a SeeRPG hivatalos kliensének forráskódját.
A szoftver nem módosított változata a hivatalos launchernek, hanem egy önálló implementáció, amely kizárólag a SeeRPG szolgáltatásaival való kompatibilitást célozza.

Ez a projekt nem hivatalos. Nem áll kapcsolatban a SEE-ONLINE Kft.-vel vagy a SeeRPG üzemeltetőivel, és nem rendelkezik azok támogatásával, jóváhagyásával vagy hivatalos partnerségével.
A projekt nem tartalmazza, nem terjeszti és nem használja fel a SeeRPG hivatalos kliensének, launcherének vagy egyéb szoftvereinek forráskódját.

A projekt célja kizárólag a kompatibilitás biztosítása és a felhasználói élmény javítása a SeeRPG közösség számára, mivel a launcher csak az anti-cheat miatt készült el, így gondoltam, hogy a launcherbe több hasznos funkciót teszek be.


<img src="https://i.imgur.com/Iim5Dk9.png">
