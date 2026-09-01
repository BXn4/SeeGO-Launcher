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
    import { Classes, Event, IDs, URLs } from "../../utils/consts";
    import { Icons } from "../../utils/icons";

    let languages: { value: string; label: string }[] = [];

    let theme = "";
    let language = "";
    let enableAnimationsValue = false;

    async function updateTheme(v: string) {
        if (theme != v) {
            theme = v;
            await Events.Emit(Event.App.updateSetting, ["theme", v]);
        }
    }

    async function updateLanguage(v: string) {
        await Events.Emit(Event.App.updateSetting, ["language", v]);

        const availableLanguages = await Config.GetLanguages();
        languages = await Promise.all(
            availableLanguages.map(async (value) => ({
                value: value,
                label: await Localization.Get(`language_${value}`, v),
            })),
        );

        await initLocalization();
    }

    async function updateAnims(v: boolean) {
        await Events.Emit(Event.App.updateSetting, ["anims", v]);
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
    });
</script>

<main>
    <div id={IDs.Views.Settings.ID}>
        <section class={Classes.Views.Settings.category}>
            <div class={Classes.Views.Settings.title}>
                {$locales[localization.settingsGeneral]}
            </div>

            <div class={Classes.Views.Settings.item}>
                <div class={Classes.Views.Settings.header}>
                    <span class={Classes.Views.Settings.name}
                        >{$locales[localization.settingLanguage]}</span
                    >
                    <span class={Classes.Views.Settings.desc}
                        >{$locales[localization.settingLanguageDesc]}</span
                    >
                </div>
                <select
                    id={IDs.Views.Settings.languageSetting}
                    class={Classes.Global.interactive}
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

            <div class={Classes.Views.Settings.item}>
                <div class={Classes.Views.Settings.header}>
                    <span class={Classes.Views.Settings.name}
                        >{$locales[localization.settingLaunchReady]}</span
                    >
                    <span class={Classes.Views.Settings.desc}
                        >{$locales[localization.settingLaunchReadyDesc]}</span
                    >
                </div>
                <label class={Classes.Views.Settings.switch}>
                    <input type="checkbox" />
                    <span
                        class="{Classes.Views.Settings.slider} {Classes.Views
                            .Settings.round} {Classes.Global.interactive} "
                    ></span>
                </label>
            </div>
            <div class={Classes.Views.Settings.item}>
                <div class={Classes.Views.Settings.header}>
                    <span class={Classes.Views.Settings.name}
                        >{$locales[localization.enableAnimations]}</span
                    >
                    <span class={Classes.Views.Settings.desc}
                        >{$locales[localization.enableAnimationsDesc]}</span
                    >
                </div>
                <label class={Classes.Views.Settings.switch}>
                    <input
                        type="checkbox"
                        bind:checked={enableAnimationsValue}
                        onchange={() => updateAnims(enableAnimationsValue)}
                    />
                    <span
                        class="{Classes.Views.Settings.slider} {Classes.Views
                            .Settings.round} {Classes.Global.interactive} "
                    ></span>
                </label>
            </div>
        </section>

        <section class={Classes.Views.Settings.category}>
            <div class={Classes.Views.Settings.title}>
                {$locales[localization.settingsAppearance]}
            </div>

            <div class={Classes.Views.Settings.item}>
                <div class={Classes.Views.Settings.header}>
                    <span class={Classes.Views.Settings.name}
                        >{$locales[localization.settingTheme]}</span
                    >
                    <span class={Classes.Views.Settings.desc}
                        >{$locales[localization.settingThemeDesc]}</span
                    >
                </div>
                <div role="group" aria-label="Theme">
                    <button
                        class="{Classes.Global.button} {Classes.Global
                            .interactive}"
                        title="{$locales[localization.themeDark]} {$locales[
                            localization.colors
                        ]}"
                        class:active={theme === "dark"}
                        onclick={() => updateTheme("dark")}
                        >{$locales[localization.themeDark]}</button
                    >
                    <button
                        class="{Classes.Global.button} {Classes.Global
                            .interactive}"
                        title="Dracula"
                        class:active={theme === "dracula"}
                        onclick={() => updateTheme("dracula")}>Dracula</button
                    >
                    <button
                        class="{Classes.Global.button} {Classes.Global
                            .interactive}"
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

        <section class={Classes.Views.Settings.category}>
            <div class={Classes.Views.Settings.title}>
                {$locales[localization.settingsInformation]}
            </div>

            <div class={Classes.Views.Settings.item}>
                <div class={Classes.Views.Settings.header}>
                    <span class={Classes.Views.Settings.name}
                        >{$locales[localization.settingVersion]}</span
                    >
                    <span class={Classes.Views.Settings.desc}
                        >SeeGO Launcher</span
                    >
                </div>
                <span
                    class="{Classes.Global.comment} {Classes.Views.Settings
                        .text}">1.0.0</span
                >
            </div>

            <div class={Classes.Views.Settings.item}>
                <div class={Classes.Views.Settings.header}>
                    <span class={Classes.Views.Settings.name}
                        >{$locales[localization.settingSGAccount]}</span
                    >
                    <span class={Classes.Views.Settings.desc}
                        >{$locales[localization.settingSGAccountDesc]}</span
                    >
                </div>
                <button
                    class="{Classes.Global.button} {Classes.Views.Settings
                        .text} {Classes.Global.interactive}"
                    >{$locales[localization.copy]}</button
                >
            </div>
        </section>
        <section class={Classes.Views.Settings.category}>
            <div class={Classes.Views.Settings.item}>
                <div class={Classes.Views.Settings.header}>
                    <span class={Classes.Views.Settings.name}
                        >SeeGo Launcher</span
                    >
                    <span class={Classes.Views.Settings.desc}
                        >by Fodor Bence</span
                    >
                </div>
                <div class={Classes.Views.Settings.iconGroup}>
                    <button
                        onclick={() => window._openURL(URLs.Repo)}
                        class="{Classes.Views.Settings.iconGroupButton} {Classes
                            .Global.interactive}"
                        title="GitHub Repo"
                    >
                        {@html Icons.GitHub}
                    </button>
                    <span class={Classes.Views.Settings.iconGroupDivider}
                    ></span>
                    <button
                        onclick={() => window._openURL(URLs.Support)}
                        class="{Classes.Views.Settings.iconGroupButton} {Classes
                            .Global.interactive}"
                        title={$locales[localization.sponsor]}
                    >
                        {@html Icons.Support}
                    </button>
                </div>
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
