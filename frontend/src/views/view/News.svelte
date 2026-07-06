<script lang="ts">
    import Titlebar from "../partials/Titlebar.svelte";
    import Navbar from "../partials/Navbar.svelte";
    import { onMount } from "svelte";
    import { Browser } from "@wailsio/runtime";
    import {
        Config,
        Localization,
    } from "../../../bindings/seegolauncher/internal/services";
    import { initLocalization } from "../../managers/localization";

    onMount(async () => {
        initLocalization();
    });
</script>

<main>
    <Titlebar />
    <Navbar />
</main>

<style>
</style>
