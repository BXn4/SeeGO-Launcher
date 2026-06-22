<script lang="ts">
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";
    import {
        Config,
        Localization,
    } from "../../../bindings/seegolauncher/internal/services";

    let theme = "";

    const languages = [
        { value: "en", label: "English" },
        { value: "hu", label: "Hungarian" },
    ];

    let language = Config.GetLanguage;

    async function updateTheme(v: string) {
        theme = v;
        await Events.Emit("app:updateTheme", v);
    }

    onMount(async () => {
        theme = await Config.GetTheme();
    });
</script>

<main>
    <div class="settings-view">
        <section class="settings-category">
            <div class="settings-title">General</div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name">Language</span>
                    <span class="setting-desc">Interface language</span>
                </div>
                <select bind:value={language}>
                    {#each languages as lang}
                        <option value={lang.value}>{lang.label}</option>
                    {/each}
                </select>
            </div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name">Launch on ready</span>
                    <span class="setting-desc"
                        >Launch and connects automatically when the client is
                        ready</span
                    >
                </div>
                <label class="switch">
                    <input type="checkbox" />
                    <span class="slider round"></span>
                </label>
            </div>
        </section>

        <section class="settings-category">
            <div class="settings-title">Appearance</div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name">Theme</span>
                    <span class="setting-desc">Color scheme preference</span>
                </div>
                <div role="group" aria-label="Theme">
                    <button
                        class="button"
                        title="Dark"
                        class:active={theme === "dark"}
                        onclick={() => updateTheme("dark")}>Dark</button
                    >
                    <button
                        class="button"
                        title="Light"
                        class:active={theme === "light"}
                        onclick={() => updateTheme("light")}>Light</button
                    >
                </div>
            </div>
        </section>

        <section class="settings-category">
            <div class="settings-title">Information</div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name">Version</span>
                    <span class="setting-desc">SeeGO Launcher</span>
                </div>
                <span class="comment setting-text">v1.0.0</span>
            </div>

            <div class="setting-item">
                <div class="setting-info">
                    <span class="setting-name">SG Account</span>
                </div>
                <span class="comment setting-text">0000 COPY</span>
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
