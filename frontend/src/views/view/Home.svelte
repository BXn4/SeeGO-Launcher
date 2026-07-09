<script lang="ts">
    import { onMount } from "svelte";
    import { Events } from "@wailsio/runtime";

    import { GetServerPlayers } from "../../../bindings/seegolauncher/internal/services/api";
    import { Config } from "../../../bindings/seegolauncher/internal/services";
    import { Icons } from "../../utils/icons";
    import { base64ToBlob } from "../../utils/helper";
    import { stripMarkup } from "../../utils/string";
    import {
        initLocalization,
        locales,
        localization,
    } from "../../managers/localization";
    import {
        GetLatestNew,
        GetNewsImage,
    } from "../../../bindings/seegolauncher/internal/services/cacheservice";

    let showDialog: Boolean = false;
    let interval: ReturnType<typeof setInterval> | undefined;

    let serverPlayersBefore = 0;
    let serverSlotsBefore = 0;
    let serverAdminsBefore = 0;
    let serverQueueBefore = 0;
    let serverPlayersNow = 0;
    let serverSlotsNow = 0;
    let serverAdminsNow = 0;
    let serverQueueNow = 0;

    let latestNewsTitle = "";
    let latestNewsContent = "";
    let latestNewsImgSrc = "";

    function start() {
        if (!interval) {
            interval = setInterval(fetchServerStatus, 60 * 1000);
        }
    }

    function stop() {
        if (interval) {
            clearInterval(interval);
            interval = undefined;
        }
    }

    Events.On("home:startInterval", async (e) => {
        start();
    });

    Events.On("home:stopInterval", async (e) => {
        stop();
    });

    onMount(() => {
        void (async () => {
            await fetchServerStatus();
        })();

        initLocalization();
        setLatestNew();
        start();

        return () => {
            stop();
        };
    });

    function closeItemDialog() {
        showDialog = false;
    }

    async function fetchServerStatus() {
        let server;
        const serverStatus = document.getElementById("server-status");
        const serverFill = document.getElementById("server-fill");
        const estimatedConnection =
            document.getElementById("estimated-connect");
        let lang = await Config.GetLanguage();

        try {
            server = await GetServerPlayers();
        } catch (err) {
            serverStatus!.textContent =
                $locales[localization.serverStatusOffline];
            serverStatus!.style.color = "var(--gray)";
            serverFill!.style.background = "";
            return;
        }
        serverPlayersNow = server.players;
        serverSlotsNow = server.slots;
        serverAdminsNow = server.admins;
        serverQueueNow = server.queue;

        let connectionIn = "";

        if (serverPlayersNow != serverPlayersBefore) {
            const element = document.getElementById("players-count");
            if (element?.classList.contains("updated")) {
                element.classList.remove("updated");
            }
            void element?.offsetWidth;
            element?.classList.add("updated");
        }

        if (serverSlotsNow != serverSlotsBefore) {
            const element = document.getElementById("players-count");
            if (element?.classList.contains("updated")) {
                element.classList.remove("updated");
            }
            void element?.offsetWidth;
            element?.classList.add("updated");
        }

        if (serverAdminsNow != serverAdminsBefore) {
            const element = document.getElementById("admins-count");
            if (element?.classList.contains("updated")) {
                element.classList.remove("updated");
            }
            void element?.offsetWidth;
            element?.classList.add("updated");
        }

        if (serverQueueNow != serverQueueBefore) {
            const element = document.getElementById("queue-count");
            if (element?.classList.contains("updated")) {
                element.classList.remove("updated");
            }
            void element?.offsetWidth;
            element?.classList.add("updated");
        }

        if (serverStatus && serverFill) {
            if (serverPlayersNow <= 0 && serverQueueNow <= 0) {
                serverStatus.textContent =
                    $locales[localization.serverStatusOffline];
                serverStatus.style.color = "var(--gray)";
                serverFill.style.background = "";
            } else if (serverPlayersNow <= 0 && serverQueueNow > 0) {
                serverStatus.textContent =
                    $locales[localization.serverStatusRestart];
                serverStatus.style.color = "var(--orange)";
                serverFill.style.background = "var(--green)";
            } else if (
                serverPlayersNow === serverAdminsNow &&
                serverAdminsNow > 0
            ) {
                serverStatus.textContent =
                    $locales[localization.serverStatusMaintenance];
                serverStatus.style.color = "var(--orange)";
                serverFill.style.background = "var(--green)";
            } else {
                serverStatus.textContent =
                    $locales[localization.serverStatusOnline];
                serverStatus.style.color = "var(--green)";
                serverFill.style.background = "var(--green)";
            }

            await Events.Emit("feedback", "Fetched server status");
        }

        // without prio
        // about 1-3 players enters from the queue to the server in every 1 mins
        // update: i connected at 14:00 with 200 queue. After 2 hours, i was 150. My pos always changed between 140 and 160, so i will incrase the queue
        if (estimatedConnection) {
            if (serverPlayersNow >= serverSlotsNow) {
                // about 30-35 seconds one player disconnects // update up (changed to 50)
                // sometimes with the gold prio theres 100 queue, so without pro its like 30-35 seconds disconnect * 10
                if (serverQueueNow > 0) {
                    // base
                    let totalSeconds = serverQueueNow * 50;
                    if (serverQueueNow >= 100) {
                        // queue * prio ratio (35%) and + 120 seconds extra
                        // its quite close, because one member in discord waited 7 hours with 500 queue without prio
                        totalSeconds =
                            totalSeconds +
                            Math.floor(serverQueueNow * 0.35) * 120;
                    }
                    if (totalSeconds >= 3600) {
                        connectionIn = `${(totalSeconds / 3600).toFixed(1)} ${$locales[localization.hours]}`;
                    } else if (totalSeconds >= 60) {
                        connectionIn = `${(totalSeconds / 60).toFixed()} ${$locales[localization.minutes]}`;
                    } else {
                        connectionIn = `${Math.round(totalSeconds)} ${$locales[localization.seconds]}`;
                    }
                }
            } else {
                if (serverQueueNow > 0) {
                    // about 2 seconds one player connects
                    const totalSeconds = serverQueueNow * 2;
                    if (totalSeconds >= 3600) {
                        connectionIn = `${(totalSeconds / 3600).toFixed(1)} ${$locales[localization.hours]}`;
                    } else if (totalSeconds >= 60) {
                        connectionIn = `${(totalSeconds / 60).toFixed()} ${$locales[localization.minutes]}`;
                    } else {
                        connectionIn = `${Math.round(totalSeconds)} ${$locales[localization.seconds]}`;
                    }
                } else {
                    estimatedConnection.textContent = "";
                }
            }

            if (serverQueueNow > 0) {
                estimatedConnection.textContent =
                    `${$locales[localization.serverStatusEstimated]}: ` +
                    connectionIn;
            }
        }

        serverPlayersBefore = serverPlayersNow;
        serverSlotsBefore = serverSlotsNow;
        serverAdminsBefore = serverAdminsNow;
        serverQueueBefore = serverQueueNow;
    }

    async function setLatestNew() {
        try {
            const news = await GetLatestNew();

            latestNewsTitle = news!.Title;
            latestNewsContent = stripMarkup(news!.Content);

            if (news?.ImageName) {
                const imageData = await GetNewsImage(news.ImageName);

                const blob = base64ToBlob(imageData);

                latestNewsImgSrc = blob;
            }
        } catch (err) {
            Events.Emit("feedback", `Failed to load latest new: ${err}`);
        }
    }
</script>

<main>
    <div id="home-view">
        <div class="feed-layout">
            <div
                id="hero-card"
                class="hero-card"
                style="background-image: url('{latestNewsImgSrc}')"
            >
                <span class="hero-overlay"></span>
                <div class="hero-content">
                    <span class="news-badge"
                        >{$locales[localization.newsLatest]}</span
                    >
                    <p id="hero-news-title" class="news-title">
                        {latestNewsTitle}
                    </p>
                    <p id="hero-news-comment" class="news-comment">
                        {latestNewsContent}
                    </p>
                    <button
                        class="button news-read interactive"
                        id="hero-news-read-latest"
                        >{$locales[localization.newsRead]}</button
                    >
                </div>
            </div>

            <div class="home-items">
                <h3 class="text items-title">
                    SeeRPG Club {$locales[localization.clubMembership]}
                </h3>
                <div id="items-container">
                    <!-->{#each membershipItems as item}
                        <div class="item-square-card">
                            <div class="item-header">
                                <span class="text item-tag">7 {day}</span>
                            </div>
                            <div class="item-image">
                                <img src={item.image} alt={item.name} />
                            </div>
                            <div class="item-name">
                                <h4 class="text">{item.name.split(" ")[0]}</h4>
                                <span class="item-price"
                                    >{item.total_price} {item.currency}</span
                                >
                            </div>
                            <button class="button add-to-card">
                                {@html Icons.Launcher.Cart}
                            </button>
                        </div>
                    {/each} <-->
                </div>
            </div>
        </div>

        <aside class="sidebar">
            <div class="widget">
                <div class="status-container">
                    <div class="status-info">
                        <span id="server-status" class="status-text"></span>
                        <span id="players-count" class="text player-count"
                            >{serverPlayersNow} / {serverSlotsNow}</span
                        >
                    </div>
                </div>

                <div id="server-fill" class="progress-bar">
                    <div
                        class="progress-bar-fill"
                        style="width: {(serverPlayersNow / serverSlotsNow) *
                            100}%"
                    ></div>
                </div>

                <div class="stats">
                    <div class="stat-box">
                        <span class="text stat-title"
                            >{$locales[localization.serverStatusAdmins]}</span
                        >
                        <span id="admins-count" class="stat-value"
                            >{serverAdminsNow}</span
                        >
                    </div>
                    <div class="stat-box">
                        <span class="text stat-title"
                            >{$locales[localization.serverStatusQueue]}</span
                        >
                        <span id="queue-count" class="stat-value"
                            >{serverQueueNow}</span
                        >
                    </div>
                </div>
                <p id="estimated-connect" class="comment estimated-connect"></p>
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
                    </button>
                    <button
                        onclick={() =>
                            window._openURL(
                                "https://www.facebook.com/seerpgofficial",
                            )}
                        class="socials interactive"
                        title="Facebook"
                        >{@html Icons.Community.FaceBook}
                    </button>
                    <button
                        onclick={() =>
                            window._openURL(
                                "https://www.facebook.com/seerpgofficial",
                            )}
                        class="socials interactive"
                        title="TikTok"
                        >{@html Icons.Community.TikTok}
                    </button>
                    <button
                        onclick={() =>
                            window._openURL(
                                "https://www.youtube.com/@seerpgofficial",
                            )}
                        class="socials interactive"
                        title="YouTube"
                        >{@html Icons.Community.YouTube}
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
