<script lang="ts">
    import { Events } from "@wailsio/runtime";
    import { Icons } from "../../utils/icons";
    import { news, loadingSuccess } from "../../managers/news";
    import { locales, localization } from "../../managers/localization";
    import { Classes, Event, IDs, URLs, View } from "../../utils/consts";
    import { stripMarkup } from "../../utils/string";
    import { sleep } from "../../utils/helper";
    import {
        serverPlayers,
        serverQueue,
        serverAdmins,
        serverSlots,
        serverOnline,
        queueTime,
    } from "../../managers/server";
</script>

<main>
    <div id={IDs.Views.Home.ID}>
        <div class={Classes.Views.Home.layout}>
            {#if $loadingSuccess == false}
                <div id={IDs.Views.Home.card} class={Classes.Views.Home.card}>
                    <div id={IDs.Views.Error.ID}>
                        {@html Icons.UI.Alert}
                        <p
                            class="{Classes.Views.Error.title} {Classes.Global
                                .text}"
                        >
                            {$locales[localization.NewsLoadFailed]}
                        </p>
                        <p
                            class="{Classes.Views.Error.alt} {Classes.Global
                                .comment}"
                        >
                            {$locales[localization.NewsLoadFailedDesc]}
                        </p>
                        <button
                            class="{Classes.Global.button} {Classes.Global
                                .interactive}"
                            onclick={() =>
                                Events.Emit(Event.Global.newsFeedUpdated, null)}
                        >
                            {$locales[localization.Retry]}
                        </button>
                    </div>
                </div>
            {:else if $loadingSuccess == true}
                {@const item = $news[0]}
                <div
                    id={IDs.Views.Home.card}
                    class={Classes.Views.Home.card}
                    style="background-image: url('{item.Image}')"
                >
                    <span class={Classes.Views.Home.overlay}></span>
                    <div class={Classes.Views.Home.content}>
                        <span class={Classes.Global.newsBadge}
                            >{$locales[localization.newsLatest]}</span
                        >
                        <p
                            id={IDs.Views.Home.title}
                            class={Classes.Global.newsTitle}
                        >
                            {item.Title}
                        </p>
                        <p
                            id={IDs.Views.Home.alt}
                            class={Classes.Global.newsAlt}
                        >
                            {stripMarkup(item.Content)}
                        </p>
                        <button
                            class="{Classes.Global.button} {Classes.Views.Home
                                .read} {Classes.Global.interactive}"
                            id={IDs.Views.Home.read}
                            onclick={async () => {
                                Events.Emit(
                                    Event.Main.Navbar.switchNavTab,
                                    View.news,
                                );
                                await sleep(1); // sleep needed, due to race conditions!
                                // first need to switch to the news view, then emit the readLatest function while on the news view.
                                // if its happens too fast, then its not going to be shown.
                                Events.Emit(Event.Main.News.readLatest, null);
                            }}>{$locales[localization.newsRead]}</button
                        >
                    </div>
                </div>
            {/if}
            <div class={Classes.Views.Home.items}>
                <h3
                    class="{Classes.Global.text} {Classes.Views.Home
                        .itemsTitle}"
                >
                    SeeRPG Club {$locales[localization.clubMembership]}
                </h3>
                <div id={IDs.Views.Home.container}></div>
            </div>
        </div>
        <aside class={Classes.Views.Home.sidebar}>
            <div class={Classes.Views.Home.widget}>
                {#if $serverOnline == false}
                    <div class={Classes.Views.Home.container}>
                        <div class={Classes.Views.Home.info}>
                            <span
                                id={IDs.Views.Home.status}
                                class="{Classes.Views.Home.text} {Classes.Views
                                    .Home.offline}"
                                >{$locales[
                                    localization.serverStatusOffline
                                ]}</span
                            >
                            <span
                                id={IDs.Views.Home.count}
                                class="{Classes.Global.text} {Classes.Views.Home
                                    .count}"
                                >{$serverPlayers} / {$serverSlots}</span
                            >
                        </div>
                    </div>
                    <div id={IDs.Views.Home.fill} class={Classes.Global.bar}>
                        <div
                            class="{Classes.Global.barFill} {Classes.Views.Home
                                .offline}"
                            style="width: {($serverPlayers / $serverSlots) *
                                100}%"
                        ></div>
                    </div>
                {:else if $serverOnline == true}
                    <div class={Classes.Views.Home.container}>
                        <div class={Classes.Views.Home.info}>
                            <span
                                id={IDs.Views.Home.status}
                                class="{Classes.Views.Home.text} {Classes.Views
                                    .Home.online}"
                                >{$locales[
                                    localization.serverStatusOnline
                                ]}</span
                            >
                            <span
                                id={IDs.Views.Home.count}
                                class="{Classes.Global.text} {Classes.Views.Home
                                    .count}"
                                >{$serverPlayers} / {$serverSlots}</span
                            >
                        </div>
                    </div>
                    <div id={IDs.Views.Home.fill} class={Classes.Global.bar}>
                        <div
                            class="{Classes.Global.barFill} {Classes.Views.Home
                                .online}"
                            style="width: {($serverPlayers / $serverSlots) *
                                100}%"
                        ></div>
                    </div>
                {/if}

                <div class={Classes.Views.Home.stats}>
                    <div class={Classes.Views.Home.statBox}>
                        <span
                            class="{Classes.Global.text} {Classes.Views.Home
                                .statTitle}"
                            >{$locales[localization.serverStatusAdmins]}</span
                        >
                        <span
                            id={IDs.Views.Home.adminsCount}
                            class={Classes.Views.Home.statValue}
                            >{$serverAdmins}</span
                        >
                    </div>
                    <div class={Classes.Views.Home.statBox}>
                        <span
                            class="{Classes.Global.text} {Classes.Views.Home
                                .statTitle}"
                            >{$locales[localization.serverStatusQueue]}</span
                        >
                        <span
                            id={IDs.Views.Home.queueCount}
                            class={Classes.Views.Home.statValue}
                            >{$serverQueue}</span
                        >
                    </div>
                </div>
                {#if $queueTime > 0}
                    <p
                        id={IDs.Views.Home.estimatedConnect}
                        class="{Classes.Global.comment} {Classes.Views.Home
                            .estimatedConnect}"
                    >
                        {$locales[localization.serverStatusEstimated]}:
                        {#if $queueTime >= 3600}
                            {($queueTime / 3600).toFixed(1)}
                            {$locales[localization.hours]}
                        {:else if $queueTime >= 60}
                            {($queueTime / 60).toFixed()}
                            {$locales[localization.minutes]}
                        {:else}
                            {Math.round($queueTime)}
                            {$locales[localization.seconds]}
                        {/if}
                    </p>
                {/if}
            </div>
            <div class={Classes.Views.Home.widget}>
                <h3
                    class="{Classes.Global.text} {Classes.Views.Home
                        .widgetTitle}"
                >
                    {$locales[localization.community]}
                </h3>
                <div class={Classes.Views.Home.socialWidget}>
                    <button
                        onclick={() => window._openURL(URLs.Discord)}
                        class="{Classes.Global.icon} {Classes.Global
                            .interactive}"
                        title={$locales[localization.discordSocialTitle]}
                        >{@html Icons.Community.Discord}
                        <span class={Classes.Global.altIcon}
                            >{@html Icons.UI.External}</span
                        >
                    </button>
                    <button
                        onclick={() => window._openURL(URLs.Facebook)}
                        class="{Classes.Global.icon} {Classes.Global
                            .interactive}"
                        title={$locales[localization.facebookSocialTitle]}
                        >{@html Icons.Community.FaceBook}
                        <span class={Classes.Global.altIcon}
                            >{@html Icons.UI.External}</span
                        >
                    </button>
                    <button
                        onclick={() => window._openURL(URLs.Tiktok)}
                        class="{Classes.Global.icon} {Classes.Global
                            .interactive}"
                        title={$locales[localization.tiktokSocialTitle]}
                        >{@html Icons.Community.TikTok}
                        <span class={Classes.Global.altIcon}
                            >{@html Icons.UI.External}</span
                        >
                    </button>
                    <button
                        onclick={() => window._openURL(URLs.Youtube)}
                        class="{Classes.Global.icon} {Classes.Global
                            .interactive}"
                        title={$locales[localization.youtubeSocialTitle]}
                        >{@html Icons.Community.YouTube}
                        <span class={Classes.Global.altIcon}
                            >{@html Icons.UI.External}</span
                        >
                    </button>
                </div>
            </div>
            <div class={Classes.Views.Home.widget}>
                <h3
                    class="{Classes.Global.text} {Classes.Views.Home
                        .widgetTitle}"
                >
                    {$locales[localization.launcherReady]}
                </h3>
                <button
                    id={IDs.Views.Home.connectButton}
                    class="{Classes.Global.text} {Classes.Global
                        .button} {Classes.Views.Home.connect} {Classes.Global
                        .interactive}"
                    ><!-->{@html Icons.Launcher.Play}<--->{$locales[
                        localization.launcherConnect
                    ]}</button
                >
            </div>
        </aside>
    </div>
</main>

<style>
    @import "../../public/styles/views/home.css";
    main {
        width: 100vw;
        height: 100vh;
        display: grid;
        grid-template-columns: var(--width, 240px) 1fr;
        grid-template-rows: var(--height, 48px) 1fr;
        overflow: hidden;
    }
</style>
