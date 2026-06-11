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
	<h4>${heading}</h4>`;
            } else if (
                trimmed.startsWith("ncl# ##") ||
                trimmed.startsWith("##")
            ) {
                const heading = trimmed.replace(/^ncl# ##|^##/, "").trim();
                parsed += `
	<h3>${heading}</h3>`;
            } else if (trimmed.startsWith("#")) {
                const heading = trimmed.replace(/^#/, "").trim();
                parsed += `
	<h2>${heading}</h2>`;
            } else if (trimmed.startsWith("c#")) {
                const text = trimmed.replace(/^c#/, "").trim();
                parsed += `
	<p class="contact">${text}</p>`;
            } else {
                const text = trimmed.replace(
                    /\[\[href\]\](.*?)\[\[href\]\]/g,
                    '<a href="$1" target="_blank">$1</a>',
                );
                parsed += `<p>${text}</p>`;
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
    <div id="content">
        <h1 id="terms-title">{title}</h1>
        <p id="terms-modified">{modified}</p>
        <hr />
        <div id="terms-content">
            <h3>
                SeeGO Launcher is an open-source project distributed under the
                MIT License. It is provided "as is", without warranty of any
                kind. You are free to use, modify, and distribute it.
            </h3>
            {@html terms}
        </div>
        <div id="actionbar">
            <button
                class="button"
                id="terms-decline-button"
                onclick={() => declineTerms()}
            >
                {decline}
            </button>
            <button
                class="button"
                id="terms-accept-button"
                onclick={() => acceptTerms()}
            >
                {accept}
            </button>
        </div>
    </div>
</main>

<style>
    main {
        width: 100vw;
        height: 100vh;
        background: var(--bg);
        overflow: hidden;
        position: relative;
        background: linear-gradient(
            -30deg,
            #1b2636,
            #1b2534,
            #1a2433,
            #1a2331,
            #1a222f,
            #1a212e,
            #19202c,
            #191f2a,
            #191e29,
            #181e27,
            #181d25,
            #171c24,
            #171b22,
            #161a21,
            #16191f
        );
    }

    #content {
        margin-top: var(--height);
        height: calc(100vh - var(--height) - 64px);
        overflow-y: auto;
        padding: 32px 48px 32px;
        scrollbar-width: thin;
    }

    #actionbar {
        position: fixed;
        bottom: 0;
        left: 0;
        right: 0;
        height: 64px;
        background: var(--surface2);
        display: flex;
        align-items: center;
        justify-content: flex-end;
        gap: 10px;
        padding: 0 48px;
        z-index: 10;
    }

    .button {
        padding: 0 20px;
        height: 34px;
        border-radius: 6px;
        border: 0.5px solid var(--button-border);
        background: transparent;
        color: var(--text);
        font-size: 12px;
        cursor: pointer;
    }

    #terms-decline-button:hover {
        background: var(--decline-hover-color);
        border-color: var(--decline-hover-border);
        color: var(--red);
    }

    #terms-accept-button:hover {
        background: var(--accept-hover-color);
        border-color: var(--accept-hover-border);
    }

    #terms-title {
        font-size: 18px;
        margin-bottom: 6px;
        letter-spacing: 0.01em;
    }

    #terms-modified {
        font-size: 11px;
        letter-spacing: 0.04em;
        margin-bottom: 20px;
    }

    #terms-content :global(h2) {
        font-size: 13px;
        color: var(--dim1);
        letter-spacing: 0.02em;
        margin: 28px 0 12px;
        padding-left: 14px;
        border-left: 2px solid var(--dim1);
    }

    #terms-content :global(h3) {
        font-size: 11px;
        color: var(--text-dim);
        letter-spacing: 0.06em;
        text-transform: uppercase;
        margin: 16px 0 6px;
        padding-left: 16px;
    }

    #terms-content :global(h4) {
        font-size: 11px;
        margin: 8px 0 4px;
        padding-left: 26px;
    }

    #terms-content :global(p) {
        font-size: 12px;
        line-height: 1.8;
        color: var(--text);
        margin: 0 0 10px;
        padding-left: 16px;
    }

    #terms-content :global(p.contact) {
        font-size: 12px;
        color: var(--comment);
        padding-left: 0;
        line-height: 1.9;
    }

    #terms-content :global(a) {
        color: var(--accent);
        text-decoration: none;
    }

    #terms-content :global(a:hover) {
        text-decoration: underline;
    }

    #terms-content :global(h2:not(:first-child)) {
        margin-top: 32px;
    }
</style>
