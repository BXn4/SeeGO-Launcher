<script lang="ts">
    import { Icons } from "../../utils/icons";
    import {
        Event,
        View,
        IDs,
        Classes,
        URLs,
        currentNavbarActive,
    } from "../../utils/consts";
    import { locales, localization } from "../../managers/localization";
    import { switchNavTab } from "../../controllers/navbar";
</script>

<main>
    <div id={IDs.Partials.Navbar.ID}>
        <div class="{Classes.Partials.Navbar.item} {Classes.Global.top}">
            <button
                class="{Classes.Partials.Navbar.button} {Classes.Global
                    .interactive}"
                class:active={$currentNavbarActive === View.home}
                id={IDs.View.home}
                title={$locales[localization.homeTitle]}
                onclick={() => switchNavTab(View.home)}
            >
                {@html Icons.Navbar.Home}
                <span class={Classes.Partials.Navbar.indicator}></span>
            </button>
            <button
                class="{Classes.Partials.Navbar.button} {Classes.Global
                    .interactive}"
                class:active={$currentNavbarActive === View.news}
                id={IDs.View.news}
                title={$locales[localization.newsTitle]}
                onclick={() => switchNavTab(View.news)}
            >
                {@html Icons.Navbar.News}
                <span class={Classes.Partials.Navbar.indicator}></span>
            </button>
            <button
                class="{Classes.Partials.Navbar.button} {Classes.Global
                    .interactive}"
                id={IDs.View.forum}
                title={$locales[localization.forumTitle]}
                onclick={() => window._openURL(URLs.Forum)}
            >
                {@html Icons.Navbar.Forum}
                <span class={Classes.Global.altIcon}
                    >{@html Icons.UI.External}</span
                >
                <span class={Classes.Partials.Navbar.indicator}></span>
            </button>
            <button
                class="{Classes.Partials.Navbar.button} {Classes.Global
                    .interactive}"
                id={IDs.View.ucp}
                title={$locales[localization.ucpTitle]}
                onclick={() => window._openURL(URLs.Forum)}
            >
                {@html Icons.Navbar.UCP}
                <span class={Classes.Global.altIcon}
                    >{@html Icons.UI.External}</span
                >
                <span class={Classes.Partials.Navbar.indicator}></span>
            </button>
            <div class={Classes.Partials.Navbar.divider}></div>
        </div>
        <div class="{Classes.Partials.Navbar.item} {Classes.Global.bottom}">
            <div class={Classes.Partials.Navbar.divider}></div>
            <button
                class="{Classes.Partials.Navbar.button} {Classes.Global
                    .interactive}"
                id={IDs.View.shop}
                title={$locales[localization.shopTitle]}
                onclick={() => window._openURL(URLs.Store)}
            >
                {@html Icons.Navbar.Shop}
                <span class={Classes.Global.altIcon}
                    >{@html Icons.UI.External}</span
                >
                <span class={Classes.Partials.Navbar.indicator}></span>
            </button>
            <button
                class="{Classes.Partials.Navbar.button} {Classes.Global
                    .interactive}"
                class:active={$currentNavbarActive === View.gallery}
                id={IDs.View.gallery}
                title={$locales[localization.galleryTitle]}
                onclick={() => switchNavTab(View.gallery)}
            >
                {@html Icons.Navbar.Gallery}
                <span class={Classes.Partials.Navbar.indicator}></span>
            </button>
            <button
                class="{Classes.Partials.Navbar.button} {Classes.Global
                    .interactive}"
                class:active={$currentNavbarActive === View.help}
                id={IDs.View.help}
                title={$locales[localization.helpTitle]}
                onclick={() => switchNavTab(View.help)}
            >
                {@html Icons.Navbar.Help}
                <span class={Classes.Partials.Navbar.indicator}></span>
            </button>
            <button
                class="{Classes.Partials.Navbar.button} {Classes.Global
                    .interactive}"
                class:active={$currentNavbarActive === View.settings}
                id={IDs.View.settings}
                title={$locales[localization.settingsTitle]}
                onclick={() => switchNavTab(View.settings)}
            >
                {@html Icons.Navbar.Settings}
                <span class={Classes.Partials.Navbar.indicator}></span>
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
