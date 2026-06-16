<script lang="ts">
    import { Events, Browser } from "@wailsio/runtime";
    import Splash from "./views/Splash.svelte";
    import Terms from "./views/Terms.svelte";
    import Main from "./views/Main.svelte";
    import Titlebar from "./views/partials/Titlebar.svelte";
    import Navbar from "./views/partials/Navbar.svelte";
    let view = $state("splash");
    let oldView = "splash";

    (window as any)._openURL = async (url: string) => {
        oldView = view;
        navigate("splash");
        Browser.OpenURL(url);

        Events.Emit("splash:setCurrentProgress", "opened-browser-window");

        await sleep(3000);
        navigate(oldView);
    };
    Events.On("app:navigate", (e) => {
        navigate(e.data);
    });

    function navigate(value: string) {
        view = value;
    }

    async function sleep(ms: number): Promise<void> {
        return new Promise((resolve) => setTimeout(resolve, ms));
    }
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
