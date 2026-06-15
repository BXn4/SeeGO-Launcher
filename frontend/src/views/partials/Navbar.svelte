<script lang="ts">
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";

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
            <button
                class="navbar-button"
                id="forum"
                title={forum}
                on:click={() => window._openURL("https://forum.see-rpg.com/")}
            >
                {@html Icons.Navbar.Forum}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button"
                id="ucp"
                title={ucp}
                on:click={() => window._openURL("https://ucp.see-rpg.com/")}
            >
                {@html Icons.Navbar.UCP}
                <span class="indicator"></span>
            </button>
            <div class="divider"></div>
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
    @import "../../public/styles/partials/navbar.css";

    * {
        box-sizing: border-box;
        margin: 0;
        padding: 0;
    }
</style>
