<script lang="ts">
    import { onMount, tick } from "svelte";
    import { Events } from "@wailsio/runtime";
    import { Icons } from "../../utils/icons";
    import { news, loadingSuccess } from "../../managers/news";
    import {
        initLocalization,
        locales,
        localization,
    } from "../../managers/localization";
    import { Event, View } from "../../utils/consts";
    import { sleep } from "../../utils/helper";
    import {
        initServerStatus,
        serverPlayers,
        serverQueue,
        serverAdmins,
        serverSlots,
        serverOnline,
        queueTime,
    } from "../../managers/server";

    let interval: ReturnType<typeof setInterval> | undefined;

    function start() {
        if (!interval) {
            interval = setInterval(setServerStatuts, 60 * 1000);
        }
    }

    function stop() {
        if (interval) {
            clearInterval(interval);
            interval = undefined;
        }
    }

    Events.On(Event.Global.startInterval, async (e) => {
        start();
    });

    Events.On(Event.Global.stopInterval, async (e) => {
        // Events.Emit(Event.Global.feedback, "stop");
        stop();
    });

    async function setServerStatuts() {
        initServerStatus();
    }

    onMount(() => {
        void (async () => {
            await setServerStatuts();
        })();

        initLocalization();
        start();

        const startInterval = Events.On(
            Event.Global.startInterval,
            async (e) => {
                start();
            },
        );

        const stopInterval = Events.On(Event.Global.stopInterval, async (e) => {
            stop();
        });

        return () => {
            stop();
            startInterval();
            stopInterval();
        };
    });
</script>

<main>
    <div id="home-view">
        <div class="feed-layout">
            {#if $loadingSuccess == false}
                <div id="hero-card" class="hero-card">
                    <div class="error-view">
                        {@html Icons.UI.Alert}
                        <p class="error-title text">
                            {$locales[localization.NewsLoadFailed]}
                        </p>
                        <p class="error-comment comment">
                            {$locales[localization.NewsLoadFailedDesc]}
                        </p>
                        <button
                            class="button interactive"
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
                    id="hero-card"
                    class="hero-card"
                    style="background-image: url('{item.Image}')"
                >
                    <span class="hero-overlay"></span>
                    <div class="hero-content">
                        <span class="news-badge"
                            >{$locales[localization.newsLatest]}</span
                        >
                        <p id="hero-news-title" class="news-title">
                            {item.Title}
                        </p>
                        <p id="hero-news-comment" class="news-comment">
                            {item.Content}
                        </p>
                        <button
                            class="button news-read interactive"
                            id="hero-news-read-latest"
                            onclick={async () => {
                                Events.Emit(
                                    Event.Main.Navbar.switchNavTab,
                                    View.news,
                                );
                                await sleep(100);
                                Events.Emit(Event.Main.News.readLatest, null);
                            }}>{$locales[localization.newsRead]}</button
                        >
                    </div>
                </div>
            {/if}
            <div class="home-items">
                <h3 class="text items-title">
                    SeeRPG Club {$locales[localization.clubMembership]}
                </h3>
                <div id="items-container"></div>
            </div>
        </div>
        <aside class="sidebar">
            <div class="widget">
                {#if $serverOnline == false}
                    <div class="status-container">
                        <div class="status-info">
                            <span id="server-status" class="status-text offline"
                                >{$locales[
                                    localization.serverStatusOffline
                                ]}</span
                            >
                            <span id="players-count" class="text player-count"
                                >{$serverPlayers} / {$serverSlots}</span
                            >
                        </div>
                    </div>
                    <div id="server-fill" class="progress-bar">
                        <div
                            class="progress-bar-fill offline"
                            style="width: {($serverPlayers / $serverSlots) *
                                100}%"
                        ></div>
                    </div>
                {:else if $serverOnline == true}
                    <div class="status-container">
                        <div class="status-info">
                            <span id="server-status" class="status-text online"
                                >{$locales[
                                    localization.serverStatusOnline
                                ]}</span
                            >
                            <span id="players-count" class="text player-count"
                                >{$serverPlayers} / {$serverSlots}</span
                            >
                        </div>
                    </div>
                    <div id="server-fill" class="progress-bar">
                        <div
                            class="progress-bar-fill online"
                            style="width: {($serverPlayers / $serverSlots) *
                                100}%"
                        ></div>
                    </div>
                {/if}

                <div class="stats">
                    <div class="stat-box">
                        <span class="text stat-title"
                            >{$locales[localization.serverStatusAdmins]}</span
                        >
                        <span id="admins-count" class="stat-value"
                            >{$serverAdmins}</span
                        >
                    </div>
                    <div class="stat-box">
                        <span class="text stat-title"
                            >{$locales[localization.serverStatusQueue]}</span
                        >
                        <span id="queue-count" class="stat-value"
                            >{$serverQueue}</span
                        >
                    </div>
                </div>
                {#if $queueTime > 0}
                    <p id="estimated-connect" class="comment estimated-connect">
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
            <div class="widget">
                <h3 class="text widget-title">
                    {$locales[localization.community]}
                </h3>
                <div class="social-links-grid">
                    <button
                        onclick={() =>
                            window._openURL(
                                "https://discord.com/invite/seerpg",
                            )}
                        class="socials interactive"
                        title="Discord"
                        >{@html Icons.Community.Discord}
                        <span class="alt-icon">{@html Icons.UI.External}</span>
                    </button>
                    <button
                        onclick={() =>
                            window._openURL(
                                "https://www.facebook.com/seerpgofficial",
                            )}
                        class="socials interactive"
                        title="Facebook"
                        >{@html Icons.Community.FaceBook}
                        <span class="alt-icon">{@html Icons.UI.External}</span>
                    </button>
                    <button
                        onclick={() =>
                            window._openURL(
                                "https://www.facebook.com/seerpgofficial",
                            )}
                        class="socials interactive"
                        title="TikTok"
                        >{@html Icons.Community.TikTok}
                        <span class="alt-icon">{@html Icons.UI.External}</span>
                    </button>
                    <button
                        onclick={() =>
                            window._openURL(
                                "https://www.youtube.com/@seerpgofficial",
                            )}
                        class="socials interactive"
                        title="YouTube"
                        >{@html Icons.Community.YouTube}
                        <span class="alt-icon">{@html Icons.UI.External}</span>
                    </button>
                </div>
            </div>
            <div class="widget">
                <h3 class="text widget-title">
                    {$locales[localization.launcherReady]}
                </h3>
                <button
                    id="connect-button"
                    class="text button connect interactive"
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
