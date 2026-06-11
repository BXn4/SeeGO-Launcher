<script lang="ts">
    import Titlebar from "./partials/Titlebar.svelte";
    import Navbar from "./partials/Navbar.svelte";
    import { onMount } from "svelte";
    import { getCategories, getItems } from "../lib/api";
    import { Events } from "@wailsio/runtime";
    import { GetServerPlayers } from "../../bindings/seegolauncher/internal/services/api";

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

    onMount(async () => {
        await fetchServerStatus();
        categories = (await getCategories()) as Category[];
        membershipCategory =
            categories.find((c) => c.slug === "seerpg-club-tagság-a") ?? null;

        if (membershipCategory) {
            membershipItems = (await getItems(membershipCategory.id)) as Item[];
        }

        setInterval(fetchServerStatus, 60 * 1000);

        await Events.Emit("app-ready", null);
    });

    async function fetchServerStatus() {
        const server = await GetServerPlayers();
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

        const serverStatus = document.getElementById("server-status");
        const serverFill = document.getElementById("server-fill");
        const estimatedConnection =
            document.getElementById("estimated-connect");

        if (serverStatus && serverFill) {
            if (serverPlayersNow <= 0 && serverQueueNow <= 0) {
                serverStatus.textContent = "Offline";
                serverStatus.style.color = "var(--gray)";
                serverFill.style.background = "";
            } else if (serverPlayersNow <= 0 && serverQueueNow > 0) {
                serverStatus.textContent = "Újraindítás";
                serverStatus.style.color = "var(--orange)";
                serverFill.style.background = "var(--green)";
            } else if (
                serverPlayersNow === serverAdminsNow &&
                serverAdminsNow > 0
            ) {
                serverStatus.textContent = "Karbantartás";
                serverStatus.style.color = "var(--orange)";
                serverFill.style.background = "var(--green)";
            } else {
                serverStatus.textContent = "Online";
                serverStatus.style.color = "var(--green)";
                serverFill.style.background = "var(--green)";
            }
        }

        if (estimatedConnection) {
            if (serverPlayersNow >= serverSlotsNow) {
                // about 30-35 seconds one player disconnects
                // sometimes with the gold prio theres 100 queue, so without pro its like 30-35 seconds disconnect * 10
                if (serverQueueNow > 0) {
                    const totalSeconds = serverQueueNow * 32;
                    if (totalSeconds >= 3600) {
                        connectionIn = `${(totalSeconds / 3600).toFixed()} óra`;
                    } else if (totalSeconds >= 60) {
                        connectionIn = `${(totalSeconds / 60).toFixed()} perc`;
                    } else {
                        connectionIn = `${Math.round(totalSeconds)} másodperc`;
                    }
                }
            } else {
                if (serverQueueNow > 0) {
                    // about 2 seconds one player connects
                    const totalSeconds = serverQueueNow * 2;
                    if (totalSeconds >= 3600) {
                        connectionIn = `${(totalSeconds / 3600).toFixed()} óra`;
                    } else if (totalSeconds >= 60) {
                        connectionIn = `${(totalSeconds / 60).toFixed()} perc`;
                    } else {
                        connectionIn = `${Math.round(totalSeconds)} másodperc`;
                    }
                } else {
                    estimatedConnection.textContent = "";
                }
            }

            estimatedConnection.textContent =
                "Becsült csatlakozás: " + connectionIn;
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
                    <span class="badge">Friss</span>
                    <h2>
                        Kényelmesebb játékélmény és hasznos javítások érkeztek!
                    </h2>
                    <p>
                        Frissítés érkezett a szerverre! A legújabb frissítésben
                        több fontos kényelmi fejlesztés, hibajavítás és tartalmi
                        bővítés is bekerült a szerverre.
                    </p>
                    <button class="hero-news-button" id="hero-news-read-latest"
                        >Elolvasom</button
                    >
                </div>
            </header>

            <div class="items">
                <h3 class="items-title">SeeRPG Club tagság</h3>
                <div id="items-container">
                    {#each membershipItems as item}
                        <div class="item-square-card" title={item.name}>
                            <div class="item-header">
                                <span class="item-tag">7 nap</span>
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
                        <span class="stat-title">Adminok</span>
                        <span id="admins-count" class="stat-value"
                            >{serverAdminsNow}</span
                        >
                    </div>
                    <div class="stat-box">
                        <span class="stat-title">Várólistán</span>
                        <span id="queue-count" class="stat-value"
                            >{serverQueueNow}</span
                        >
                    </div>
                </div>
                <p id="estimated-connect" class="estimated-connect"></p>
            </div>
            <div class="widget">
                <h3 class="widget-title">Közösség</h3>
                <div class="social-links-grid">
                    <a href="/" class="social-box" title="Discord">
                        <svg viewBox="0 0 24 24" fill="currentColor"
                            ><path
                                d="M18.59 5.89c-1.23-.57-2.54-.99-3.92-1.23c-.17.3-.37.71-.5 1.04c-1.46-.22-2.91-.22-4.34 0c-.14-.33-.34-.74-.51-1.04c-1.38.24-2.69.66-3.92 1.23c-2.48 3.74-3.15 7.39-2.82 10.98c1.65 1.23 3.24 1.97 4.81 2.46c.39-.53.73-1.1 1.03-1.69c-.57-.21-1.11-.48-1.62-.79c.14-.1.27-.21.4-.31c3.13 1.46 6.52 1.46 9.61 0c.13.11.26.21.4.31c-.51.31-1.06.57-1.62.79c.3.59.64 1.16 1.03 1.69c1.57-.49 3.17-1.23 4.81-2.46c.39-4.17-.67-7.78-2.82-10.98Zm-9.75 8.78c-.94 0-1.71-.87-1.71-1.94s.75-1.94 1.71-1.94s1.72.87 1.71 1.94c0 1.06-.75 1.94-1.71 1.94m6.31 0c-.94 0-1.71-.87-1.71-1.94s.75-1.94 1.71-1.94s1.72.87 1.71 1.94c0 1.06-.75 1.94-1.71 1.94"
                            ></path></svg
                        >
                    </a>
                    <a href="/" class="social-box" title="Facebook">
                        <svg viewBox="0 0 24 24" fill="currentColor"
                            ><path
                                d="M9.198 21.5h4v-8.01h3.604l.396-3.98h-4V7.5a1 1 0 0 1 1-1h3v-4h-3a5 5 0 0 0-5 5v2.01h-2l-.396 3.98h2.396z"
                            ></path></svg
                        >
                    </a>
                    <a href="/" class="social-box" title="TikTok">
                        <svg viewBox="0 0 16 16" fill="currentColor"
                            ><path
                                d="M8.3 1.01c.75-.01 1.5 0 2.25-.01c.05.89.36 1.8 1.01 2.43c.64.65 1.55.94 2.44 1.04v2.35c-.83-.03-1.66-.2-2.42-.56c-.33-.15-.63-.34-.93-.54c0 1.7 0 3.41-.01 5.1c-.04.82-.31 1.63-.78 2.3c-.75 1.12-2.06 1.85-3.4 1.87c-.82.05-1.65-.18-2.35-.6c-1.16-.69-1.98-1.97-2.1-3.33q-.03-.435 0-.87c.1-1.11.65-2.17 1.49-2.89c.95-.84 2.29-1.24 3.54-1c.01.86-.02 1.73-.02 2.59c-.57-.19-1.24-.13-1.74.22c-.37.24-.64.6-.79 1.02c-.12.3-.09.62-.08.94c.14.96 1.05 1.76 2.01 1.67c.64 0 1.26-.39 1.59-.94c.11-.19.23-.39.24-.62c.06-1.04.03-2.08.04-3.13c0-2.35 0-4.7.01-7.04"
                            ></path></svg
                        >
                    </a>
                    <a href="/" class="social-box" title="YouTube">
                        <svg viewBox="0 0 24 24" fill="currentColor"
                            ><path
                                d="M12.006 19.012h-.02c-.062 0-6.265-.012-7.83-.437a2.5 2.5 0 0 1-1.764-1.765A26.494 26.494 0 0 1 1.986 12a26.646 26.646 0 0 1 .417-4.817A2.564 2.564 0 0 1 4.169 5.4c1.522-.4 7.554-.4 7.81-.4H12c.063 0 6.282.012 7.831.437c.859.233 1.53.904 1.762 1.763c.29 1.594.427 3.211.407 4.831a26.568 26.568 0 0 1-.418 4.811a2.51 2.51 0 0 1-1.767 1.763c-1.52.403-7.553.407-7.809.407Zm-2-10.007l-.005 6l5.212-3l-5.207-3Z"
                            ></path></svg
                        >
                    </a>
                </div>
            </div>
            <div class="widget"></div>
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

    .hero-news-button {
        background: var(--orange);
        color: var(--black);
        border: none;
        padding: 10px 20px;
        border-radius: 12px;
        cursor: pointer;
        font-size: 14px;
    }

    .hero-news-button:hover {
        transform: scale(1.05);
    }

    .widget {
        background: var(--surface1);
        border: 1px solid var(--border);
        border-radius: 16px;
        padding: 24px;
        display: flex;
        flex-direction: column;
        gap: 16px;
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
        margin-top: -4px;
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

    .social-box {
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

    .social-box svg {
        width: 22px;
        height: 22px;
    }

    .social-box:hover {
        color: var(--dim1);
    }

    .social-box:hover svg {
        transform: scale(1.1);
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
