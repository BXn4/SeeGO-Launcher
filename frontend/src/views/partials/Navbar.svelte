<script lang="ts">
    import { onMount } from "svelte";
    import { Browser, Events } from "@wailsio/runtime";

    import {
        Config,
        Localization,
    } from "../../../bindings/seegolauncher/internal/services";

    import { Icons } from "../../utils/icons";
    onMount(() => {
        setLocales();
    });

    let home: string = "";
    let news: string = "";
    let forum: string = "";
    let ucp: string = "";
    let shop: string = "";
    let gallery: string = "";
    let players: string = "";
    let help: string = "";
    let settings: string = "";
    let active: string = "home";

    async function setLocales() {
        const lang = await Config.GetLanguage();

        [home, news, forum, ucp, shop, gallery, players, help, settings] =
            await Promise.all([
                Localization.Get("home-title", lang),
                Localization.Get("news-title", lang),
                Localization.Get("forum-title", lang),
                Localization.Get("ucp-title", lang),
                Localization.Get("shop-title", lang),
                Localization.Get("gallery-title", lang),
                Localization.Get("players-title", lang),
                Localization.Get("help-title", lang),
                Localization.Get("settings-title", lang),
            ]);
    }
</script>

<main>
    <div id="navbar">
        <div class="navbar-item top">
            <button
                class="navbar-button"
                class:active={active === "home"}
                id="home"
                title={home}
                on:click={() => {
                    active = "home";
                    Events.Emit("main:navigate", "home");
                }}
            >
                {@html Icons.Navbar.Home}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button"
                class:active={active === "news"}
                id="news"
                title={news}
                on:click={() => {
                    active = "news";
                    Events.Emit("main:navigate", "news");
                }}
            >
                {@html Icons.Navbar.News}
                <span class="indicator"></span>
            </button>
            <button class="navbar-button" id="forum" title={forum}>
                {@html Icons.Navbar.Forum}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button"
                id="ucp"
                title={ucp}
                on:click={() => Browser.OpenURL("https://ucp.see-rpg.com/")}
            >
                {@html Icons.Navbar.UCP}
                <span class="indicator"></span>
            </button>
        </div>
        <div class="navbar-item bottom">
            <div class="divider"></div>
            <button
                class="navbar-button"
                class:active={active === "shop"}
                id="shop"
                title={shop}
                on:click={() => {
                    active = "shop";
                    Events.Emit("main:navigate", "shop");
                }}
            >
                {@html Icons.Navbar.Shop}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button"
                class:active={active === "gallery"}
                id="gallery"
                title={gallery}
                on:click={() => {
                    active = "gallery";
                    Events.Emit("main:navigate", "gallery");
                }}
            >
                {@html Icons.Navbar.Gallery}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button"
                class:active={active === "players"}
                id="players"
                title={players}
                on:click={() => {
                    active = "players";
                    Events.Emit("main:navigate", "players");
                }}
            >
                {@html Icons.Navbar.ServerStatus}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button"
                class:active={active === "help"}
                id="help"
                title={help}
                on:click={() => {
                    active = "help";
                    Events.Emit("main:navigate", "help");
                }}
            >
                {@html Icons.Navbar.Help}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button"
                class:active={active === "settings"}
                id="settings"
                title={settings}
                on:click={() => {
                    active = "settings";
                    Events.Emit("main:navigate", "settings");
                }}
            >
                {@html Icons.Navbar.Settings}
                <span class="indicator"></span>
            </button>
        </div>
    </div>
</main>

<style>
    * {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }

    #navbar {
        position: fixed;
        top: var(--height);
        left: 0;
        width: var(--width);
        height: calc(100vh - var(--height));
        background: var(--surface1);
        border-right: 0.5px solid var(--border);
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        padding: 16px 0;
        z-index: 5;
    }

    .navbar-item {
        display: flex;
        flex-direction: column;
        align-items: center;
        gap: 4px;
    }

    .divider {
        width: 24px;
        height: 1px;
        background: var(--border);
        margin-bottom: 8px;
    }

    .navbar-button {
        position: relative;
        width: 40px;
        height: 40px;
        display: flex;
        align-items: center;
        justify-content: center;
        border-radius: 8px;
        border: none;
        background: transparent;
        color: var(--navbar-icon);
        cursor: pointer;
    }

    :global(.navbar-button svg) {
        width: 18px;
        height: 18px;
        flex-shrink: 0;
    }

    .navbar-button:hover {
        color: var(--navbar-icon-hover);
        background: var(--navbar-icon-hover-bg);
        transform: translateY(-1px);
    }

    .navbar-button:active {
        transform: scale(0.94);
    }

    .navbar-button.active {
        color: var(--accent);
        background: var(--navbar-icon-active);
    }

    .navbar-button .indicator {
        position: absolute;
        left: -1px;
        top: 50%;
        transform: translateY(-50%) scaleY(0);
        width: 2px;
        height: 16px;
        border-radius: 0 2px 2px 0;
        background: var(--accent);
    }

    .navbar-button.active .indicator {
        transform: translateY(-50%) scaleY(1);
    }

    #shop {
        color: var(--navbar-icon-highlighted);
    }
</style>
