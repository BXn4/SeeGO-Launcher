<script lang="ts">
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";
    import {
        Config,
        Localization,
    } from "../../../bindings/seegolauncher/internal/services";
    import {
        initLocalization,
        locales,
        localization,
    } from "../../managers/localization";

    let languages: { value: string; label: string }[] = [];

    let theme = "";
    let language = "";
    let enableAnimationsValue = false;

    async function updateTheme(v: string) {
        if (theme != v) {
            theme = v;
            await Events.Emit("app:updateSetting", ["theme", v]);
        }
    }

    async function updateLanguage(v: string) {
        await Events.Emit("app:updateSetting", ["language", v]);

        const availableLanguages = await Config.GetLanguages();
        languages = await Promise.all(
            availableLanguages.map(async (value) => ({
                value: value,
                label: await Localization.Get(`language_${value}`, v),
            })),
        );

        initLocalization();
    }

    async function updateAnims(v: boolean) {
        await Events.Emit("app:updateSetting", ["anims", v]);
    }

    onMount(async () => {
        theme = await Config.GetTheme();
        language = await Config.GetLanguage();
        enableAnimationsValue = await Config.GetEnableAnimations();

        const availableLanguages = await Config.GetLanguages();
        languages = await Promise.all(
            availableLanguages.map(async (value) => ({
                value: value,
                label: await Localization.Get(`language_${value}`, language),
            })),
        );

        initLocalization();
    });
</script>

<main>
    <div class="settings-view">
        <section class="settings-category">
            <div class="settings-title">
                {$locales[localization.settingsGeneral]}
            </div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name"
                        >{$locales[localization.settingLanguage]}</span
                    >
                    <span class="setting-desc"
                        >{$locales[localization.settingLanguageDesc]}</span
                    >
                </div>
                <select
                    id="setting-language"
                    class="interactive"
                    bind:value={language}
                    onchange={() => updateLanguage(language)}
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
                    <span class="setting-name"
                        >{$locales[localization.settingLaunchReady]}</span
                    >
                    <span class="setting-desc"
                        >{$locales[localization.settingLaunchReadyDesc]}</span
                    >
                </div>
                <label class="switch">
                    <input type="checkbox" />
                    <span class="slider round interactive"></span>
                </label>
            </div>
            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name"
                        >{$locales[localization.enableAnimations]}</span
                    >
                    <span class="setting-desc"
                        >{$locales[localization.enableAnimationsDesc]}</span
                    >
                </div>
                <label class="switch">
                    <input
                        type="checkbox"
                        bind:checked={enableAnimationsValue}
                        onchange={() => updateAnims(enableAnimationsValue)}
                    />
                    <span class="slider round interactive"></span>
                </label>
            </div>
        </section>

        <section class="settings-category">
            <div class="settings-title">
                {$locales[localization.settingsAppearance]}
            </div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name"
                        >{$locales[localization.settingTheme]}</span
                    >
                    <span class="setting-desc"
                        >{$locales[localization.settingThemeDesc]}</span
                    >
                </div>
                <div role="group" aria-label="Theme">
                    <button
                        class="button interactive"
                        title="{$locales[localization.themeDark]} {$locales[
                            localization.colors
                        ]}"
                        class:active={theme === "dark"}
                        onclick={() => updateTheme("dark")}
                        >{$locales[localization.themeDark]}</button
                    >
                    <button
                        class="button interactive"
                        title="{$locales[localization.themeLight]} {$locales[
                            localization.colors
                        ]}"
                        class:active={theme === "light"}
                        onclick={() => updateTheme("light")}
                        >{$locales[localization.themeLight]}</button
                    >
                </div>
            </div>
        </section>

        <section class="settings-category">
            <div class="settings-title">
                {$locales[localization.settingsInformation]}
            </div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name"
                        >{$locales[localization.settingVersion]}</span
                    >
                    <span class="setting-desc">SeeGO Launcher</span>
                </div>
                <span class="comment setting-text">1.0.0</span>
            </div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name"
                        >{$locales[localization.settingSGAccount]}</span
                    >
                    <span class="setting-desc"
                        >{$locales[localization.settingSGAccountDesc]}</span
                    >
                </div>
                <button class="button setting-text interactive"
                    >{$locales[localization.copy]}</button
                >
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
