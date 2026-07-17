<script lang="ts">
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";

    import {
        Config,
        Localization,
    } from "../../../bindings/seegolauncher/internal/services";

    import { Icons } from "../../utils/icons";
    import { State, Event, View } from "../../utils/consts";
    onMount(() => {
        setLocales();
    });

    let home: string = "";
    let news: string = "";
    let forum: string = "";
    let ucp: string = "";
    let shop: string = "";
    let gallery: string = "";
    let help: string = "";
    let settings: string = "";

    async function setLocales() {
        const lang = await Config.GetLanguage();

        [home, news, forum, ucp, shop, gallery, help, settings] =
            await Promise.all([
                Localization.Get("home-title", lang),
                Localization.Get("news-title", lang),
                Localization.Get("forum-title", lang),
                Localization.Get("ucp-title", lang),
                Localization.Get("shop-title", lang),
                Localization.Get("gallery-title", lang),
                Localization.Get("help-title", lang),
                Localization.Get("settings-title", lang),
            ]);
    }
</script>

<main>
    <div id="navbar">
        <div class="navbar-item top">
            <button
                class="navbar-button interactive"
                class:active={State.currentNavbarActive === View.home}
                id="home"
                title={home}
                onclick={() => {
                    State.currentNavbarActive = View.home;
                    Events.Emit(Event.Main.navigate, View.home);
                }}
            >
                {@html Icons.Navbar.Home}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button interactive"
                class:active={State.currentNavbarActive === View.news}
                id="news"
                title={news}
                onclick={() => {
                    State.currentNavbarActive = View.news;
                    Events.Emit(Event.Main.navigate, View.news);
                }}
            >
                {@html Icons.Navbar.News}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button interactive"
                id="forum"
                title={forum}
                onclick={() => window._openURL("https://forum.see-rpg.com/")}
            >
                {@html Icons.Navbar.Forum}
                <span class="alt-icon">{@html Icons.UI.External}</span>
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button interactive"
                id="ucp"
                title={ucp}
                onclick={() => window._openURL("https://ucp.see-rpg.com/")}
            >
                {@html Icons.Navbar.UCP}
                <span class="alt-icon">{@html Icons.UI.External}</span>
                <span class="indicator"></span>
            </button>
            <div class="divider"></div>
        </div>
        <div class="navbar-item bottom">
            <div class="divider"></div>
            <button
                class="navbar-button interactive"
                class:active={State.currentNavbarActive === View.shop}
                id="shop"
                title={shop}
                onclick={() => {
                    State.currentNavbarActive = View.shop;
                    Events.Emit(Event.Main.navigate, View.shop);
                }}
            >
                {@html Icons.Navbar.Shop}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button interactive"
                class:active={State.currentNavbarActive === View.gallery}
                id="gallery"
                title={gallery}
                onclick={() => {
                    State.currentNavbarActive = View.gallery;
                    Events.Emit(Event.Main.navigate, View.gallery);
                }}
            >
                {@html Icons.Navbar.Gallery}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button interactive"
                class:active={State.currentNavbarActive === View.help}
                id="help"
                title={help}
                onclick={() => {
                    State.currentNavbarActive = View.help;
                    Events.Emit(Event.Main.navigate, View.help);
                }}
            >
                {@html Icons.Navbar.Help}
                <span class="indicator"></span>
            </button>
            <button
                class="navbar-button interactive"
                class:active={State.currentNavbarActive === View.settings}
                id="settings"
                title={settings}
                onclick={() => {
                    State.currentNavbarActive = View.settings;
                    Events.Emit(Event.Main.navigate, View.settings);
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
