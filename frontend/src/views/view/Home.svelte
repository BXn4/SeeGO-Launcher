<script lang="ts">
    import { onMount } from "svelte";
    import { GetServerPlayers } from "../../../bindings/seegolauncher/internal/services/api";
    import {
        Config,
        Localization,
    } from "../../../bindings/seegolauncher/internal/services";
    import { Icons } from "../../utils/icons";

    let serverPlayersBefore = -1;
    let serverSlotsBefore = -1;
    let serverAdminsBefore = -1;
    let serverQueueBefore = -1;
    let serverPlayersNow = -1;
    let serverSlotsNow = -1;
    let serverAdminsNow = -1;
    let serverQueueNow = -1;

    let latest: string = "";
    let read: string = "";
    let membership: string = "";
    let day: string = "";
    let serverStatus: string = "";
    let admins: string = "";
    let queue: string = "";
    let estimated: string = "";
    let hours: string = "";
    let minutes: string = "";
    let seconds: string = "";
    let community: string = "";
    let launcherStatus: string = "";
    let launcherConnect: string = "";

    let itemDialogName: string = "";
    let itemDialogDesc: string = "";
    let itemDialogImage: string = "";
    let itemDialogPrice: string = "";
    let itemDialogCurrency: string = "";

    let showDialog: Boolean = false;

    onMount(() => {
        const interval = setInterval(fetchServerStatus, 60 * 1000);

        void (async () => {
            await setLocales();
            await fetchServerStatus();
        })();

        return () => {
            clearInterval(interval);
        };
    });

    async function setLocales() {
        let lang = await Config.GetLanguage();

        [
            latest,
            read,
            membership,
            day,
            serverStatus,
            admins,
            queue,
            estimated,
            hours,
            minutes,
            seconds,
            community,
            launcherStatus,
            launcherConnect,
        ] = await Promise.all([
            Localization.Get("news-latest", lang),
            Localization.Get("news-read", lang),
            Localization.Get("club-membership", lang),
            Localization.Get("club-membership-day", lang),
            Localization.Get("server-status-online", lang),
            Localization.Get("server-status-admins", lang),
            Localization.Get("server-status-queue", lang),
            Localization.Get("server-status-estimated", lang),
            Localization.Get("hours", lang),
            Localization.Get("minutes", lang),
            Localization.Get("seconds", lang),
            Localization.Get("community", lang),
            Localization.Get("launcher-ready", lang),
            Localization.Get("launcher-connect", lang),
        ]);
    }

    function closeItemDialog() {
        showDialog = false;
    }

    async function fetchServerStatus() {
        const server = await GetServerPlayers();
        serverPlayersNow = server.players;
        serverSlotsNow = server.slots;
        serverAdminsNow = server.admins;
        serverQueueNow = server.queue;

        let connectionIn = "";
        let lang = await Config.GetLanguage();

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

        const serverStatus = document.getElementById("server-status");
        const serverFill = document.getElementById("server-fill");
        const estimatedConnection =
            document.getElementById("estimated-connect");

        if (serverStatus && serverFill) {
            if (serverPlayersNow <= 0 && serverQueueNow <= 0) {
                serverStatus.textContent = await Localization.Get(
                    "server-status-offline",
                    lang,
                );
                serverStatus.style.color = "var(--gray)";
                serverFill.style.background = "";
            } else if (serverPlayersNow <= 0 && serverQueueNow > 0) {
                serverStatus.textContent = await Localization.Get(
                    "server-status-restart",
                    lang,
                );
                serverStatus.style.color = "var(--orange)";
                serverFill.style.background = "var(--green)";
            } else if (
                serverPlayersNow === serverAdminsNow &&
                serverAdminsNow > 0
            ) {
                serverStatus.textContent = await Localization.Get(
                    "server-status-maintenance",
                    lang,
                );
                serverStatus.style.color = "var(--orange)";
                serverFill.style.background = "var(--green)";
            } else {
                serverStatus.textContent = await Localization.Get(
                    "server-status-online",
                    lang,
                );
                serverStatus.style.color = "var(--green)";
                serverFill.style.background = "var(--green)";
            }
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
                        connectionIn = `${(totalSeconds / 3600).toFixed(1)} ${hours}`;
                    } else if (totalSeconds >= 60) {
                        connectionIn = `${(totalSeconds / 60).toFixed()} ${minutes}`;
                    } else {
                        connectionIn = `${Math.round(totalSeconds)} ${seconds}`;
                    }
                }
            } else {
                if (serverQueueNow > 0) {
                    // about 2 seconds one player connects
                    const totalSeconds = serverQueueNow * 2;
                    if (totalSeconds >= 3600) {
                        connectionIn = `${(totalSeconds / 3600).toFixed(1)} ${hours}`;
                    } else if (totalSeconds >= 60) {
                        connectionIn = `${(totalSeconds / 60).toFixed()} ${minutes}`;
                    } else {
                        connectionIn = `${Math.round(totalSeconds)} ${seconds}`;
                    }
                } else {
                    estimatedConnection.textContent = "";
                }
            }

            if (serverQueueNow > 0) {
                estimatedConnection.textContent =
                    `${estimated}: ` + connectionIn;
            }
        }

        serverPlayersBefore = serverPlayersNow;
        serverSlotsBefore = serverSlotsNow;
        serverAdminsBefore = serverAdminsNow;
        serverQueueBefore = serverQueueNow;
    }
</script>

<main>
    <div id="home-view">
        <div class="feed-layout">
            <header class="hero-card">
                <span class="hero-overlay"></span>
                <div class="hero-content">
                    <span class="badge">{latest}</span>
                    <p class="news-title">
                        Kényelmesebb játékélmény és hasznos javítások érkeztek!
                    </p>
                    <p class="news-comment">
                        Frissítés érkezett a szerverre! A legújabb frissítésben
                        több fontos kényelmi fejlesztés, hibajavítás és tartalmi
                        bővítés is bekerült a szerverre.
                    </p>
                    <button
                        class="button news-read interactive"
                        id="hero-news-read-latest">{read}</button
                    >
                </div>
            </header>

            <div class="home-items">
                <h3 class="text items-title">SeeRPG Club {membership}</h3>
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
                        <span class="text stat-title">{admins}</span>
                        <span id="admins-count" class="stat-value"
                            >{serverAdminsNow}</span
                        >
                    </div>
                    <div class="stat-box">
                        <span class="text stat-title">{queue}</span>
                        <span id="queue-count" class="stat-value"
                            >{serverQueueNow}</span
                        >
                    </div>
                </div>
                <p id="estimated-connect" class="comment estimated-connect"></p>
            </div>
            <div class="widget">
                <h3 class="text widget-title">{community}</h3>
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
                <h3 class="text widget-title">{launcherStatus}</h3>
                <button
                    id="connect-button"
                    class="text button connect interactive"
                    ><!-->{@html Icons.Launcher.Play}<--->{launcherConnect}</button
                >
            </div>
        </aside>
    </div>
    {#if showDialog}
        <div
            class="dialog-backdrop"
            role="button"
            tabindex="0"
            onclick={closeItemDialog}
            onkeydown={(e) => {
                if (e.key === "Enter" || e.key === " ") {
                    closeItemDialog();
                }
            }}
        >
            <div class="dialog">
                <img src={itemDialogImage} alt={itemDialogName} />
                <h2>{itemDialogName}</h2>
                {@html itemDialogDesc}
                <span class="item-price">
                    {itemDialogPrice}
                    {itemDialogCurrency}
                </span>
            </div>
        </div>
    {/if}
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
