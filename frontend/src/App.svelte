<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import Splash from "./views/Splash.svelte";
    import Main from "./views/Main.svelte";

    let view = $state("splash");

    Events.On("navigate", (e) => {
        view = e.data;
    });
</script>

{#if view === "splash"}
    <Splash />
{:else if view === "main"}
    <Main />
{/if}
