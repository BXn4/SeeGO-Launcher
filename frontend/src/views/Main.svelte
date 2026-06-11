<script lang="ts">
    import Titlebar from "./partials/Titlebar.svelte";
    import Navbar from "./partials/Navbar.svelte";
    import { onMount } from "svelte";
    import { getCategories, getItems } from "../lib/api";
    import { Browser } from "@wailsio/runtime";
    import { GetServerPlayers } from "../../bindings/seegolauncher/internal/services/api";
    import {
        Config,
        Localization,
    } from "../../bindings/seegolauncher/internal/services";
    import { Icons } from "../utils/icons";

    let serverPlayersBefore = -1;
    let serverSlotsBefore = -1;
    let serverAdminsBefore = -1;
    let serverQueueBefore = -1;
    let serverPlayersNow = -1;
    let serverSlotsNow = -1;
    let serverAdminsNow = -1;
    let serverQueueNow = -1;

    interface Category {
        id: number;
        name: string;
        slug: string;
    }

    interface Item {
        id: number;
        name: string;
        description: string;
        image: string;
        total_price: number;
        currency: string;
        created_at: string;
    }

    let categories: Category[] = [];
    let membershipItems: Item[] = [];
    let membershipCategory: Category | null = null;

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

    onMount(async () => {
        await setLocales();
        await fetchServerStatus();
        categories = (await getCategories()) as Category[];
        membershipCategory =
            categories.find((c) => c.slug === "seerpg-club-tagság-a") ?? null;

        if (membershipCategory) {
            membershipItems = (await getItems(membershipCategory.id)) as Item[];
        }

        setInterval(fetchServerStatus, 60 * 1000);

        // await Events.Emit("app-ready", null);
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
        if (estimatedConnection) {
            if (serverPlayersNow >= serverSlotsNow) {
                // about 30-35 seconds one player disconnects
                // sometimes with the gold prio theres 100 queue, so without pro its like 30-35 seconds disconnect * 10
                if (serverQueueNow > 0) {
                    // base
                    let totalSeconds = serverQueueNow * 32;
                    if (serverQueueNow >= 100) {
                        // queue * prio ratio (35%) and + 60 seconds extra
                        // its quite close, because one member in discord waited 7 hours with 500 queue without prio
                        totalSeconds =
                            totalSeconds +
                            Math.floor(serverQueueNow * 0.35) * 60;
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

            estimatedConnection.textContent = `${estimated} ` + connectionIn;
        }

        serverPlayersBefore = serverPlayersNow;
        serverSlotsBefore = serverSlotsNow;
        serverAdminsBefore = serverAdminsNow;
        serverQueueBefore = serverQueueNow;
    }
</script>

<main>
    <Titlebar />
    <Navbar />

    <div id="home-view">
        <div class="feed-layout">
            <header class="hero-card">
                <span class="hero-overlay"></span>
                <div class="hero-content">
                    <span class="badge">{latest}</span>
                    <h2>
                        Kényelmesebb játékélmény és hasznos javítások érkeztek!
                    </h2>
                    <p>
                        Frissítés érkezett a szerverre! A legújabb frissítésben
                        több fontos kényelmi fejlesztés, hibajavítás és tartalmi
                        bővítés is bekerült a szerverre.
                    </p>
                    <button class="button hero-news" id="hero-news-read-latest"
                        >{read}</button
                    >
                </div>
            </header>

            <div class="items">
                <h3 class="items-title">SeeRPG Club {membership}</h3>
                <div id="items-container">
                    {#each membershipItems as item}
                        <div class="item-square-card" title={item.name}>
                            <div class="item-header">
                                <span class="item-tag">7 {day}</span>
                            </div>
                            <div class="item-image">
                                <img src={item.image} alt={item.name} />
                            </div>
                            <div class="item-name">
                                <h4>{item.name.split(" ")[0]}</h4>
                                <span class="item-price"
                                    >{item.total_price} {item.currency}</span
                                >
                            </div>
                        </div>
                    {/each}
                </div>
            </div>
        </div>

        <aside class="sidebar">
            <div class="widget">
                <div class="status-container">
                    <div class="status-info">
                        <span id="server-status" class="status-text"></span>
                        <span id="players-count" class="player-count"
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
                        <span class="stat-title">{admins}</span>
                        <span id="admins-count" class="stat-value"
                            >{serverAdminsNow}</span
                        >
                    </div>
                    <div class="stat-box">
                        <span class="stat-title">{queue}</span>
                        <span id="queue-count" class="stat-value"
                            >{serverQueueNow}</span
                        >
                    </div>
                </div>
                <p id="estimated-connect" class="estimated-connect"></p>
            </div>
            <div class="widget">
                <h3 class="widget-title">{community}</h3>
                <div class="social-links-grid">
                    <button
                        on:click={() =>
                            Browser.OpenURL(
                                "https://discord.com/invite/seerpg",
                            )}
                        class="button social"
                        title="Discord"
                        >{@html Icons.Community.Discord}
                    </button>
                    <button
                        on:click={() =>
                            Browser.OpenURL(
                                "https://www.facebook.com/seerpgofficial",
                            )}
                        class="button social"
                        title="Facebook"
                        >{@html Icons.Community.FaceBook}
                    </button>
                    <button
                        on:click={() =>
                            Browser.OpenURL(
                                "https://www.facebook.com/seerpgofficial",
                            )}
                        class="button social"
                        title="TikTok"
                        >{@html Icons.Community.TikTok}
                    </button>
                    <button
                        on:click={() =>
                            Browser.OpenURL(
                                "https://www.youtube.com/@seerpgofficial",
                            )}
                        class="button social"
                        title="YouTube"
                        >{@html Icons.Community.YouTube}
                    </button>
                </div>
            </div>
            <div class="widget">
                <h3 class="widget-title">{launcherStatus}</h3>
                <button id="connect-button" class="button connect"
                    ><!-->{@html Icons.Launcher.Play}<--->{launcherConnect}</button
                >
            </div>
        </aside>
    </div>
</main>

<style>
    main {
        width: 100vw;
        height: 100vh;
        display: grid;
        grid-template-columns: var(--width, 240px) 1fr;
        grid-template-rows: var(--height, 48px) 1fr;
        overflow: hidden;
        color: var(--text);
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

    #home-view {
        grid-column: 2;
        grid-row: 2;
        padding: 32px;
        display: grid;
        grid-template-columns: 1fr 280px;
        gap: 32px;
        overflow-y: auto;
    }

    .feed-layout {
        display: flex;
        flex-direction: column;
        gap: 32px;
    }

    .hero-card {
        position: relative;
        min-height: 200px;
        border-radius: 24px;
        background: url("https://news.see-rpg.com/img/S2YAXN7DQbg7loLhgpMl.png")
            center/cover;
        border: 1px solid var(--border);
        overflow: hidden;
        display: flex;
        align-items: flex-end;
        padding: 48px;
        box-shadow: 0 20px 40px -15px rgba(0, 0, 0, 0.6);
    }

    .hero-overlay {
        position: absolute;
        inset: 0;
        background: linear-gradient(
            to top,
            rgba(11, 15, 23, 0.8) 40%,
            rgba(11, 15, 23, 0.4) 100%
        );
        z-index: 1;
    }

    .hero-content {
        position: relative;
        z-index: 2;
        max-width: 600px;
    }

    .badge {
        background: var(--green);
        color: var(--black);
        font-size: 12px;
        padding: 4px 8px;
        border-radius: 8px;
        text-transform: uppercase;
        letter-spacing: 0.5px;
        display: inline-block;
        margin-bottom: 12px;
        font-weight: bold;
    }

    .hero-content h2 {
        margin: 0 0 8px 0;
        font-size: 32px;
    }

    .hero-content p {
        margin: 0 0 20px 0;
        color: var(--comment);
    }

    .button.hero-news {
        background: var(--orange);
        color: var(--black);
        border: none;
        padding: 10px 20px;
        border-radius: 12px;
        cursor: pointer;
        font-size: 14px;
    }

    .widget {
        background: var(--surface1);
        border: 1px solid var(--border);
        border-radius: 16px;
        padding: 24px;
        display: flex;
        flex-direction: column;
        gap: 27px;
    }

    .status-container {
        display: flex;
        align-items: center;
        gap: 12px;
    }

    .status-info {
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-grow: 1;
    }

    .status-text {
        font-size: 16px;
    }

    .player-count {
        font-size: 14px;
        color: var(--text);
    }

    .progress-bar {
        width: 100%;
        height: 6px;
        background: rgba(255, 255, 255, 0.05);
        border-radius: 10px;
        overflow: hidden;
    }

    .progress-bar-fill {
        height: 100%;
        border-radius: 10px;
    }

    .stats {
        display: grid;
        grid-template-columns: 1fr 1fr;
        gap: 12px;
        margin-top: 4px;
    }

    .stat-box {
        background: rgba(255, 255, 255, 0.02);
        border: 1px solid var(--border);
        padding: 12px;
        border-radius: 10px;
        display: flex;
        flex-direction: column;
        gap: 4px;
    }

    .stat-title {
        font-size: 12px;
        color: var(--text);
        text-transform: uppercase;
    }

    .stat-value {
        font-size: 14px;
        color: var(--comment);
    }

    .estimated-connect {
        font-size: 12px;
        margin-top: -12px;
        margin-bottom: -8px;
        color: var(--comment);
    }

    .items {
        display: flex;
        flex-direction: column;
        gap: 16px;
    }

    .items-title {
        margin: 0;
        font-size: 18px;
        text-transform: uppercase;
        letter-spacing: 1.5px;
        color: var(--text);
    }

    #items-container {
        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
        gap: 16px;
    }

    .item-square-card {
        background: var(--surface1);
        border: 1px solid var(--border);
        border-radius: 14px;
        padding: 16px;
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        aspect-ratio: 1 / 1;
        box-sizing: border-box;
        cursor: pointer;
    }

    .item-tag {
        font-size: 10px;
        color: var(--text);
        background: rgba(255, 255, 255, 0.06);
        padding: 4px 8px;
        border-radius: 8px;
    }

    .item-image {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-grow: 1;
        margin: 8px 0;
    }

    .item-image img {
        max-height: 65px;
        max-width: 100%;
        object-fit: contain;
        filter: drop-shadow(0 2px 2px rgba(0, 0, 0, 0.5));
        transition: transform 0.2s;
    }

    .item-square-card:hover .item-image img {
        transform: scale(1.2);
    }

    .item-name {
        display: flex;
        flex-direction: column;
        gap: 2px;
    }

    .item-name h4 {
        margin: 0;
        font-size: 14px;
        color: var(--text);
    }

    .item-price {
        font-size: 12px;
        color: var(--green);
    }

    .sidebar {
        display: flex;
        flex-direction: column;
        gap: 24px;
    }

    .widget-title {
        margin: 0;
        font-size: 12px;
        text-transform: uppercase;
        letter-spacing: 1px;
        color: var(--text-muted);
    }

    .social-links-grid {
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        gap: 10px;
    }

    .button.social {
        aspect-ratio: 1;
        background: rgba(255, 255, 255, 0.03);
        border: 1px solid var(--border);
        border-radius: 12px;
        display: flex;
        align-items: center;
        justify-content: center;
        color: var(--comment);
        text-decoration: none;
        width: 48px;
        height: 48px;
    }

    :global(.button.social svg) {
        width: 22px;
        height: 22px;
    }

    .button.social:hover {
        color: var(--dim1);
    }

    .button.connect {
        border: none;
        padding: 24px;
        background-color: var(--darker-green);
        border-radius: 16px;
    }

    .button:hover {
        cursor: pointer;
        transform: scale(1.05);
    }

    @media (max-width: 900px) {
        #home-view {
            grid-template-columns: 1fr;
        }
        .sidebar {
            order: 0;
        }
    }
    @keyframes refresh {
        0% {
            opacity: 0;
            transform: scale(0.97);
        }
        100% {
            opacity: 1;
            transform: scale(1);
        }
    }

    :global(.updated) {
        animation: refresh 0.5s ease-out;
    }
</style>
