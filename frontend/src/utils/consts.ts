import { writable } from "svelte/store";

export const URLs = {
  Forum: "https://forum.see-rpg.com/",
  UCP: "https://ucp.see-rpg.com/",
  Discord: "https://discord.com/invite/seerpg",
  Facebook: "https://www.facebook.com/seerpgofficial",
  Tiktok: "https://www.tiktok.com/@seerpg",
  Youtube: "https://www.youtube.com/@seerpgofficial",
  Repo: "https://github.com/BXn4/SeeGO-Launcher",
  Support: "https://github.com/sponsors/BXn4"
}

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
    infoFeedback: "infoFeedback",
    debugFeedback: "debugFeedback",
    updateLanguage: "updateLanguage"
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

export const IDs = {
  Partials: {
    Navbar: {
      ID: "navbar"
    },
    Titlebar: {
      ID: "titlebar",
      windowTitle: "window-title",
      windowControls: "window-controls",
      minimizeButton: "min-button",
      closeButton: "close-button",
      controls: "window-controls",
      title: "window-title",
    },
  },
  Views: {
    Error: {
      ID: "error-view",
    },
    Splash: {
      ID: "splash-view",
      logo: "splash-logo",
      alt: "splash-alt",
    },
    Home: {
      ID: "home-view",
      card: "hero-card",
      title: "hero-news-title",
      alt: "hero-news-alt",
      read: "hero-news-read-latest",
      container: "items-container",
      status: "server-status",
      count: "players-count",
      fill: "server-fill",
      adminsCount: "admins-count",
      queueCount: "queue-count",
      estimatedConnect: "estimated-connect",
      connectButton: "connect-button"
    },
    Terms: {
      ID: "terms-view",
      content: "terms-content",
      actionbar: "terms-actionbar"
    },
    News: {
      ID: "news-view",
    },
    Settings: {
      ID: "settings-view",
      languageSetting: "setting-language",
    }
  },
  View: {
    ...View,
    forum: "forum",
    ucp: "ucp",
  }
}

export const Classes = {
  Partials: {
    Navbar: {
      Class: "navbar",
      item: "navbar-item",
      button: "navbar-button",
      indicator: "indicator",
      divider: "divider"
    },
    Titlebar: {
      dim1: "dim1",
      dim2: "dim2",
      windowControl: "window-control"
    }
  },
  Views: {
    Error: {
      title: "error-title",
      alt: "error-alt"
    },
    Splash: {
      title: "splash-title",
      alt: "splash-alt",
      loader: "loader"
    },
    Home: {
      layout: "feed-layout",
      card: "hero-card",
      overlay: "hero-overlay",
      content: "hero-content",
      read: "news-read",
      items: "home-items",
      itemsTitle: "items-title",
      sidebar: "sidebar",
      widget: "widget",
      widgetTitle: "widget-title",
      container: "status-container",
      info: "status-info",
      text: "status-text",
      offline: "offline",
      count: "player-count",
      online: "online",
      stats: "stats",
      statBox: "stat-box",
      statTitle: "stat-title",
      statValue: "stat-value",
      estimatedConnect: "estimated-connect",
      socialWidget: "social-links-grid",
      connect: "connect",
    },
    Terms: {
      title: "terms-title",
      alt: "terms-alt"
    },
    News: {
      layout: "news-layout",
      card: "news-card",
      overlay: "news-overlay",
      content: "news-content",
      popupOverlay: "news-popup-overlay",
      popup: "news-popup",
      popupImage: "news-popup-image",
      popupHeader: "news-popup-header",
      popupContent: "news-popup-content"
    },
    Settings: {
      category: "settings-category",
      title: "settings-title",
      item: "setting-item",
      header: "setting-info",
      name: "setting-name",
      desc: "setting-desc",
      switch: "switch",
      slider: "slider",
      round: "round",
      text: "setting-text",
      iconGroup: "icon-group",
      iconGroupButton: "icon-group-button",
      iconGroupDivider: "icon-group-divider"
    }
  },
  Global: {
    top: "top",
    bottom: "bottom",
    interactive: "interactive",
    altIcon: "alt-icon",
    text: "text",
    comment: "comment",
    button: "button",
    bar: "progress-bar",
    barFill: "progress-bar-fill",
    icon: "icon",
    cancel: "cancel",
    ok: "ok",
    newsBadge: "news-badge",
    newsTitle: "news-title",
    newsAlt: "news-alt",
    old: "old",
    newsRead: "news-read",
  }
}

export const currentAppView = writable(View.splash);
export const currentMainView = writable(View.home);
export const currentNavbarActive = writable(View.home);
export const isTermsView = writable(false);
