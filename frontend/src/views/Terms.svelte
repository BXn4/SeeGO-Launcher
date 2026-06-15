<script lang="ts">
    import { onMount } from "svelte";
    import { GetCachedTerms } from "../../bindings/seegolauncher/internal/services/cacheservice";
    import { Events } from "@wailsio/runtime";
    import Titlebar from "./partials/Titlebar.svelte";
    import {
        Config,
        Localization,
    } from "../../bindings/seegolauncher/internal/services";

    let title = "";
    let modified = "";
    let terms = "";

    let accept: string = "";
    let decline: string = "";
    let lang: string = "";

    function parse(raw: string): string {
        const lines = raw.split("\n");
        let parsed = "";
        let headerParsed = false;
        for (const line of lines) {
            const trimmed = line.trim();
            if (!headerParsed) {
                if (trimmed === "ind#") {
                    headerParsed = true;
                    continue;
                }
                if (!title && trimmed) {
                    title = trimmed;
                    continue;
                }
                if (!modified && trimmed) {
                    modified = trimmed;
                    continue;
                }
                continue;
            }
            if (!trimmed) continue;
            if (trimmed.startsWith("ncl# ###") || trimmed.startsWith("###")) {
                const heading = trimmed.replace(/^ncl# ###|^###/, "").trim();
                parsed += `
    <p class="highlighted terms-category-sub">${heading}</p>`;
            } else if (
                trimmed.startsWith("ncl# ##") ||
                trimmed.startsWith("##")
            ) {
                const heading = trimmed.replace(/^ncl# ##|^##/, "").trim();
                parsed += `
	<p class="highlighted terms-category-sub">${heading}</p>`;
            } else if (trimmed.startsWith("#")) {
                const heading = trimmed.replace(/^#/, "").trim();
                parsed += `
	<p class="highlighted terms-category">${heading}<p>`;
            } else if (trimmed.startsWith("c#")) {
                const text = trimmed.replace(/^c#/, "").trim();
                parsed += `
	<p class="comment terms-contact">${text}</p>`;
            } else {
                const text = trimmed.replace(
                    /\[\[href\]\](.*?)\[\[href\]\]/g,
                    '<a class="highlighted link" onclick="window._openURL(\'$1\')">$1</a>',
                );
                parsed += `<p class="text terms-text">${text}</p>`;
            }
        }
        return parsed;
    }
    onMount(async () => {
        const raw = await GetCachedTerms();
        terms = parse(raw);

        setLocales();
    });

    async function setLocales() {
        lang = await Config.GetLanguage();

        [accept, decline] = await Promise.all([
            Localization.Get("accept", lang),
            Localization.Get("decline", lang),
        ]);
    }

    async function acceptTerms() {
        await Events.Emit("terms-accepted", null);
    }
    async function declineTerms() {
        await Events.Emit("terms-declined", null);
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
            <button class="button cancel" onclick={() => declineTerms()}>
                {decline}
            </button>
            <button class="button ok" onclick={() => acceptTerms()}>
                {accept}
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
