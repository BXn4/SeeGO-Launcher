<script lang="ts">
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";
    import {
        Config,
        Localization,
    } from "../../../bindings/seegolauncher/internal/services";

    let theme = "";
    let language = "";
    let settingsGeneral = "";
    let settingLanguage = "";
    let settingLanguageDesc = "";
    let settingLaunchReady = "";
    let settingLaunchReadyDesc = "";
    let settingsApppearance = "";
    let settingTheme = "";
    let settingThemeDesc = "";
    let themeDark = "";
    let themeLight = "";
    let colors = "";
    let settingsInformation = "";
    let settingVersion = "";
    let settingSGAccount = "";
    let settingSGAccountDesc = "";
    let copy = "";

    let languages: { value: string; label: string }[] = [];

    async function setLocales() {
        let lang = language;
        [
            settingsGeneral,
            settingLanguage,
            settingLanguageDesc,
            settingLaunchReady,
            settingLaunchReadyDesc,
            settingsApppearance,
            settingTheme,
            settingThemeDesc,
            themeDark,
            themeLight,
            colors,
            settingsInformation,
            settingVersion,
            settingSGAccount,
            settingSGAccountDesc,
            copy,
        ] = await Promise.all([
            Localization.Get("settings-general", lang),
            Localization.Get("setting-language", lang),
            Localization.Get("setting-language-desc", lang),
            Localization.Get("setting-launch-ready", lang),
            Localization.Get("setting-launch-ready-desc", lang),
            Localization.Get("settings-appearance", lang),
            Localization.Get("setting-theme", lang),
            Localization.Get("setting-theme-desc", lang),
            Localization.Get("theme-dark", lang),
            Localization.Get("theme-light", lang),
            Localization.Get("colors", lang),
            Localization.Get("settings-information", lang),
            Localization.Get("setting-version", lang),
            Localization.Get("setting-sg-account", lang),
            Localization.Get("setting-sg-account-desc", lang),
            Localization.Get("copy", lang),
        ]);
    }

    async function updateTheme(v: string) {
        theme = v;
        await Events.Emit("app:updateTheme", v);
    }

    async function updateLanguage() {
        await Events.Emit("app:updateLanguage", language);

        const availableLanguages = await Config.GetLanguages();
        languages = await Promise.all(
            availableLanguages.map(async (value) => ({
                value: value,
                label: await Localization.Get(`language_${value}`, language),
            })),
        );

        setLocales();
    }

    onMount(async () => {
        theme = await Config.GetTheme();
        language = await Config.GetLanguage();
        const availableLanguages = await Config.GetLanguages();
        languages = await Promise.all(
            availableLanguages.map(async (value) => ({
                value: value,
                label: await Localization.Get(`language_${value}`, language),
            })),
        );

        setLocales();
    });
</script>

<main>
    <div class="settings-view">
        <section class="settings-category">
            <div class="settings-title">{settingsGeneral}</div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name">{settingLanguage}</span>
                    <span class="setting-desc">{settingLanguageDesc}</span>
                </div>
                <select
                    id="setting-language"
                    bind:value={language}
                    onchange={() => updateLanguage()}
                >
                    {#each languages as lang}
                        <option
                            selected={lang.value == language}
                            value={lang.value}>{lang.label}</option
                        >
                    {/each}
                </select>
            </div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name">{settingLaunchReady}</span>
                    <span class="setting-desc">{settingLaunchReadyDesc}</span>
                </div>
                <label class="switch">
                    <input type="checkbox" />
                    <span class="slider round"></span>
                </label>
            </div>
        </section>

        <section class="settings-category">
            <div class="settings-title">{settingsApppearance}</div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name">{settingTheme}</span>
                    <span class="setting-desc">{settingThemeDesc}</span>
                </div>
                <div role="group" aria-label="Theme">
                    <button
                        class="button"
                        title="{themeDark} {colors}"
                        class:active={theme === "dark"}
                        onclick={() => updateTheme("dark")}>{themeDark}</button
                    >
                    <button
                        class="button"
                        title="{themeLight} {colors}"
                        class:active={theme === "light"}
                        onclick={() => updateTheme("light")}
                        >{themeLight}</button
                    >
                </div>
            </div>
        </section>

        <section class="settings-category">
            <div class="settings-title">{settingsInformation}</div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name">{settingVersion}</span>
                    <span class="setting-desc">SeeGO Launcher</span>
                </div>
                <span class="comment setting-text">1.0.0</span>
            </div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name">{settingSGAccount}</span>
                    <span class="setting-desc">{settingSGAccountDesc}</span>
                </div>
                <span class="comment setting-text">{copy}</span>
            </div>
        </section>
    </div>
</main>

<style>
    @import "../../public/styles/views/settings.css";
    main {
        width: 100vw;
        height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
    }
</style>
