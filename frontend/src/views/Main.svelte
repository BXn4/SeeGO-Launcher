<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import { currentMainView, Event, View } from "../utils/consts";
    import Home from "./view/Home.svelte";
    import News from "./view/News.svelte";
    import Settings from "./view/Settings.svelte";
    import { onMount } from "svelte";
    import {
        startFetchServerStatusInterval,
        stopFetchServerStatusInterval,
    } from "../managers/server";
    import { get } from "svelte/store";

    let navigateEvent: () => void;

    onMount(() => {
        navigateEvent = Events.On(Event.Main.navigate, (e) => {
            currentMainView.set(e.data);

            if (get(currentMainView) != View.home) {
                stopFetchServerStatusInterval();
            } else {
                startFetchServerStatusInterval();
            }
        });

        return () => {
            navigateEvent();
        };
    });
</script>

{#if $currentMainView === View.home}
    <Home />
{:else if $currentMainView === View.news}
    <News />
{:else if $currentMainView === View.settings}
    <Settings />
{/if}
