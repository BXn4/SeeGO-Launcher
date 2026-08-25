<script lang="ts">
    import { onMount } from "svelte";
    import { GetCachedTerms } from "../../bindings/seegolauncher/internal/services/cacheservice";
    import { Events } from "@wailsio/runtime";
    import Titlebar from "./partials/Titlebar.svelte";
    import {
        initLocalization,
        locales,
        localization,
    } from "../managers/localization";
    import { parseTerms } from "../utils/string";
    import { Event } from "../utils/consts";
    let terms = "";
    let title = "";
    let modified = "";

    onMount(async () => {
        const raw = await GetCachedTerms();
        [terms, title, modified] = parseTerms(raw);

        initLocalization();
    });

    async function acceptTerms() {
        await Events.Emit(Event.Terms.accept, null);
    }
    async function declineTerms() {
        await Events.Emit(Event.Terms.decline, null);
    }
</script>

<main>
    <Titlebar></Titlebar>
    <div id="terms">
        <h1 class="text terms-title">{title}</h1>
        <p class="comment terms-modified">{modified}</p>
        <hr />
        <div id="terms-content">
            <p class="comment terms-comment">
                SeeGO Launcher is an open-source project distributed under the
                MIT License. It is provided "as is", without warranty of any
                kind. You are free to use, modify, and distribute it.
            </p>
            {@html terms}
        </div>
        <div id="terms-actionbar">
            <button
                class="button cancel interactive"
                onclick={() => declineTerms()}
            >
                {$locales[localization.decline]}
            </button>
            <button class="button ok interactive" onclick={() => acceptTerms()}>
                {$locales[localization.accept]}
            </button>
        </div>
    </div>
</main>

<style>
    @import "../public/styles/views/terms.css";
    main {
        width: 100vw;
        height: 100vh;
        display: flex;
        align-items: center;
        justify-content: center;
    }
</style>
