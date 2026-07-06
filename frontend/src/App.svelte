<script lang="ts">
    import { Events, Browser } from "@wailsio/runtime";
    import Splash from "./views/Splash.svelte";
    import Terms from "./views/Terms.svelte";
    import Main from "./views/Main.svelte";
    import Titlebar from "./views/partials/Titlebar.svelte";
    import Navbar from "./views/partials/Navbar.svelte";
    import { Config } from "../bindings/seegolauncher/internal/services";
    import { onMount } from "svelte";

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

    Events.On("app:updateSetting", async (e) => {
        const [key, value] = e.data;
        switch (key) {
            case "theme":
                setTheme(value);
                Config.SetTheme(value);
                return;
            case "language":
                Config.SetLanguage(value);
                return;
            case "anims":
                setAnimations(value);
                Config.SetEnableAnimations(value);
                return;
        }
    });

    Events.On("app:notActive", async (e) => {
        Events.Emit("home:stopInterval", null);
    });

    Events.On("app:active", async (e) => {
        Events.Emit("home:startInterval", null);
    });

    function navigate(value: string) {
        view = value;
    }

    async function sleep(ms: number): Promise<void> {
        return new Promise((resolve) => setTimeout(resolve, ms));
    }

    function setTheme(theme: string) {
        document.documentElement.classList.remove("dark", "dracula", "light");
        document.documentElement.classList.add(theme);
    }

    function setAnimations(value: boolean) {
        const link = document.getElementById("anims") as HTMLLinkElement;
        link.disabled = !value;
    }

    onMount(async () => {
        setTheme(await Config.GetTheme());

        Events.Emit("app:ready");
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
