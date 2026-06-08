<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import Splash from "./views/Splash.svelte";
    import Terms from "./views/Terms.svelte";
    import Main from "./views/Main.svelte";
    import {
        Localization,
        Config,
    } from "../bindings/seegolauncher/internal/services";

    let view = $state("splash");

    Events.On("navigate", (e) => {
        view = e.data;
    });

    Events.On("update-text", async (e) => {
        const el = document.getElementById(e.data.id);
        if (el) {
            el.textContent = await Localization.Get(
                e.data.value,
                await Config.GetLanguage(),
            );
        }
    });
</script>

{#if view === "splash"}
    <Splash />
{:else if view === "terms"}
    <Terms />
{:else if view === "main"}
    <Main />
{/if}
