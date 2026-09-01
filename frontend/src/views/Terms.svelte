<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import { locales, localization } from "../managers/localization";
    import { Classes, Event, IDs } from "../utils/consts";
    import { terms, title, modified, initTerms } from "../managers/terms";
    import { onMount } from "svelte";
    async function acceptTerms() {
        await Events.Emit(Event.Terms.accept, null);
    }
    async function declineTerms() {
        await Events.Emit(Event.Terms.decline, null);
    }

    onMount(async () => await initTerms());
</script>

<main>
    <div id={IDs.Views.Terms.ID}>
        <h1 class="{Classes.Global.text} {Classes.Views.Terms.title}">
            {$title}
        </h1>
        <p class="{Classes.Global.comment} {Classes.Views.Terms.alt}">
            {$modified}
        </p>
        <hr />
        <div id={IDs.Views.Terms.content}>
            <p class="{Classes.Global.comment} {Classes.Views.Terms.alt}">
                SeeGO Launcher is an open-source project distributed under the
                MIT License. It is provided "as is", without warranty of any
                kind. You are free to use, modify, and distribute it.
            </p>
            {@html $terms}
        </div>
        <div id={IDs.Views.Terms.actionbar}>
            <button
                class="{Classes.Global.button} {Classes.Global.cancel} {Classes
                    .Global.interactive}"
                onclick={() => declineTerms()}
            >
                {$locales[localization.decline]}
            </button>
            <button
                class="{Classes.Global.button} {Classes.Global.ok} {Classes
                    .Global.interactive}"
                onclick={() => acceptTerms()}
            >
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
