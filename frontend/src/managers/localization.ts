import { writable } from "svelte/store"
import { Config, Localization } from "../../bindings/seegolauncher/internal/services";

export const localization = {
  lang: "lang",
  languageEn: "language_en",
  languageHu: "language_hu",

  launcherMinimized: "notification-launcher-minimized",

  splashLoading: "splash-loading",
  splashLoadingNews: "splash-loading-news",
  splashLoadingSerial: "splash-loading-serial",
  openedInBrowserWindow: "opened-browser-window",

  termsDeclinedTitle: "terms-declined-dialog-title",
  termsDeclinedContent: "terms-declined-dialog-content",
  accept: "accept",
  decline: "decline",

  seeGOInfo1: "seego-info1",
  seeGOInfo2: "seego-info2",
  seeGOInfo3: "seego-info3",
  seeGOInfo4: "seego-info4",

  windowMinimize: "window-minimize",
  windowClose: "window-close",

  playTitleReady: "play-title-ready",
  playTitleOffline: "play-title-offline",

  homeTitle: "home-title",
  newsTitle: "news-title",
  forumTitle: "forum-title",
  ucpTitle: "ucp-title",
  shopTitle: "shop-title",
  galleryTitle: "gallery-title",
  playersTitle: "players-title",
  helpTitle: "help-title",
  settingsTitle: "settings-title",

  newsLatest: "news-latest",
  newsRead: "news-read",

  clubMembership: "club-membership",
  clubMembershipDay: "club-membership-day",

  serverStatusOnline: "server-status-online",
  serverStatusOffline: "server-status-offline",
  serverStatusRestart: "server-status-restart",
  serverStatusMaintenance: "server-status-maintenance",
  serverStatusAdmins: "server-status-admins",
  serverStatusQueue: "server-status-queue",
  serverStatusEstimated: "server-status-estimated",

  hours: "hours",
  minutes: "minutes",
  seconds: "seconds",

  community: "community",
  launcherReady: "launcher-ready",
  launcherConnect: "launcher-connect",

  settingsGeneral: "settings-general",
  settingLanguage: "setting-language",
  settingLanguageDesc: "setting-language-desc",
  settingLaunchReady: "setting-launch-ready",
  settingLaunchReadyDesc: "setting-launch-ready-desc",

  settingsAppearance: "settings-appearance",
  settingTheme: "setting-theme",
  settingThemeDesc: "setting-theme-desc",
  themeDark: "theme-dark",
  themeLight: "theme-light",
  colors: "colors",

  settingsInformation: "settings-information",
  settingVersion: "setting-version",
  settingSGAccount: "setting-sg-account",
  settingSGAccountDesc: "setting-sg-account-desc",
  copy: "copy",

  enableAnimations: "enable-animations",
  enableAnimationsDesc: "enable-animations-desc",

  NewsLoadFailed: "news-load-failed",
  NewsLoadFailedDesc: "news-load-failed-desc",
  Retry: "retry"
} as const;

export const locales = writable<Record<string, string>>({});

export async function initLocalization() {
  const lang = await Config.GetLanguage()
  const keys = Object.values(localization);
  const values = await Promise.all(keys.map((k) => Localization.Get(k, lang)));
  const map: Record<string, string> = {};
  keys.forEach((k, i) => (map[k] = values[i]));
  locales.set(map);
}
