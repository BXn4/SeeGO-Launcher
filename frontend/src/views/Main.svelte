<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import Home from "./view/Home.svelte";
    import News from "./view/News.svelte";
    import Settings from "./view/Settings.svelte";

    let view = $state("home");

    Events.On("main:navigate", (e) => {
        if (view != e.data) {
            view = e.data;
        }
    });
</script>

{#if view === "home"}
    <Home />
{:else if view === "news"}
    <News />
{:else if view === "settings"}
    <Settings />
{/if}
