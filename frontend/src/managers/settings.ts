import { writable, get } from "svelte/store";
import { Events } from "@wailsio/runtime";
import { Event } from "../utils/consts";
import {
  Config,
  Localization,
} from "../../bindings/seegolauncher/internal/services";
import {
  initLocalization,
} from "../managers/localization";

export const languages = writable<{ value: string; label: string }[]>([]);

export let theme = writable<string>("")
export let language = writable<string>("")
export let enableAnimations = writable<boolean>(false);

export async function initSettings() {
  theme.set(await Config.GetTheme());
  language.set(await Config.GetLanguage());
  enableAnimations.set(await Config.GetEnableAnimations());

  const availableLanguages = await Config.GetLanguages();
  languages.set(await Promise.all(
    availableLanguages.map(async (value) => ({
      value: value,
      label: await Localization.Get(`language_${value}`, get(language)),
    })),
  ));
}

export async function updateTheme(v: string) {
  if (get(theme) != v) {
    theme.set(v);
    await Events.Emit(Event.App.updateSetting, ["theme", v]);
  }
}

export async function updateLanguage(v: string) {
  await Events.Emit(Event.App.updateSetting, ["language", v]);
  const availableLanguages = await Config.GetLanguages();
  languages.set(await Promise.all(
    availableLanguages.map(async (value) => ({
      value: value,
      label: await Localization.Get(`language_${value}`, v),
    })),
  ));

  await initLocalization();
}

export async function updateAnims(v: boolean) {
  await Events.Emit(Event.App.updateSetting, ["anims", v]);
}
