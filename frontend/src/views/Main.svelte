<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import { State, Event, View } from "../utils/consts";
    import Home from "./view/Home.svelte";
    import News from "./view/News.svelte";
    import Settings from "./view/Settings.svelte";

    Events.On(Event.Main.navigate, (e) => {
        // State.currentNavbarActive = e.data; not works
        State.currentMainView = e.data;
    });
</script>

{#if State.currentMainView === View.home}
    <Home />
{:else if State.currentMainView === View.news}
    <News />
{:else if State.currentMainView === View.settings}
    <Settings />
{/if}
