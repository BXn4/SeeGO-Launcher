<script lang="ts">
    import { Events, Browser } from "@wailsio/runtime";
    import Splash from "./views/Splash.svelte";
    import Terms from "./views/Terms.svelte";
    import Main from "./views/Main.svelte";
    import Titlebar from "./views/partials/Titlebar.svelte";
    import Navbar from "./views/partials/Navbar.svelte";
    let view = $state("splash");
    Events.On("app:navigate", (e) => {
        (window as any)._openURL = (url: string) => Browser.OpenURL(url);
        view = e.data;
    });
</script>

{#if view === "splash"}
    <Splash />
{:else if view === "terms"}
    <Terms />
{:else if view === "main"}
    <Main />
    <Titlebar />
    <Navbar />
{/if}
