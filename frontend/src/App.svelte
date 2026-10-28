<script lang="ts">
    import { Events, Browser } from "@wailsio/runtime";
    import Splash from "./views/Splash.svelte";
    import Terms from "./views/Terms.svelte";
    import Main from "./views/Main.svelte";
    import Titlebar from "./views/partials/Titlebar.svelte";
    import Navbar from "./views/partials/Navbar.svelte";
    import { Config } from "../bindings/seegolauncher/internal/services";
    import { onMount } from "svelte";
    import {
        Event,
        View,
        currentAppView,
        currentMainView,
        isTermsView,
    } from "./utils/consts";
    import { initNews } from "./managers/news";
    import {
        fetchServerStatus,
        startFetchServerStatusInterval,
        stopFetchServerStatusInterval,
    } from "./managers/server";
    import { initLocalization } from "./managers/localization";
    import { get } from "svelte/store";
    import { initSettings } from "./managers/settings";

    let navigateEvent: () => void;
    let updateSettingEvent: () => void;
    let notActiveEvent: () => void;
    let activeEVent: () => void;
    let appReadyEvent: () => void;

    (window as any)._openURL = async (url: string) => {
        navigate(View.splash);
        Browser.OpenURL(url);

        Events.Emit(
            Event.Splash.setCurrentProgress,
            Event.Splash.openedBrowserWindow,
        );

        await sleep(1500);
        if (get(isTermsView)) {
            navigate(View.terms);
        } else {
            navigate(View.main);
        }

        Events.Emit(Event.Main.navigate, get(currentMainView));
    };

    function navigate(view: string) {
        currentAppView.set(view);
    }

    async function sleep(ms: number): Promise<void> {
        return new Promise((resolve) => setTimeout(resolve, ms));
    }

    function setTheme(theme: string) {
        document.documentElement.classList.remove("dark", "dracula", "light");
        document.documentElement.classList.add(theme);
    }

    function setAnimations(value: boolean) {
        document.documentElement.classList.toggle("disable-animations", !value);
    }

    onMount(() => {
        (async () => {
            await initLocalization();
            setTheme(await Config.GetTheme());
            setAnimations(await Config.GetEnableAnimations());

            navigateEvent = Events.On(Event.App.navigate, (e) => {
                if (e.data == View.terms) {
                    isTermsView.set(true);
                } else if (e.data == View.main) {
                    isTermsView.set(false);
                }
                navigate(e.data);
            });

            updateSettingEvent = Events.On(
                Event.App.updateSetting,
                async (e) => {
                    const [key, value] = e.data;
                    switch (key) {
                        case "theme":
                            setTheme(value);
                            Config.SetTheme(value);
                            return;
                        case "language":
                            Config.SetLanguage(value);
                            Events.Emit(Event.Global.updateLanguage);
                            return;
                        case "anims":
                            setAnimations(value);
                            Config.SetEnableAnimations(value);
                            return;
                    }
                },
            );

            notActiveEvent = Events.On(Event.App.notActive, async (e) => {
                stopFetchServerStatusInterval();
            });

            activeEVent = Events.On(Event.App.active, async (e) => {
                startFetchServerStatusInterval();
            });

            appReadyEvent = Events.On(Event.App.ready, async (e) => {
                fetchServerStatus();
                initNews();
                initSettings();
                startFetchServerStatusInterval();
            });

            Events.Emit(Event.App.domReady);
        })();

        return () => {
            navigateEvent();
            updateSettingEvent();
            notActiveEvent();
            activeEVent();
            appReadyEvent();
        };
    });
</script>

{#if $currentAppView === View.splash}
    <Splash />
{:else if $currentAppView === View.terms}
    <Terms />
    <Titlebar />
{:else if $currentAppView === View.main}
    <Main />
    <Titlebar />
    <Navbar />
{/if}
