<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import {
        Config,
        Localization,
    } from "../../../bindings/seegolauncher/internal/services";
    import { onMount } from "svelte";

    async function closeWindow() {
        await Events.Emit("close", null);
    }
    async function minimizeWindow() {
        await Events.Emit("minimize", null);
    }
    async function toggleMaximize() {
        await Events.Emit("toggle-maximize", null);
    }

    onMount(() => {
        setLocales();
    });

    async function setLocales() {
        const lang = await Config.GetLanguage();
        const minButton = document.getElementById("min-button");
        const closeButton = document.getElementById("close-button");

        if (minButton) {
            minButton.setAttribute(
                "title",
                await Localization.Get("window-minimize", lang),
            );
        }

        if (closeButton) {
            closeButton.setAttribute(
                "title",
                await Localization.Get("window-close", lang),
            );
        }
    }
</script>

<main>
    <div
        id="titlebar"
        role="button"
        tabindex="0"
        ondblclick={() => toggleMaximize()}
    >
        <div id="drag-region">
            <div id="titlebar-hover-area"></div>
            <div id="window-title">
                <span class="dim1">See</span><span class="dim2">Go</span>
            </div>
            <div id="window-controls">
                <span
                    class="window-control"
                    id="min-button"
                    title=""
                    role="button"
                    tabindex="0"
                    onclick={() => minimizeWindow()}
                    onkeydown={(e) => e.key === "Enter" && minimizeWindow()}
                >
                    <span>−</span>
                </span>
                <span
                    class="window-control"
                    id="close-button"
                    title=""
                    role="button"
                    tabindex="0"
                    onclick={() => closeWindow()}
                    onkeydown={(e) => e.key === "Enter" && closeWindow()}
                >
                    <span>×</span>
                </span>
            </div>
        </div>
    </div>
</main>

<style>
    #titlebar {
        position: fixed;
        top: 0;
        left: 0;
        right: 0;
        height: var(--height);
        background: var(--surface2);
        border-bottom: 0.5px solid var(--border);
        z-index: 10;
    }

    #drag-region {
        height: 100%;
        display: flex;
        align-items: center;
        position: relative;
        --wails-draggable: drag;
    }

    #titlebar-hover-area {
        position: fixed;
        top: calc(var(--height) - 8px);
        left: 0;
        right: 0;
        height: 14px;
        z-index: 11;
        pointer-events: all;
    }

    #window-title {
        position: absolute;
        left: 50%;
        transform: translateX(-50%);
        font-size: 11px;
        letter-spacing: 0.12em;
        color: var(--text);
        font-weight: 300;
    }

    #window-title .dim1 {
        color: var(--dim1);
    }

    #window-title .dim2 {
        color: var(--dim2);
    }

    #window-controls {
        position: absolute;
        right: 0;
        top: 0;
        height: 100%;
        display: flex;
    }

    .window-control {
        width: 44px;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        font-size: 15px;
        color: var(--text);
        cursor: pointer;
    }

    .window-control:hover {
        color: var(--text);
        background: var(--window-icon-hover);
    }

    #close-button:hover {
        color: var(--red);
    }
</style>
