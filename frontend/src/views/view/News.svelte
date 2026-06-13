<script lang="ts">
    import Titlebar from "../partials/Titlebar.svelte";
    import Navbar from "../partials/Navbar.svelte";
    import { onMount } from "svelte";
    import { Browser } from "@wailsio/runtime";
    import {
        Config,
        Localization,
    } from "../../../bindings/seegolauncher/internal/services";
    let latest: string = "";
    let read: string = "";

    onMount(async () => {
        await setLocales();
    });

    async function setLocales() {
        let lang = await Config.GetLanguage();

        [latest, read] = await Promise.all([
            Localization.Get("news-latest", lang),
            Localization.Get("news-read", lang),
        ]);
    }
</script>

<main>
    <Titlebar />
    <Navbar />
</main>

<style>
</style>
