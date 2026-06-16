<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import Home from "./view/Home.svelte";
    import News from "./view/News.svelte";

    let view = $state(localStorage.getItem("view") ?? "home");

    Events.On("main:navigate", (e) => {
        if (view != e.data) {
            view = e.data;
            localStorage.setItem("view", view);
        }
    });
</script>

{#if view === "home"}
    <Home />
{:else if view === "news"}
    <News />
{/if}
