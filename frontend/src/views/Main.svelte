<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import { State } from "../utils/consts";
    import Home from "./view/Home.svelte";
    import News from "./view/News.svelte";
    import Settings from "./view/Settings.svelte";

    Events.On("main:navigate", (e) => {
        State.currentMainView = e.data;
    });
</script>

{#if State.currentMainView === "home"}
    <Home />
{:else if State.currentMainView === "news"}
    <News />
{:else if State.currentMainView === "settings"}
    <Settings />
{/if}
