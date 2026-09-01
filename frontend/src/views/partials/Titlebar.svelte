<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import { Classes, Event, IDs } from "../../utils/consts";
    import { localization, locales } from "../../managers/localization";

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
</script>

<main>
    <div
        id={IDs.Partials.Titlebar.ID}
        role="button"
        tabindex="0"
        ondblclick={() => toggleMaximize()}
    >
        <div id={IDs.Partials.Titlebar.windowTitle}>
            <span class={Classes.Partials.Titlebar.dim1}>See</span><span
                class={Classes.Partials.Titlebar.dim2}>Go</span
            >
        </div>
        <div id={IDs.Partials.Titlebar.windowControls}>
            <span
                class="{Classes.Partials.Titlebar.windowControl} {Classes.Global
                    .interactive}"
                id={IDs.Partials.Titlebar.minimizeButton}
                title={$locales[localization.windowMinimize]}
                role="button"
                tabindex="0"
                onclick={() => minimizeWindow()}
                onkeydown={(e) => e.key === "Enter" && minimizeWindow()}
            >
                <span>−</span>
            </span>
            <span
                class="{Classes.Partials.Titlebar.windowControl} {Classes.Global
                    .interactive}"
                id={IDs.Partials.Titlebar.closeButton}
                title={$locales[localization.windowClose]}
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
