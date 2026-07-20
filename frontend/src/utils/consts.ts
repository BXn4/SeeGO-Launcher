export const State = {
  currentAppView: "splash",
  currentMainView: "home",
  currentNavbarActive: "home"
};

export const Event = {
  App: {
    ready: "app:ready",
    domReady: "app:domReady",
    navigate: "app:navigate",
    close: "app:close",
    minimize: "app:minimize",
    notActive: "app:notActive",
    active: "app:active",
    updateSetting: "app:updateSetting",
  },
  Splash: {
    setCurrentProgress: "splash:setCurrentProgress",
    openedBrowserWindow: "opened-browser-window",
  },
  Terms: {
    accept: "terms:accept",
    decline: "terms:decline"
  },
  Main: {
    navigate: "main:navigate",
    Navbar: {
      switchNavTab: "navbar:switchNavTab"
    },
    News: {
      readLatest: "news:readLatest"
    }
  },
  Global: {
    startInterval: "startInterval",
    stopInterval: "stoptInterval",
    newsFeedUpdated: "newsFeedUpdated",
    feedback: "feedback"
  }
}

export const View = {
  splash: "splash",
  terms: "terms",
  main: "main",
  home: "home",
  news: "news",
  shop: "shop",
  gallery: "gallery",
  help: "help",
  settings: "settings",
}
