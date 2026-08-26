<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import {
        Config,
        Localization,
    } from "../../../bindings/seegolauncher/internal/services";
    import { onMount } from "svelte";
    import { Event } from "../../utils/consts";

    async function closeWindow() {
        await Events.Emit(Event.App.notActive, null);
        await Events.Emit(Event.App.close, null);
    }

    async function minimizeWindow() {
        await Events.Emit(Event.App.notActive, null);
        await Events.Emit(Event.App.minimize, null);
    }

    async function toggleMaximize() {
        // await Events.Emit("toggle-maximize", null);
    }

    onMount(() => {
        setLocales();
    });

    Events.On(Event.Global.updateLanguage, async (e) => {
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
        <div id="window-title">
            <span class="dim1">See</span><span class="dim2">Go</span>
        </div>
        <div id="window-controls">
            <span
                class="window-control interactive"
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
                class="window-control interactive"
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
</main>

<style>
    @import "../../public/styles/partials/titlebar.css";
</style>
